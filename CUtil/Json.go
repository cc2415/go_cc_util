package CUtil

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// JsonEncode 转为json
func JsonEncode(data interface{}) (string, error) {
	//return json.Marshal(data)
	if jsonStr, err := json.Marshal(data); err != nil {
		fmt.Println("Error =", err)
		return "", err
	} else {
		return string(jsonStr), nil
	}
}

// JsonGet 解析获得json字符串里面的数据 jsonStr：json字符串。path为层级目录 data, tool, 0
func JsonGet(jsonStr string, path ...interface{}) jsoniter.Any {
	return jsoniter.Get([]byte(jsonStr), path...)
}
