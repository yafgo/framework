package logger

import (
	"errors"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/yafgo/framework/facades"
	"github.com/yafgo/framework/log/formatter"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Daily struct {
}

func (daily *Daily) Handle(channel string) (logrus.Hook, error) {
	var hook logrus.Hook
	logPath := facades.Config.GetString(channel + ".path")
	if logPath == "" {
		return hook, errors.New("error log path")
	}

	// lumberjack
	_rotatelogs := &lumberjack.Logger{
		Filename:   logPath, // 日志文件位置
		MaxSize:    1,       // 单文件最大容量,单位是MB
		MaxBackups: 3,       // 最大保留过期文件个数
		MaxAge:     1,       // 保留过期文件的最大时间间隔,单位是天
		Compress:   true,    // 是否需要压缩滚动日志, 使用的 gzip 压缩
	}

	levels := getLevels(facades.Config.GetString(channel + ".level"))
	writerMap := lfshook.WriterMap{}
	for _, level := range levels {
		writerMap[level] = _rotatelogs
	}

	return lfshook.NewHook(
		writerMap,
		&formatter.General{},
	), nil
}
