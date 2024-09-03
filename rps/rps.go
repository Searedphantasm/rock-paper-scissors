package rps

import "math/rand/v2"

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
)

var (
	drawMap = map[int]string{
		1: "مساوی مساوی",
		2: "برابر برابر",
		3: "سگ سارونه",
	}
	winnerMap = map[int]string{
		1: "بردی بردی",
		2: "خیلی بردی بردی",
		3: "سگ سارونسسسس",
	}
	losserMap = map[int]string{
		1: "خیلی یاخت باخت",
		2: "باختی باختی",
		3: "سگ سارونهسسیلسلب",
	}
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	computerValue := rand.IntN(3)
	computerChoice := ""
	roundResult := ""
	message := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer choice ROCK"
		break
	case PAPER:
		computerChoice = "Computer choice PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer choice SCISSORS"
		break
	default:
	}

	randIdx := rand.IntN(3) + 1

	if playerValue == computerValue {
		roundResult = "It's draw!"
		message = drawMap[randIdx]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = winnerMap[randIdx]
	} else {
		roundResult = "Computer wins!"
		message = losserMap[randIdx]
	}

	var result = Round{message, computerChoice, roundResult}

	return result
}
