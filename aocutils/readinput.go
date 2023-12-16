package aocutils

import (
	"log"
	"os"
)

func ReadInput(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
