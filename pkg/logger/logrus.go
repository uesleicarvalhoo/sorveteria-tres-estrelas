package logger

import "github.com/sirupsen/logrus"

type Logrus struct {
	logger *logrus.Logger
}

func NewLogrus(level string) (*Logrus, error) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})

	logger.SetLevel(lvl)

	return &Logrus{
		logger: logger,
	}, nil
}

func (l Logrus) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l Logrus) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l Logrus) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l Logrus) Warningf(format string, args ...interface{}) {
	l.logger.Warningf(format, args...)
}

func (l Logrus) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l Logrus) DebugJSON(e interface{}) {
	l.logger.Debug(e)
}

func (l Logrus) InfoJSON(e interface{}) {
	l.logger.Info(e)
}

func (l Logrus) ErrorJSON(e interface{}) {
	l.logger.Error(e)
}

func (l Logrus) WarningJSON(e interface{}) {
	l.logger.Warning(e)
}

func (l Logrus) FatalJSON(e interface{}) {
	l.logger.Fatal(e)
}
