package logger

import (
	"log"
)

type CustomLogger struct{}

func NewCustomLogger() (*CustomLogger, error) {
	return &CustomLogger{}, nil
}

func (l *CustomLogger) Debug(message string, fields map[string]interface{}) {
	log.Printf("[DEBUG] %s: %v\n", message, fields)
}

func (l *CustomLogger) Info(message string, fields map[string]interface{}) {
	log.Printf("[INFO] %s: %v\n", message, fields)
}

func (l *CustomLogger) Warn(message string, fields map[string]interface{}) {
	log.Printf("[WARN] %s: %v\n", message, fields)
}

func (l *CustomLogger) Error(message string, fields map[string]interface{}) {
	log.Printf("[ERROR] %s: %v\n", message, fields)
}
