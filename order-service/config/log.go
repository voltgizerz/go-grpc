package config

import (
	logrus "github.com/sirupsen/logrus"
	nested "github.com/antonfisher/nested-logrus-formatter"
)

// SetupLog - return logrus.
func SetupLog() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	return log
}
