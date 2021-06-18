package imetrics

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
)

type PromNet struct {
	LISTEN string
	ESTABLISHED string
	CLOSEWAIT string
	CLOSED string
}

func (pn *PromNet) ToString() string{
	ReNewNet(pn)
	return "" + pn.LISTEN + "\n" +
	 "" + pn.ESTABLISHED + "\n" +
	 "" + pn.CLOSEWAIT + "\n" +
	 "" + pn.CLOSED +"\n"

}

type Nets struct {
	Count int
	Object []*NetObject
}
type NetObject struct {
	fd uint32
	pid int32
	localaddr string
	remoteaddr string
}

func ReNewNet(PromNet *PromNet)  {
	nc, err := net.Connections("all")
	if err!=nil{
		panic(err)
	}
	ESTABLISHED :=  &Nets{}
	CLOSEWAIT := &Nets{}
	LISTEN := &Nets{}
	CLOSED := &Nets{}
	for i := range nc {
		switch nc[i].Status {
		 case "ESTABLISHED":
		 	ESTABLISHED.Count++
		 case "CLOSE_WAIT":
		 	CLOSEWAIT.Count++
		 case "LISTEN":
		 	LISTEN.Count++
		 case "CLOSED":
		 	CLOSED.Count++
		}
	}
	PromNet.LISTEN = fmt.Sprintf("easy_prometheus_system_net_LISTEN%s %d",BuildFmt("LISTEN",Application,"LISTEN"),LISTEN.Count)
	PromNet.CLOSEWAIT = fmt.Sprintf("easy_prometheus_system_net_CLOSE_WAIT%s %d",BuildFmt("CLOSE_WAIT",Application,"CLOSE_WAIT"),CLOSEWAIT.Count)
	PromNet.CLOSED = fmt.Sprintf("easy_prometheus_system_net_CLOSED%s %d",BuildFmt("CLOSED",Application,"CLOSED"),CLOSED.Count)
	PromNet.ESTABLISHED = fmt.Sprintf("easy_prometheus_system_net_ESTABLISHED%s %d",BuildFmt("ESTABLISHED",Application,"ESTABLISHED"),ESTABLISHED.Count)
}