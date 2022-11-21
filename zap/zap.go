// Package zap implements a mapper from 	"go.uber.org/zap"
package zap

import (
	"go.uber.org/zap"

	log "github.com/adamhassel/logger-abstract"
)

// Logger is an Contextual logger wrapper over Logrus's logger.
type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(z *zap.SugaredLogger) log.Contextual {
	var l Logger
	l.SugaredLogger = z
	l.Info("Using mapped zap logger")
	return &l
}

// With wraps the zap sugared logger's With-method, which returns an explicit
// *zap.Sugared, which can't be abstracted without a wrapper :(
func (l Logger) With(fields ...interface{}) log.Leveled {
	return l.SugaredLogger.With(fields)
}
