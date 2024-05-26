package pkg

import (
	"encoding/hex"
	"io"
	"lab2/gost341112"
	"log"
	"os"
	"time"
)

func calculateFileHash() (string, error) {
	filePath, err := os.Executable()
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := gost341112.New256()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func readHashFromFile() (string, error) {
	file, err := os.Open("files/hash.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := make([]byte, 64) // SHA-256 produces 64-character hash
	_, err = file.Read(hash)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func PeriodCheckIntegrity() {
	for {
		CheckIntegrity()
		time.Sleep(1 * time.Nanosecond)
	}
}

func CheckIntegrity() {
	currentHash, err := calculateFileHash()
	if err != nil {
		log.Println("Ошибка при рассчёте хеша:", err)
		return
	}
	savedHash, err := readHashFromFile()
	if err != nil {
		log.Println("Ошибка при чтении сохранённого хеша из файла:", err)
		return
	}
	if currentHash != savedHash {
		log.Println("Целостность файла нарушена!")
		// return
	} else {
		log.Println("Целостность файла проверена успешно.")
	}
}
