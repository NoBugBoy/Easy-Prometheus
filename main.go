package main

import (
	"flag"
	"log"
	"metrics/imetrics"
	"net/http"
)


func main() {
	port := flag.String("port", "8089", "port")
	flag.Parse()
	http.HandleFunc("/", all)
	log.Printf("metrics started on port %s !",*port)
	err := http.ListenAndServe(":"+ *port, nil)
	log.Fatal(err)
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
