package convert

import (
	"fmt"
	"sync"
)

var converts = make(map[uint64]interface{})
var mu sync.RWMutex

type Converter[From, To any] interface {
	Convert(From) To
	Reverse(To) From
}

func Register[From, To any](i uint64, h Converter[From, To]) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := converts[i]; exists {
		panic(fmt.Sprintf("converter already register i:%d", i))
	}
	converts[i] = h
}

func Get[From, To any](i uint64) (Converter[From, To], error) {
	mu.RLock()
	v, exists := converts[i]
	if !exists {
		return nil, fmt.Errorf("converter not found key:%s", i)
	}
	mu.RUnlock()
	h, ok := v.(Converter[From, To])
	if !ok {
		return nil, fmt.Errorf("converter type mismatch key:%d got:%T", i, v)
	}
	return h, nil
}
