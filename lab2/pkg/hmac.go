package pkg

import (
	"lab2/gost341112"
)

func Hmac(key []byte, data []byte) [64]byte {
	b := 128
	k0 := make([]byte, b)

	if len(key) > b {
		hash := gost341112.Sum512(key)
		copy(k0, hash[:])
	} else {
		copy(k0, key)
	}

	if len(key) < b {
		padding := make([]byte, b-len(key))
		k0 = append(k0, padding...)
	}
	key = ClearData()
	ipad := make([]byte, b)
	opad := make([]byte, b)
	for i := 0; i < b; i++ {
		ipad[i] = k0[i] ^ 0x36
		opad[i] = k0[i] ^ 0x5c
	}

	ikeypad := gost341112.Sum512(append(ipad, data...))

	return gost341112.Sum512(append(opad, ikeypad[:]...))
}
