package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCsvFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	return records, nil
}

func CreateCsvFile(fileName string) ([][]string, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	header := []string{"ID", "TASK", "STATUS"}

	writer := csv.NewWriter(file)
	writer.Write(header)

	defer writer.Flush()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WrietCsvFile() {
}

func AddTodo(task string, fileName string) ([][]string, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	todo := []string{strconv.Itoa(len(data)), task, "undoen"}

	writer := csv.NewWriter(file)

	writer.Write(todo)

	defer writer.Flush()

	if writer.Error() != nil {
		return nil, writer.Error()
	}
	return data, nil
}

func DeleteTodo(number string, fileName string) (string, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}

	defer file.Close()

	return "deleted sussufull", nil
}
