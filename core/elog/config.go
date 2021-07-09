package elog

import (
	"fmt"

	"github.com/gotomicro/ego/core/eapp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// DefaultLoggerName 业务日志名
	DefaultLoggerName = "default.log"
	// EgoLoggerName 系统文件名
	EgoLoggerName = "ego.sys"
)

// Config ...
type Config struct {
	Debug           bool   // 是否双写至文件控制日志输出到终端
	Level           string // 日志初始等级，默认info级别
	Dir             string // [fileWriter]日志输出目录，默认logs
	Name            string // [fileWriter]日志文件名称，默认框架日志ego.sys，业务日志default.log
	EnableAddCaller bool   // 是否添加调用者信息，默认不加调用者信息
	EnableAsync     bool   // 是否异步，默认异步
	Writer          string // 使用哪种Writer，可选[file|ali|stderr]，默认file
	core            zapcore.Core
	asyncStopFunc   func() error
	fields          []zap.Field // 日志初始化字段
	CallerSkip      int
	encoderConfig   *zapcore.EncoderConfig
}

// Filename ...
func (config *Config) Filename() string {
	return fmt.Sprintf("%s/%s", config.Dir, config.Name)
}

// DefaultConfig ...
func DefaultConfig() *Config {
	dir := "./logs"
	if eapp.EgoLogPath() != "" {
		dir = eapp.EgoLogPath()
	}
	return &Config{
		Name:            DefaultLoggerName,
		Dir:             dir,
		Level:           "info",
		CallerSkip:      1,
		EnableAddCaller: false,
		EnableAsync:     true,
		asyncStopFunc:   func() error { return nil },
		encoderConfig:   defaultZapConfig(),
		Writer:          "file",
	}
}

func (c *Config) GetEncoderConfig() *zapcore.EncoderConfig {
	return c.encoderConfig
}
