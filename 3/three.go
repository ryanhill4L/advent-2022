package three

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	challengeThreeInput = "3/three_03_input.txt"

	priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Three() {
	fmt.Println()
	fmt.Println("Day Three")

	content, err := ioutil.ReadFile(challengeThreeInput)
	if err != nil {
		log.Fatalf("failed to read file: %s | error: %e", challengeThreeInput, err)
	}

	rucksacks := parseTextToRucksacks(string(content))

	// Part One
	var total int
	for i := range rucksacks {
		rucksacks[i].FindSharedItem()

		total = total + GetPriorityOfItem(rucksacks[i].SharedItem)
	}

	fmt.Printf("The sum of the priorities of shared items is %v\n", total)

	// Part Two
	var groups []Groups
	for i := 0; i < len(rucksacks); i = i + 3 {
		groups = append(groups, Groups{
			Rucksacks: []Rucksack{
				rucksacks[i],
				rucksacks[i+1],
				rucksacks[i+2],
			},
		})
	}

	var badgeTotal int
	for i := range groups {
		groups[i].FindBadge()

		badgeTotal = badgeTotal + GetPriorityOfItem(groups[i].Badge)
	}

	fmt.Printf("The sum of the priorities of the group badges are %v\n", badgeTotal)
}

func parseTextToRucksacks(in string) (resp []Rucksack) {
	sacks := strings.Split(in, "\n")

	for _, sack := range sacks {
		if sack == "" {
			continue
		}

		compartmentLength := len(sack) / 2

		compartments := []string{
			sack[:compartmentLength],
			sack[compartmentLength:],
		}

		resp = append(resp, Rucksack{
			All:         sack,
			Compartment: compartments,
		})
	}

	return resp
}

type Rucksack struct {
	All         string
	Compartment []string
	SharedItem  string
}

func (r *Rucksack) FindSharedItem() {
	var resp []string
	var count int

	for i := range r.Compartment[0] {
		for j := range r.Compartment[1] {
			if fmt.Sprintf("%c", r.Compartment[0][i]) == fmt.Sprintf("%c", r.Compartment[1][j]) {
				resp = append(resp, fmt.Sprintf("%c", r.Compartment[0][i]))
				count++
				break
			}
		}
	}

	if len(resp) > 0 {
		r.SharedItem = resp[0]
	}
}

func GetPriorityOfItem(in string) int {
	if in == "" {
		return 0
	}

	return strings.Index(priority, in) + 1
}

type Groups struct {
	Rucksacks []Rucksack
	Badge     string
}

func (g *Groups) FindBadge() {
	var resp []string
	var count int

	for i := range g.Rucksacks[0].All {
	double:
		for j := range g.Rucksacks[1].All {
			for k := range g.Rucksacks[2].All {
				if fmt.Sprintf("%c", g.Rucksacks[0].All[i]) == fmt.Sprintf("%c", g.Rucksacks[1].All[j]) &&
					fmt.Sprintf("%c", g.Rucksacks[0].All[i]) == fmt.Sprintf("%c", g.Rucksacks[2].All[k]) {
					resp = append(resp, fmt.Sprintf("%c", g.Rucksacks[0].All[i]))
					count++
					break double
				}
			}
		}
	}

	if len(resp) > 0 {
		g.Badge = resp[0]
	}
}
