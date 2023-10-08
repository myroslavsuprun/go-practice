package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("./mock",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return nil
			}
			if info.IsDir() {
				return nil
			}

			name := info.Name()
			parentDir := filepath.Base(filepath.Dir(path))
			newName := bytes.NewBufferString("")
			for i := 0; i < len(name); i++ {
				if name[i] == byte('.') {
					break
				}

				if withinUppercase(name[i]) && !withinUppercase(name[i-1]) && name[i-1] != byte('-') {
					newName.WriteByte('-')
				}

				char := bytes.ToLower([]byte(string(name[i])))

				if withinUppercase(name[i]) && (withinUppercase(name[i+1]) || withinUppercase(name[i-1])) {
					char = []byte(string(name[i]))
				}

				newName.Write(char)
			}
			newName.WriteString(fmt.Sprintf(".%s%s", parentDir, filepath.Ext(name)))

			err = os.Rename(path, filepath.Join(filepath.Dir(path), newName.String()))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(filepath.Dir(path), newName.String())
			return nil
		})
}

func withinUppercase(c byte) bool {
	return c >= 'A' && c <= 'Z'
}
