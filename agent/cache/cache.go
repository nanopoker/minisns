package cache

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"time"
)

const REDIS_INTERFACE_NAME = "Redis"

//RConn contains a redis connection
type RConn struct {
	conn redis.Conn
}

//GetRConn returns a RConn
func GetRConn() *RConn {
	rConn := &RConn{}
	rConn.conn = gPool.Get()
	return rConn
}

//Set sets a value in cache
func (rConn *RConn) Set(ctx context.Context, key string, value []byte) (err error) {
	conn := rConn.conn

	_, err = conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		err = errors.WithMessage(err, fmt.Sprintf("Set key %s to %s fail", key, v))
		return
	}
	return
}

//SetEX sets key to hold the string value and set key to timeout after a given number of seconds
func (rConn *RConn) SetEX(ctx context.Context, key string, timeout int, value []byte) (err error) {
	conn := rConn.conn

	_, err = conn.Do("SETEX", key, timeout, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return errors.WithMessage(err, fmt.Sprintf("SetEX key %s to %s fail", key, v))
	}
	return nil
}

//SetNXEX sets key to hold string value if key does not exist with TTL
//return ErrNil if key exists
//return nil if key not exists
func (rConn *RConn) SetNXEX(ctx context.Context, key string, value []byte, ttl int) (data []byte, err error) {
	conn := rConn.conn

	data, err = redis.Bytes(conn.Do("SET", key, value, "NX", "EX", ttl))
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return data, errors.WithMessage(err, fmt.Sprintf("SetNXEX key %s to %s fail", key, v))
	}
	return data, nil
}

// Get gets the value of a key in []byte
// if key does not exist，error is not nil.
func (rConn *RConn) Get(ctx context.Context, key string) (by []byte, err error) {
	conn := rConn.conn

	var data []byte
	data, err = redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, errors.WithMessage(err, fmt.Sprintf("Get key %s fail", key))
	}
	return data, nil
}

//Delete deletes a key
func (rConn *RConn) Delete(ctx context.Context, key string) (err error) {
	conn := rConn.conn

	_, err = conn.Do("DEL", key)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("DEL key %s fail", key))
	}
	return nil
}

//LPush inserts all the specified values at the tail of the list stored at key
func (rConn *RConn) LPush(ctx context.Context, key string, value []byte) (err error) {
	connection := rConn.conn
	_, err = connection.Do("LPUSH", key, value)
	return errors.WithMessage(err, fmt.Sprintf("LPUSH failed, key:%v, value:%v", key, value))
}

//BRPop is a blocking list pop primitive
func (rConn *RConn) BRPop(ctx context.Context, key string, queueBlockingTime int) (by []byte, err error) {
	connection := rConn.conn
	timer := time.NewTimer(time.Duration(queueBlockingTime) * time.Second)
	for {
		select {
		case <-timer.C:
			return nil, nil
		default:
			val, err := redis.Bytes(connection.Do("RPOP", key))
			if err != nil {
				if errors.Cause(err) != redis.ErrNil {
					return nil, errors.WithMessage(err, "BRPOP failed")
				}
				time.Sleep(100 * time.Millisecond)
				continue
			}
			return val, nil
		}
	}
}

//LLen returns the length of the list stored at key
func (rConn *RConn) LLen(ctx context.Context, key string) (len int, err error) {
	connection := rConn.conn
	data, err := redis.Int(connection.Do("LLEN", key))
	return data, errors.WithMessage(err, fmt.Sprintf("LLEN failed, key:%v", key))
}

func (rConn *RConn) Close() error {
	if rConn.conn != nil {
		return rConn.conn.Close()
	}
	return nil
}
