// @Title  logger.go
// @Description	日志模块，zinx框架中的日志模块太丑了，不符合自己的口味，就自己弄一个新的
// @Author  SiriusWilliam  2021-01-19 16:37
// @Update  2021-01-19 16:37
package logger

import (
	"encoding/json"
	"github.com/aceld/zinx/utils"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

/*
	操作句柄
*/
type Logger struct {
	lock sync.Mutex
}

/*
	配置文件读取
*/
type LoggerConfig struct {
	LogDir       string
	LogFileName  string
	LogLevel     int
	LogToConsole bool
	LogToFile    bool
}

var LogConf LoggerConfig = LoggerConfig{}

var Log Logger

func init() {
	Log = Logger{}
	file, err := ioutil.ReadFile(utils.GlobalObject.ConfFilePath)
	if err != nil {
		LogConf.LogDir = ""
		LogConf.LogFileName = ""
		LogConf.LogToConsole = true
		LogConf.LogToFile = false
	} else {
		_ = json.Unmarshal(file, &LogConf)

		if LogConf.LogDir == "" && LogConf.LogFileName == "" {
			LogConf.LogToFile = false
		}
	}
}
// 检查日志文件是否存在
func checkFile(){
	var file *os.File
	var err error
	filePath := LogConf.LogDir + "\\" + LogConf.LogFileName
	// 首先看目录的路径是否存在
	_, err = os.Stat(filePath)
	// 如果文件不存在，先创建目录，再创建文件
	if err != nil {
		// 创建多级目录，如果目录存在，则下面一行不会完全执行成功
		_ = os.MkdirAll(LogConf.LogDir, os.ModePerm)
		// 创建文件
		file, err = os.Create(filePath)
	}
	// 关闭文件，不使用这个句柄，而用OpenFile的追加模式
	file.Close()
}
/*
ALL 最低等级的，用于打开所有日志记录。

TRACE designates finer-grained informational events than the DEBUG.Since:1.2.12，很低的日志级别，一般不会使用。

DEBUG 指出细粒度信息事件对调试应用程序是非常有帮助的，主要用于开发过程中打印一些运行信息。

INFO 消息在粗粒度级别上突出强调应用程序的运行过程。打印一些你感兴趣的或者重要的信息，这个可以用于生产环境中输出程序运行的一些重要信息，但是不能滥用，避免打印过多的日志。

WARN 表明会出现潜在错误的情形，有些信息不是错误信息，但是也要给程序员的一些提示。

ERROR 指出虽然发生错误事件，但仍然不影响系统的继续运行。打印错误和异常信息，如果不想输出太多的日志，可以使用这个级别。

FATAL 指出每个严重的错误事件将会导致应用程序的退出。这个级别比较高了。重大错误，这种级别你可以直接停止程序了。

OFF 最高等级的，用于关闭所有日志记录。
*/
const (
	ALL = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

func (l *Logger) Error(msg string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now()
	str := "[" + now.Format("2006-01-02 15:04:05") + "]\tERROR\t" + msg + "\n"
	if LogConf.LogLevel > ERROR {
		return
	}
	if LogConf.LogToConsole {
		color.Red(str)
	}
	if LogConf.LogToFile {
		var file *os.File
		var err error
		// 拼接路径
		filePath := LogConf.LogDir + "\\" + LogConf.LogFileName

		checkFile()
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 777)
		if err != nil {
			color.Red("[%s]\t%s\n", now.Format("2006-01-02 15:04:05"), err.Error())
			return
		}
		defer file.Close()
		length, err := file.WriteString(str)
		if err != nil || length != len(str) {
			color.Red("Log File Write Error:", err)
			return
		}
	}
}
