package main

import (
	"fmt"
)

func decimalToBinary(val uint64) []uint64 {
	bits := []uint64{}

	for val != 0 {
		mod := val % 2
		bits = append(bits, mod)

		val = val / 2
	}

	for i := 0; i < len(bits); i++ {
		fmt.Print(bits)
	}

	return bits
}

