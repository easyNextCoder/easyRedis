package more

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"time"
)

func redisType() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	//args := redis.Args{}

	c1.Do("hset", "m", 1, 1)
	c1.Do("HINCRBY", "n", 2, 2)
	do, err := c1.Do("hmget", "n", 2)
	if err != nil {
		return
	}
	fmt.Println(do)
}

func bigKey() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	//args := redis.Args{}

	startTime := time.Now()
	outMax := 1
	max := 1000000
	for i := 0; i < outMax; i++ {
		key := "key" + strconv.Itoa(i)
		for j := 0; j < max; j++ {
			innerKey := "innerKey" + strconv.Itoa(j)
			val := "val" + strconv.Itoa(j)
			c1.Do("hset", key, innerKey, val)
		}
	}
	endTime := time.Now()
	fmt.Println("time cost", endTime.Sub(startTime).Milliseconds())

}

func getBigKey() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	start := time.Now()

	stringMap, err := c1.Do("hgetall", "key0")
	if err != nil {
		return
	}
	end := time.Now()

	fmt.Println("getBigKey time", end.Sub(start).Milliseconds())

	cnt := 0
	for i, v := range stringMap.([]interface{}) {
		fmt.Println(i, string(v.([]byte)))
		cnt++
		if cnt > 10 {
			break
		}
	}

}

func runWork() {
	getBigKey()
}
