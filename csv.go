package yalzo

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ReadCSV(r io.Reader) ([]Todo, []Todo, error) {
	scanner := bufio.NewScanner(r)
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	todos := make([]Todo, 0, 100)
	archs := make([]Todo, 0, 100)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		arch := false

		no, err := strconv.Atoi(strings.TrimSpace(items[0]))

		if err != nil {
			return nil, nil, err
		}

		if strings.TrimSpace(items[3]) == "true" {
			arch = true
		}

		todo := &Todo{
			no:         no,
			label:      strings.TrimSpace(items[1]),
			title:      strings.TrimSpace(items[2]),
			isArchived: arch,
		}
		if arch {
			todos = append(todos, (*todo))
		} else {
			archs = append(archs, (*todo))
		}
	}
	return todos, archs, nil
}
