package log

import (
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelToZapLevel = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"error": zapcore.ErrorLevel,
}

// NewZapLog zap log
func NewZapLog(c []OutputConfig) Logger {
	var cores []zapcore.Core
	callerSkip := 2
	for _, o := range c {
		if o.CallerSkip > 0 {
			callerSkip = o.CallerSkip
		}
		cores = append(cores, newCore(&o))
	}
	return &zapLog{
		l: zap.New(
			zapcore.NewTee(cores...),
			zap.AddCallerSkip(callerSkip),
			zap.AddCaller(),
		),
	}
}

func newCore(c *OutputConfig) zapcore.Core {
	level, ok := levelToZapLevel[c.Level]
	if !ok {
		level = zapcore.DebugLevel
	}
	switch c.Writer {
	case "file":
		writeSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   filepath.Join(c.WriterConfig.LogPath, c.WriterConfig.Filename),
			MaxSize:    c.WriterConfig.MaxSize,
			MaxAge:     c.WriterConfig.MaxAge,
			MaxBackups: c.WriterConfig.MaxBackups,
			Compress:   c.WriterConfig.Compress,
		})
		return zapcore.NewCore(newEncoder(), zapcore.AddSync(writeSyncer), level)
	case "console":
		return zapcore.NewCore(newEncoder(), zapcore.AddSync(os.Stdout), level)
	default:
		return zapcore.NewCore(newEncoder(), zapcore.AddSync(os.Stdout), level)
	}
}

func newEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeTime: func(t time.Time, encode zapcore.PrimitiveArrayEncoder) {
				encode.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	)
}

type zapLog struct {
	l *zap.Logger
}

func (z *zapLog) Debug(args ...interface{}) {
	z.l.Sugar().Debug(args...)
}

func (z *zapLog) Debugf(format string, args ...interface{}) {
	z.l.Sugar().Debugf(format, args...)
}

func (z *zapLog) Info(args ...interface{}) {
	z.l.Sugar().Info(args...)
}

func (z *zapLog) Infof(format string, args ...interface{}) {
	z.l.Sugar().Infof(format, args...)
}

func (z *zapLog) Error(args ...interface{}) {
	z.l.Sugar().Error(args...)
}

func (z *zapLog) Errorf(format string, args ...interface{}) {
	z.l.Sugar().Errorf(format, args...)
}

func (z *zapLog) Sync() error {
	return z.l.Sugar().Sync()
}
