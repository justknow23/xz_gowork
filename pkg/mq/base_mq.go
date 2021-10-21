package mq

import (
	mqc "gitlab.idc.xiaozhu.com/xz-go/message-center"
	"xz_gowork/pkg/errors"
)

// BaseMQ baseMQ 写队列基础对象
type BaseMQ struct {
	Service string
	Topic   string
	Opts    []func(message *mqc.Message)
}

// Send 发消息
func (bc *BaseMQ) Send(body interface{}) error {
	message, err := mqc.NewMessage(bc.Service, bc.Topic, body, bc.Opts...)
	if err != nil {
		return err
	}
	err = message.Send(func(response *mqc.MessageCenterResponse) error {
		if !response.IsOK() {
			return errors.New(response.GetErrorMessage())
		}
		return nil
	})
	return err
}
