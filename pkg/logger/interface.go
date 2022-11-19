package logger

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	DebugJSON(e interface{})
	InfoJSON(e interface{})
	ErrorJSON(e interface{})
	WarningJSON(e interface{})
	FatalJSON(e interface{})
}
