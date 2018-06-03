package redis

import (
	"github.com/convee/goboot/conf"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func New(name string) (pool *redis.Pool) {
	redisConfig := conf.Get().Redis[name]
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisConfig.Address)
		},
	}
	return
}
