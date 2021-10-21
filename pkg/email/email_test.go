package email_test

import (
	test2 "insurance/test"
	"os"
	"testing"
	"time"

	"insurance/pkg/email"
	"insurance/pkg/global"
	"insurance/pkg/tools"
)

func TestMain(m *testing.M) {
	test2.InitConfig()
	code := m.Run()
	os.Exit(code)
}

func Test_SendMail(t *testing.T) {
	ec := email.NewMailClient()

	err := ec.SendMail("test send mail", "<h1>test mail</>", "gengtao@xiaozhu.com", "gengtao@xiaozhu.com")
	if err != nil {
		t.Errorf("Expected err nil,actual %v", err)
	}
}

func Test_SendAttachMail(t *testing.T) {
	ec := email.NewMailClient()
	date := time.Now().Format(global.DateFmtYMD)
	file, err := tools.CreateCSVFile([]string{
		"11111",
		"22222",
	}, "level-calc-failed-%s", date)
	if err != nil {
		t.Errorf("Expected err nil,actual %v", err)
	}
	err = ec.SendAttachMail("level_calc_monitor", "<h1>失败房东ID详见附件</h1>", file, "房东等级计算失败", date)
	if err != nil {
		t.Errorf("Expected err nil,actual %v", err)
	}
}
