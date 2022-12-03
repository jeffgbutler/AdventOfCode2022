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
