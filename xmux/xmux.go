package xmux

import (
	"fmt"
	"sync"
)

type Mux[K comparable, V any] interface {
	Register(key K, val V)
	Get(key K) (V, error)
}

type mux[K comparable, V any] struct {
	lock    sync.RWMutex
	options *Options
	m       map[K]V
}

func New[K comparable, V any](opts ...Option) Mux[K, V] {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	return &mux[K, V]{
		m:       make(map[K]V),
		options: opt,
	}
}

func (mux *mux[K, V]) Get(key K) (V, error) {
	mux.lock.RLock()
	defer mux.lock.RUnlock()
	if v, ok := mux.m[key]; ok {
		return v, nil
	}
	var v V
	return v, fmt.Errorf("mux key[%s] not found", key)
}

func (mux *mux[K, V]) Register(key K, val V) {
	mux.lock.Lock()
	defer mux.lock.Unlock()
	if _, ok := mux.m[key]; ok && !mux.options.AllowsDuplicate {
		panic(fmt.Sprintf("mux key[%v] already regiater", key))
	}
	mux.m[key] = val
}
