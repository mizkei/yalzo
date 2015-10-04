package yalzo

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

func ReadCSV(fp *os.File) ([]Todo, []Todo, error) {
	scanner := bufio.NewScanner(fp)
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

func SaveCSV(todos []Todo, archs []Todo, fp *os.File) {
	w := bufio.NewWriter(fp)
	csv_list := append(createTodoCSV(todos), createTodoCSV(archs)...)
	buf := bytes.NewBufferString("")
	for i := 0; i < len(csv_list); i++ {
		buf.WriteString(csv_list[i])
		// if not last offset, append return '\n' to tail.
		if i+1 != len(csv_list[i]) {
			buf.WriteString("\n")
		}
	}
}

func createTodoCSV(todos []Todo) []string {
	csv := make([]string, 0, len(todos))
	for i := 0; i < len(todos); i++ {
		todo := todos[i]
		buf := bytes.NewBufferString(strconv.Itoa(todo.no))
		buf.WriteString(",")
		buf.WriteString(todo.label)
		buf.WriteString(",")
		buf.WriteString(todo.title)
		buf.WriteString(",")
		if todo.isArchived {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}
		csv[i] = buf.String()
	}
	return csv
}
