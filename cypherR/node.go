package cypherr

import (
	"fmt"
	"strings"
)

type Node struct {
	*Label
	Variable   string
	Properties Props
}

func NewNode() *Node {
	return &Node{
		Label:      NewLabel(),
		Properties: Props{},
		Variable:   "",
	}
}

func (n *Node) SetVariable(name string) *Node {
	n.Variable = name
	return n
}

func (n *Node) SetProps(props ...Prop) *Node {
	for _, p := range props {
		n.Properties[p.Key] = p.Value
	}
	return n
}

func (n *Node) SetLabel(label string) *Node {
	n.Names = append(n.Names, label)
	return n
}

func (n *Node) SetLabels(condition Condition, labels ...string) *Node {
	n.Names = append(n.Names, labels...)
	n.Condition = condition
	return n
}

func (n *Node) ToPattern() QueryPattern {
	return QueryPattern{OnlyNode: OnlyNode{Node: n}}
}

func (n *Node) ToCypher() string {
	node := "("

	if n.Variable != "" {
		node += n.Variable
	}

	if len(n.Label.Names) > 0 {
		condition := ""
		if n.Label.Condition != "" {
			condition = fmt.Sprintf("%v", n.Label.Condition)
		}
		node += fmt.Sprintf(":%v", strings.Join(n.Label.Names, condition)) + " "
	}

	if len(n.Properties) > 0 {
		node += n.Properties.ToCypher()
	}

	return node + ")"
}
