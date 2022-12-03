package Day02

import (
	"bufio"
	"os"
	"strings"
)

type PlayOption int
type Outcome int

const (
	UnknownPlay PlayOption = -1
	Rock                   = 1
	Paper                  = 2
	Scissors               = 3
)

const (
	UnknownOutcome Outcome = -1
	Win                    = 6
	Lose                   = 0
	Draw                   = 3
)

type round struct {
	opponentPlay PlayOption
	myPlay       PlayOption
}

func readInputFilePart1(f string) []round {
	var rounds []round
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		rounds = append(rounds, calculateRoundPart1(fields))
	}

	return rounds
}

func readInputFilePart2(f string) []round {
	var rounds []round
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		rounds = append(rounds, calculateRoundPart2(fields))
	}

	return rounds
}

func calculateRoundPart1(someRound []string) round {
	return round{translatePlay(someRound[0]), translatePlay(someRound[1])}
}

func calculateRoundPart2(someRound []string) round {
	opponentPlay := translatePlay(someRound[0])
	return round{opponentPlay, translateMyPlayPart2(opponentPlay, someRound[1])}
}

func calculateMatchOutcome(inputs []round) int {
	total := 0
	for _, someRound := range inputs {
		total += calculateRoundValue(someRound)
	}

	return total
}

func translatePlay(play string) PlayOption {
	switch play {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		return UnknownPlay
	}
}

func translateMyPlayPart2(opponentPlay PlayOption, myPlay string) PlayOption {
	switch myPlay {
	case "X": // lose
		return findLosePlay(opponentPlay)
	case "Y": // draw
		return opponentPlay
	case "Z": // win
		return findWinPlay(opponentPlay)
	default:
		return UnknownPlay
	}
}

func calculateRoundValue(theRound round) int {
	return int(calculateOutCome(theRound)) + int(theRound.myPlay)
}

func calculateOutCome(theRound round) Outcome {
	switch theRound.opponentPlay {
	case Rock:
		return outcomeVsRock(theRound.myPlay)
	case Paper:
		return outcomeVsPaper(theRound.myPlay)
	case Scissors:
		return outcomeVsScissors(theRound.myPlay)
	default:
		return UnknownOutcome
	}
}

func outcomeVsRock(myPlay PlayOption) Outcome {
	switch myPlay {
	case Rock:
		return Draw
	case Paper:
		return Win
	case Scissors:
		return Lose
	default:
		return UnknownOutcome
	}
}

func outcomeVsPaper(myPlay PlayOption) Outcome {
	switch myPlay {
	case Rock:
		return Lose
	case Paper:
		return Draw
	case Scissors:
		return Win
	default:
		return UnknownOutcome
	}
}

func outcomeVsScissors(myPlay PlayOption) Outcome {
	switch myPlay {
	case Rock:
		return Win
	case Paper:
		return Lose
	case Scissors:
		return Draw
	default:
		return UnknownOutcome
	}
}

func findLosePlay(opponentPlay PlayOption) PlayOption {
	switch opponentPlay {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		return UnknownPlay
	}
}

func findWinPlay(opponentPlay PlayOption) PlayOption {
	switch opponentPlay {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		return UnknownPlay
	}
}
