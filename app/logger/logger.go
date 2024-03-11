package logger

const (
	_ int = iota
	LevelError
	LevelInfo
	LevelDebug
)

type Log interface {
	Info(format string, a ...any)  // log level is 2
	Error(format string, a ...any) // log level is 1
	Debug(format string, a ...any) // log level is 3

	// system will only print log which level is less than or equal to (<=) applied level
	SetLogLevel(level int) // default will be 3
	LogLevel() int

	Close()
}
