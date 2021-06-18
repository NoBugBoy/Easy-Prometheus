package imetrics

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"log"
	"time"
)
//# HELP jvm_gc_pause_seconds Time spent in GC pause
//# TYPE jvm_gc_pause_seconds summary
//jvm_gc_pause_seconds_count{action="end of major GC",application="test",cause="Metadata GC Threshold",} 1.0
//jvm_gc_pause_seconds_sum{action="end of major GC",application="test",cause="Metadata GC Threshold",} 0.064
func GetHost() string {
	host, err := host.Info()
	if err != nil {
		log.Println(err)
		return ""
	}
	return "system_host_hostname" + "{action=\"主机名称\",application=\"go\",cause=\"主机名称\",} " + host.Hostname
}

func GetCpu() string {
	cpus, err := cpu.Info()
	if err != nil {
		log.Println(err)
		return ""
	}
	str, _ := json.Marshal(cpus)
	return string(str)
}
func Cpu() []float64 {
	percent, _ := cpu.Percent(time.Second,false)
	return percent
}
func Avg() string {
	info, _ := load.Avg()
	str, _ := json.Marshal(info)
	return string(str)
}
func Memory() string {
	v, _ := mem.VirtualMemory()
	str, _ := json.Marshal(v)
	return string(str)
}
func Disk(){
	infos, _ := disk.Partitions(false)
	for _, info := range infos {
		stat, _ := disk.Usage(info.Mountpoint)
		fmt.Println(stat.Total / 1024 / 1024 )
		fmt.Println(stat.Free / 1024 / 1024 )
	}

}
