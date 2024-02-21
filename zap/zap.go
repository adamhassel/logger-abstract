// Package zap implements a mapper from 	"go.uber.org/zap"
package zap

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	log "github.com/adamhassel/logger-abstract"
)

type Kind int

const (
	DEV Kind = iota
	PROD
	CONSOLE
)

// Logger is an Contextual logger wrapper over Logrus's logger.
type Logger struct {
	*zap.SugaredLogger
}

func NewDevLogger() log.ContextualExtendedLogger {
	return New(DEV)
}

func NewProdLogger() log.ContextualExtendedLogger {
	return New(PROD)
}

func New(kind Kind) log.ContextualExtendedLogger {
	var l Logger
	var z *zap.Logger
	var err error

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:          "msg",
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "logger",
		CallerKey:           "caller",
		FunctionKey:         "fn",
		StacktraceKey:       "stack",
		SkipLineEnding:      false,
		LineEnding:          "",
		EncodeLevel:         zapcore.CapitalColorLevelEncoder,
		EncodeTime:          zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.00000"),
		EncodeDuration:      zapcore.StringDurationEncoder,
		EncodeCaller:        nil,
		EncodeName:          nil,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    " ",
	}

	switch kind {
	case DEV:
		z, err = zap.NewDevelopment()
	case PROD:
		z, err = zap.NewProduction()
	case CONSOLE:
		z = zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel))
	}
	if err != nil {
		panic(err)
	}

	l.SugaredLogger = z.Sugar()
	l.Info("using zap logger")
	return &l
}

// With wraps the zap sugared logger's With-method, which returns an explicit
// *zap.Sugared, which can't be abstracted without a wrapper :(
func (l Logger) With(fields ...interface{}) log.LeveledExtended {
	return Logger{l.SugaredLogger.With(fields...)}
}

func (l Logger) Log(lvl log.Level, args ...interface{}) {
	switch lvl {
	case log.DEBUG:
		l.Debug(args...)
	case log.INFO:
		l.Info(args...)
	case log.WARN:
		l.Warn(args...)
	case log.ERROR:
		l.Error(args...)
	case log.PANIC:
		l.Panic(args...)
	case log.FATAL:
		l.Fatal(args...)
	}
}

func (l Logger) Print(args ...interface{}) {
	l.Info(args...)
}

func (l Logger) Println(args ...interface{}) {
	l.Infoln(args...)
}
func (l Logger) Printf(fmt string, args ...interface{}) {
	l.Infof(fmt, args...)
}

func (l Logger) Zap() *zap.SugaredLogger {
	return l.SugaredLogger
}

func (l Logger) Output() io.Writer {
	return l
}

func (l Logger) Write(b []byte) (int, error) {
	l.Print(string(b))
	return len(b), nil
}

func (l Logger) SetOutput(io.Writer) {}
