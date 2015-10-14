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

	inproc := make([]Todo, 0, 100)
	archs := make([]Todo, 0, 100)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		isArchived, err := strconv.ParseBool(record[2])
		if err != nil {
			return nil, nil, err
		}

		todo := &Todo{
			label: strings.TrimSpace(record[0]),
			title: strings.TrimSpace(record[1]),
		}

		if isArchived {
			archs = append(archs, *todo)
		} else {
			inproc = append(inproc, *todo)
		}
	}

	return inproc, archs, nil
}

func SaveCSV(todos []Todo, isArchived bool, w io.Writer) {
	writer := csv.NewWriter(w)

	for _, v := range todos {
		writer.Write([]string{
			v.label,
			v.title,
			strconv.FormatBool(isArchived),
		})
	}
	writer.Flush()
}
