package http

type ClientInterface[T any, L any] interface {
	Resources() ([]string, error)
	All(args ...string) ([]L, error, bool)
	Get(args ...string) (T, error, bool)
}
