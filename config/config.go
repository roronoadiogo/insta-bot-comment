package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	loggerLevelFile    = zap.ErrorLevel
	loggerLevelConsole = zap.InfoLevel
)

func InitializeLogger() *zap.Logger {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	log, _ := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(log)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, loggerLevelFile),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), loggerLevelConsole),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func LoadEnv(logger *zap.Logger) error {

	logger.Info("Getting the Enviroment definitions")

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load(fmt.Sprintf("config/.env.%s", env))

	if err != nil {
		logger.Error("Error in env params definitions, application should be stopped", zapcore.Field{Key: "error: ", Interface: err})
		panic(err)
	}

	return nil
}
