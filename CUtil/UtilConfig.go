package CUtil

type UtilConfig struct {
	// 根目录
	LogRootPath string
	// 日志文件扩展名
	LogFileExtensionName string
}

var MyUtilConfig UtilConfig

// InitUtil 初始化工具类
func InitUtil(LogRootPath string, logFileExtensionName string) {
	MyUtilConfig = UtilConfig{
		LogRootPath:          LogRootPath,
		LogFileExtensionName: logFileExtensionName,
	}
}
