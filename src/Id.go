package src

import (
	"encoding/hex"
	"log"
)


func DecodeId(id string) string {
	decodedBytes, err := hex.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}

	return string(decodedBytes)
}

func EncodeId(id string) string {
	b := []byte(id)
	return hex.EncodeToString(b)
}
