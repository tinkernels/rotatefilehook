# RotateFileHook

This is a simple hook for logrus to write rotated log files using https://github.com/natefinch/lumberjack


# Example

```go
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tinkernels/rotatefilehook"
)

rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
    Filename:   "logfile.log",
    MaxSize:    5, // the maximum size in megabytes
    MaxBackups: 7, // the maximum number of old log files to retain
    MaxAge:     7, // the maximum number of days to retain old log files
    LocalTime:  true,
    Levels:     []logrus.Level{logrus.InfoLevel, logrus.WarnLevel},
    Formatter:  &logrus.TextFormatter{FullTimestamp: true},
})
if err != nil {
    panic(err)
}
var logger = logrus.New()
logger.AddHook(rotateFileHook)
```