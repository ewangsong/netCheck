package main

import (
	"fmt"

	"github.com/ewangsong/netcheck/model"
)

func main() {
	model.AddCheckNet()
	fmt.Println("-------------------------正在检测请稍等.......................")
	fmt.Println()
	s := model.NewSystemInfo()
	fmt.Println("-------------------------本机信息---------------------------")
	s.GetSystemInfo()
	switch s.Ostype {
	case "windows":
		printPinginfo()
		model.GetWindowsPingInfo("-n")
		printHttpInfo()
		model.GetHttpInfo()
		model.HoldOn()
	case "linux", "darwin":
		printPinginfo()
		model.GetUnixPingInfo("-c")
		printHttpInfo()
		model.GetHttpInfo()
		// case "darwin":
		// 	printPinginfo()
		// 	model.GetUnixPingInfo("-c")
		// 	printHttpInfo()
		// 	model.GetHttpInfo()
	}

}

func printPinginfo() {
	fmt.Println()
	fmt.Println("-----------------------ping信息------------------------------")
}

func printHttpInfo() {
	fmt.Println()
	fmt.Println("----------------------HTTP信息------------------------------")
}
