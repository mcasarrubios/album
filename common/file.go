package common

import (
	"io/ioutil"
	"os"
)

// ReadFile reads a whole file
func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}
