package server

import (
	"easyRedis/redisCom"
	"easyRedis/redisPool"
	dlock2 "easyRedis/redisTool/dlock"
	"github.com/gomodule/redigo/redis"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

func init() {
	conn := redisPool.GetRedis()
	defer conn.Close()
	do, err := conn.Do(redisCom.Set, "money", 1000)
	if err != nil {
		log.Println("set money err ", err)
		return
	}

	log.Println("redis set done do val ", do)
}

func serve1() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":6376", nil)
	if err != nil {
		log.Printf("serve err %s\n", err)
		return
	}
	log.Println("server1 started")
}

func serve2() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":6375", nil)
	if err != nil {
		log.Printf("serve err %s\n", err)
		return
	}
	log.Println("server2 started")
}

func forward(w http.ResponseWriter, r *http.Request) {
	w.Header()
}

var totalMoney = 1000
var ptm = &totalMoney
var mutex sync.Mutex

func handler(w http.ResponseWriter, r *http.Request) {

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("read req Body err %s\n", err)
		return
	}

	//got := workWithMutex()
	//got := workWithRedis()
	got := workWithRedisAndDlock()
	w.Write([]byte(strconv.Itoa(got)))

	log.Println(string(bytes), r.URL.Host, r.URL.Path)
}

func workWithMutex() int {
	mutex.Lock()
	defer mutex.Unlock()
	got := min(rand.Intn(10), *ptm)
	*ptm -= got
	return got
}

func workWithRedis() int {
	mutex.Lock()
	defer mutex.Unlock()
	conn := redisPool.GetRedis()
	defer conn.Close()
	v, err := redis.Int(conn.Do(redisCom.Get, "money"))
	if err != nil {
		log.Println("get money err ", err)
		return 0
	}

	got := min(rand.Intn(10), v)
	v -= got
	conn.Do(redisCom.Set, "money", v)
	return got
}

func workWithRedisAndDlock() int {

	dlock := dlock2.NewRedisDLock(0, rand.Int63())
	defer dlock.Close()

	ok, err := dlock.Lock()
	if err != nil {
		log.Println("workWithRedisAndDlock err ", err)
		return 0
	}
	if !ok {
		log.Println("workWithRedisAndDlock lock failed")
		return 0
	}
	defer dlock.Unlock()

	conn := redisPool.GetRedis()
	defer conn.Close()
	v, err := redis.Int(conn.Do(redisCom.Get, "money"))
	if err != nil {
		log.Println("get money err ", err)
		return 0
	}

	//发1-10块的红包
	got := min(rand.Intn(9)+1, v)
	v -= got
	conn.Do(redisCom.Set, "money", v)
	return got
}
