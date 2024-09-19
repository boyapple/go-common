package log

import (
	"sync"
)

const defaultLogName = "default"

var (
	DefaultLogger = NewZapLog(defaultLogger)
	mu            sync.RWMutex
	loggers       = make(map[string]Logger)
)

func init() {
	Register(defaultLogName, DefaultLogger)
}

// Register 日志注册
func Register(name string, l Logger) {
	mu.Lock()
	defer mu.Unlock()
	loggers[name] = l
}

// Logger 日志接口
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Sync() error
}

// GetDefaultLogger 默认日志
func GetDefaultLogger() Logger {
	return Get(defaultLogName)
}

// Get 获取日志接口
func Get(name string) Logger {
	mu.RLock()
	defer mu.RUnlock()
	l := loggers[name]
	return l
}

func Sync() error {
	for _, l := range loggers {
		if err := l.Sync(); err != nil {
			return err
		}
	}
	return nil
}
