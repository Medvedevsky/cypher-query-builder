package cypher

import (
	"fmt"
	"reflect"
	"strings"
)

// Prop ...
type Prop struct {
	Key   string
	Value interface{}
}

// Props ...
type Props map[string]interface{}

func (p Props) ToCypher() string {
	var props string

	if len(p) > 0 {
		var propsArr []string
		for key, prop := range p {

			t := reflect.TypeOf(prop)
			k := t.Kind()

			if k == reflect.Bool {
				propsArr = append(propsArr, fmt.Sprintf("%v: %t", key, prop.(bool)))
			}

			if k == reflect.String {
				propsArr = append(propsArr, fmt.Sprintf("%v: '%v'", key, prop))
				break
			}

			if k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 ||
				k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 ||
				k == reflect.Float32 || k == reflect.Float64 {
				propsArr = append(propsArr, fmt.Sprintf("%v: %v", key, prop))
			}

		}
		props = fmt.Sprintf("{%v}", strings.Join(propsArr, ", "))
	}

	return props
}
