package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("Could not read file")
	}
	// convert byte[] into string
	valueString := string(data)
	// convert string to float
	value, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return defaultValue, errors.New("Could not parse stored value")
	}
	// only return float
	return value, nil
}

func WriteValueToFile(fileName string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
