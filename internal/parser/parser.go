package parser

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

func ParseFile(path string) (map[string]any, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return map[string]any{}, err
	}

	fileData, err := os.ReadFile(absPath)
	if err != nil {
		return map[string]any{}, err
	}

	var parsedData map[string]any
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".json":
		err = json.Unmarshal(fileData, &parsedData)
	case ".yml", ".yaml":
		err = yaml.Unmarshal(fileData, &parsedData)
	default:
		return map[string]any{}, fmt.Errorf("unsupported file extension: %s", ext)
	}

	if err != nil {
		return map[string]any{}, err
	}

	return parsedData, nil
}
