package one

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	challengeOneInput = "one/one_01_input.txt"
)

type Elf struct {
	foodCalories []int
}

func One() {
	content, err := ioutil.ReadFile(challengeOneInput)
	if err != nil {
		log.Fatalf("failed to read file: %s | error: %e", challengeOneInput, err)
	}

	elves := parseTextToElves(string(content))

	// Part One
	var total int
	for _, elf := range elves {
		if elf.GetTotalCalories() > total {
			total = elf.GetTotalCalories()
		}
	}

	fmt.Printf("The top elf is carrying %v calories\n", total)

	// Part Two
	var elfTotals []int
	for _, elf := range elves {
		elfTotals = append(elfTotals, elf.GetTotalCalories())
	}

	sort.Ints(elfTotals)

	topElves := elfTotals[len(elfTotals)-3:]
	var topTotal int

	for _, elfTotal := range topElves {
		topTotal = topTotal + elfTotal
	}

	fmt.Printf("The top 3 elves are carrying %v calories\n", topTotal)
}

func parseTextToElves(in string) (resp []Elf) {
	splitIn := strings.Split(in, "\n")

	elfIndexer := 0
	temp := Elf{}

	for i := range splitIn {
		if splitIn[i] == "" {
			resp = append(resp, temp)

			elfIndexer++
			temp = Elf{}
			continue
		}

		calorie, _ := strconv.ParseUint(splitIn[i], 10, 32)
		temp.foodCalories = append(temp.foodCalories, int(calorie))
	}

	return resp
}

func (e Elf) GetTotalCalories() (resp int) {
	for _, item := range e.foodCalories {
		resp = resp + item
	}

	return resp
}
