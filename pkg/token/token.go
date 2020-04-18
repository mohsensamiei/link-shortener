package token

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"unsafe"
)

var src = rand.NewSource(time.Now().UnixNano())

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func States(length int) int {
	return int(math.Pow(float64(len(LetterBytes)), float64(length)))
}
func StatesWithPrefix(length int, prefix string) int {
	left := length - len(prefix)
	if left < 0 {
		left = 0
	}
	return int(math.Pow(float64(len(LetterBytes)), float64(left)))
}

func Generate(length int) string {
	bytes := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for index, cache, remain := length-1, src.Int63(), letterIdxMax; index >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(LetterBytes) {
			bytes[index] = LetterBytes[idx]
			index--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&bytes))
}
func GenerateWithPrefix(length int, prefix string) string {
	left := length - len(prefix)
	if left > 0 {
		return fmt.Sprint(prefix, Generate(left))
	}
	return prefix[0:length]
}
