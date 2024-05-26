package main

import (
	"encoding/binary"
	"fmt"
	"lab3/pkg"
	"time"
)

func main() {
	seed := []byte{
		0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0x17, 0x28,
		0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0x17, 0x28,
		0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0x17, 0x28,
		0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0x17, 0x28,
	}
	binary.BigEndian.PutUint64(seed, uint64(time.Now().UnixMilli()))
	// Создание экземпляра генератора с начальным seed
	generator := pkg.New(seed)
	for i := 0; i < 10; i++ {
		number := generator.Next()
		if pkg.BitsTest(number) == true {
			fmt.Printf("[%d] %d\n", i+1, number)
		} else {
			fmt.Printf("[%d] Bad number\n", i+1)
			i--
		}

	}

	//// ---------------------------Создание 1 МБ -------------------- 285
	//pkg.WriteFileForSizeTest(generator, 1, "1MB_file.bin", time.Now())
	//
	//// ---------------------------Создание 100 МБ -------------------- 28157
	//pkg.WriteFileForSizeTest(generator, 100, "100MB_file.bin", time.Now())
	//
	//// ---------------------------Создание 1000 МБ -------------------- 282717
	//
	//pkg.WriteFileForSizeTest(generator, 1000, "1000MB_file.bin", time.Now())
	//
	//// ---------------------------Создание 10^3-10^4 МБ --------------------16
	//
	//pkg.GenerateRandomAmount(generator, "RandomAmount_file.bin", time.Now())
	//
	//// ---------------------------Создание для NIST --------------------38
	//pkg.WriteFileForNISTTest(generator, time.Now())
}
