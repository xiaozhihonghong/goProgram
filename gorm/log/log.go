package log

import (
	"log"
	"os"
)

// New 创建一个新的 Logger。out 参数设置日志数据将被写入的目的地
// 参数 prefix 会在生成的每行日志的最开始出现
// 参数 flag 定义日志记录包含哪些属性
var (
	errorlog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)  //红色
	infolog = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)  //蓝色
)

var (
	Error = errorlog.Println
	Errorf = errorlog.Printf
	Info = infolog.Println
	Infof = infolog.Printf
)
