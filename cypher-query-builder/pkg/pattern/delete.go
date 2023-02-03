package pattern

import (
	"errors"
	"fmt"
)

type RemoveConfig struct {
	Name   string
	Field  string
	Labels []string
}

func (r *RemoveConfig) ToString() (string, error) {
	if r.Name == "" {
		return "", errors.New("RemoveConfig - name must be defined")
	}

	if (r.Labels != nil && len(r.Labels) > 0) && r.Field != "" {
		return "", errors.New("RemoveConfig - labels and field cannot both be defined")
	}

	query := r.Name

	if r.Field != "" {
		return query + fmt.Sprintf(".%s", r.Field), nil
	} else {
		for _, label := range r.Labels {
			query += fmt.Sprintf(":%s", label)
		}

		return query, nil
	}
}
