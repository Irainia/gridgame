package util

import (
	"errors"

	"github.com/irainia/gridgame/util"
)

// DecryptString will decrypt string
func DecryptString(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}
	if len(input) < 2 {
		return "", errors.New("input length should be at least 2 (two) characters")
	}

	value := 0
	output := make([]byte, len(input)-1)
	for i := 1; i < len(input); i++ {
		value = int(input[i]) - int(input[i-1])
		if value < 0 {
			value += util.MaxByte + 1
		}
		output[i-1] = byte(value)
	}

	return string(output), nil
}
