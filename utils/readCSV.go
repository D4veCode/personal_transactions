package utils

import (
	"encoding/csv"
	"mime/multipart"
)

func ReadCSV(file multipart.File, delimiter string) ([][]string, error) {

	reader := csv.NewReader(file)
	transcations, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return transcations, nil
}