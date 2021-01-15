//获取系统信息
package model

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	Ostype    string
	Osarch    string
	HostName  string
	CpuNumber int
	CpuUse    string
	MemInfo   string
	LanIP     string
	WlanIP    string
}

//返回一个系统信息结构体
func NewSystemInfo() *SystemInfo {
	ostype := getOsType()
	osarch := getCpuArch()
	hostname := getHostName()
	cpunumber := getCpuNumber()
	cpuuse := getCpuUse()
	meminfo := getMemInfo()
	lanip := getLanIP()
	wlanip := getWlanIP()
	return &SystemInfo{
		Ostype:    ostype,
		Osarch:    osarch,
		HostName:  hostname,
		CpuNumber: cpunumber,
		CpuUse:    cpuuse,
		MemInfo:   meminfo,
		LanIP:     lanip,
		WlanIP:    wlanip,
	}
}

//操作系统类型
func getOsType() string {
	return runtime.GOOS
}

//cpu 架构
func getCpuArch() string {
	return runtime.GOARCH
}

//获取电脑名称
func getHostName() string {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	return name
}

//cpu核数
func getCpuNumber() int {
	cpuNumber := runtime.GOMAXPROCS(0)
	return cpuNumber
}

//cpu使用率
func getCpuUse() string {
	cpuinfo, _ := cpu.Percent(time.Duration(time.Second), false)
	return fmt.Sprintf("CPU使用率:%.0f%%", cpuinfo[0])
}

//内存信息
func getMemInfo() string {
	memInfo, _ := mem.VirtualMemory()
	total := BturnG(memInfo.Total)
	free := BturnG(memInfo.Free)
	return fmt.Sprintf("总共:%vG,使用:%.0f%%,剩余%vG", total, memInfo.UsedPercent, free)
}

//内存大小转换
func BturnG(b uint64) int {
	return int(b / 1024 / 1024 / 1024)
}

//ip 信息
func getLanIP() (arrip string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())

	}
	//去除非活动IP
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			//去除回环地址
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						arrip = arrip + ipnet.IP.String() + " "
					}
				}
			}
		}
	}

	return arrip
}

//获取公网IP看是否开了代理

func getWlanIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

//检查操作系统的类型
func (s *SystemInfo) CheckOsType() string {
	return s.Ostype
}

func (s *SystemInfo) GetSystemInfo() {
	fmt.Printf("系统:%s\n", s.Ostype)
	fmt.Printf("架构:%s\n", s.Osarch)
	fmt.Printf("主机:%s\n", s.HostName)
	fmt.Printf("核数:%d\n", s.CpuNumber)
	fmt.Printf("CPU:%s\n", s.CpuUse)
	fmt.Printf("内存:%s\n", s.MemInfo)
	fmt.Printf("内网IP:%s\n", s.LanIP)
	fmt.Printf("公网IP:%s\n", s.WlanIP)

}

//windows 执行完不推出终端
func HoldOn() {
	fmt.Println("检测执行完成输入任意字符退出......")
	fmt.Scanln()

}
