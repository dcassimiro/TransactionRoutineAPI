package logger

import (
	"context"

	logrus "github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

// SetLevel change the logger level
func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)
}

// GetLevel recovers the logger level
func GetLevel() logrus.Level {
	return logrus.GetLevel()
}

// Error displays error details
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// ErrorContext displays error details with context
func ErrorContext(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Error(args...)
}

// Info displays log info details
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// InfoContext displays log info details with context
func InfoContext(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Info(args...)
}

// Debug displays debug log details
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// DebugContext displays debug log details with context
func DebugContext(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Debug(args...)
}

// Trace displays log trace details
func Trace(args ...interface{}) {
	logrus.Trace(args...)
}

// TraceContext displays log trace details with context
func TraceContext(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Trace(args...)
}

// Fatal displays error details
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
