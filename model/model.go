package model

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8     = Charset("UTF-8")
	GB18030  = Charset("GB18030")
	Pingerr1 = "最短 = 不可达，最长 = 不可达，平均 = 不可达 丢失率是已发送 = 5，已接收 = 0，丢失 = 5 (100% 丢失)"
	Pingerr2 = "请求找不到主机 fdsfsdf。请检查该名称，然后重试。"
)

var CheckPing = map[string]string{
	"核心":   "192.168.1.1",
	"防火墙":  "192.168.1.2",
	"网关出口": "114.255.81.251",
	"DNS":  "172.30.16.13",
	"114":  "114.114.114.114",
}

var CheckHttp = map[string]string{
	//"googel": "https://www.google.com/",
	"百度": "https://www.baidu.com",
	"网易": "https://www.163.com",
	"万顺": "http://www.wsecar.com",
	"石墨": "http://shimo.im",
}

//字符转换
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

//输入自己要检测的网址
func AddCheckNet() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入要检测的地址默认请回车：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	if text != "" {
		text = checkInPut(text)
		CheckHttp["测试"] = text
		astr := strings.Split(text, "/")
		CheckPing[astr[2]] = astr[2]
	}
}
func checkInPut(s string) string {
	//判断是否有前缀http://没有则加上，有则不修改
	if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
		return s
	} else {
		s = "http://" + s
	}
	return s
}
