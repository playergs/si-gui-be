package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB        DBConfig        `json:"DB"`
	JwtAuth   JwtAuthConfig   `json:"JwtAuth"`
	BizRedis  redis.RedisConf `json:"BizRedis"`
	RateLimit RateLimitConfig `json:"RateLimit"`
}

type DBConfig struct {
	DataSource string `json:"DataSource"`
}

type JwtAuthConfig struct {
	AccessSecret  string `json:"AccessSecret"`
	AccessExpire  int64  `json:"AccessExpire"`
	RefreshSecret string `json:"RefreshSecret"`
	RefreshExpire int64  `json:"RefreshExpire"`
}

type RateLimitConfig struct {
	RequestsPerSecond int `json:"RequestsPerSecond"`
	BurstSize         int `json:"BurstSize"`
}
