package imetrics

import (
	"fmt"
	"strings"
)

const (
	Application = "easy_prometheus"
)

func BuildFmt(action ,application,cause string) string {
	return fmt.Sprintf("{action=\"%s\",application=\"%s\",cause=\"%s\",} ",action,application,cause)
}
func BuildFmtNet(action ,application string, net []*NetObject) string {
	str := ""
	for i := range net {
		str += fmt.Sprintf("pid=%d,fd=%d,laddr=%s,raddr=%s,",net[i].pid,net[i].fd,strings.Replace(net[i].localaddr,"\"","",-1),strings.Replace(net[i].remoteaddr,"\"","",-1))
	}
	return fmt.Sprintf("{action=\"%s\",application=\"%s\",cause=\"%s\",} ",action,application,str)
}
