package inputfile

import (
	"bufio"
	"os"
)

func Read(f string) ([]string, error) {
	var lines []string
	file, err := os.Open(f)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	bufScanner := bufio.NewScanner(file)

	for bufScanner.Scan() {
		lines = append(lines, bufScanner.Text())
	}

	return lines, bufScanner.Err()
}

func ReadGroups(f string) ([][]string, error) {
	lines, err := Read(f)
	if err != nil {
		return nil, err
	}

	var answer [][]string
	var currentGroup []string
	for _, v := range lines {
		if len(v) == 0 {
			if len(currentGroup) > 0 {
				answer = append(answer, currentGroup)
				currentGroup = []string{}
			} else {
				continue
			}
		} else {
			currentGroup = append(currentGroup, v)
		}
	}

	if len(currentGroup) > 0 {
		answer = append(answer, currentGroup)
	}

	return answer, nil
}
