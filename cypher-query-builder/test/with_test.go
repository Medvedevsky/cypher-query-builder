package test

import (
	"testing"

	"github.com/Medvedevsky/cypher-query-builder/pkg/cypher"
	"github.com/Medvedevsky/cypher-query-builder/pkg/pattern"

	"github.com/stretchr/testify/require"
)

func TestWithConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var res string

	//must define a function or name
	t1 := pattern.WithConfig{}
	_, err = t1.ToString()
	req.NotNil(err)

	//pattern
	t4 := pattern.WithConfig{
		Name:  "n",
		Field: "m",
	}
	res, err = t4.ToString()
	req.Nil(err)
	req.EqualValues("n.m", res)

	//pattern
	t5 := pattern.WithConfig{
		Name:  "n",
		Field: "m",
		As:    "abc",
	}
	res, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.m AS abc", res)

	// WITH clause
	t6, err := cypher.NewQueryBuilder().With(pattern.WithConfig{
		Name:  "n",
		Field: "m",
		As:    "abc",
	}).Execute()
	req.Nil(err)
	req.EqualValues("WITH n.m AS abc", t6)
}
