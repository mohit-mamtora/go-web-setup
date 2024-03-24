package logger

import (
	"fmt"
	"path/filepath"

	"github.com/mohit-mamtora/go-web-setup/app/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileLogger struct {
	logLevel int
	logger   *zap.Logger
	rotator  *rotatingFile
}

var _ logger.Log = (*FileLogger)(nil)

func NewFileLogger(filePath, fileName string, maxSizeMB int, jsonEncoder bool) (*FileLogger, error) {

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
	core := zapcore.NewCore(encoder, sink, zapcore.InfoLevel)

	return &FileLogger{
		rotator:  rotator,
		logLevel: 3,
		logger:   zap.New(core),
	}, nil
}

func (f *FileLogger) SetLogLevel(level int) {
	f.logLevel = level
}

func (f *FileLogger) Info(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	f.logger.Info(msg)
}

func (f *FileLogger) Debug(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	f.logger.Debug(msg)
}

func (f *FileLogger) Error(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	f.logger.Debug(msg)
}

func (f *FileLogger) LogLevel() int {
	return f.logLevel
}

func (f *FileLogger) Close() {
	f.logger.Sync()
	f.rotator.Close()
}
