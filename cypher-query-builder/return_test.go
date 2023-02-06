package cypher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReturnConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var cypher string

	//name not defined
	t1 := ReturnConfig{
		Type: "a",
		As:   "abc",
	}
	_, err = t1.ToString()
	req.NotNil(err)

	//pattern
	t4 := ReturnConfig{
		Name: "n",
		Type: "m",
	}
	cypher, err = t4.ToString()
	req.Nil(err)
	req.EqualValues("n.m", cypher)

	//pattern
	t5 := ReturnConfig{
		Name: "n",
		Type: "m",
		As:   "abc",
	}
	cypher, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.m AS abc", cypher)
}
