package cypher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExampleMatch_AllNodes(t *testing.T) {
	req := require.New(t)
	var err error

	/*
		MATCH (n)
		RETURN n
	*/
	t1, err := NewQueryBuilder().
		Match(NewNode().SetVariable("n").AsPattern()).
		Return(ReturnConfig{Name: "n"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (n)\nRETURN n", t1)
}

func TestExampleMatch_Node(t *testing.T) {
	req := require.New(t)
	var err error

	/*
		MATCH (n:`My Label`)
		RETURN n
	*/
	node := NewNode().SetVariable("n").SetLabel("My Label").AsPattern()

	t1, err := NewQueryBuilder().
		Match(node).
		Return(ReturnConfig{Name: "n"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (n:`My Label`)\nRETURN n", t1)
}

func TestExampleMatch_NodeByLabelAndAttributes(t *testing.T) {
	req := require.New(t)
	var err error

	/*
		MATCH (n:`My Label`:`Our Label`)
		WHERE n.attr1 = 'value 1' AND n.attr2 = 'value 2'
		RETURN n
	*/
	node := NewNode().SetVariable("n").SetLabels(Сolon, "My Label", "Our Label").AsPattern()

	t1, err := NewQueryBuilder().
		Match(node).
		Where(ConditionalConfig{
			Name:              "n",
			Field:             "attr1",
			ConditionOperator: EqualToOperator,
			Check:             "value 1",
			Condition:         AND}, ConditionalConfig{
			Name:              "n",
			Field:             "attr2",
			ConditionOperator: EqualToOperator,
			Check:             "value 2",
		}).
		Return(ReturnConfig{Name: "n"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (n:`My Label`:`Our Label`)\nWHERE n.attr1 = 'value 1' AND n.attr2 = 'value 2'\nRETURN n", t1)
}

func TestExampleMatch_Optional(t *testing.T) {
	req := require.New(t)
	var err error

	/*
		MATCH (n:`My Label`)
		OPTIONAL MATCH (n)-[]->(m)
		RETURN n, m
	*/
	edge := NewEdge().SetPath(Outgoing).Relationship(FullRelationship{
		LeftNode:  NewNode().SetVariable("n"),
		RightNode: NewNode().SetVariable("m"),
	})

	t1, err := NewQueryBuilder().
		Match(NewNode().SetVariable("n").SetLabel("My Label").AsPattern()).
		OptionlMath(edge).
		Return(ReturnConfig{Name: "n"}, ReturnConfig{Name: "m"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (n:`My Label`)\nOPTIONAL MATCH (n)-[]->(m)\nRETURN n, m", t1)
}

func TestExampleReturn_Property(t *testing.T) {
	req := require.New(t)
	var err error

	/*
		MATCH (n:`My Label`)
		RETURN n.fitstProperty
	*/
	node := NewNode().SetVariable("n").SetLabel("My Label").AsPattern()

	t1, err := NewQueryBuilder().
		Match(node).
		Return(ReturnConfig{Name: "n", Type: "fitstProperty"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (n:`My Label`)\nRETURN n.fitstProperty", t1)
}
