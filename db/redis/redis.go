package redis

import (
	"time"

	"github.com/convee/goboot/conf"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func New(name string) *redis.Pool {
	redisConfig := conf.Get().Redis[name]
	pool = &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,                                  //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   redisConfig.MaxActive,                                //最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: time.Duration(redisConfig.IdleTimeout) * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) { //建立连接
			return redis.Dial("tcp", redisConfig.Address)
		},
	}
	return pool
}

func Set(k, v string) (bool, error) {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Setnx(k, v string) (bool, error) {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("SETNX", k, v)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Setex(k string, v string, ex int) (bool, error) {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("SETEX", k, ex, v)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Get(k string) (string, error) {
	c := pool.Get()
	defer c.Close()
	return redis.String(c.Do("GET", k))
}

func Expire(k string, ex int) (bool, error) {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		return false, err
	}
	return true, nil
}
