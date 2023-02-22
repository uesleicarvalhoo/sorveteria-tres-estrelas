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

func (l Logrus) DebugJSON(e map[string]interface{}) {
	l.logger.WithFields(e).Debug()
}

func (l Logrus) InfoJSON(e map[string]interface{}) {
	l.logger.WithFields(e).Info()
}

func (l Logrus) ErrorJSON(e map[string]interface{}) {
	l.logger.WithFields(e).Error()
}

func (l Logrus) WarningJSON(e map[string]interface{}) {
	l.logger.WithFields(e).Warning(e)
}

func (l Logrus) FatalJSON(e map[string]interface{}) {
	l.logger.WithFields(e).Fatal(e)
}
