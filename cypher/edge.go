package cypher

import (
	"fmt"
	"strings"
)

// Edge ...
type Edge struct {
	variable   string
	label      Label
	properties Props
	path       Path
	condition  Condition
}

// NewEdge ...
func NewEdge() *Edge {
	return &Edge{
		variable:   "",
		label:      Label{},
		properties: Props{},
		path:       "",
	}
}

func (e *Edge) SetVariable(name string) *Edge {
	e.variable = name
	return e
}

func (e *Edge) SetLabel(label string) *Edge {
	e.label.Names = append(e.label.Names, label)
	return e
}

func (e *Edge) SetLabels(condition Condition, labels ...string) *Edge {
	e.label.Names = append(e.label.Names, labels...)
	e.condition = condition
	return e
}

func (e *Edge) SetProps(props ...Prop) *Edge {
	for _, p := range props {
		e.properties[p.Key] = p.Value
	}
	return e
}

func (e *Edge) SetPath(path Path) *Edge {
	e.path = path
	return e
}

// Relationship ...
func (e Edge) Relationship(f FullRelationship) QueryPattern {
	// f.Edge = e
	return QueryPattern{
		FullRelationship: f,
		Edge:             e,
	}
}

func (e Edge) PartialRelationship(p PartialRelationship) QueryPattern {
	// p.Edge = e
	return QueryPattern{
		PartialRelationship: p,
		Edge:                e,
	}
}

func (e Edge) PartialRelationshipBuild(p PartialRelationship) string {
	//switch
	if p.LeftDirection {
		return fmt.Sprintf("%v%v", p.Node.ToCypher(), e.ToCypher())
	}

	if p.RightDirection {
		return fmt.Sprintf("%v%v", e.ToCypher(), p.Node.ToCypher())
	}

	fmt.Println("error not set type direction")
	return ""
}

func (e Edge) RelationshipBuild(f FullRelationship) string {

	if f.LeftNode == nil || f.RightNode == nil {
		//error
		fmt.Println("error RelationshipBuild not have nodes")
		return ""
	}

	leftNode := f.LeftNode
	rightNode := f.RightNode

	return fmt.Sprintf("%v%v%v", leftNode.ToCypher(), e.ToCypher(), rightNode.ToCypher())
}

func (e Edge) RelationshipMulti(edges ...string) string {
	res := ""
	for _, edge := range edges {
		res += fmt.Sprintf("%v", edge)
	}
	return res
}

func (e Edge) ToCypher() string {
	edge := ""

	if e.variable != "" {
		edge += e.variable
	}

	if len(e.label.Names) > 0 {
		condition := ""
		if e.label.Condition != "" {
			condition = fmt.Sprintf("%v", e.label.Condition)
		}
		edge += fmt.Sprintf(":%v", strings.Join(e.label.Names, condition)) + " "
	}

	if len(e.properties) > 0 {
		edge += e.properties.ToCypher()
	}

	edge = fmt.Sprintf("-[%v]-", edge)

	switch e.path {
	case Outgoing:
		edge += ">"
	case Incoming:
		edge = "<" + edge
	case Bidirectional:
		edge = "<" + edge + ">"
	case Plain:
	default:
	}

	return edge
}
