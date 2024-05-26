package lab2

func format(zi []byte, ci uint64, L []byte, P []byte, A []byte, U []byte) []byte {
	retLen := 1 + 32 + 64 + len(L) + len(P) + len(A) + len(U)
	retString := make([]byte, retLen)
	pos := 0
	retString[pos] = byte(0xFC)
	pos++
	C := make([]byte, 32)
	for i := 0; i < 32; i++ {
		C[i] = byte((ci >> (uint(i) * 8)) & 0xff)
	}
	for i := 0; i < 32; i++ {
		retString[pos] = C[i]
		pos++
	}
	for i := 0; i < 64; i++ {
		retString[pos] = zi[i]
		pos++
	}
	for i := 0; i < len(L); i++ {
		retString[pos] = L[i]
		pos++
	}
	for i := 0; i < len(P); i++ {
		retString[pos] = P[i]
		pos++
	}
	for i := 0; i < len(U); i++ {
		retString[pos] = U[i]
		pos++
	}
	for i := 0; i < len(A); i++ {
		retString[pos] = A[i]
		pos++
	}
	return retString
}

func first_Key(key []byte, salt []byte) []byte {
	k1_512 := Hmac(key, salt)
	return k1_512[:32]
}
func setL(L_s uint64) []byte {
	var L [8]byte
	for i := 0; i < 8; i++ {
		L[i] = byte((L_s >> (i * 8)) & 0xff)
	}
	return L[:]
}
func Generator(key []byte, L_str uint64, T []byte, P []byte, A []byte, U []byte) []byte {

	K1 := first_Key(key, T)
	key = ClearData()
	C := uint64(1)

	finalLen := L_str / 64
	if L_str%64 != 0 {
		finalLen++
	}
	zi := make([]byte, 64)
	retString := make([]byte, finalLen*64)

	for i := uint64(0); i < finalLen; i++ {
		formatStr := format(zi, C, setL(L_str), P, A, U)
		new_str := Hmac(K1, formatStr)
		copy(zi, new_str[:])
		copy(retString[i*32:], zi)
		C++
	}

	finalString := make([]byte, L_str)
	copy(finalString, retString)

	return finalString
}
