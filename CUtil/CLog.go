package CUtil

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func init() {
	fmt.Println("log初始化")
}

// AddInfoLog 添加日志
func AddInfoLog(data map[string]interface{}, msg string) error {
	now := time.Now()
	logFileName := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
	dirName := fmt.Sprintf("%d%02d%02d", now.Year(), now.Month(), now.Day())
	rootPath := MyUtilConfig.LogRootPath + "/runtime/log/" + dirName
	success, _err := checkDir(rootPath)
	if !success {
		fmt.Println(_err)
		return _err
	}
	LogFile, err := os.OpenFile(rootPath+"/"+logFileName+"."+MyUtilConfig.LogFileExtensionName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return errors.New("打开日志文件错误")
	}

	writers := []io.Writer{
		LogFile, os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		DataKey:         "data",
	})
	logrus.SetOutput(fileAndStdoutWriter)
	if data != nil {
		logrus.WithFields(data).Info(msg)
	} else {
		logrus.Info(msg)
	}
	return nil
}

func checkDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if os.IsNotExist(_err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, errors.New("创建文件目录失败")
		}
	}
	return true, _err

}
