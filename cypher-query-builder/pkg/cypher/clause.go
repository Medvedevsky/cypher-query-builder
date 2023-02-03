package cypher

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"test/neo4j/pkg/pattern"
)

type QueryBuilder struct {
	query  string
	errors []error
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

// подобие полиморфизма подтипов
func (qb *QueryBuilder) mapConfigToString(clauses ...pattern.QueryConfig) string {
	query := ""

	for _, clause := range clauses {
		res, error := clause.ToString()

		if error != nil {
			qb.addError(error)
		}
		query += res
	}

	return query
}

func (qb *QueryBuilder) queryPatternMap(pattern pattern.QueryPattern) string {
	query := ""

	if reflect.ValueOf(pattern).IsZero() {
		qb.addError(errors.New("error match QueryPattern null"))
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

func (qb *QueryBuilder) queryPatternUsage(clauses string, patterns ...pattern.QueryPattern) string {
	if len(patterns) == 0 {
		error := fmt.Sprintf("error %s patterns null", clauses)
		qb.addError(errors.New(error))
		return ""
	}
	query := clauses + " "
	for _, pattern := range patterns {
		query += qb.queryPatternMap(pattern)
	}
	query += "\n"

	return query
}

func (qb *QueryBuilder) Match(patterns ...pattern.QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("MATCH", patterns...)
	return qb
}

func (qb *QueryBuilder) OptionlMath(patterns ...pattern.QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("OPTIONAL MATH", patterns...)
	return qb
}

func (qb *QueryBuilder) Merge(patterns ...pattern.QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("MERGE", patterns...)
	return qb
}

func (qb *QueryBuilder) Create(patterns ...pattern.QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("CREATE", patterns...)
	return qb
}

func (qb *QueryBuilder) Delete(detchDelete bool, deleteClause pattern.RemoveConfig) *QueryBuilder {
	if reflect.ValueOf(deleteClause).IsZero() {
		qb.addError(errors.New("error empty Delete clause"))
		return qb
	}

	if detchDelete {
		qb.query += "DETACH DELETE "
	} else {
		qb.query += "DELETE "
	}

	res := qb.mapConfigToString(&deleteClause)
	qb.query += res

	qb.query += "\n"

	return qb
}

func (qb *QueryBuilder) Where(whereClauses ...pattern.ConditionalConfig) *QueryBuilder {
	if len(whereClauses) == 0 {
		qb.addError(errors.New("error empty where clause"))
		return qb
	}

	qb.query += "WHERE "
	for _, clause := range whereClauses {
		res := qb.mapConfigToString(&clause)
		qb.query += res
	}
	qb.query += "\n"

	return qb
}

func (qb *QueryBuilder) Return(returnClauses ...pattern.ReturnConfig) *QueryBuilder {
	if len(returnClauses) == 0 {
		qb.addError(errors.New("error empty where clause"))
		return qb
	}

	query := "RETRUN "
	for _, clause := range returnClauses {
		res := qb.mapConfigToString(&clause)
		query += res
		query += ", "
	}
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query

	return qb
}

func (qb *QueryBuilder) Remove(removeClauses pattern.RemoveConfig) *QueryBuilder {
	if reflect.ValueOf(removeClauses).IsZero() {
		qb.addError(errors.New("error empty where clause"))
		return qb
	}

	query := "REMOVE "
	qb.mapConfigToString(&removeClauses)
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query

	return qb
}

func (qb *QueryBuilder) Union(all bool) *QueryBuilder {
	if all {
		qb.query += "UNION ALL\n"
		return qb
	}

	qb.query += "UNION\n"
	return qb
}

func (qb *QueryBuilder) With(withClauses ...pattern.WithConfig) *QueryBuilder {
	if len(withClauses) == 0 {
		qb.addError(errors.New("error empty WITH clause"))
		return qb
	}

	query := "WITH "
	for _, clause := range withClauses {
		res := qb.mapConfigToString(&clause)
		query += res
		query += ", "
	}
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query

	return qb
}

func (qb *QueryBuilder) OrderBy(orderByClause pattern.OrderByConfig) *QueryBuilder {
	if reflect.ValueOf(orderByClause).IsZero() {
		qb.addError(errors.New("error empty OrderBy clause"))
		return qb
	}

	qb.query += "ORDER BY "
	res := qb.mapConfigToString(&orderByClause)
	qb.query += res
	qb.query += "\n"

	return qb
}

func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.query += "LIMIT " + strconv.Itoa(limit) + "\n"

	return qb
}

// CALL {subquery}
func (qb *QueryBuilder) CALL(nqb *QueryBuilder) *QueryBuilder {
	res := "CALL {\n"
	subquery, error := nqb.Execute()
	if error != nil {
		qb.addError(error)
	}

	var buffer bytes.Buffer

	for i, rune := range subquery {
		buffer.WriteRune(rune)
		char := string(rune)

		if char == "\n" {
			if i != len(subquery)-1 {
				buffer.WriteRune('\t')
			}
		}
	}

	subquery = buffer.String()
	res += "\t" + subquery + "}\n"
	qb.query += res

	return qb
}

func (qb *QueryBuilder) Execute() (string, error) {
	return qb.query, qb.errorBuild()
}

func (q *QueryBuilder) addError(e error) {
	if q.errors == nil {
		q.errors = []error{e}
	} else {
		q.errors = append(q.errors, e)
	}
}

func (qb *QueryBuilder) errorBuild() error {
	if len(qb.errors) > 0 {
		str := "errors found: "
		for _, err := range qb.errors {
			str += err.Error() + ";"
		}

		str = strings.TrimSuffix(str, ";") + fmt.Sprintf(" -- total errors (%v)", len(qb.errors))
		return errors.New(str)
	}

	return nil
}
