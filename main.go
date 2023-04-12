package main

import (
	"fmt"
	CUtil2 "go_cc_util/CUtil"
	"os"
)

func init() {
	root, _ := os.Getwd()
	CUtil2.InitUtil(root, "log")
	fmt.Println("初始化")
}

type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

type student struct {
	name string
	age  int
}

func main() {
	//a:=map[string]interface{}{"fff":23}
	//CUtil.AddInfoLog(map[string]interface{}{"fff":23}, "测试")

	fmt.Println(CUtil2.GetTimeStamp())
	fmt.Println(CUtil2.FormatTimeStamp(CUtil2.GetTimeStamp(), "Y-m-d"))
	//fmt.Println(CUtil.HttpGet())
	//fmt.Println(CUtil.HttpPost())
	fmt.Println(CUtil2.JsonGet("", "data", "list", 0))
	fmt.Println(CUtil2.JsonEncode(map[string]string{"data": "wahaha"}))
}
