package test

import (
	"test/neo4j/pkg/pattern"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConditionalConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var cypher string

	//must define a function or name
	t1 := pattern.ConditionalConfig{
		Field:             "a",
		Label:             "b",
		ConditionFunction: "c",
	}
	_, err = t1.ToString()
	req.NotNil(err)

	//only one of field or label can be set
	t2 := pattern.ConditionalConfig{
		Name:  "a",
		Field: "b",
		Label: "c",
	}
	_, err = t2.ToString()
	req.NotNil(err)

	//condition operator can not be empty with var check
	t3 := pattern.ConditionalConfig{
		Name:  "a",
		Field: "b",
		Check: 2,
	}
	_, err = t3.ToString()
	req.NotNil(err)

	//only one of label or condition function can be set
	t4 := pattern.ConditionalConfig{
		Name:              "a",
		Label:             "b",
		ConditionFunction: "tfunc",
	}
	_, err = t4.ToString()
	req.NotNil(err)

	//pattern
	t5 := pattern.ConditionalConfig{
		Name:  "n",
		Field: "m",
	}
	cypher, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.m", cypher)

	//pattern
	t6 := pattern.ConditionalConfig{
		Name:  "n",
		Label: "L",
	}
	cypher, err = t6.ToString()
	req.Nil(err)
	req.EqualValues("n:L", cypher)

	//pattern
	t7 := pattern.ConditionalConfig{
		Name:              "n",
		Field:             "l",
		ConditionOperator: pattern.EqualToOperator,
		Check:             21,
	}
	cypher, err = t7.ToString()
	req.Nil(err)
	req.EqualValues("n.l = 21", cypher)

	//pattern
	t8 := pattern.ConditionalConfig{
		Name:              "n",
		Field:             "l",
		ConditionFunction: "tfunc",
	}
	cypher, err = t8.ToString()
	req.Nil(err)
	req.EqualValues("tfunc(n.l)", cypher)
}
