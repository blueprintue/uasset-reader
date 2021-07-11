package internal

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
)

func ReadNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func Read(file *os.File, number int, res interface{}) error {
	data := ReadNextBytes(file, number)
	buffer := bytes.NewBuffer(data)

	return binary.Read(buffer, binary.LittleEndian, res)
}
