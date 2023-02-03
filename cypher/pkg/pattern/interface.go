package pattern

type QueryConfig interface {
	ToString() (string, error)
}
