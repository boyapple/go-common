package convert

import (
	"fmt"
	"sync"
)

var converts = make(map[string]interface{})
var mu sync.RWMutex

type Converter[From, To any] interface {
	Convert(From) To
	Reverse(To) From
}

func Register[From, To any](key string, h Converter[From, To]) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := converts[key]; exists {
		panic(fmt.Sprintf("converter already register key:%s", key))
	}
	converts[key] = h
}

func Get[From, To any](key string) (Converter[From, To], error) {
	mu.RLock()
	v, exists := converts[key]
	if !exists {
		return nil, fmt.Errorf("converter not found key:%s", key)
	}
	mu.RUnlock()
	h, ok := v.(Converter[From, To])
	if !ok {
		return nil, fmt.Errorf("converter type mismatch key:%s got:%T", key, v)
	}
	return h, nil
}
