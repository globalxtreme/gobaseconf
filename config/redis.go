package config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	RedisPool              *redis.Pool
	RedisAsyncWorkflowPool *redis.Pool
)

func InitRedis() {
	var err error

	maxActive := 5000
	if maxActiveEnv := os.Getenv("REDIS_MAX_ACTIVE"); maxActiveEnv != "" {
		maxActive, err = strconv.Atoi(maxActiveEnv)
		if err != nil {
			log.Panicf("Failed to parse REDIS_MAX_ACTIVE env var: %s", err)
		}
	}

	maxIdle := 100
	if maxIdleEnv := os.Getenv("REDIS_MAX_IDLE"); maxIdleEnv != "" {
		maxIdle, err = strconv.Atoi(maxIdleEnv)
		if err != nil {
			log.Panicf("Failed to parse REDIS_MAX_IDLE env var: %s", err)
		}
	}

	idleTimeout := 180
	if idleTimeoutEnv := os.Getenv("REDIS_IDLE_TIMEOUT"); idleTimeoutEnv != "" {
		idleTimeout, err = strconv.Atoi(idleTimeoutEnv)
		if err != nil {
			log.Panicf("Failed to parse REDIS_IDLE_TIMEOUT env var: %s", err)
		}
	}

	RedisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
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

func InitRedisAsyncWorkflowPool() {
	var err error

	maxActive := 5000
	if maxActiveEnv := os.Getenv("REDIS_ASYNC_WORKFLOW_MAX_ACTIVE"); maxActiveEnv != "" {
		maxActive, err = strconv.Atoi(maxActiveEnv)
		if err != nil {
			log.Panicf("Failed to parse REDIS_ASYNC_WORKFLOW_MAX_ACTIVE env var: %s", err)
		}
	}

	maxIdle := 100
	if maxIdleEnv := os.Getenv("REDIS_ASYNC_WORKFLOW_MAX_IDLE"); maxIdleEnv != "" {
		maxIdle, err = strconv.Atoi(maxIdleEnv)
		if err != nil {
			log.Panicf("Failed to parse REDIS_ASYNC_WORKFLOW_MAX_IDLE env var: %s", err)
		}
	}

	idleTimeout := 180
	if idleTimeoutEnv := os.Getenv("REDIS_ASYNC_WORKFLOW_IDLE_TIMEOUT"); idleTimeoutEnv != "" {
		idleTimeout, err = strconv.Atoi(idleTimeoutEnv)
		if err != nil {
			log.Panicf("Failed to parse REDIS_ASYNC_WORKFLOW_IDLE_TIMEOUT env var: %s", err)
		}
	}

	RedisAsyncWorkflowPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
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
