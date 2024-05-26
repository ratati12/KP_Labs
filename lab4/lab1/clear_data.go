package lab1

import (
	"math/rand"
	"time"
)

func ClearData() [8]uint32 {
	var data [8]uint32
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(data); i++ {
		data[i] = rand.Uint32()
	}
	//log.Println("Очищена переменная. (MagmaOMAC)")
	return data
}
