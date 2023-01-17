package common

import (
	"log"
	"os"
)

func ReadFile(fname string) string {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
