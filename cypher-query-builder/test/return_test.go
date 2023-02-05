package test

import (
	"testing"

	"github.com/Medvedevsky/cypher-query-builder/pkg/pattern"

	"github.com/stretchr/testify/require"
)

func TestReturnConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var cypher string

	//name not defined
	t1 := pattern.ReturnConfig{
		Type: "a",
		As:   "abc",
	}
	_, err = t1.ToString()
	req.NotNil(err)

	//pattern
	t4 := pattern.ReturnConfig{
		Name: "n",
		Type: "m",
	}
	cypher, err = t4.ToString()
	req.Nil(err)
	req.EqualValues("n.m", cypher)

	//pattern
	t5 := pattern.ReturnConfig{
		Name: "n",
		Type: "m",
		As:   "abc",
	}
	cypher, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.m AS abc", cypher)
}
