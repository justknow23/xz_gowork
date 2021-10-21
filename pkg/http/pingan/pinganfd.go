package pingan

import "context"

// PingAnFdClient 需调用 NewOrderClient 获得初始化
type PingAnFdClient struct {
	PingAnClient
}

// NewPingAnFdClient -
func NewPingAnFdClient(ctx context.Context) *PingAnFdClient {
	bc := &PingAnFdClient{}
	bc.Name = "service_pingan_gateway"
	bc.Ctx = ctx
	return bc
}