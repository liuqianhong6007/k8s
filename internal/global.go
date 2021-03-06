package internal

import (
	"k8s.io/client-go/kubernetes"

	"github.com/liuqianhong6007/k8s/util"
)

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Inner   bool   `yaml:"inner"`
	CfgPath string `yaml:"cfg_path"`
}

// 全局配置参数
var gConfig Config

func ReadConfig(filepath string) {
	util.ReadConfig(filepath, &gConfig)
}

func Cfg() *Config { return &gConfig }

// 全局 kubernetes 客户端
var gClientset *kubernetes.Clientset

func InitK8sClientset(inner bool, cfgPath string) {
	gClientset = util.NewKubernetesClientset(inner, cfgPath)
}

func K8sClientset() *kubernetes.Clientset { return gClientset }
