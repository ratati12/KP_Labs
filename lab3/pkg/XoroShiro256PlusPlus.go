package pkg

import (
	"encoding/binary"
	"math/bits"
)

type XoroShiro256PlusPlus struct {
	seed [4]uint64
}

func New(seed []byte) *XoroShiro256PlusPlus {
	if len(seed) != 32 {
		panic("seed should be 32 bytes")
	}
	s0 := seed[0:8]
	s1 := seed[8:16]
	s2 := seed[16:24]
	s3 := seed[24:32]

	v0 := binary.BigEndian.Uint64(s0)
	v1 := binary.BigEndian.Uint64(s1)
	v2 := binary.BigEndian.Uint64(s2)
	v3 := binary.BigEndian.Uint64(s3)

	return &XoroShiro256PlusPlus{seed: [4]uint64{v0, v1, v2, v3}}
}

func (x *XoroShiro256PlusPlus) Next() uint64 {
	s0, s1, s2, s3 := x.seed[0], x.seed[1], x.seed[2], x.seed[3]

	x.seed[0] = s0 ^ s3 ^ s1
	x.seed[1] = s1 ^ s2 ^ s0
	x.seed[2] = s2 ^ s0 ^ (s1 << 17)
	x.seed[3] = bits.RotateLeft64(s1^s3, 45)

	return bits.RotateLeft64(s0+s3, 23) + s0
}
