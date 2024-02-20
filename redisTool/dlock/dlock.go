package dlock

import (
	"easyRedis/redisCom"
	"easyRedis/redisPool"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math/rand"
)

type DLock interface {
	Lock()
	Unlock()
}

type redisDLock struct {
	redis.Conn
	pid        int    //进程的标识
	key        string //分布式锁的key
	randNumber int64  //标识锁的随机数
}

func NewRedisDLock(pid int, randN int64) *redisDLock {
	return &redisDLock{
		Conn:       redisPool.GetRedis(),
		key:        "dlock",
		pid:        pid,
		randNumber: randN,
	}
}

func (r *redisDLock) Lock() (bool, error) {
	do, err := redis.String(r.Do(redisCom.Set, r.key, r.randNumber, "px", 5000, "nx"))
	if err != nil {
		fmt.Printf("redisDLock lock err %s\n", err)
		return false, fmt.Errorf("redisDLock lock err %s", err)
	}
	fmt.Printf("redisDLock lock do %+v\n", do)
	return do == "OK", nil
}

func (r *redisDLock) Unlock() {
	do, err := r.Do(redisCom.Eval, redisCom.ScriptCompAndDel, 1, r.key, r.randNumber)
	if err != nil {
		fmt.Printf("redisDLock unlock err %s\n", err)
		return
	}
	fmt.Printf("redisDLock unlock do %+v\n", do)
}

func runRedisDLock() {

	rd := NewRedisDLock(0, rand.Int63())
	rd.Lock()
	defer rd.Unlock()
}
