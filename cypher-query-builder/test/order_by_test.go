package test

import (
	"test/neo4j/pkg/cypher"
	"test/neo4j/pkg/pattern"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderByConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var res string

	//name not defined
	t1 := pattern.OrderByConfig{
		Member: "abc",
	}
	_, err = t1.ToString()
	req.NotNil(err)

	//member not defined
	t2 := pattern.OrderByConfig{
		Name: "abc",
	}
	_, err = t2.ToString()
	req.NotNil(err)

	//both member and name not defined
	t3 := pattern.OrderByConfig{}
	_, err = t3.ToString()
	req.NotNil(err)

	//pattern
	t4 := pattern.OrderByConfig{
		Name:   "n",
		Member: "m",
	}
	res, err = t4.ToString()
	req.Nil(err)
	req.EqualValues("n.m", res)

	//pattern
	t5 := pattern.OrderByConfig{
		Name:   "n",
		Member: "m",
		Desc:   true,
	}
	res, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.m DESC", res)

	//clause ORDER BY
	t6, err := cypher.NewQueryBuilder().OrderBy(pattern.OrderByConfig{
		Name:   "n",
		Member: "m",
		Desc:   true,
	}).Execute()
	req.Nil(err)
	req.EqualValues("ORDER BY n.m DESC", t6)
}
