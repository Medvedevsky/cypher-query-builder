package cypher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNode_ToString(t *testing.T) {
	req := require.New(t)
	var cypher string
	var err error

	//Node must have a variable with at least one label
	_, err = NewNode().SetLabel("TEST").ToCypher()
	req.NotNil(err)

	//pattern
	cypher, err = NewNode().SetVariable("n").ToCypher()
	req.NoError(err)
	req.EqualValues("(n)", cypher)

	//pattern
	cypher, err = NewNode().
		SetVariable("n").
		SetLabel("TEST").
		ToCypher()

	req.NoError(err)
	req.EqualValues("(n:TEST)", cypher)

	//pattern
	cypher, err = NewNode().
		SetVariable("n").
		SetLabels(And, "TEST", "TEST2").
		ToCypher()

	req.NoError(err)
	req.EqualValues("(n:TEST&TEST2)", cypher)

	//pattern
	cypher, err = NewNode().
		SetVariable("n").
		SetLabel("TEST").
		SetProps(Prop{Key: "p", Value: "name"}).
		ToCypher()

	req.NoError(err)
	req.EqualValues("(n:TEST {p: 'name'})", cypher)
}
