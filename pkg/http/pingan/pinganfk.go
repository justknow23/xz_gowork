package pingan

import (
	"context"
)


// PingAnFkClient 需调用 NewOrderClient 获得初始化
type PingAnFkClient struct {
	PingAnClient
}

// NewPingAnFkClient -
func NewPingAnFkClient(ctx context.Context) *PingAnFkClient {
	bc := &PingAnFkClient{}
	bc.Name = "service_pingan_gateway"
	bc.Ctx = ctx
	return bc
}
