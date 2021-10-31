package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _ log.Logger = (*ZapLogger)(nil)

// Zap 结构体
type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

//
func NewZapLogger() *ZapLogger {
	logger := initZapLogger(
		// Level
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		// Options
		// 将记录器配置为记录给定的级别及以上的所有信息的堆栈追踪
		zap.AddStacktrace(zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		// 跳过的callers数量
		zap.AddCallerSkip(2),
		// 设置为开发模式，DPanic级别的日志会panic而不是简单的记录错误
		zap.Development(),
	)
	return logger
}

// 创建一个 ZapLogger 实例
func initZapLogger(level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	// zapcore的编码器配置
	encoder := zapcore.EncoderConfig{
		LevelKey:      "level", //结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:       "ts",    //结构化（json）输出：时间的key（INFO，WARN，ERROR等）
		NameKey:       "logger",
		CallerKey:     "caller", //结构化（json）输出：打印日志的文件对应的Key
		MessageKey:    "msg",    //结构化（json）输出：msg的key
		StacktraceKey: "stack",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		LineEnding:     zapcore.DefaultLineEnding,      // 写入日志的默认行结束
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志等级转换成大写（INFO
		EncodeDuration: zapcore.SecondsDurationEncoder, // 秒
		EncodeCaller:   zapcore.ShortCallerEncoder,     //采用短文件路径编码输出（test/main.go:14 ）
	}

	// io.Writer
	writeSyncer := getLogWriter()

	// 设置 zapcore
	core := zapcore.NewCore(
		// 编码器配置
		zapcore.NewConsoleEncoder(encoder),
		// 写入缓冲区
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			writeSyncer,
		),
		// 日志级别
		level,
	)
	// // 实现多个输出
	// core := zapcore.NewTee(
	// 	zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(infoWriter), infoLevel),                         //将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
	// 	zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(warnWriter), warnLevel),                         //warn及以上写入errPath
	// 	zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel), //同时将日志输出到控制台，NewJSONEncoder 是结构化输出
	// )

	//  使用设置好的 core 和 option 构建一个新的记录器
	zapLogger := zap.New(core, opts...)

	return &ZapLogger{
		log:  zapLogger,
		Sync: zapLogger.Sync,
	}
}

// Log 方法实现了 kratos中的 Logger interface
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {

	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	// 按照 KV 传入的时候,使用的 zap.Field
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}

	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}

// 日志自动切割，采用 lumberjack 实现的
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./Comet.log",
		MaxSize:    10,    // 最大M数，超过则切割
		MaxBackups: 5,     // 最大文件保留数，超过则删除老的日志文件
		MaxAge:     30,    // 最大的保存时间
		Compress:   false, // 是否压缩日志文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
