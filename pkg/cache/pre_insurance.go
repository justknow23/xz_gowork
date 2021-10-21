package cache

import (
	"context"
	"fmt"
	"gitlab.idc.xiaozhu.com/xz-go/common/log"
	"gitlab.idc.xiaozhu.com/xz-go/common/redis"
	"insurance/pkg/errors"
	"insurance/pkg/tools"
	"time"
)

// CacheExpire cacheExpire
const CacheExpire = 3 * time.Second

// PreInsuranceCache 缓存
type PreInsuranceCache struct {
	Expired time.Duration
	Prefix  string
	Key     string
}

// GetKey 获取缓存key
func (p *PreInsuranceCache) GetKey() string {
	return fmt.Sprintf(p.Prefix + p.Key)
}

// Get 获取缓存
func (p *PreInsuranceCache) Get(ctx context.Context) string {
	key := p.GetKey()
	return redis.Cluster.Get(ctx, key).Val()
}

// Set 设置缓存
func (p *PreInsuranceCache) Set(ctx context.Context, data string) error {
	key := p.GetKey()
	return redis.Cluster.Set(ctx, key, data, p.Expired).Err()
}

// Del 删除缓存
func (p *PreInsuranceCache) Del(ctx context.Context) error {
	key := p.GetKey()
	return redis.Cluster.Del(ctx, key).Err()
}

// GetExpired 获取过期时间
func (p *PreInsuranceCache) GetExpired() (time.Duration, error) {
	monthNow := time.Now()
	monthEnd, err := tools.GetSpecialEndDate(monthNow, "5m")
	if err != nil {
		return time.Second, err
	}
	expired := time.Duration(monthEnd.Sub(monthNow).Seconds())
	return expired, nil
}

func (p *PreInsuranceCache) CacheCheck(ctx context.Context) error {
	//并发
	o := p.Get(ctx)
	if len(o) > 0 {
		log.Infof("redis check bookOrderId has exist %v", p)
		return errors.New("重复消费！" + p.GetKey())
	}
	if err := p.Set(ctx, "1"); err != nil {
		log.Info("redis set error: ", err)
		return err
	}
	//幂等  todo id+time 幂等处理 event

	return nil
}
