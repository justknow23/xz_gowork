package cache

import (
	"context"
	"gitlab.idc.xiaozhu.com/xz-go/common/log"
	"xz_gowork/pkg/errors"
)

// CrmVisitCache 缓存
type CrmVisitCache struct {
	Base
	Ck, Ik string
}

func (p *CrmVisitCache) Check(ctx context.Context) error {
	p.Ck = p.GetCacheKey()
	p.Ik = p.GetIdempotentKey()
	//重复
	co := p.Get(ctx, p.Ck)
	if len(co) > 0 {
		log.Infof("checkcachehasexist %v", p)
		return errors.ErrorRepeat
	}
	if err := p.Set(ctx, p.Ck, "1"); err != nil {
		return err
	}
	//幂等
	io := p.Get(ctx, p.Ik)
	if len(io) > 0 {
		log.Infof("checkidempotenthasexist %v", p)
		return errors.ErrorIdempotent
	}
	if err := p.Set(ctx, p.Ik, "1"); err != nil {
		return err
	}

	return nil
}

func (p *CrmVisitCache) Dela(ctx context.Context) error {
	_ = p.Del(ctx, p.Ck)
	_ = p.Del(ctx, p.Ik)
	return nil
}
