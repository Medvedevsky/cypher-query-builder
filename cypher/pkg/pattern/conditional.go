package pattern

import (
	"errors"
	"fmt"
)

type ConditionalConfig struct {
	ConditionOperator BooleanOperator
	Condition         Condition

	ConditionFunction string
	Name              string
	Field             string
	Label             string
	Check             interface{}
}

func (condition *ConditionalConfig) ToString() (string, error) {
	if condition.Name == "" {
		return "", errors.New("var name can not be empty")
	}

	query := ""

	//build the fields
	if condition.Field != "" {
		query += fmt.Sprintf("%s.%s", condition.Name, condition.Field)
	} else if condition.Label != "" {
		//we're done here
		return fmt.Sprintf("%s:%s", condition.Name, condition.Label), nil
	} else {
		query = condition.Name
	}

	//build the operators
	if condition.ConditionOperator != "" {
		query += fmt.Sprintf(" %s", condition.ConditionOperator)
	} else if condition.ConditionFunction != "" {
		//if its a condition function, we're done
		return fmt.Sprintf("%s(%s)", condition.ConditionFunction, query), nil
	}
	query += fmt.Sprintf(" %v", condition.Check)

	if condition.Condition != "" {
		query += fmt.Sprintf(" %s ", condition.Condition)
	}

	return query, nil
}
