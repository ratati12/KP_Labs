package pkg

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func WriteFileForSizeTest(generator *XoroShiro256PlusPlus, size int, filename string, start time.Time) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	for i := 0; i < size*131072; i++ {
		value := generator.Next()

		// Преобразование uint64 в байты
		buffer := make([]byte, 8) // uint64 занимает 8 байт
		binary.LittleEndian.PutUint64(buffer, value)

		// Запись байтов в файл
		_, err = file.Write(buffer)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
	end := time.Now()
	elapsedTime := end.Sub(start).Milliseconds()
	fmt.Printf("Записан файл: %s за %d мс\n", filename, elapsedTime)
}

func WriteFileForNISTTest(generator *XoroShiro256PlusPlus, start time.Time) {
	filename := "ForNIST.bin"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var nextBatch [8]byte
	for i := 0; i < 1000; i++ {
		binary.LittleEndian.PutUint64(nextBatch[:], generator.Next())
		for _, bt := range nextBatch {
			b := fmt.Sprintf("%08b", bt)
			_, err := file.WriteString(b)
			if err != nil {
				fmt.Println("Ошибка при записи в файл:", err)
			}
		}
	}
	end := time.Now()
	elapsedTime := end.Sub(start).Milliseconds()
	fmt.Printf("Записан файл: %s за %d мс\n", filename, elapsedTime)
}

func GenerateRandomAmount(generator *XoroShiro256PlusPlus, filename string, start time.Time) {
	rand.Seed(time.Now().UnixNano())

	// Генерация случайного числа в промежутке от 10^3 до 10^4
	amount := rand.Intn(9000) + 1000
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	for i := 0; i < amount; i++ {
		value := generator.Next()

		// Преобразование uint64 в байты
		buffer := make([]byte, 8) // uint64 занимает 8 байт
		binary.LittleEndian.PutUint64(buffer, value)

		// Запись байтов в файл
		_, err = file.Write(buffer)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
	end := time.Now()
	elapsedTime := end.Sub(start).Milliseconds()
	fmt.Printf("Записан файл: %s с кол-вом %d ключей за %d мс\n", filename, amount, elapsedTime)

}
