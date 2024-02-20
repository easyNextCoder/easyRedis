package string

import (
	. "easyRedis/redisCom"
	"easyRedis/redisPool"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sort"
	"strconv"
	"time"
)

func stringOperate() {
	conn := redisPool.GetRedis()
	defer conn.Close()

	//set
	conn.Do(Set, "key", 1)
	//get
	do, err := conn.Do(Get, "key")
	if err != nil {
		return
	}
	fmt.Println("set", do, string(do.([]byte)))

	//setex Redis Setex 命令为指定的 key 设置值及其过期时间。如果 key 已经存在， SETEX 命令将会替换旧的值。
	conn.Do(Setex, "setex_key", 3, "setex_val") //3
	do, err = conn.Do(Get, "setex_key")
	if err != nil {
		return
	}
	fmt.Println("setex get", do, string(do.([]byte)))
	time.Sleep(time.Second * 3)
	do, err = conn.Do(Get, "setex_key")
	if err != nil {
		return
	}
	fmt.Println("setex get after 3s", do, err)

	//del
	conn.Do(Del, "key")
	do2, err := conn.Do(Get, "key")
	if err == redis.ErrNil {
		fmt.Println("redis.ErrNil")
	}
	if err != nil {
		return
	}
	fmt.Println("del", do2)

	//incr
	conn.Do(Set, "key", "hello1")
	conn.Do(Incr, "key")
	val, err := redis.String(conn.Do(Get, "key"))
	if err != nil {
		return
	}
	fmt.Println("incr 对字符串，浮点不生效不生效，<但是对于字符串数字是生效的>", val)

	conn.Do(Set, "key", "1")
	conn.Do(Incr, "key")
	val2, err := redis.String(conn.Do(Get, "key"))
	if err != nil {
		return
	}
	fmt.Println("incr 只对整型的数生效", val2)

	//decrby
	conn.Do(Del, "key")
	conn.Do(Set, "key", "1")
	conn.Do(Decrby, "key", 10)
	val3, err := redis.String(conn.Do(Get, "key"))
	if err != nil {
		return
	}
	fmt.Println("decrby 能减到负数", val3)

	//setnx
	conn.Do(Del, "key")
	conn.Do(Setnx, "key", "1")
	conn.Do(Setnx, "key", "2")
	val4, err := redis.String(conn.Do(Get, "key"))
	if err != nil {
		return
	}
	val42, _ := conn.Do(Get, "key")
	fmt.Println("setnx 不会发生覆盖", val4, string(val42.([]byte)))

	//mset mget
	conn.Do(Del, "key")
	conn.Do(Mset, "key", "val1", "key2", "val2", "key3", "val3")
	val5, err := redis.Strings(conn.Do(Mget, "key", "key2", "key3"))
	if err != nil {
		return
	}
	val52, _ := conn.Do(Mget, "key", "key2", "key3")
	fmt.Println("mset mget", val5, string(val52.([]interface{})[2].([]byte)))

	//APPEND
	conn.Do(Del, "key")
	conn.Do(Set, "key", "1")
	conn.Do(APPEND, "key", "1234")
	val6, err := redis.String(conn.Do(Get, "key"))
	if err != nil {
		return
	}
	val62, _ := conn.Do(Get, "key")
	fmt.Println("append结果", val6, string(val62.([]byte)))

	//strlen
	conn.Do(Del, "key")
	conn.Do(Set, "key", "12345")
	val7, err := redis.Int64(conn.Do(Strlen, "key"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	val72, _ := conn.Do(Get, "key")
	fmt.Println("strlen", val7, string(val72.([]byte))) //strlen

	//incrbyfloat
	conn.Do(Del, "key")
	conn.Do(Set, "key", "3.5")
	val8, err := redis.Float64(conn.Do(Incrbyfloat, "key", 1.2))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	val82, _ := conn.Do(Get, "key")
	fmt.Println("incrbyfloat", val8, string(val82.([]byte)))

	//incrbyfloat
	conn.Do(Del, "key")
	conn.Do(Set, "key", "123456789")
	val9, err := redis.String(conn.Do(Getrange, "key", 0, 4))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("getrange", val9)

	//setrange
	conn.Do(Del, "key")
	conn.Do(Set, "key", "123456789")
	conn.Do(Setrange, "key", 3, "hello")
	val10, err := redis.String(conn.Do(Get, "key"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("setrange", val10)

	//scan
	conn.Do("flushall")
	conn.Do("hmset", "key", "field1", "val1") //scan既可以扫描出字符串，对于集合和哈希等一样能够扫描出来
	conn.Do("hmset", "key", "field2", "val2")
	conn.Do("hmset", "key", "field3", "val3")
	conn.Do("sadd", "set", "val1")
	for i := 0; i < 100; i++ {
		conn.Do(Set, fmt.Sprintf("key%02d", i), "val"+strconv.Itoa(i))
		conn.Do(Set, fmt.Sprintf("name%02d", i), "val"+strconv.Itoa(i))
	}
	var pos int = 0
	var res []string
	scan_(conn, &pos, 10, 30, "name*", &res) //scan的过程中会出现重复，count*n会大于总的key的个数
	sort.StringSlice(res).Sort()
	fmt.Println("scan res:", len(res), res)

}
