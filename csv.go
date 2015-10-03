package yalzo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadCSV(fp *os.File) ([]Todo, error) {
	scanner := bufio.NewScanner(fp)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	todos := make([]Todo, 0, 100)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")

		no, err := strconv.Atoi(strings.TrimSpace(items[0]))
		if err != nil {
			return nil, err
		}

		todo := &Todo{
			no:    no,
			label: strings.TrimSpace(items[1]),
			title: strings.TrimSpace(items[2]),
		}
		todos = append(todos, (*todo))
	}
	return todos, nil
}
