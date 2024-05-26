package pkg

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"lab4/lab1"
	"lab4/lab2"
	"lab4/lab3"
	"log"
	"os"
	"strings"
	"time"
)

var Keys [][8]uint32

var seed = []byte{
	0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0x17, 0x28,
	0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0x17, 0x28,
	0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0x17, 0x28,
	0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0x17, 0x28,
}

func WriteToFile(msg []byte, file *os.File) {
	_, err := file.WriteString(hex.EncodeToString(msg) + "\n")
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}
}

func GetICV(msg []byte, key [8]uint32) []byte {
	var ICV []byte
	byteSlice := make([]byte, 8)
	var uint64Slice []uint64
	if len(msg)%8 != 0 {
		newBytes := make([]byte, (len(msg)/8+1)*8-len(msg))
		msg = append(msg, newBytes...)
	}
	for i := 0; i < len(msg); i += 8 {
		// Преобразование каждой группы байт в uint64
		uint64Value := binary.BigEndian.Uint64(msg[i : i+8])
		// Добавление uint64 в срез
		uint64Slice = append(uint64Slice, uint64Value)
	}
	binary.BigEndian.PutUint64(byteSlice, lab1.Omac(uint64Slice, key, int64(len(uint64Slice))))
	ICV = append(ICV, byteSlice[4:]...)
	return ICV
}

func SetHeader(SeqNum uint64) []byte {
	var EVCK uint32
	var Header []byte
	EVCK += 0 << 31 // ExternalKeyIdFlag
	// Version 0
	EVCK += 0xf3 << 8 // CS
	EVCK += 0x80      // KeyID
	byteSlice4 := make([]byte, 4)
	byteSlice8 := make([]byte, 8)
	binary.BigEndian.PutUint32(byteSlice4, EVCK)
	binary.BigEndian.PutUint64(byteSlice8, SeqNum)
	Header = append(Header, byteSlice4...)
	Header = append(Header, byteSlice8[2:]...)
	return Header // 10 Byte
}

func MakeMessage(Payload []byte, SeqNum uint64, key [8]uint32) []byte {
	var msg []byte
	Header := SetHeader(SeqNum)
	fmt.Printf("[#%d] Header [%d bytes]: %x\n", SeqNum, len(Header), Header)
	msg = append(msg, Header...)
	fmt.Printf("[#%d] Payload [%d bytes | %d blocks]: %x\n", SeqNum, len(Payload), len(Payload)/8, Payload)
	msg = append(msg, Payload...)
	ICV := GetICV(msg, key)
	fmt.Printf("[#%d] ICV [%d bytes]: %x\n", SeqNum, len(ICV), ICV)
	msg = append(msg, ICV...)
	return msg
}

func GetHeaderParam(msg []byte) (uint32, uint32, uint32, uint32, uint64) {
	Header := msg[:10]
	EVCK := binary.BigEndian.Uint32(Header[:4])
	SeqNumByte := make([]byte, 2)
	SeqNumByte = append(SeqNumByte, Header[4:]...)
	SeqNum := binary.BigEndian.Uint64(SeqNumByte)
	fmt.Printf("[#%d] Header [%d bytes]: %x\n", SeqNum, len(Header), Header)
	EKIF := EVCK >> 31
	//fmt.Printf("[#%d] ExternalKeyIdFlag: %x\n", SeqNum, EKIF)

	Version := EVCK >> 16 & ^uint32(1<<15)
	//fmt.Printf("[#%d] Version: %x\n", SeqNum, Version)
	CS := EVCK >> 8 & uint32(0xff)
	//fmt.Printf("[#%d] CS: %x\n", SeqNum, CS)
	KeyId := EVCK & uint32(0xff)
	//fmt.Printf("[#%d] KeyID: %x\n", SeqNum, KeyId)
	return EKIF, Version, CS, KeyId, SeqNum
}

func CheckConsistency(msg []byte, key [8]uint32) bool {
	MessageICV := make([]byte, 4)
	copy(MessageICV, msg[len(msg)-4:])
	msgnoICV := make([]byte, len(msg)-4)
	copy(msgnoICV, msg[:len(msg)-4])
	if !bytes.Equal(MessageICV, GetICV(msgnoICV, key)) {
		return false
	}
	return true
}
func GetDecryptedPayloadData(msg []byte, key [8]uint32) []uint64 {
	var Payload []byte
	Payload = msg[10 : len(msg)-4]
	var PayloadData []uint64

	// Проходим по срезу байт с шагом в 8 байт
	for i := 0; i < len(Payload); i += 8 {
		// Преобразуем каждые 8 байт в uint64
		uint64Value := binary.BigEndian.Uint64(Payload[i : i+8])
		// Добавляем uint64 в срез
		PayloadData = append(PayloadData, uint64Value)
	}
	var Data []uint64
	for i := 0; i < len(PayloadData); i++ {
		Data = append(Data, lab1.Decrypt(PayloadData[i], key))
	}
	return Data
}

func ConvertUint64Byte(num uint64) []byte {
	byteSlice := make([]byte, 8)
	binary.BigEndian.PutUint64(byteSlice, num)
	return byteSlice
}

func CreateKey(generator *lab3.XoroShiro256PlusPlus) {
	var temp [8]uint32
	kdfResult := lab2.Generator(ConvertUint64Byte(generator.Next()), 32, ConvertUint64Byte(generator.Next()), ConvertUint64Byte(generator.Next()), ConvertUint64Byte(generator.Next()), ConvertUint64Byte(generator.Next()))

	for i := 0; i < 8; i++ {
		startIndex := i * 4
		endIndex := (i + 1) * 4
		uint32Value := binary.BigEndian.Uint32(kdfResult[startIndex:endIndex])
		temp[i] = uint32Value
	}

	Keys = append(Keys, temp)
}

func ToCrisp(data []uint64) {

	file, err := os.OpenFile("files/channel", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	var Payload []byte
	var SeqNum uint64
	binary.BigEndian.PutUint64(seed, uint64(time.Now().UnixMilli()))
	generator := lab3.New(seed)
	CreateKey(generator)
	fmt.Printf("[#%d] Using Key: %x\n", SeqNum, Keys[SeqNum])
	for i := 0; i < len(data); i++ {
		byteSlice := make([]byte, 8)
		binary.BigEndian.PutUint64(byteSlice, lab1.Encrypt(data[i], Keys[SeqNum]))
		Payload = append(Payload, byteSlice...)
		if len(Payload) > 2026 {
			msg := MakeMessage(Payload, SeqNum, Keys[SeqNum])
			WriteToFile(msg, file)
			fmt.Printf("[#%d] %x\n\n", SeqNum, msg)
			SeqNum++
			CreateKey(generator)
			fmt.Printf("[#%d] Using Key: %x\n", SeqNum, Keys[SeqNum])
			Payload = nil
			msg = nil
		}
	}
	if len(Payload) > 0 {
		msg := MakeMessage(Payload, SeqNum, Keys[SeqNum])
		WriteToFile(msg, file)
		fmt.Printf("[#%d] %x\n\n", SeqNum, msg)
	}
}

func FromCrisp(msg []byte) {
	EKIF, Version, CS, KeyId, SeqNum := GetHeaderParam(msg)
	if Version != 0 && EKIF != 0 && KeyId != 0x80 && CS != 0xf3 && SeqNum > 0xfffffffffffff {
		log.Println("Error header Params")
		return
	}
	if CheckConsistency(msg, Keys[SeqNum]) != true {
		//fmt.Println("[#%d] Consistency check failed")
		return
	}
	fmt.Printf("[#%d] Consistency check passed\n", SeqNum)
	fmt.Printf("[#%d] Using Key: %x\n", SeqNum, Keys[SeqNum])
	Data := GetDecryptedPayloadData(msg, Keys[SeqNum])

	fmt.Printf("[#%d] ", SeqNum)
	for i := 0; i < len(Data); i++ {
		fmt.Printf("%x", Data[i])
	}
	fmt.Printf("\n\n")

	//fmt.Printf("[#%d] ", SeqNum)
	//for i := 0; i < len(Data); i++ {
	//	reconstructedString := ""
	//	reconstructedString += uint64ToString(Data[i])
	//	fmt.Printf("%s", reconstructedString)
	//}
	//fmt.Printf("\n")

}

func uint64ToString(num uint64) string {
	var result []string
	for num > 0 {
		result = append([]string{string(num & 0xFF)}, result...)
		num >>= 8
	}
	return strings.Join(result, "")
}
