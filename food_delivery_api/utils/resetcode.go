package utils

import (
	"math/rand"
	"time"
)

const resetCodeLength = 6
const resetCodeCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateResetCode() string {
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, resetCodeLength)
	for i := 0; i < resetCodeLength; i++ {
		code[i] = resetCodeCharset[rand.Intn(len(resetCodeCharset))]
	}

	return string(code)
}
