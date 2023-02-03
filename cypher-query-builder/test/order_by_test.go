package test

import (
	"test/neo4j/pkg/pattern"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderByConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var cypher string

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
	cypher, err = t4.ToString()
	req.Nil(err)
	req.EqualValues("n.m", cypher)

	//pattern
	t5 := pattern.OrderByConfig{
		Name:   "n",
		Member: "m",
		Desc:   true,
	}
	cypher, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.m DESC", cypher)
}
