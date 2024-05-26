package pkg

import (
	"encoding/hex"
	"lab4/gost341112"
	"lab4/lab1"
	"log"
	"os"
)

var sum512 = "1b54d01a4af5b9d5cc3d86d68d285462b19abc2475222f35c085122be4ba1ffa00ad30f8767b3a82384c6574f024c311e2a481332b08ef7f41797891c1646f48"
var sum256 = "9d151eefd8590b89daa6ba6cb74af9275dd051026bb149a452fd84e5e57b5500"

func MechanismCheck() {
	key := [8]uint32{0xffeeddcc, 0xbbaa9988, 0x77665544, 0x33221100, 0xf0f1f2f3, 0xf4f5f6f7, 0xf8f9fafb, 0xfcfdfeff}
	message_magma := uint64(0xfedcba9876543210)
	message_omac := []uint64{0x92def06b3c130a59, 0xdb54c704f8189d20, 0x4a98fb2e67a8024c, 0x8912409b17b57e41}
	message_streebog := []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32}
	hash256 := gost341112.Sum256(message_streebog)
	result256 := hex.EncodeToString(hash256[:])
	hash512 := gost341112.Sum512(message_streebog)
	result512 := hex.EncodeToString(hash512[:])
	resultMagma := lab1.Encrypt(message_magma, key)
	if resultMagma == uint64(0x4EE901E5C2D8CA3D) && lab1.Omac(message_omac, key, 4) == uint64(0x154e7210) && (result512 == sum512) && (result256 == sum256) {
		log.Println("Механизм контроля работоспособности криптографических алгоритмов пройден!")
	} else {
		log.Println("Механизм контроля работоспособности криптографических алгоритмов НЕ пройден!")
		os.Exit(2)
	}

}
