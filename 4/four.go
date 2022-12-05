package four

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	challengeFourInput = "4/four_04_input.txt"
)

func Four() {
	fmt.Println()
	fmt.Println("Day Four")

	content, err := ioutil.ReadFile(challengeFourInput)
	if err != nil {
		log.Fatalf("failed to read file: %s | error: %e", challengeFourInput, err)
	}

	fmt.Println(string(content))
}
