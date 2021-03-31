package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/liuqianhong6007/k8s/api"
	"github.com/liuqianhong6007/k8s/internal"
	"github.com/liuqianhong6007/k8s/util"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "./config.yaml", "config file path")
	flag.Parse()
}

func main() {
	internal.ReadConfig(cfgPath)

	// init k8s grpc
	internal.InitK8sClientset(internal.Cfg().Inner, internal.Cfg().CfgPath)

	r := gin.Default()
	util.RegisterRoute(r)

	serverAddr := fmt.Sprintf("%s:%d", internal.Cfg().Host, internal.Cfg().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
