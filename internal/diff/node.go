package diff

const (
	Added     NodeType = "added"
	Removed   NodeType = "deleted"
	Changed   NodeType = "changed"
	Unchanged NodeType = "unchanged"
	Nested    NodeType = "nested"
)

type NodeType string

type Node struct {
	Key      string
	Type     NodeType
	OldValue any
	NewValue any
	Children []Node
}
