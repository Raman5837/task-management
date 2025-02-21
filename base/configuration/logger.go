package configuration

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/Raman5837/task-management/base/constants"
	"github.com/rs/zerolog"
)

// Get log level from .env if present, else use `INFO` as the log level
func __get_level() zerolog.Level {

	var err error
	var logLevel zerolog.Level

	level := constants.GetEnv("LOG_LEVEL")

	if level == "" {
		return zerolog.InfoLevel
	}
	if logLevel, err = zerolog.ParseLevel(level); err != nil {
		return zerolog.InfoLevel
	}
	return logLevel
}

var Logger *NewLogger
var singleton sync.Once

// Logger struct to have some extra functionality on-top of `zerolog.Logger`
type NewLogger struct {
	Logger zerolog.Logger
}

// Returns NewLogger instance
func newLoggerInstance() *NewLogger {

	level := __get_level()
	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Int("PID", os.Getpid()).Logger()
	return &NewLogger{Logger: logger}

}

// Returns singleton instance of `NewLogger`
func GetLogger() *NewLogger {
	singleton.Do(func() {
		Logger = newLoggerInstance()
	})
	return Logger
}

func getCallerInfo() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", file, line)
}

func (instance *NewLogger) Debug(message string, values ...any) {
	callerInfo := getCallerInfo()
	instance.Logger.Debug().Str("caller", callerInfo).Msg(fmt.Sprintf(message, values...))
}

func (instance *NewLogger) Info(message string, values ...any) {
	callerInfo := getCallerInfo()
	instance.Logger.Info().Str("caller", callerInfo).Msg(fmt.Sprintf(message, values...))
}

func (instance *NewLogger) Warn(message string, values ...any) {
	callerInfo := getCallerInfo()
	instance.Logger.Warn().Str("caller", callerInfo).Msg(fmt.Sprintf(message, values...))
}

func (instance *NewLogger) Error(exception error, message string, values ...any) {
	callerInfo := getCallerInfo()
	instance.Logger.Error().Str("caller", callerInfo).Err(exception).Msg(fmt.Sprintf(message, values...))
}

func (instance *NewLogger) Fatal(exception error, message string, values ...any) {
	callerInfo := getCallerInfo()
	instance.Logger.Fatal().Str("caller", callerInfo).Err(exception).Msg(fmt.Sprintf(message, values...))
}
