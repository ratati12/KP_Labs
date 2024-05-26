package main

import (
	"encoding/binary"
	"fmt"
	"lab2/pkg"
	"log"
	"os"
	"time"
)

func main() {

	//key := []byte{0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79,
	//	0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f, 0x80, 0x81, 0x82, 0x83}
	pkg.SetLogger()
	pkg.CheckIntegrity()
	pkg.MechanismCheck()
	//pkg.AccessCheck()

	go pkg.PeriodCheckIntegrity()

	// -------------------------------------------------------------------------
	// ----------------------- Чтение ключа из файла ---------------------------
	// -------------------------------------------------------------------------

	key_file, err := os.OpenFile("files/key.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Println("Ошибка открытия файла:", err)
		return
	}
	defer key_file.Close()
	// Получаем размер файла ключа
	key_fileInfo, err := key_file.Stat()
	if err != nil {
		log.Println("Ошибка получения информации о файле:", err)
		return
	}
	key_fileSize := key_fileInfo.Size()
	if key_fileSize != 32 {
		log.Println("Неправильный размер", err)
		return
	}
	// Создаем массив для хранения uint32 для ключа
	key := make([]byte, key_fileSize)
	// Читаем данные из файла в массив []byte
	err = binary.Read(key_file, binary.LittleEndian, &key)
	if err != nil {
		return
	}
	//pkg.KeyUseCheck(key)
	var i int
	seed := make([]byte, 64)

	start := time.Now()
	for i < 1000000 {

		binary.BigEndian.PutUint64(seed, uint64(time.Now().UnixMilli()))
		fmt.Printf("%x\n", pkg.Generator(key, 32, seed, []byte("1111"), []byte("2222"), []byte("3333")))
		i++
	}
	end := time.Now()
	elapsedTime := end.Sub(start).Milliseconds()
	fmt.Printf("За %d мс\n", elapsedTime)
}
