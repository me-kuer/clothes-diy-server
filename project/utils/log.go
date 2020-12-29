package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func initLogger() {
	Log = logrus.New()
	// 只输出不低于当前级别的是日志数据
	Log.SetLevel(logrus.InfoLevel)
	// 输出文件
	logfile, _ := os.OpenFile("logs/app.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	logrus.SetOutput(logfile)

	// 日志定位(显示输出该日志的 文件:行号)
	logrus.SetReportCaller(true)
}
