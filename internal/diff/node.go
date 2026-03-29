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
	Key      string   `json:"key"`
	Type     NodeType `json:"type"`
	OldValue any      `json:"oldValue,omitempty"`
	NewValue any      `json:"newValue,omitempty"`
	Children []Node   `json:"children,omitempty"`
}
