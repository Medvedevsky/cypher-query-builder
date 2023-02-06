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
		MATCH (n:`My Label``)
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
