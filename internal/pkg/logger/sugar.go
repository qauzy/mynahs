package logger

var (
	prefix = ""
)

func Info(args ...interface{}) {
	args = append([]interface{}{prefix}, args...)
	Logger.Sugar().Info(args...)
}

func Infof(template string, args ...interface{}) {
	Logger.Sugar().Infof(prefix+template, args...)
}

func Warn(args ...interface{}) {
	args = append([]interface{}{prefix}, args...)
	Logger.Sugar().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Logger.Sugar().Warnf(prefix+template, args...)
}

func Error(args ...interface{}) {
	args = append([]interface{}{prefix}, args...)
	Logger.Sugar().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Logger.Sugar().Errorf(prefix+template, args...)
}
