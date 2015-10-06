package yalzo

import (
	"bufio"
	"fmt"
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
			archs = append(archs, (*todo))
		} else {
			todos = append(todos, (*todo))
		}
	}

	return todos, archs, nil
}

func SaveCSV(todos []Todo, w io.Writer) {
	size := len(todos)
	for i := 0; i < size; i++ {
		no := strconv.Itoa(todos[i].no)
		l := todos[i].label
		t := todos[i].title

		if todos[i].isArchived {
			fmt.Fprintf(w, "%s,%s,%s,%s", no, l, t, "true")
		} else {
			fmt.Fprintf(w, "%s,%s,%s,%s", no, l, t, "false")
		}

		// if not last offset, append return '\n' to tail.
		if i+1 != size {
			fmt.Fprintf(w, "\n")
		}
	}
}
