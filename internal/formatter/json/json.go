package json

import (
	"encoding/json"

	"code/internal/diff"
)

func Format(diffTree []diff.Node) (string, error) {
	data := map[string]any{"diff": diffTree}
	jsonb, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonb), nil
}
