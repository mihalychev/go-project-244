package stylish

import (
	"fmt"
	"slices"
	"strings"

	"code/internal/diff"
)

const (
	initialDepth = 1
	indentSize   = 4

	lineTemplate = "%s%s: %s"
)

func Format(diffTree []diff.Node) string {
	return formatNodes(diffTree, initialDepth)
}

func formatNodes(nodes []diff.Node, depth int) string {
	lines := []string{"{"}

	for _, node := range nodes {
		switch node.Type {
		case diff.Added:
			lines = append(lines,
				fmt.Sprintf(lineTemplate,
					nodePrefix(depth, "+"),
					node.Key,
					formatValue(node.NewValue, depth+1),
				),
			)

		case diff.Removed:
			lines = append(lines,
				fmt.Sprintf(lineTemplate,
					nodePrefix(depth, "-"),
					node.Key,
					formatValue(node.OldValue, depth+1),
				),
			)

		case diff.Changed:
			lines = append(lines,
				fmt.Sprintf(lineTemplate,
					nodePrefix(depth, "-"),
					node.Key,
					formatValue(node.OldValue, depth+1),
				),
			)
			lines = append(lines,
				fmt.Sprintf(lineTemplate,
					nodePrefix(depth, "+"),
					node.Key,
					formatValue(node.NewValue, depth+1),
				),
			)

		case diff.Unchanged:
			lines = append(lines,
				fmt.Sprintf(lineTemplate,
					nodePrefix(depth, " "),
					node.Key,
					formatValue(node.OldValue, depth+1),
				),
			)

		case diff.Nested:
			lines = append(lines,
				fmt.Sprintf(lineTemplate,
					nodePrefix(depth, " "),
					node.Key,
					formatNodes(node.Children, depth+1),
				),
			)

		}
	}

	lines = append(lines, closingIndent(depth)+"}")
	return strings.Join(lines, "\n")
}

func formatValue(value any, depth int) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case map[string]any:
		return formatMap(v, depth)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func formatMap(m map[string]any, depth int) string {
	lines := []string{"{"}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, key := range keys {
		lines = append(lines,
			fmt.Sprintf("%s%s: %s",
				nodePrefix(depth, " "),
				key,
				formatValue(m[key], depth+1),
			),
		)
	}

	lines = append(lines, closingIndent(depth)+"}")
	return strings.Join(lines, "\n")
}

func nodePrefix(depth int, sign string) string {
	return strings.Repeat(" ", depth*indentSize-2) + sign + " "
}

func closingIndent(depth int) string {
	return strings.Repeat(" ", (depth-1)*indentSize)
}
