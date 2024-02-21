package log

// Logger is the basic interface that logs a message with a level. This is separate, since many loggers don't actually implement this method.
type Logger interface {
	Log(l Level, args ...interface{})
}

// Standard is the interface used by Go's standard library's log package, except
// the Print functions, since those are often omitted by loggers. For a version
// that includes that, see thr "Extended" versions.
type Standard interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
}

// StandardExtended is like Standard, but also with print functions. The stdlib logger implements StandardExtended.
type StandardExtended interface {
	Standard

	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

// Leleved is an interface with commonly used log level methods.
type Leveled interface {
	Standard

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
}

// Contextual is an interface that allows context addition to a log statement before
// calling the final print (message/level) method.
type Contextual interface {
	Leveled

	With(fields ...interface{}) Leveled
}

// LeveledExtended extends Leveled with Print-methods
type LeveledExtended interface {
	StandardExtended
	Leveled
}

// ContextualExtended extends Contextual with Print-methods
type ContextualExtended interface {
	LeveledExtended

	With(fields ...interface{}) LeveledExtended
}

type ContextualLogger interface {
	Logger
	Contextual
}

type ContextualExtendedLogger interface {
	Logger
	ContextualExtended
}
