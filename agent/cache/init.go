package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/nanopoker/minisns/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (

	// gPool is a redis connection pool
	gPool *redis.Pool
)

// Init initilizes redis cache
func Init() error {
	gPool = &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: config.REDIS_MAXIDLECONNS,
		// max number of connections
		MaxActive:   config.REDIS_MAXACTIVECONNS,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(config.REDIS_DIALNETWORK, config.REDIS_HOST, redis.DialDatabase(config.REDIS_DATABASE))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			elapse := time.Now().Sub(t)
			if elapse.Seconds() < 10 {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	// test connections after initializing redis connection pool
	conn := getConn()
	err := ping(conn)
	if err != nil {
		return err
	}
	return nil
}

func ping(c redis.Conn) error {
	_, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}
	return nil
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		gPool.Close()
		os.Exit(0)
	}()
}

func getConn() redis.Conn {
	conn := gPool.Get()
	return conn
}

func GetPool() *redis.Pool {
	return gPool
}
