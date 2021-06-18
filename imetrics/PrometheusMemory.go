package imetrics

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

type PromMem struct {
	Total string
	Available string
	Used string
	Free string
	UsedPercent string
	SwapTotal string
	SwapUsed string
	SwapFree string
	SwapUsedPercent string

}

func (pm *PromMem) ToString() string{
	ReNewMemory(pm)
	return pm.Available + "\n" +
		pm.Total + "\n" +
		pm.Free + "\n" +
		pm.Used + "\n" +
		pm.UsedPercent + "\n" +
		pm.SwapTotal + "\n" +
		pm.SwapUsed + "\n" +
		pm.SwapFree + "\n" +
		pm.SwapUsedPercent + "\n"
}

func ReNewMemory(IPromMen *PromMem)  {
	iMem, err := mem.VirtualMemory()
	if err!=nil{
		panic(err)
	}
	iSwapMemMem, err := mem.SwapMemory()
	if err!=nil{
		panic(err)
	}
	IPromMen.Total = fmt.Sprintf("easy_prometheus_system_mem_total%s %d",BuildFmt("总内存",Application,"总内存"),iMem.Total / 1024 / 1024 / 1024)
	IPromMen.Available = fmt.Sprintf("easy_prometheus_system_mem_available%s %d",BuildFmt("可用内存",Application,"可用内存"),iMem.Available / 1024 / 1024)
	IPromMen.Used = fmt.Sprintf("easy_prometheus_system_mem_used%s %d",BuildFmt("已用内存",Application,"已用内存"),iMem.Used / 1024 / 1024)
	IPromMen.Free = fmt.Sprintf("easy_prometheus_system_mem_free%s %d",BuildFmt("空闲内存",Application,"空闲内存"),iMem.Free / 1024 / 1024)
	IPromMen.UsedPercent = fmt.Sprintf("easy_prometheus_system_mem_used_percent%s %.2f",BuildFmt("占用百分比",Application,"占用百分比"),iMem.UsedPercent)
	IPromMen.SwapTotal = fmt.Sprintf("easy_prometheus_system_mem_swap_total%s %d",BuildFmt("swap总内存",Application,"swap总内存"),iSwapMemMem.Total / 1024 / 1024 /1024)
	IPromMen.SwapUsed = fmt.Sprintf("easy_prometheus_system_mem_swap_used%s %d",BuildFmt("swap已用内存",Application,"swap已用内存"),iSwapMemMem.Used / 1024 / 1024)
	IPromMen.SwapFree = fmt.Sprintf("easy_prometheus_system_mem_swap_free%s %d",BuildFmt("swap空闲内存",Application,"swap空闲内存"),iSwapMemMem.Free / 1024 / 1024)
	IPromMen.SwapUsedPercent = fmt.Sprintf("easy_prometheus_system_mem_swap_used_percent%s %.2f",BuildFmt("swap占用百分比",Application,"swap占用百分比"),iSwapMemMem.UsedPercent)

}
