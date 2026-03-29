package plain

import (
	"fmt"
	"strings"

	"code/internal/diff"
)

func Format(diffTree []diff.Node) string {
	return formatNodes(diffTree, "")
}

func formatNodes(nodes []diff.Node, parents string) string {
	lines := []string{}

	for _, node := range nodes {
		switch node.Type {
		case diff.Added:
			lines = append(lines,
				fmt.Sprintf("Property '%s' was added with value: %v",
					formatKey(node.Key, parents),
					formatValue(node.NewValue),
				),
			)

		case diff.Removed:
			lines = append(lines,
				fmt.Sprintf("Property '%s' was removed",
					formatKey(node.Key, parents),
				),
			)

		case diff.Changed:
			lines = append(lines,
				fmt.Sprintf("Property '%s' was updated. From %v to %v",
					formatKey(node.Key, parents),
					formatValue(node.OldValue),
					formatValue(node.NewValue),
				),
			)

		case diff.Unchanged:
			continue

		case diff.Nested:
			lines = append(lines,
				formatNodes(node.Children, formatKey(node.Key, parents)),
			)
		}
	}

	return strings.Join(lines, "\n")
}

func formatKey(key, parents string) string {
	if parents == "" {
		return key
	}

	return fmt.Sprintf("%s.%s", parents, key)
}

func formatValue(value any) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case string:
		return fmt.Sprintf("'%s'", v)
	case map[string]any:
		return "[complex value]"
	default:
		return fmt.Sprintf("%v", v)
	}
}
