package diff

import (
	"maps"
	"reflect"
	"slices"
)

func BuildDiffTree(data1, data2 map[string]any) []Node {
	keys := unionSortedKeys(data1, data2)
	tree := []Node{}

	for _, key := range keys {
		val1, ok1 := data1[key]
		val2, ok2 := data2[key]

		if !ok1 {
			tree = append(tree, Node{Key: key, Type: Added, NewValue: val2})
			continue
		}

		if !ok2 {
			tree = append(tree, Node{Key: key, Type: Removed, OldValue: val1})
			continue
		}

		map1, isMap1 := val1.(map[string]any)
		map2, isMap2 := val2.(map[string]any)

		switch {
		case isMap1 && isMap2:
			tree = append(tree, Node{
				Key:      key,
				Type:     Nested,
				Children: BuildDiffTree(map1, map2),
			})
		case reflect.DeepEqual(val1, val2):
			tree = append(tree, Node{Key: key, Type: Unchanged, OldValue: val1})
		default:
			tree = append(tree, Node{Key: key, Type: Changed, OldValue: val1, NewValue: val2})
		}
	}

	return tree
}

func unionSortedKeys(data1, data2 map[string]any) []string {
	combinedMap := maps.Clone(data1)
	maps.Copy(combinedMap, data2)

	sortedKeys := slices.Collect(maps.Keys(combinedMap))
	slices.Sort(sortedKeys)

	return sortedKeys
}
