package cypher

import (
	"fmt"
	"strings"
)

type Label struct {
	Names     []string
	Condition Condition
}

func (l Label) ToCypher() string {
	res := ""
	if len(l.Names) > 0 {
		res += ":"
		condition := ""
		if l.Condition != "" {
			condition = fmt.Sprintf("%v", l.Condition)
		}

		for _, name := range l.Names {
			if strings.Contains(name, " ") {
				res += fmt.Sprintf("`%v`%v", name, condition)
			} else {
				res += fmt.Sprintf("%v%v", name, condition)
			}
		}
		res = strings.TrimSuffix(res, condition)
	}
	return res
}
