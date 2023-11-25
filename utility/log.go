package utility

import (
	"context"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

type LogLevel int

const (
	Info    LogLevel = 0
	Error   LogLevel = 8
	Warning LogLevel = 4
	Debug   LogLevel = -4
)

type CustomLogger struct {
	logging.Logger
	logger *log.Logger
}

func InterceptorLogger() *CustomLogger {
	return &CustomLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (c *CustomLogger) Log(ctx context.Context, level logging.Level, msg string, fields ...any) {
	var colorFunc func(format string, a ...interface{}) string
	switch level {
	case logging.Level(Info):
		colorFunc = color.New(color.FgGreen).SprintfFunc()
	case logging.Level(Error):
		colorFunc = color.New(color.FgRed).SprintfFunc()
	case logging.Level(Warning):
		colorFunc = color.New(color.FgYellow).SprintfFunc()
	case logging.Level(Debug):
		colorFunc = color.New(color.FgBlue).SprintfFunc()
	}

	coloredMessage := colorFunc(msg)
	c.logger.Println(coloredMessage)
}

func (c *CustomLogger) Info(message string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c.Log(ctx, logging.Level(Info), message)
}

func (c *CustomLogger) Error(message string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c.Log(ctx, logging.Level(Error), message)
}

func (c *CustomLogger) Warning(message string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c.Log(ctx, logging.Level(Warning), message)
}

func (c *CustomLogger) Debug(message string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c.Log(ctx, logging.Level(Debug), message)
}
