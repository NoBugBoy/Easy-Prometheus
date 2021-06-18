package imetrics

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

type PromDisk struct {
	Mountpoint string
	Total string
	Free string
	Used string
	UsedPercent string
}

func (pd *PromDisk) ToString() string{
	return  "" +
	    pd.Total + "\n" +
		pd.Free + "\n" +
		pd.Used + "\n" +
		pd.UsedPercent + "\n"

}
func DiskString() string {
	infos, err := disk.Partitions(false)
	if err!=nil{
		panic(err)
	}
	IPromDisk := make([]*PromDisk,0)
	for i, info := range infos {
		stat, err := disk.Usage(info.Mountpoint)
		if err!=nil{
			panic(err)
		}
		pd := &PromDisk{}
		pd.Mountpoint = info.Mountpoint
		pd.Free = fmt.Sprintf("easy_prometheus_system_disk_%d_free%s %d",i,BuildFmt(info.Mountpoint+"剩余空间",Application,info.Mountpoint+"剩余空间"),stat.Free / 1024 / 1024)
		pd.Total = fmt.Sprintf("easy_prometheus_system_disk_%d_total%s %d",i,BuildFmt(info.Mountpoint+"总空间",Application,info.Mountpoint+"总空间"),stat.Total / 1024 / 1024)
		pd.Used = fmt.Sprintf("easy_prometheus_system_disk_%d_used%s %d",i,BuildFmt(info.Mountpoint+"已用空间",Application,info.Mountpoint+"已用空间"),stat.Used /1024 /1024)
		pd.UsedPercent = fmt.Sprintf("easy_prometheus_system_disk_%d_used_percent%s %.2f",i,BuildFmt(info.Mountpoint,Application,info.Mountpoint+"已用占比"),stat.UsedPercent)
		IPromDisk = append(IPromDisk, pd)
	}
	var str string
	for i := range IPromDisk {
		str += IPromDisk[i].ToString()
	}
	return str
}
