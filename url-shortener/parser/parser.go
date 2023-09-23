package parser

import (
	"os"
	"shortener/flagReader"
)

// Accepts a YAML file location relative to the current working directory
func GetBytes(fileName string) ([]byte, error) {
	if fileName == flagReader.DefaultFileName {
		return nil, nil
	}

	file, err := os.Open("./" + fileName)
	if err != nil {
		return nil, err
	}
	data := make([]byte, 1000)

	file.Read(data)
	for i, b := range data {
		if b == 0 {
			data = data[:i]
			break
		}
	}

	defer file.Close()

	return data, nil
}
