package convert

import (
	"fmt"
)

var handlers = make(map[string]interface{})

type Handler[A any, B any] interface {
	Convert(A) (B, error)
	Reverse(B) (A, error)
}

func Register[A any, B any](key string, h Handler[A, B]) {
	if _, exists := handlers[key]; exists {
		panic(fmt.Sprintf("converter already register key:%s", key))
	}
	handlers[key] = h
}

func Get[A any, B any](key string) (Handler[A, B], error) {
	v, exists := handlers[key]
	if !exists {
		return nil, fmt.Errorf("converter not found key:%s", key)
	}
	h, ok := v.(Handler[A, B])
	if !ok {
		return nil, fmt.Errorf("converter type mismatch key:%s got:%T", key, v)
	}
	return h, nil
}
