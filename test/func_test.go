package test

import (
	"fmt"
	"github.com/cc2415/go_cc_util/CUtil"
	"os"
	"testing"
)

func init() {
	root, _ := os.Getwd()
	fmt.Printf("初始化日志位置:%s\n", root)
	CUtil.InitUtil(root, "log")
}

func TestLog(t *testing.T) {

	fmt.Println("写入日志")
	a := map[string]interface{}{"fff": 23}
	CUtil.AddInfoLog(map[string]interface{}{"fff": 23}, "测试")
	CUtil.AddInfoLog(a, "测试2")
}

func TestTime(t *testing.T) {
	//res :=CUtil.Format(CUtil.GetTimeStamp(), "Y-m-d H:i:s")
	//fmt.Println(res)
	ts := CUtil.GetTimestamp()

	layout := "Y-m-d H:i:s"
	fmt.Printf("%s  %s", "格式化：	", CUtil.Format(ts, layout))
	fmt.Println("")

	layout = "Y"
	fmt.Printf("%s  %s", "年：   	", CUtil.Format(ts, layout))
	fmt.Println("")
	layout = "m"
	fmt.Printf("%s  %s", "月：   	", CUtil.Format(ts, layout))
	fmt.Println("")

}
