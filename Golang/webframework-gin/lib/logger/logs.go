package logger

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLog(conf string) {
	rawJSON := []byte(conf)
	var cfg zap.Config
	if configErr := json.Unmarshal(rawJSON, &cfg); configErr != nil {
		panic(configErr)
	}
	var buildErr error
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	log, buildErr = cfg.Build(zap.AddCallerSkip(1))
	if buildErr != nil {
		panic(buildErr)
	}
	defer log.Sync()
}

func Info(args ...interface{}) {
	log.Sugar().Info(args)
}
func InfoF(template string, args ...interface{}) {
	log.Sugar().Infof(template, args)
}

func InfoW(template string, paramKey string, paramVal interface{}) {
	log.Sugar().Infow(template, paramKey, paramVal, paramKey, paramVal)
}
func Warn(args ...interface{}) {
	log.Sugar().Warn(args)
}

func WarnF(template string, args ...interface{}) {
	log.Sugar().Warnf(template, args)
}

func Error(args ...interface{}) {
	log.Sugar().Error(args)
}

func ErrorF(template string, args ...interface{}) {
	log.Sugar().Errorf(template, args)
}

func Debug(args ...interface{}) {
	log.Sugar().Debug(args)
}

func DebugF(template string, args ...interface{}) {
	log.Sugar().Debugf(template, args)
}

func DebugW(template string, paramKey string, paramVal interface{}) {
	log.Sugar().Debugw(template, paramKey, paramVal, paramKey, paramVal)
}
