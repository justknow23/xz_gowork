package cache

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"insurance/pkg/json"

	"github.com/sirupsen/logrus"
	"gitlab.idc.xiaozhu.com/xz-go/common/redis"
	"gitlab.idc.xiaozhu.com/xz-go/common/server"
)

// Options 分布式锁配置
type Options struct {
	expired time.Duration
	prefix  string
	keys    interface{}
	data    string
	ctx     context.Context
	debug   bool
	logger  *logrus.Entry
}

// Option 分布式锁配置方法
type Option func(*Options)

// Timeout 超时时间默认:5s
func Timeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.expired = timeout
	}
}

// Prefix 前缀默认:""
func Prefix(prefix string) Option {
	return func(o *Options) {
		o.prefix = prefix
	}
}

// Keys 唯一标识
func Keys(keys interface{}) Option {
	return func(o *Options) {
		o.keys = keys
	}
}

// SetData 存储值
func SetData(data string) Option {
	return func(o *Options) {
		o.data = data
	}
}

// WithContext 上下文,是否调试
func WithContext(ctx context.Context, debug ...bool) Option {
	return func(o *Options) {
		o.ctx = ctx
		if logger := server.FromContext(ctx).Log(); len(debug) > 0 && debug[0] && logger != nil {
			o.debug = true
			o.logger = logger.WithField("name", "DistributedLock")
		} else {
			o.debug = false
		}
	}
}

// NewOptions 构建配置
func NewOptions(options ...Option) Options {
	opts := Options{
		expired: 5 * time.Second,
		prefix:  "",
		keys:    "",
		data:    "1",
		ctx:     context.Background(),
		debug:   false,
		logger:  &logrus.Entry{},
	}
	for _, opt := range options {
		opt(&opts)
	}
	return opts
}

// DistributedLock 分布式锁
type DistributedLock struct {
	opts Options
}

// NewDistributedLock 创建分布式锁
func NewDistributedLock(opt ...Option) (*DistributedLock, error) {
	opts := NewOptions(opt...)
	return &DistributedLock{
		opts: opts,
	}, nil
}

// Lock 加锁
func (dl *DistributedLock) Lock() bool {
	key := dl.getKey()
	if dl.opts.debug {
		dl.opts.logger.Infof("lock value:%v  key: %s", dl.opts.keys, key)
	}
	return redis.Cluster.SetNX(dl.opts.ctx, key, dl.opts.data, dl.opts.expired).Val()
}

// Unlock 解锁
func (dl *DistributedLock) Unlock() error {
	key := dl.getKey()
	if dl.opts.debug {
		dl.opts.logger.Infof("unlock value:%v key: %s", dl.opts.keys, key)
	}
	return redis.Cluster.Del(dl.opts.ctx, key).Err()
}

// getKey 获取缓存key
func (dl *DistributedLock) getKey() string {
	return fmt.Sprintf("%s%s", dl.opts.prefix, dl.hash(dl.serialize(dl.opts.keys)))
}

// serialize 序列化函数
func (dl *DistributedLock) serialize(keys interface{}) string {
	str, _ := json.Marshal(keys)
	return string(str)
}

// hash Hash函数
func (dl *DistributedLock) hash(input string) string {
	tokenByte := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", tokenByte)
}
