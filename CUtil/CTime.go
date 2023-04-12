package CUtil

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func init() {

}

// GetTimeStamp 获得秒时间戳
func GetTimeStamp() int64 {
	return time.Now().Unix()
}

// GetTimestamp 返回当前时间戳，以秒为单位
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// FormatTimeStamp 格式化时间戳 format: Y-m-d H:i:s 或着 ""
func FormatTimeStamp(timeStamp int64, format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	//Y-m-d H:i:s 2006年01月02 15:04:05
	format = strings.Trim(format, " ")
	format = strings.Replace(format, "Y", "2006", 1)
	format = strings.Replace(format, "y", "06", 1)
	format = strings.Replace(format, "m", "01", 1)
	format = strings.Replace(format, "n", "1", 1)
	format = strings.Replace(format, "d", "02", 1)
	format = strings.Replace(format, "H", "15", 1)
	format = strings.Replace(format, "i", "04", 1)
	format = strings.Replace(format, "s", "05", 1)

	return time.Unix(timeStamp, 0).Format(format)
}

// Format 格式化时间戳为指定格式的字符串
func Format(ts int64, layout string) string {
	t := time.Unix(ts, 0)
	// 将 PHP 格式化字符串转换为 Go 格式化字符串
	goLayout := phpToGoDateFormat(layout)
	return t.Format(goLayout)
}

// phpToGoDateFormat 将 PHP 格式化字符串转换为 Go 格式化字符串
func phpToGoDateFormat(phpFormat string) string {
	//replacements := []struct {
	//	from, to string
	//}{
	//	{"Y", "2006"},
	//	{"y", "06"},
	//	{"m", "01"},
	//	{"n", "1"},
	//	{"d", "02"},
	//	{"j", "2"},
	//	{"H", "15"},
	//	{"h", "03"},
	//	{"i", "04"},
	//	{"s", "05"},
	//}
	//
	//for _, r := range replacements {
	//	layout = strings.Replace(layout, r.from, r.to, -1)
	//}
	//
	//return layout

	replacements := map[string]string{
		"d": "02",
		"D": "Mon",
		"j": "2",
		"l": "Monday",
		"N": "Monday", // Not a perfect match, but close enough for most cases
		"S": "th",
		"w": "0",
		"z": "0",
		"W": "00",
		"F": "January",
		"m": "01",
		"M": "Jan",
		"n": "1",
		//"t": "31",
		"L": "false",
		"o": "2006",
		"Y": "2006",
		"y": "06",
		"a": "pm",
		"A": "PM",
		"B": "",
		"g": "3",
		"G": "15",
		"h": "03",
		"H": "15",
		"i": "04",
		"s": "05",
		"u": "000000",
		"e": "MST",
		"I": "false",
		"O": "-0700",
		"P": "-07:00",
		"T": "MST",
		"Z": "0",
		"c": "",
		"r": "",
		"U": "0",
	}
	var buffer bytes.Buffer
	for i := 0; i < len(phpFormat); i++ {
		if replacement, ok := replacements[string(phpFormat[i])]; ok {
			buffer.WriteString(replacement)
		} else {
			buffer.WriteByte(phpFormat[i])
		}
	}
	return buffer.String()
}

// replaceWithPlaceholder 将指定字符替换为 Go 格式化字符串的占位符
func replaceWithPlaceholder(layout, from, to string) string {
	return fmt.Sprintf("%s%s%s", "%", to, layoutReplace(layout, from, ""))
}

// layoutReplace 替换 PHP 格式化字符串中的字符
func layoutReplace(layout, from, to string) string {
	return string([]rune(layoutReplaceAll(layout, from, to)))
}

// layoutReplaceAll 替换 PHP 格式化字符串中的所有匹配字符
func layoutReplaceAll(layout, from, to string) string {
	return string(bytesReplaceAll([]byte(layout), []byte(from), []byte(to)))
}

// bytesReplaceAll 替换字节片中的所有匹配字节
func bytesReplaceAll(b, from, to []byte) []byte {
	return bytes.ReplaceAll(b, from, to)
}
