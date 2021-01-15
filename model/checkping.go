//检查ping信息
package model

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetWindowsPingInfo(t string) {
	for _, local := range CheckPing {
		cmd := exec.Command("ping", t, "5", local)
		byte, _ := cmd.CombinedOutput()

		info := WindowsPing(byte)

		fmt.Printf("%s  %s\n", local, info)
	}

}
func GetUnixPingInfo(t string) {
	for _, local := range CheckPing {
		cmd := exec.Command("ping", t, "5", local)
		byte, _ := cmd.CombinedOutput()

		info := UnixPing(byte)

		fmt.Printf("%s  %s\n", local, info)
	}

}

func WindowsPing(b []byte) (info string) {
	str := ConvertByte2String(b, GB18030)
	if strings.Index(str, "主机") != -1 {
		return Pingerr2
	}
	if strings.Index(str, "最短") == -1 {
		return Pingerr1
	}
	tempbyte := strings.SplitAfter(str, "数据包:")
	tempstr := strings.TrimSpace(tempbyte[1])
	tempbyte = strings.SplitAfter(tempstr, "往返行程的估计时间(以毫秒为单位):")
	timestr := strings.TrimSpace(tempbyte[1])
	tempbyte = strings.Split(tempbyte[0], ",")
	loststr := tempbyte[0]
	astr := strings.Split(loststr, "，")
	loststr = strings.Join(astr[0:3], "，")

	return timestr + " " + loststr
}

func UnixPing(b []byte) string {
	str := ConvertByte2String(b, UTF8)
	tempstr := strings.Split(str, "statistics ---")
	return tempstr[1]
}
