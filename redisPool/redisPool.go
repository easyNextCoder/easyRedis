package redisPool

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type EasyRedis struct {
	conn redis.Conn
}

func (zpr *EasyRedis) Close() error { return zpr.conn.Close() }

func (zpr *EasyRedis) Err() error {
	return zpr.conn.Err()
}

func (zpr *EasyRedis) Do(commandName string, args ...interface{}) (interface{}, error) {
	start := time.Now()
	reply, err := zpr.conn.Do(commandName, args...)
	end := time.Now()
	diff := end.Sub(start)
	timeout := time.Second

	if diff > timeout {
		log.Printf("EasyRedis slow log. diff(%s) start(%d) end(%d) err(%s) cmd(%s) args(%+v)", diff, start.Unix(), end.Unix(), err, commandName, args)
	}

	return reply, err
}

func (zpr *EasyRedis) Send(commandName string, args ...interface{}) error {
	return zpr.conn.Send(commandName, args...)
}

func (zpr *EasyRedis) Flush() error {
	return zpr.conn.Flush()
}

func (zpr *EasyRedis) Receive() (interface{}, error) {
	return zpr.conn.Receive()
}

func easyRedisDial(network, address string, options ...redis.DialOption) (redis.Conn, error) {
	conn, err := redis.Dial(network, address, options...)
	if err != nil {
		return nil, err
	}
	return &EasyRedis{conn: conn}, nil
}

func makePool(addr string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: time.Minute * 10,
		MaxActive:   20,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return easyRedisDial("tcp", addr, redis.DialReadTimeout(time.Millisecond*200))
		},
	}

}

var defaultPool *redis.Pool

func init() {
	defaultPool = makePool("127.0.0.1:6379")
	log.Println("init redis pool done")
}

func GetRedis() redis.Conn {
	return defaultPool.Get()
}
