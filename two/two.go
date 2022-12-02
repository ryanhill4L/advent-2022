package two

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	challengeTwoInput = "two/two_02_input.txt"

	RockPoints     = 1
	PaperPoints    = 2
	ScissorsPoints = 3

	LostPoints = 0
	DrawPoints = 3
	WinPoints  = 6
)

type Player struct {
}

func Two() {
	fmt.Println()
	fmt.Println("Day Two")

	content, err := ioutil.ReadFile(challengeTwoInput)
	if err != nil {
		log.Fatalf("failed to read file: %s | error: %e", challengeTwoInput, err)
	}

	rounds := parseTextToStrategy(string(content))

	// Part One
	var totalPoints uint
	for _, round := range rounds {
		totalPoints = totalPoints + round.GetRoundPoints()
	}

	fmt.Printf("The outcome of the strategy guide is %v points\n", totalPoints)

	// Part Two
	var neededPoints uint

	fmt.Printf("The outcome of the decrypted strategy guide is %v points\n", neededPoints)
}

type Round struct {
	Opponent uint
	Self     uint
}

func parseTextToStrategy(in string) (resp []Round) {
	rounds := strings.Split(in, "\n")

	for _, round := range rounds {
		strategy := strings.Split(round, " ")

		if len(strategy) == 2 {
			resp = append(resp, Round{
				Opponent: StringToSelection(strategy[0]),
				Self:     StringToSelection(strategy[1]),
			})
			continue
		}
	}

	return resp
}

func (r Round) GetRoundPoints() uint {
	return r.GetSelection() + r.GetOutcome()
}

func (r Round) GetSelection() uint {
	return r.Self
}

func (r Round) GetOutcome() uint {
	switch r.Self {
	case 1: // Rock
		switch r.Opponent {
		case 1: // Rock
			return DrawPoints
		case 2: // Paper
			return LostPoints
		case 3: // Scissors
			return WinPoints
		default:
			log.Fatalf("round outcome not possible | Self: %v, Opponent: %v", r.Self, r.Opponent)

			return 0
		}
	case 2: // Paper
		switch r.Opponent {
		case 1: // Rock
			return WinPoints
		case 2: // Paper
			return DrawPoints
		case 3: // Scissors
			return LostPoints
		default:
			log.Fatalf("round outcome not possible | Self: %v, Opponent: %v", r.Self, r.Opponent)

			return 0
		}
	case 3: // Scissors
		switch r.Opponent {
		case 1: // Rock
			return LostPoints
		case 2: // Paper
			return WinPoints
		case 3: // Scissors
			return DrawPoints
		default:
			log.Fatalf("round outcome not possible | Self: %v, Opponent: %v", r.Self, r.Opponent)

			return 0
		}
	default:
		log.Fatalf("round outcome not possible | Self: %v, Opponent: %v", r.Self, r.Opponent)

		return 0
	}
}

func StringToSelection(in string) uint {
	switch in {
	case "A", "X":
		return RockPoints
	case "B", "Y":
		return PaperPoints
	case "C", "Z":
		return ScissorsPoints
	default:
		log.Fatalf("selection input invalid %s", in)

		return 0
	}
}
