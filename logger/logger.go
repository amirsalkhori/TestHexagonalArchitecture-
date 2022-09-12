package logger

import "go.uber.org/zap"

var log *zap.Logger
func init() {
	config := zap.NewProductionConfig()
	encodConfig := zap.NewProductionEncoderConfig()
	encodConfig.StacktraceKey = ""
	encodConfig.TimeKey = "timeStamp"
	config.EncoderConfig = encodConfig
	
	var err error
	log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field){
	log.Info(message)
}

func Debug(message string, fields ...zap.Field){
	log.Debug(message)
}

func Error(message string, fields ...zap.Field){
	log.Error(message)
}