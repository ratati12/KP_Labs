package pkg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"lab1/gost341112"
	"log"
	"os"
)

func KeyUseCheck(array [8]uint32) {
	buf := new(bytes.Buffer)

	for _, v := range array {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("Ошибка записи:", err)
			return
		}
	}
	// Получаем срез байтов из буфера
	key := buf.Bytes()

	hash := gost341112.New256()
	// Записываем данные в хэш
	hash.Write(key)
	// Получаем вычисленный хэш в виде среза байтов
	hashkey := hash.Sum(nil)

	file, err := os.Open("files/keyhashes.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Создаем сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes() // Считываем строку как срез байтов
		if bytes.Equal(line, hashkey) {
			log.Println("Ключ уже использовался")
			os.Exit(1)
		}
	}
	file, err = os.OpenFile("files/keyhashes.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	hashkey = append(hashkey, '\n')
	_, err = file.Write(hashkey)
	if err != nil {
		return
	}
}
