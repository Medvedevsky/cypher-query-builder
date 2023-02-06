package cypher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProp_ToCypher(t *testing.T) {
	req := require.New(t)
	var cypher string

	//pattern
	t1 := Props{}
	t1["key"] = "value"
	cypher = t1.ToCypher()
	req.EqualValues("{key: 'value'}", cypher)

	//pattern
	t2 := Props{}
	t2["key"] = true
	cypher = t2.ToCypher()
	req.EqualValues("{key: true}", cypher)

	//pattern
	t3 := Props{}
	t3["key"] = 123
	cypher = t3.ToCypher()
	req.EqualValues("{key: 123}", cypher)
}
