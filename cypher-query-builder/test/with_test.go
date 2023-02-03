package test

import (
	"test/neo4j/pkg/pattern"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var cypher string

	//must define a function or name
	t1 := pattern.WithConfig{}
	_, err = t1.ToString()
	req.NotNil(err)

	//pattern
	t4 := pattern.WithConfig{
		Name:  "n",
		Field: "m",
	}
	cypher, err = t4.ToString()
	req.Nil(err)
	req.EqualValues("n.m", cypher)

	//pattern
	t5 := pattern.WithConfig{
		Name:  "n",
		Field: "m",
		As:    "abc",
	}
	cypher, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.m AS abc", cypher)
}
