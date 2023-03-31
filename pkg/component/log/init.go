package log

import (
	"fmt"
	log "go.uber.org/zap"
	logcore "go.uber.org/zap/zapcore"
)

type (
	Field         = logcore.Field
	Logger        = log.Logger
	Option        = log.Option
	SugaredLogger = log.SugaredLogger
)

var logger *log.Logger

func InitLogger() {
	logger = GetLogger([]string{"./logs/info.log"}, []string{"./logs/error.log"})
	defer func() {
		if err := Sync(); err != nil {
			fmt.Println(err.Error())
		}
	}()
}
func GetLogger(outputPath, errorPath []string) *log.Logger {
	encoderConfig := logcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     logcore.DefaultLineEnding,
		EncodeLevel:    logcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     logcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: logcore.SecondsDurationEncoder,
		EncodeCaller:   logcore.FullCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	atom := log.NewAtomicLevelAt(log.DebugLevel)

	config := log.Config{
		Level:             atom,
		Development:       true,
		DisableStacktrace: false,
		DisableCaller:     false,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		InitialFields:     map[string]interface{}{"service": "notice"},
		OutputPaths:       outputPath,
		ErrorOutputPaths:  errorPath,
	}

	return log.Must(config.Build())
}
