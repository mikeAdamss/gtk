package parsers

import (
	"encoding/json"
	"fmt"
	"log"
)

func ArrayOfLines(bytes []byte) {

	var lines []string
	err := json.Unmarshal(bytes, &lines)
	if err != nil {
		log.Fatal("Unable to unmarshall expected array of string from bytes", err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
