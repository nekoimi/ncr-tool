package logs

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"strconv"
)

// CustomTextFormatter 自定义格式化器，继承自 logrus.TextFormatter
type CustomTextFormatter struct {
	logrus.TextFormatter
	ForceColors   bool
	ColorDebug    *color.Color
	ColorInfo     *color.Color
	ColorWarning  *color.Color
	ColorError    *color.Color
	ColorCritical *color.Color
}

func New() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&CustomTextFormatter{
		ForceColors:   true,
		ColorDebug:    color.New(color.FgWhite),
		ColorInfo:     color.New(color.FgGreen),
		ColorWarning:  color.New(color.FgYellow),
		ColorError:    color.New(color.FgRed),
		ColorCritical: color.New(color.BgRed, color.FgWhite),
	})
	return logger
}

// Format 格式化方法，用于将日志条目格式化为字节数组
func (f *CustomTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.ForceColors {
		switch entry.Level {
		case logrus.DebugLevel:
			f.ColorDebug.Println(entry.Message) // 使用蓝色打印信息日志
		case logrus.InfoLevel:
			f.ColorInfo.Println(entry.Message) // 使用蓝色打印信息日志
		case logrus.WarnLevel:
			f.ColorWarning.Println(entry.Message) // 使用黄色打印警告日志
		case logrus.ErrorLevel:
			f.ColorError.Println(entry.Message) // 使用红色打印错误日志
		case logrus.FatalLevel, logrus.PanicLevel:
			f.ColorCritical.Println(entry.Message) // 使用带有红色背景和白色文本的样式打印严重日志
		default:
			f.PrintColored(entry)
		}
		return nil, nil
	} else {
		return f.TextFormatter.Format(entry)
	}
}

// PrintColored 自定义方法，用于将日志条目以带颜色的方式打印出来
func (f *CustomTextFormatter) PrintColored(entry *logrus.Entry) {
	levelColor := color.New(color.FgCyan, color.Bold)             // 定义蓝色和粗体样式
	levelText := levelColor.Sprintf("%-6s", entry.Level.String()) // 格式化日志级别文本

	msg := levelText + " " + entry.Message
	if entry.HasCaller() {
		msg += " (" + entry.Caller.File + ":" + strconv.Itoa(entry.Caller.Line) + ")" // 添加调用者信息
	}

	fmt.Fprintln(color.Output, msg) // 使用有颜色的方式打印消息到终端
}
