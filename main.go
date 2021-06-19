package main

import (
	"encoding/json"
	"flag"
	"fmt"
	httpgo "github.com/NoBugBoy/httpgo/http"
	"io/ioutil"
	"log"
	"metrics/imetrics"
	"net/http"
)


func main() {
	port := flag.String("port", "8089", "port")
	flag.Parse()
	http.HandleFunc("/", all)
	http.HandleFunc("/webhook", dingding)
	log.Printf("metrics started on port %s !",*port)
	err := http.ListenAndServe(":"+ *port, nil)
	log.Fatal(err)
}

type Ding struct {
	Alerts []struct{
		Annotations struct{
			Summary string `json:"summary"`
		}  `json:"annotations"`
	}  `json:"alerts"`
}

func dingding(w http.ResponseWriter, r *http.Request)  {
	s, _ := ioutil.ReadAll(r.Body)
	ding := &Ding{}
	fmt.Println(string(s))
	json.Unmarshal(s,ding)
	anno := ding.Alerts[0]
	req :=&httpgo.Req{}
	x, err := req.Header("Content-Type", "application/json").
		Method(http.MethodPost).
		Url("https://oapi.dingtalk.com/robot/send?access_token=xxxxx").
		Params(httpgo.Query{
			"link": map[string]interface{}{
				"title": "AlertManager通知",
				"text": "通知" + anno.Annotations.Summary,
				"picUrl": "https://photo.16pic.com/00/65/09/16pic_6509905_b.png",
			    "messageUrl":"http://localhost:9090/alerts",
			},
		    "msgtype": "link",
	    }).Go().Body()
	if err!=nil {
		log.Println(err)
	}
	fmt.Println(x)
}

func all(w http.ResponseWriter, r *http.Request) {
	var res string
	host := &imetrics.PromHost{}
	res += host.ToString()
	mem := imetrics.PromMem{}
	res += mem.ToString()
	cpu := imetrics.PromCpu{}
	res += cpu.ToString()
	net := imetrics.PromNet{}
	res += net.ToString()
	res += imetrics.DiskString()
	_, _ = w.Write([]byte(res))
}
