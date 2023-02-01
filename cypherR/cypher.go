package cypherr

import (
	"fmt"
	"reflect"
)

// QueryBuilder ...
type QueryBuilder struct {
	query  string
	errors []error
}

// NewQueryBuilder ...
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) queryPatternMap(pattern QueryPattern) string {
	query := ""

	if reflect.ValueOf(pattern).IsZero() {
		// error
		qb.addError(fmt.Errorf("error match QueryPattern null"))
		return ""
	}

	if !reflect.ValueOf(pattern.OnlyNode).IsZero() {
		query += pattern.OnlyNode.Node.ToCypher()

		return query
	}

	if !reflect.ValueOf(pattern.PartialRelationship).IsZero() {
		p := pattern.PartialRelationship

		query += pattern.Edge.PartialRelationshipBuild(p)
		return query
	}

	if !reflect.ValueOf(pattern.FullRelationship).IsZero() {
		f := pattern.FullRelationship
		query += pattern.Edge.RelationshipBuild(f)

		return query
	}

	return ""
}

func (qb *QueryBuilder) queryPatternUsage(clauses string, patterns ...QueryPattern) string {
	if len(patterns) == 0 {
		// error
		error := fmt.Sprintf("error %s patterns null", clauses)
		qb.addError(fmt.Errorf(error))
		return ""
	}
	query := clauses + " "
	for _, pattern := range patterns {
		query += qb.queryPatternMap(pattern)
	}
	query += "\n"

	return query
}

func (qb QueryBuilder) Match(patterns ...QueryPattern) QueryBuilder {
	qb.query += qb.queryPatternUsage("MATCH", patterns...)
	return qb
}

func (qb QueryBuilder) OptionlMath(patterns ...QueryPattern) QueryBuilder {
	qb.query += qb.queryPatternUsage("OPTIONAL MATH", patterns...)
	return qb
}

func (qb QueryBuilder) Where(whereClauses ...WhereQuery) QueryBuilder {
	if len(whereClauses) == 0 {
		qb.addError(fmt.Errorf("error empty where clause"))
		return qb
	}

	query := qb.query + "WHERE "
	query += whereMap(whereClauses...)
	query += "\n"

	return QueryBuilder{query: query}
}

// With ...
// func (qb QueryBuilder) With(withClauses ...string) QueryBuilder {
// 	return QueryBuilder{
// 		qb.query + `
// 		WITH
// 			` + strings.Join(withClauses, ", "),
// 	}
// }

// Return ...
func (qb QueryBuilder) Return(returnClauses ...ReturnQuery) QueryBuilder {

	if len(returnClauses) == 0 {
		qb.addError(fmt.Errorf("error empty where clause"))
		return qb
	}
	query := qb.query + "RETRUN "
	query += returnMap(returnClauses...)
	query += "\n"

	return QueryBuilder{query: query}
	// return QueryBuilder{
	// 	qb.query + "RETURN " + strings.Join(returnClauses, ", "),
	// }
}

// OrderBy ...
// func (qb QueryBuilder) OrderBy(orderByClause string) QueryBuilder {
// 	return QueryBuilder{
// 		qb.query + `
// 		ORDER BY
// 		` + orderByClause,
// 	}
// }

// OrderByDesc ...
// func (qb QueryBuilder) OrderByDesc(orderByDescClause string) QueryBuilder {
// 	return QueryBuilder{
// 		qb.query + `
// 		ORDER BY
// 			` + orderByDescClause + ` DESC`,
// 	}
// }

// Limit ...
// func (qb QueryBuilder) Limit(limit int) QueryBuilder {
// 	return QueryBuilder{
// 		qb.query + `
// 		LIMIT	` + strconv.Itoa(limit),
// 	}
// }

// Execute ...
func (qb QueryBuilder) Execute() (string, []error) {
	return qb.query, qb.errors
}

// As ...
func As(initial, alias string) string {
	return fmt.Sprintf("%v AS %v", initial, alias)
}

// Assign ...
func Assign(name, pattern string) string {
	return fmt.Sprintf("%v = %v", name, pattern)
}

func whereMap(conditions ...WhereQuery) string {
	query := ""
	for _, condition := range conditions {
		query += fmt.Sprintf("%v.%v %v %v %v ", condition.Name, condition.Field, condition.BooleanOperator, condition.Check, condition.Condition)
	}

	return query
}

func returnMap(conditions ...ReturnQuery) string {
	res := ""
	for i, condition := range conditions {
		if i != len(conditions)-1 {
			if condition.Field != "" {
				res += fmt.Sprintf("%v.%v", condition.Name, condition.Field) + ", "
			} else {
				res += fmt.Sprintf("%v", condition.Name) + ", "
			}
			continue
		}

		if condition.Field != "" {
			res += fmt.Sprintf("%v.%v", condition.Name, condition.Field)
		} else {
			res += fmt.Sprintf("%v", condition.Name)
		}
	}
	return res
}

func (q *QueryBuilder) addError(e error) {
	if q.errors == nil {
		q.errors = []error{e}
	} else {
		q.errors = append(q.errors, e)
	}
}
