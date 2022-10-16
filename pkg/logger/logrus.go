package logger

import "github.com/sirupsen/logrus"

func NewLogrus(level string) (*logrus.Logger, error) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetLevel(lvl)

	return logger, nil
}
