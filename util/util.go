package util

import (
	"log"
	"os"
)

func GetDataFromeFile(file string) []byte {
	fp, err := os.OpenFile(file, os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 1000)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return data[:n]
}
