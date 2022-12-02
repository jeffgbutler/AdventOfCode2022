package Day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PlayOption int
type Outcome int

const (
	Rock     PlayOption = 1
	Paper               = 2
	Scissors            = 3
)

const (
	Win  Outcome = 6
	Lose         = 0
	Draw         = 3
)

type round struct {
	opponentPlay PlayOption
	myPlay       PlayOption
}

func readInputFilePart1(f string) ([]round, error) {
	var lines []round
	file, err := os.Open(f)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	bufScanner := bufio.NewScanner(file)

	for bufScanner.Scan() {
		line := bufScanner.Text()
		s := strings.Fields(line)
		r, err := translatePart1(s)
		if err != nil {
			return lines, err
		}
		lines = append(lines, r)
	}

	return lines, bufScanner.Err()
}

func readInputFilePart2(f string) ([]round, error) {
	var lines []round
	file, err := os.Open(f)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	bufScanner := bufio.NewScanner(file)

	for bufScanner.Scan() {
		line := bufScanner.Text()
		s := strings.Fields(line)
		r, err := translatePart2(s)
		if err != nil {
			return lines, err
		}
		lines = append(lines, r)
	}

	return lines, bufScanner.Err()
}

func translatePart1(someRound []string) (round, error) {
	opponentPlay, err := translatePlay(someRound[0])
	if err != nil {
		return round{}, err
	}

	myPlay, err := translatePlay(someRound[1])
	if err != nil {
		return round{}, err
	}

	theRound := round{opponentPlay, myPlay}
	return theRound, nil
}

func translatePart2(someRound []string) (round, error) {
	opponentPlay, err := translatePlay(someRound[0])
	if err != nil {
		return round{Rock, Rock}, err
	}

	myPlay, err := translateMyPlayPart2(opponentPlay, someRound[1])
	if err != nil {
		return round{}, err
	}

	theRound := round{opponentPlay, myPlay}
	return theRound, nil
}

func calculateMatchOutcome(inputs []round) (int, error) {
	total := 0
	for _, someRound := range inputs {
		roundValue, err := calculateRoundValue(someRound)
		if err != nil {
			return 0, err
		}

		total += roundValue
	}

	return total, nil
}

func translatePlay(play string) (PlayOption, error) {
	switch play {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	default:
		return 0, fmt.Errorf("unexpected PlayOption %s", play)
	}
}

func translateMyPlayPart2(opponentPlay PlayOption, myPlay string) (PlayOption, error) {
	switch myPlay {
	case "X": // lose
		return findLosePlay(opponentPlay)
	case "Y": // draw
		return opponentPlay, nil
	case "Z": // win
		return findWinPlay(opponentPlay)
	default:
		return Rock, fmt.Errorf("unexpected PlayOption %s", myPlay)
	}
}

func calculateRoundValue(theRound round) (int, error) {
	outcome, err := calculateOutCome(theRound)
	if err != nil {
		return 0, err
	}

	return int(outcome) + int(theRound.myPlay), nil
}

func calculateOutCome(theRound round) (Outcome, error) {
	switch theRound.opponentPlay {
	case Rock:
		return outcomeVsRock(theRound.myPlay)
	case Paper:
		return outcomeVsPaper(theRound.myPlay)
	case Scissors:
		return outcomeVsScissors(theRound.myPlay)
	default:
		return 0, fmt.Errorf("unexpected PlayOption %d", int(theRound.opponentPlay))
	}
}

func outcomeVsRock(myPlay PlayOption) (Outcome, error) {
	switch myPlay {
	case Rock:
		return Draw, nil
	case Paper:
		return Win, nil
	case Scissors:
		return Lose, nil
	default:
		return 0, fmt.Errorf("unexpected PlayOption %d", int(myPlay))
	}
}

func outcomeVsPaper(myPlay PlayOption) (Outcome, error) {
	switch myPlay {
	case Rock:
		return Lose, nil
	case Paper:
		return Draw, nil
	case Scissors:
		return Win, nil
	default:
		return 0, fmt.Errorf("unexpected PlayOption %d", int(myPlay))
	}
}

func outcomeVsScissors(myPlay PlayOption) (Outcome, error) {
	switch myPlay {
	case Rock:
		return Win, nil
	case Paper:
		return Lose, nil
	case Scissors:
		return Draw, nil
	default:
		return 0, fmt.Errorf("unexpected PlayOption %d", int(myPlay))
	}
}

func findLosePlay(opponentPlay PlayOption) (PlayOption, error) {
	switch opponentPlay {
	case Rock:
		return Scissors, nil
	case Paper:
		return Rock, nil
	case Scissors:
		return Paper, nil
	default:
		return 0, fmt.Errorf("unexpected PlayOption %d", int(opponentPlay))
	}
}

func findWinPlay(opponentPlay PlayOption) (PlayOption, error) {
	switch opponentPlay {
	case Rock:
		return Paper, nil
	case Paper:
		return Scissors, nil
	case Scissors:
		return Rock, nil
	default:
		return 0, fmt.Errorf("unexpected PlayOption %d", int(opponentPlay))
	}
}
