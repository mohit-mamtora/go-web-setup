package logger

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/mohit-mamtora/go-web-setup/app/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileLogger struct {
	logLevel int8
	logger   *zap.Logger
	rotator  *rotatingFile
}

var _ logger.Log = (*FileLogger)(nil)

func NewFileLogger(filePath, fileName string, maxSizeMB int, logLevel int8, jsonEncoder bool) (*FileLogger, error) {

	rotator := &rotatingFile{
		filename:  filepath.Join(filePath, fileName),
		maxSizeMB: maxSizeMB,
	}
	// Create a file sink
	sink := zapcore.AddSync(rotator)

	// Configure encoder
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if jsonEncoder {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}
	// Create encoder

	// Create a core with the rotating file sink
	core := zapcore.NewCore(encoder, sink, zapcore.Level(logLevel))

	return &FileLogger{
		rotator:  rotator,
		logLevel: logLevel,
		logger:   zap.New(core),
	}, nil
}

func (f *FileLogger) SetLogLevel(logLevel int8) {
	f.logLevel = logLevel
	f.logger.Core().Enabled(zapcore.Level(logLevel))
}

func (f *FileLogger) Info(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	f.logger.Info(msg)
}

func (f *FileLogger) Debug(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(msg)
	f.logger.Debug(msg)
}

func (f *FileLogger) Error(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	f.logger.Debug(msg)
}

func (f *FileLogger) Fatal(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	log.Println(msg) // print to shell i/o
	f.logger.Fatal(msg)
}

func (f *FileLogger) LogLevel() int8 {
	return f.logLevel
}

func (f *FileLogger) Close() {
	f.logger.Sync()
	f.rotator.Close()
}
