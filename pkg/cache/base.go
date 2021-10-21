package cache

import (
	"context"
	"fmt"
	"gitlab.idc.xiaozhu.com/xz-go/common/redis"
	"time"
	"xz_gowork/pkg/global"
	"xz_gowork/pkg/tools"
)

// Expire cacheExpire
const Expire = 5 * time.Second

// Base 缓存
type Base struct {
	Expired time.Duration
	Prefix  string
	Key     string
	Time    time.Time
}

// GetCacheKey 获取缓存key
func (p *Base) GetCacheKey() string {
	return fmt.Sprintf(global.Settings.PrefixCache + ":" + p.Prefix + ":" + p.Key)
}

// GetIdempotentKey 获取缓存key
func (p *Base) GetIdempotentKey() string {
	return fmt.Sprintf(global.Settings.PrefixCache + ":" + p.Prefix + ":" + p.Key + ":" + p.Time.Format(global.DateFmtYMDHISS))
}

// Get 获取缓存
func (p *Base) Get(ctx context.Context, key string) string {
	return redis.Cluster.Get(ctx, key).Val()
}

// Set 设置缓存
func (p *Base) Set(ctx context.Context, key string, data string) error {
	return redis.Cluster.Set(ctx, key, data, p.Expired).Err()
}

// Del 删除缓存
func (p *Base) Del(ctx context.Context, key string) error {
	return redis.Cluster.Del(ctx, key).Err()
}

// GetExpired 获取过期时间
func (p *Base) GetExpired() (time.Duration, error) {
	monthNow := time.Now()
	monthEnd, err := tools.GetSpecialEndDate(monthNow, "5m")
	if err != nil {
		return time.Second, err
	}
	expired := time.Duration(monthEnd.Sub(monthNow).Seconds())
	return expired, nil
}
