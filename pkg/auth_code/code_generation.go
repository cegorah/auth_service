/*
Utils for helping generating fixed length code
*/

package auth_code

import (
	"log"
	"math/rand"
)

type numRange struct {
	low int
	hi  int
}

var ranges = []numRange{
	{0, 0},
	{0, 9},
	{10, 99},
	{100, 999},
	{1000, 9999},
	{10000, 99999},
	{100000, 999999},
}

func fixedLengthCode(codeLen int) int {
	rangesLen := len(ranges)
	if codeLen > rangesLen {
		codeLen = rangesLen - 1
		log.Printf("codeLen out of range using %d instead\n", codeLen)
	}
	l, hi := ranges[codeLen].low, ranges[codeLen].hi
	return l + rand.Intn(hi-l)
}
