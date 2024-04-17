package logger

/* reference: https://github.com/uber-go/zap/blob/v1.27.0/zapcore/level.go#L34 */
const (
	DebugLevel int8 = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

type Log interface {
	Info(format string, a ...any)  // log level is 2
	Error(format string, a ...any) // log level is 1
	Debug(format string, a ...any) // log level is 3
	Fatal(format string, a ...any) // log level is 4

	// system will only print log which level is less than or equal to (<=) applied level
	SetLogLevel(level int8) // default will be 3
	LogLevel() int8

	Close()
}
