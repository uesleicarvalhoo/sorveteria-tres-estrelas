package logger

import "github.com/sirupsen/logrus"

func Configure(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})

	logrus.SetLevel(lvl)

	return nil
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Warning(args ...interface{}) {
	logrus.Warning(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func DebugJSON(e map[string]interface{}) {
	logrus.WithFields(e).Debug()
}

func InfoJSON(e map[string]interface{}) {
	logrus.WithFields(e).Info()
}

func ErrorJSON(e map[string]interface{}) {
	logrus.WithFields(e).Error()
}

func WarningJSON(e map[string]interface{}) {
	logrus.WithFields(e).Warning(e)
}

func FatalJSON(e map[string]interface{}) {
	logrus.WithFields(e).Fatal(e)
}
