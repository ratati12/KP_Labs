package pkg

import (
	"encoding/binary"
	"math"
)

const (
	numBitsInSeq = 8 * 16
)

func BitsTest(number uint64) bool {
	seq := make([]byte, 8)
	binary.BigEndian.PutUint64(seq, uint64(number))
	count := 0
	for i := range seq {
		count += countBits(seq[i])
	}

	s := float64(count - (numBitsInSeq - count))
	s_obs := s / math.Sqrt(numBitsInSeq)
	p_value := math.Erfc(s_obs / math.Sqrt(2))
	if p_value > 0.01 {
		return true
	} else {
		return false
	}
}

func countBits(n byte) int {
	count := 0
	for n != 0 {
		count++
		n = n & (n - 1)
	}
	return count
}
