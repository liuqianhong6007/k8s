package main

import (
	"flag"

	"github.com/liuqianhong6007/k8s/cmd/watch_service/grpc"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8083", "grpc server addr")
	flag.Parse()
}

func main() {
	server := grpc.NewServer(addr)
	server.SetPodLabels()
	server.Listen()
}
