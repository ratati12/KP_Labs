package lab2

import (
	"math/rand"
	"time"
)

func ClearData() []byte {
	data := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i++ {
		data[i] = byte(rand.Intn(256))
	}
	//log.Println("Очищена переменная (KDF).")
	return data
}
