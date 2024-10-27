package applogger

import (
	"time"

	"github.com/samber/lo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func encodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}

func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02T15:04:05.000000+09:00"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}

var LogTypeKeyName = "logType"

func Init() {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.DisableStacktrace = true

	var err error
	Log, err = zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic("Error zap init.")
	}

	defer Log.Sync()
}

func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, addAppLogType(fields...)...)
}

func Info(message string, fields ...zap.Field) {
	Log.Info(message, addAppLogType(fields...)...)
}

func Warn(message string, fields ...zap.Field) {
	Log.Warn(message, addAppLogType(fields...)...)
}

func Error(message string, fields ...zap.Field) {
	Log.Error(message, addAppLogType(fields...)...)
}

func addAppLogType(fields ...zap.Field) []zap.Field {
	_, found := lo.Find(fields, func(field zap.Field) bool {
		return field.Key == LogTypeKeyName
	})

	if found {
		return fields
	} else {
		return append(fields, zap.String(LogTypeKeyName, "app-log"))
	}
}
