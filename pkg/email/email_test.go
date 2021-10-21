package email_test

import (
	innerTest "xz_gowork/test"
	"os"
	"testing"
	"time"

	"xz_gowork/pkg/email"
	"xz_gowork/pkg/global"
	"xz_gowork/pkg/tools"
)

func TestMain(m *testing.M) {
	innerTest.InitConfig()
	code := m.Run()
	os.Exit(code)
}

func Test_SendMail(t *testing.T) {
	ec := email.NewMailClient()
	err := ec.SendMail("test send mail", "<h1>test mail</h1>", "social.noreply@xiaozhu.com", "gengtao@xiaozhu.com")
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
		"汉字乱码???????",
	}, "scene-demo-test-%s", date)
	if err != nil {
		t.Errorf("Expected err nil,actual %v", err)
	}
	err = ec.SendAttachMail("scene-demo", "<h1>详见附件</h1>", file, "测试附件文件")
	if err != nil {
		t.Errorf("Expected err nil,actual %v", err)
	}
}
