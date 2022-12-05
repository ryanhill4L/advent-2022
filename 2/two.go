package two

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	challengeTwoInput = "2/two_02_input.txt"

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
	newRounds := ToFullRound(rounds)

	for _, round := range newRounds {
		neededPoints = neededPoints + round.GetRoundPoints()
	}

	fmt.Printf("The outcome of the decrypted strategy guide is %v points\n", neededPoints)
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

type Round struct {
	Opponent uint
	Self     uint
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

func ToFullRound(in []Round) []FullRound {
	resp := make([]FullRound, len(in))

	for i, round := range in {
		resp[i] = FullRound{
			Opponent: round.Opponent,
			Outcome:  round.Self,
		}

		resp[i].GetSelfFromOutcome()
	}

	return resp
}

type FullRound struct {
	Opponent uint
	Self     uint
	Outcome  uint
}

func (f *FullRound) GetSelfFromOutcome() {
	switch f.Outcome {
	case 1: // Lose
		f.Self = (f.Opponent - 1) % 3
		if f.Self == 0 {
			f.Self = 3
		}
	case 2: // Draw
		f.Self = f.Opponent
	case 3: // Win
		f.Self = (f.Opponent + 1) % 3
		if f.Self == 0 {
			f.Self = 1
		}
	}
}

func (f FullRound) GetRoundPoints() uint {
	return f.Self + f.OutcomeToPoints()
}

func (f FullRound) OutcomeToPoints() uint {
	switch f.Outcome {
	case 1: // Lose
		return 0
	case 2: // Draw
		return 3
	case 3: // Win
		return 6
	default:
		log.Fatal("returned not possible outcome")

		return 0
	}
}
