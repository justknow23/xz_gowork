package email

import (
	"fmt"
	"gitlab.idc.xiaozhu.com/xz-go/common/config"
	"log"
)

//InsuranceTos 保险接收邮件用户列表
var InsuranceTos []string

type MailClient struct {
}

func NewMailClient() *MailClient {
	c := &MailClient{}
	ConfigInit()

	return c
}

func ConfigInit() {
	// 邮件配置读取
	var mailConfig *mail.Config
	if err := config.Load(mail.ConfigKey, &mailConfig); err != nil {
		log.Println(err)
	}
	fmt.Printf("email config : %v\r\n", mailConfig)
	mail.Setup(mailConfig)
}

// SendMail 普通邮件
func (c *MailClient) SendMail(sub string, HTMLContent string, from string, to ...string) error {
	m, err := mail.NewMail(
		mail.From(from),
		mail.To(to),
		mail.Subject(sub),
	)
	if err != nil {
		log.Fatal(err)
	}
	m.SetBody(mail.TypeHTML, HTMLContent)

	if err := mail.Send(m); err != nil {
		return err
	}
	return nil
}

// SendSceneMail 预置场景邮件
func (c *MailClient) SendSceneMail(scene string, HTMLContent string, topic ...interface{}) error {
	// 使用场景邮件时, 后面的参数将会按照顺序替换到场景邮件主题内格式符
	m, err := mail.NewSceneMail(scene, topic...)
	if err != nil {
		log.Fatal(err)
	}
	m.SetBody(mail.TypeHTML, HTMLContent)
	if err := mail.Send(m); err != nil {
		return err
	}
	return nil
}

// SendAttachMail 带附件邮件
func (c *MailClient) SendAttachMail(scene string, HTMLContent string, attach string, topic ...interface{}) error {
	// 使用场景邮件时, 后面的参数将会按照顺序替换到场景邮件主题内格式符
	m, err := mail.NewSceneMail(scene, topic...)
	if err != nil {
		log.Fatal(err)
	}
	m.SetBody(mail.TypeHTML, HTMLContent)
	if attach != "" {
		m.Attach(attach) // 参考 mail_test.go
	}

	if err := mail.Send(m); err != nil {
		return err
	}
	return nil
}