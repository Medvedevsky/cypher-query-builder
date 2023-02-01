package cypherr

type QueryPattern struct {
	PartialRelationship PartialRelationship
	FullRelationship    FullRelationship
	OnlyNode            OnlyNode
	Edge                Edge
}

type PartialRelationship struct {
	LeftDirection  bool
	RightDirection bool
	Node           *Node
}

type FullRelationship struct {
	LeftNode  *Node
	RightNode *Node
}

type OnlyNode struct {
	Node *Node
}

type WhereQuery struct {
	Name            string
	Field           string
	Check           interface{}
	Condition       Condition
	BooleanOperator BooleanOperator
}

type ReturnQuery struct {
	Name  string
	Field string
}

type Condition string

const (
	// And symbol condition "&"
	And Condition = "&"

	// Or symbol condition "|"
	Or Condition = "|"

	AND Condition = "AND"
	OR  Condition = "OR"
)

type BooleanOperator string

const (
	EqualToOperator BooleanOperator = "="
	InOperator      BooleanOperator = "IN"
	IsOperator      BooleanOperator = "IS"
)
