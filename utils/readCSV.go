package utils

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filePath string, delimiter string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	transcations, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return transcations, nil
}