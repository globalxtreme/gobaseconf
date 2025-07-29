package config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

var (
	RedisPool              *redis.Pool
	RedisAsyncWorkflowPool *redis.Pool
)

func InitRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   500,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")))
			if err != nil {
				return nil, err
			}

			if os.Getenv("REDIS_PASSWORD") != "" {
				if _, err = c.Do("AUTH", os.Getenv("REDIS_PASSWORD")); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, nil
		},
	}
}

func InitRedisAsyncWorkflow() {
	RedisAsyncWorkflowPool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   500,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_ASYNC_WORKFLOW_HOST"), os.Getenv("REDIS_ASYNC_WORKFLOW_PORT")))
			if err != nil {
				return nil, err
			}

			if os.Getenv("REDIS_ASYNC_WORKFLOW_PASSWORD") != "" {
				if _, err = c.Do("AUTH", os.Getenv("REDIS_ASYNC_WORKFLOW_PASSWORD")); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, nil
		},
	}
}
