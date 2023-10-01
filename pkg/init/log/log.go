package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	return logger
}

var sugarLogger *zap.SugaredLogger

func GetSugarLogger() *zap.SugaredLogger {
	return sugarLogger
}

// 初始log
func InitLogger(servicename string) {
	encoder := getEncoder()
	core1 := zapcore.NewCore(encoder, getLogWriter(servicename), zapcore.DebugLevel)
	core2 := zapcore.NewCore(encoder, getLogErrWriter(servicename), zapcore.ErrorLevel)

	core := zapcore.NewTee(core1, core2)
	logger = zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(servicename string) zapcore.WriteSyncer {
	file, _ := os.Create(fmt.Sprintf("D:/studysystem_micro/internal/service/%s/logs/study_system.log", servicename))
	return zapcore.AddSync(file)
}
func getLogErrWriter(servicename string) zapcore.WriteSyncer {
	file, _ := os.Create(fmt.Sprintf("D:/studysystem_micro/internal/service/%s/logs/study_system.err.log", servicename))
	return zapcore.AddSync(file)
}
