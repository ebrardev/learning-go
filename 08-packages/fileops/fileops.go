package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("could not read file: " + err.Error())
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("could not stored  to value: " + err.Error())
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprintf("%f", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
