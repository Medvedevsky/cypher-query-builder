package cypher

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
	LessThanOperator             BooleanOperator = "<"
	GreaterThanOperator          BooleanOperator = ">"
	LessThanOrEqualToOperator    BooleanOperator = "<="
	GreaterThanOrEqualToOperator BooleanOperator = ">="
	EqualToOperator              BooleanOperator = "="
	InOperator                   BooleanOperator = "IN"
	IsOperator                   BooleanOperator = "IS"
	RegexEqualToOperator         BooleanOperator = "=~"
	StartsWithOperator           BooleanOperator = "STARTS WITH"
	EndsWithOperator             BooleanOperator = "ENDS WITH"
	ContainsOperator             BooleanOperator = "CONTAINS"
)

type OrderByOperator string

const (
	Asc  OrderByOperator = "ASC"
	Desc OrderByOperator = "DESC"
)

// Path ...
type Path string

const (
	// Plain --
	Plain Path = "--"

	// Outgoing -->
	Outgoing Path = "-->"

	// Incoming <--
	Incoming Path = "<--"

	// Bidirectional <-->
	Bidirectional Path = "<-->"
)
