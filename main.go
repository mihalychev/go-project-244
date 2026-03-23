package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Gendiff(filepath1, filepath2 string) (string, error) {
	fileData1, err := parseFile(filepath1)
	if err != nil {
		return "", err
	}
	fileData2, err := parseFile(filepath2)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v\n%v", fileData1, fileData2), nil
}

func parseFile(path string) (map[string]interface{}, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return map[string]interface{}{}, err
	}

	fileData, err := os.ReadFile(absPath)
	if err != nil {
		return map[string]interface{}{}, err
	}

	parsedData, err := parseJSON(fileData)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return parsedData, nil
}

func parseJSON(fileData []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(fileData, &data)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil
}
