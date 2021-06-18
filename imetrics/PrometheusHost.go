package imetrics

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
)

type PromHost struct {
	HostName string
	UpTime string
	KernelVersion string
	KernelArch string
	Platform string
}

func (ph *PromHost) ToString() string{
	ReNewHost(ph)
	return "" + ph.UpTime + "\n"

}

func ReNewHost(IPromHost *PromHost)  {
	hostInfo, err := host.Info()
	if err!=nil{
		panic(err)
	}
	IPromHost.UpTime = fmt.Sprintf("easy_prometheus_system_host_uptime%s %d",BuildFmt("启动时间",Application,"启动时间"),hostInfo.Uptime)
	IPromHost.HostName = fmt.Sprintf("easy_prometheus_system_host_name%s %s",BuildFmt("OS名称",Application,"OS名称"),hostInfo.Hostname)
	IPromHost.Platform = fmt.Sprintf("easy_prometheus_system_host_platform%s %s",BuildFmt("平台",Application,"平台"),hostInfo.Platform)
	IPromHost.KernelVersion = fmt.Sprintf("easy_prometheus_system_host_kernel_version%s %s",BuildFmt("内核版本",Application,"内核版本"),hostInfo.KernelVersion)
	IPromHost.KernelArch = fmt.Sprintf("easy_prometheus_system_host_kernel_arch%s %s",BuildFmt("内核架构",Application,"内核架构"),hostInfo.KernelVersion)
}
