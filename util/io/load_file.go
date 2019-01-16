package util

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"

	util "github.com/irainia/gridgame/util/endec"
)

// LoadScore will load score
func LoadScore(path string) (int, error) {
	if path == "" {
		return -1, errors.New("path cannot be empty")
	}
	var outputString string

	stat, err := os.Stat(path)
	if err != nil {
		return -1, err
	}
	if mode := stat.Mode(); mode.IsDir() {
		return -1, errors.New("path is not file")
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return -1, err
	}

	decrypted, err := util.DecryptString(string(content))
	if err != nil {
		return -1, err
	}

	outputString = string(decrypted)
	return strconv.Atoi(outputString)
}
