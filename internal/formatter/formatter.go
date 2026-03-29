package formatter

import (
	"code/internal/diff"
	"code/internal/formatter/json"
	"code/internal/formatter/plain"
	"code/internal/formatter/stylish"
	"fmt"
)

const (
	JSON    = "json"
	PLAIN   = "plain"
	STYLISH = "stylish"
)

func Format(format string, diffTree []diff.Node) (string, error) {
	switch format {
	case JSON:
		return json.Format(diffTree)
	case PLAIN:
		return plain.Format(diffTree), nil
	case STYLISH, "":
		return stylish.Format(diffTree), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}
