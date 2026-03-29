package formatter

import (
	"code/internal/diff"
	"fmt"
)

const (
	PLAIN   = "plain"
	STYLISH = "stylish"
)

func Format(format string, diffTree []diff.Node) (string, error) {
	switch format {
	case PLAIN:
		return FormatPlain(), nil
	case STYLISH, "":
		return FormatStylish(diffTree), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}
