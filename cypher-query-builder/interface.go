package cypher

type QueryConfig interface {
	ToString() (string, error)
}
