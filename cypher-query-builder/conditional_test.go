package cypher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConditionalConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var res string

	//must define a function or name
	t1 := ConditionalConfig{
		Field:             "a",
		Label:             "b",
		ConditionFunction: "c",
	}
	_, err = t1.ToString()
	req.NotNil(err)

	//only one of field or label can be set
	t2 := ConditionalConfig{
		Name:  "a",
		Field: "b",
		Label: "c",
	}
	_, err = t2.ToString()
	req.NotNil(err)

	//condition operator can not be empty with var check
	t3 := ConditionalConfig{
		Name:  "a",
		Field: "b",
		Check: 2,
	}
	_, err = t3.ToString()
	req.NotNil(err)

	//only one of label or condition function can be set
	t4 := ConditionalConfig{
		Name:              "a",
		Label:             "b",
		ConditionFunction: "tfunc",
	}
	_, err = t4.ToString()
	req.NotNil(err)

	//pattern
	t5 := ConditionalConfig{
		Name:  "n",
		Field: "m",
	}
	res, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.m", res)

	//pattern
	t6 := ConditionalConfig{
		Name:  "n",
		Label: "L",
	}
	res, err = t6.ToString()
	req.Nil(err)
	req.EqualValues("n:L", res)

	//pattern
	t7 := ConditionalConfig{
		Name:              "n",
		Field:             "l",
		ConditionOperator: EqualToOperator,
		Check:             21,
	}
	res, err = t7.ToString()
	req.Nil(err)
	req.EqualValues("n.l = 21", res)

	//pattern
	t8 := ConditionalConfig{
		Name:              "n",
		Field:             "l",
		ConditionFunction: "tfunc",
	}
	res, err = t8.ToString()
	req.Nil(err)
	req.EqualValues("tfunc(n.l)", res)

	// where condition
	t9, err := NewQueryBuilder().
		Where(ConditionalConfig{
			Name:              "p",
			Field:             "age",
			ConditionOperator: EqualToOperator,
			Check:             12,
			Condition:         AND,
		}, ConditionalConfig{
			Name:              "p",
			Field:             "height",
			ConditionOperator: EqualToOperator,
			Check:             150,
		}).Execute()
	req.NoError(err)
	req.EqualValues("WHERE p.age = 12 AND p.height = 150", t9)
}
