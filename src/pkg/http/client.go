package http

type ClientInterface[T any] interface {
	Get(args ...string) (T, error)
}
