package common

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const contentDir string = "inputfiles/"

func ReadContent(day int) (content string) {
	filename := fmt.Sprintf("%sinput_%02d.txt", contentDir, day)
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %s", filename)
	}
	content = strings.Trim(string(bytes), "\n")
	return
}
