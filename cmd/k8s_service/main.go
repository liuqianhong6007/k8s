package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

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

	setPodLabels(internal.Cfg().Port)

	serverAddr := fmt.Sprintf("%s:%d", internal.Cfg().Host, internal.Cfg().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}

// 将地址信息写入 kubernetes labels
func setPodLabels(port int) {
	clientset := util.NewKubernetesClientset(true, "")
	namespace := os.Getenv("NAMESPACE")
	podIP := os.Getenv("POD_IP")
	podName := os.Getenv("HOSTNAME")

	data := fmt.Sprintf(`{ "metadata": { "annotations": { "addr": "%s:%d" } } }`, podIP, port)
	_, err := clientset.CoreV1().Pods(namespace).Patch(context.Background(), podName, types.MergePatchType, []byte(data), metav1.PatchOptions{})
	if err != nil {
		log.Println(err)
	}
}
