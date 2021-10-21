package global

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.idc.xiaozhu.com/xz-go/common/config"
)

// ProjectName -
const ProjectName = "insurance"

// Settings 业务配置
var Settings Config

// Config 配置
type Config struct {
	// Stdout 终端输出
	Stdout bool `yaml:"stdout"`
	// LogLevel 日志级别 panic<fatal<error<warn<info<debug<trace
	LogLevel string `yaml:"log_level"`
	// LoggerLevel 日志级别
	LoggerLevel logrus.Level `yaml:"-"`
	// PingAnDepartmentCode -
	PinganGrantType string `yaml:"pingan_grant_type"`
	// PingAnDepartmentCode -
	PinganClientId string `yaml:"pingan_client_id"`
	// PingAnDepartmentCode -
	PinganClientSecret string `yaml:"pingan_client_secret"`
	// PingAnDepartmentCode -
	PinganFKPrefix string `yaml:"pingan_fk_prefix"`
	// PingAnDepartmentCode -
	PinganFDPrefix string `yaml:"pingan_fd_prefix"`
	// PingAnDepartmentCode -
	PingAnDepartmentCode string `yaml:"department_code"`
	// PingAnPrepaidAccountId -
	PingAnPrepaidAccountId string `yaml:"prepaid_account_id"`
	// PingAnPrepaidAccountType -
	PingAnPrepaidAccountType string `yaml:"prepaid_account_type"`
	// PingAnTenantProductCode -
	PingAnTenantProductCode string `yaml:"tenant_product_code"`
	// PingAnTenantProductPackageCode -
	PingAnTenantProductPackageCode string `yaml:"tenant_product_package_code"`
	// PingAnLandlordProductCode -
	PingAnLandlordProductCode string `yaml:"landlord_product_code"`
	// PingAnLandlordProductPackageCode -
	PingAnLandlordProductPackageCode string `yaml:"landlord_product_package_code"`
	// EmailPeopleList -
	EmailPeopleList string `yaml:"email_people_list"`
	// ContentNoticeString -
	ContentNoticeString string `yaml:"content_notice_string"`
	// ContentNoticeUrl -
	ContentNoticeUrl string `yaml:"content_notice_url"`
	// FailCodeList -
	FailCodeList string `yaml:"fail_code_list"`
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
