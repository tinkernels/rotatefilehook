package rotatefilehook

import (
    "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
    "io"
)

type RotateFileConfig struct {
    Filename   string
    MaxSize    int
    MaxBackups int
    MaxAge     int
    LocalTime  bool
    Compress   bool
    Levels     []logrus.Level
    Formatter  logrus.Formatter
}

type RotateFileHook struct {
    Config    RotateFileConfig
    logWriter io.Writer
}

func NewRotateFileHook(config RotateFileConfig) (logrus.Hook, error) {

    hook := RotateFileHook{
        Config: config,
    }
    hook.logWriter = &lumberjack.Logger{
        Filename:   config.Filename,
        MaxSize:    config.MaxSize,
        MaxBackups: config.MaxBackups,
        MaxAge:     config.MaxAge,
        Compress:   config.Compress,
        LocalTime:  config.LocalTime,
    }
    return &hook, nil
}

func (hook *RotateFileHook) Levels() []logrus.Level {
    return hook.Config.Levels
}

func (hook *RotateFileHook) Fire(entry *logrus.Entry) (err error) {
    b, err := hook.Config.Formatter.Format(entry)
    if err != nil {
        return err
    }
    _, err = hook.logWriter.Write(b)
    if err != nil {
        return err
    }
    return nil
}
