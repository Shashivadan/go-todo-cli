package util

import (
	"encoding/csv"
	"errors"
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

func DeleteTodo(id string, fileName string) (string, error) {
	readFile, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}

	reader := csv.NewReader(readFile)

	data, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	idx, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	if len(data) < idx || idx == 0 {
		return "", errors.New("this id is not present in todos")
	}

	updatedData := make([][]string, 0)
	for idx, val := range data {
		if val[0] != id {
			if idx == 0 {
				updatedData = append(updatedData, []string{val[0], val[1], val[2]})
			} else {
				updatedData = append(updatedData, []string{strconv.Itoa(len(updatedData)), val[1], val[2]})
			}
		}
	}
	readFile.Close()
	writeFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}

	overwriteCsv := csv.NewWriter(writeFile)

	err = overwriteCsv.WriteAll(updatedData)
	if err != nil {
		return "", err
	}
	defer overwriteCsv.Flush()
	return "deleted sussufull", nil
}

func DoenTodo(id string, fileName string) (string, error) {
	readFile, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	reader := csv.NewReader(readFile)

	data, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	idx, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}
	if len(data) < idx || idx == 0 {
		return "", errors.New("this id is not present in todos")
	}
	updatedData := make([][]string, 0)

	for _, val := range data {
		if val[0] == id {
			updatedData = append(updatedData, []string{val[0], val[1], "doen"})
			continue
		}
		updatedData = append(updatedData, val)
	}
	readFile.Close()
	writeFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}
	overWriteFile := csv.NewWriter(writeFile)
	err = overWriteFile.WriteAll(updatedData)
	if err != nil {
		return "", err
	}
	return "task completed doen", nil
}
