package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"lab4/pkg"
	"log"
	"os"
)

func main() {
	pkg.SetLogger()
	pkg.CheckIntegrity()
	pkg.AccessCheck()
	pkg.MechanismCheck()
	go pkg.PeriodCheckIntegrity()
	//key := [8]uint32{0xffeeddcc, 0xbbaa9988, 0x77665544, 0x33221100, 0xf0f1f2f3, 0xf4f5f6f7, 0xf8f9fafb, 0xfcfdfeff}
	//message := []uint64{0xfedcba9876543210, 0xfedcba9876543210}
	//key := []byte{
	//	0xcc, 0xdd, 0xee, 0xff,
	//	0x88, 0x99, 0xaa, 0xbb,
	//	0x44, 0x55, 0x66, 0x77,
	//	0x00, 0x11, 0x22, 0x33,
	//	0xf3, 0xf2, 0xf1, 0xf0,
	//	0xf7, 0xf6, 0xf5, 0xf4,
	//	0xfb, 0xfa, 0xf9, 0xf8,
	//	0xff, 0xfe, 0xfd, 0xfc}
	//message := [1024]uint64{0xfedcba9876543210}
	//message := []uint64{0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210,
	//	0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210,
	//	0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210,
	//	0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210,
	//	0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210, 0xfedcba9876543210}

	filein, errin := os.Open("files/plaintext2.txt")
	if errin != nil {
		log.Println("Ошибка открытия файла:", errin)
		return
	}
	defer filein.Close()

	// Определяем размер файла
	fileInfo, err := filein.Stat()
	if err != nil {
		log.Println("Ошибка получения информации о файле:", err)
		return
	}
	fileSize := fileInfo.Size()

	// Создаем срез для хранения данных
	data := make([]byte, fileSize)

	// Читаем данные из файла
	_, err = filein.Read(data)
	if err != nil {
		log.Println("Ошибка чтения данных из файла:", err)
		return
	}

	if len(data)%8 != 0 {
		newBytes := make([]byte, (len(data)/8+1)*8-len(data))
		data = append(data, newBytes...)
	}
	// Преобразуем считанные данные в срез uint64
	var message []uint64
	for i := 0; i < len(data); i += 8 {
		uint64Value := binary.BigEndian.Uint64(data[i : i+8])
		message = append(message, uint64Value)
	}

	fmt.Println("---------------- To Crisp -----------------")
	// -----------------  ToCrisp  --------------
	pkg.ToCrisp(message[:])

	fmt.Println("---------------- From Crisp -----------------")
	// -----------------  FromCrisp  --------------
	file, err := os.Open("files/channel")
	if err != nil {
		log.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Создаем новый сканнер для чтения файла
	scanner := bufio.NewScanner(file)

	// Читаем файл построчно
	for scanner.Scan() {
		line, err := hex.DecodeString(scanner.Text())
		if err != nil {
			log.Println("Ошибка декодирования строки:", err)
			return
		}
		pkg.FromCrisp(line)
	}

	// Проверяем ошибки сканнера
	if err := scanner.Err(); err != nil {
		log.Println("Ошибка чтения файла:", err)
	}

}
