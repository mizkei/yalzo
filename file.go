package yalzo

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

func ReadCSV(r io.Reader) ([]Todo, []Todo, error) {
	reader := csv.NewReader(r)
	reader.LazyQuotes = true
	reader.Comment = '#'

	todos := make([]Todo, 0, 100)
	archs := make([]Todo, 0, 100)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		arch := false

		no, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, nil, err
		}

		if record[3] == "true" {
			arch = true
		}

		todo := &Todo{
			no:         no,
			label:      strings.TrimSpace(record[1]),
			title:      strings.TrimSpace(record[2]),
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
	writer := csv.NewWriter(w)

	for i := 0; i < size; i++ {
		writer.Write([]string{
			strconv.Itoa(todos[i].no),
			todos[i].label,
			todos[i].title,
			strconv.FormatBool(todos[i].isArchived),
		})
	}
	writer.Flush()
}
