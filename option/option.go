package option

type Option[T any] interface {
	apply(T)
}

type option[T any] struct {
	f func(T)
}

func New[T any](f func(T)) Option[T] {
	return &option[T]{
		f: f,
	}
}

func (o *option[T]) apply(t T) {
	o.f(t)
}
