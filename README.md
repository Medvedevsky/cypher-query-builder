# Cypher Query Builder

Cypher Query Builder for Neo4j

## Example usage

``` go
pNode := pattern.NewNode().SetVariable("p").SetLabel("Person").AsPattern()

callCypher, err := cypher.NewQueryBuilder().
	Call(cypher.NewQueryBuilder().
		Match(pNode).
		Return(pattern.ReturnConfig{Name: "p"}).
		OrderBy(pattern.OrderByConfig{Name: "p", Member: "age", Asc: true}).
		Limit(1).
	Union(false).
		Match(pNode).
		Return(pattern.ReturnConfig{Name: "p"}).
		OrderBy(pattern.OrderByConfig{Name: "p", Member: "age", Desc: true}).
		Limit(1)).
	Return(pattern.ReturnConfig{Name: "p", Type: "name"}, pattern.ReturnConfig{Name: "p", Type: "age"}).
	OrderBy(pattern.OrderByConfig{Name: "p", Member: "name"}).
	Execute()
```

```
CALL {
  MATCH (p:Person)
  RETURN p
  ORDER BY p.age ASC
  LIMIT 1
UNION
  MATCH (p:Person)
  RETURN p
  ORDER BY p.age DESC
  LIMIT 1
}
RETURN p.name, p.age
ORDER BY p.name
```

## Implemented Query Clauses
    + Match
    + Optional Match
    + Merge
    + Where
    + With
    + Return
    + OrderBy
    + Limit
    + Create
    + Delete
    + Remove
    + Union
    + Call {subquery}
