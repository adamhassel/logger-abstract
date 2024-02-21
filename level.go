package log

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	PANIC
	FATAL
)
