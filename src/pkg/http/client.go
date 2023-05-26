package http

type ClientInterface[T any, L any] interface {
	All(args ...string) ([]L, error, bool)
	Get(args ...string) (T, error, bool)
}
