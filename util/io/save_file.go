package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	endecUtil "github.com/Irainia/gridgame/util/endec"
)

// SaveScore will save score
func SaveScore(path string, score int) error {
	stat, err := os.Stat(path)
	if err == nil {
		if mode := stat.Mode(); mode.IsDir() {
			return errors.New("path is not file")
		}

		err = os.Remove(stat.Name())
		if err != nil {
			return err
		}
	}

	encrypted, err := endecUtil.EncryptString(fmt.Sprintf("%0256d", score))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, []byte(encrypted), os.ModePerm)
}
