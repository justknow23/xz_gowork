package global

import (
	"fmt"
	"gitlab.idc.xiaozhu.com/xz-go/common/config"
)

// ProjectName -
const ProjectName = "xz_gowork"

// Settings 业务配置
var Settings Config

// Config 配置
type Config struct {
	// ConfigDemo -
	ConfigDemo  string `yaml:"configDemo"`
	PrefixCache string `yaml:"prefixCache"`
	Ak          string `yaml:"accessKey"`
	Sk          string `yaml:"secretKey"`
	BucketName  string `yaml:"bucketName"`
}

// SetupByConfig 初始化
func SetupByConfig(configKey string, loadFunc func(name string)) {
	if err := config.LoadWithCallback(configKey, &Settings, loadFunc); err != nil {
		fmt.Printf("load config [%s] error:%s\n", configKey, err.Error())
		return
	}
	loadFunc(ProjectName)
	fmt.Printf("config [%s] change callback registered\n", configKey)
}

func reload(name string) {
	if err := config.Load(name, &Settings); err != nil {
		return
	}
	fmt.Printf("load config [%s] success:\n %#v\n", ProjectName, Settings)
}

// Setup 初始化
func Setup() {
	SetupByConfig(ProjectName, reload)
}
