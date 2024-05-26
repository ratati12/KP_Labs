package pkg

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"lab2/gost341112"
	"log"
	"os"
	"strings"
)

var SavedHashLog = "8d74ce474fff1c4a77ba83845b39e4b8756ad7ae85d0e1f4acff19954faf579f"
var SavedHashPass = "cf1943dd1495f66bf104bf37e1f40e08e59c0d146570e35f16839ff6d05d2d29"

func AccessCheck() {

	// Запрашиваем логин у пользователя
	fmt.Print("Введите логин: ")
	reader := bufio.NewReader(os.Stdin)
	login, _ := reader.ReadString('\n')
	// Удаляем символ новой строки из пароля
	login = strings.TrimSuffix(login, "\n")

	// Вычисляем хеш SHA256
	hasher := gost341112.New256()
	hasher.Write([]byte(login))
	loginHash := hex.EncodeToString(hasher.Sum(nil))

	// Сравниваем полученный хеш с существующим хешем
	if loginHash != SavedHashLog {
		fmt.Println("Доступ запрещен.")
		log.Printf("Доступ запрещен. %s\n", login)
		os.Exit(1)
	}
	// Запрашиваем пароль у пользователя
	fmt.Print("Введите пароль: ")

	password, _ := reader.ReadString('\n')
	// Удаляем символ новой строки из пароля
	password = strings.TrimSuffix(password, "\n")
	hasher = gost341112.New256()
	// Вычисляем хеш SHA256 для введенного пароля
	hasher.Write([]byte(password))
	passwordHash := hex.EncodeToString(hasher.Sum(nil))

	// Сравниваем полученный хеш с существующим хешем пароля
	if passwordHash == SavedHashPass {
		fmt.Println("Доступ разрешен.")
		log.Printf("Доступ разрешен пользователю %s\n", login)
	} else {
		fmt.Println("Доступ запрещен.")
		log.Printf("Доступ запрещен пользователю %s\n", login)
		os.Exit(1)
	}
}
