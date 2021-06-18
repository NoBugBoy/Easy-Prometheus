package imetrics

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"time"
)

type PromCpu struct {
	Family string
	ModelName string
	Percent string
	Load1 string
	Load5 string
	Load15 string
}

func (pc *PromCpu) ToString() string{
	ReNewCpu(pc)
	return pc.Family + "\n" +
		pc.Percent + "\n" +
		pc.Load1 + "\n" +
		pc.Load5 + "\n" +
		pc.Load15 + "\n"
}
func ReNewCpu(IPromCpu *PromCpu)  {
	cpuInfo, err := cpu.Info()
	if err!=nil{
		panic(err)
	}
	percent, err := cpu.Percent(time.Second,false)
	if err!=nil{
		panic(err)
	}
	info, err := load.Avg()
	if err!=nil{
		panic(err)
	}
	if len(cpuInfo) > 0 {
		ci := cpuInfo[0]
		IPromCpu.Family = fmt.Sprintf("easy_prometheus_system_cpu_family%s %s",BuildFmt("核心数",Application,"核心数"),ci.Family)
		IPromCpu.ModelName = fmt.Sprintf("easy_prometheus_system_cpu_model_name%s %s",BuildFmt("CPU名称",Application,"CPU名称"),ci.ModelName)
	}
	if len(percent) > 0 {
		cpuSum := percent[0]
		IPromCpu.Percent = fmt.Sprintf("easy_prometheus_system_cpu_percent%s %.4f",BuildFmt("Cpu利用率",Application,"Cpu利用率"),cpuSum)
	}
	IPromCpu.Load1 = fmt.Sprintf("easy_prometheus_system_cpu_load1%s %.4f",BuildFmt("1分钟Cpu平均负载",Application,"1分钟Cpu平均负载"),info.Load1)
	IPromCpu.Load5 = fmt.Sprintf("easy_prometheus_system_cpu_load5%s %.4f",BuildFmt("5分钟Cpu平均负载",Application,"5分钟Cpu平均负载"),info.Load5)
	IPromCpu.Load15 = fmt.Sprintf("easy_prometheus_system_cpu_load15%s %.4f",BuildFmt("15分钟Cpu平均负载",Application,"15分钟Cpu平均负载"),info.Load15)

}
