package cypher

import (
	"fmt"
	"strings"
)

type Prop struct {
	Key   string
	Value interface{}
}

type Props map[string]interface{}

func (p Props) ToCypher() string {
	var props string

	if len(p) > 0 {
		var propsArr []string
		for key, prop := range p {
			switch prop.(type) {
			case string:
				propsArr = append(propsArr, fmt.Sprintf("%v: '%s'", key, prop))
			default:
				propsArr = append(propsArr, fmt.Sprintf("%v: %v", key, prop))
			}
		}
		props = fmt.Sprintf("{%v}", strings.Join(propsArr, ", "))
	}

	return props
}
