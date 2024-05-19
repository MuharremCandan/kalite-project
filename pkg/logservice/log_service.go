package logservice

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

type LogService struct {
	logger *logrus.Logger
}

func NewLogService() *LogService {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{} // Use JSON format for logs
	logHook, err := logrustash.New(
		&logrustash.LogrusFields{
			Type:     "appName",       // Type can be any string that helps identify your logs in Logstash
			Fields:   logrus.Fields{}, // You can add default fields here if needed
			Host:     "logstash",      // Logstash host address
			Port:     5044,            // Logstash port
			Protocol: "tcp",           // Protocol can be "tcp" or "udp" based on your Logstash configuration
		})
	if err != nil {
		logger.Fatalf("Failed to create Logstash hook: %v", err)
	}
	logger.Hooks.Add(logHook)
	logger.Hooks.Add(logHook)

	return &LogService{logger: logger}
}

func (ls *LogService) Info(message string, fields logrus.Fields) {
	ls.logger.WithFields(fields).Info(message)
}

func (ls *LogService) Warn(message string, fields logrus.Fields) {
	ls.logger.WithFields(fields).Warn(message)
}

func (ls *LogService) Error(message string, fields logrus.Fields) {
	ls.logger.WithFields(fields).Error(message)
}

// You can define more logging functions as needed (e.g., Debug, Fatal, etc.)
