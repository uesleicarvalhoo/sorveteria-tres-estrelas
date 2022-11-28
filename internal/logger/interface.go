package logger

// TODO: Mover para /internal/logger
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	DebugJSON(e map[string]interface{})
	InfoJSON(e map[string]interface{})
	ErrorJSON(e map[string]interface{})
	WarningJSON(e map[string]interface{})
	FatalJSON(e map[string]interface{})
}
