package util

import (
	"errors"
	"math/rand"

	"github.com/irainia/gridgame/util"
)

// EncryptString will encrypt string
func EncryptString(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	value := rand.Intn(util.MaxByte+1) + 1
	output := make([]byte, len(input)+1)
	output[0] = byte(value)
	for i := 0; i < len(input); i++ {
		value = value + int(input[i])
		output[i+1] = byte(value)
	}

	return string(output), nil
}
