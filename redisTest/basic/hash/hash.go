package hash

import (
	. "easyRedis/redisCom"
	"easyRedis/redisPool"
	"easyRedis/util"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

var conn = redisPool.GetRedis()

type Hash struct {
	funcMap map[string]func()
	key     string
}

func (h *Hash) GetFuncMap() map[string]func() {
	return h.funcMap
}

func (h *Hash) GetKey() string {
	return h.key
}

var hash = &Hash{funcMap: make(map[string]func()), key: "user_hash"}

func init() {
	BindFuncMap(hash, conn)
}

func hasht() {
	conn := redisPool.GetRedis()
	defer conn.Close()

	//hscan
	conn.Do("flushall")
	for i := 0; i < 1000; i++ {
		str := util.RandString(80)
		k1, _, k2, _ := fmt.Sprintf("key%02d%s", i, str), "val"+strconv.Itoa(i), fmt.Sprintf("name%02d", i), "val"+strconv.Itoa(i)
		conn.Do("hmset", "rootKey", k1, "val1") //scan既可以扫描出字符串，对于集合和哈希等一样能够扫描出来
		conn.Do("hmset", "rootKey", k2, "val2")
	}
	//var pos int = 0
	//var res []string
	//xscan_(conn, Hscan, "rootKey", &pos, 10, 10, "name*", &res)
	////当hash对象所保存的键的数量小于512且key和value的长度都小于64个字节的时候hscan的count会失效
	////https://blog.csdn.net/dianxiaoer20111/article/details/120241141
	//sort.StringSlice(res).Sort()
	//fmt.Println("hscan res:", len(res), res)

}

func (h *Hash) Hash_hset(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Hset, h.key, "age", 23))

}

func (h *Hash) Hash_hget(name string) {

	w := &Wrapper{name}

	//int
	w.Run(conn.Do(Hset, h.key, "age", 23))
	ret := w.Run(conn.Do(Hget, h.key, "age"))

	s := string(ret.([]uint8))
	i, _ := strconv.Atoi(s)
	fmt.Printf("%s int %d\n", name, i)

	conn.Do(Del, h.key)

	//float
	w.Run(conn.Do(Hset, h.key, "age", 2.3))
	ret = w.Run(conn.Do(Hget, h.key, "age"))

	s = string(ret.([]uint8))
	f, _ := strconv.ParseFloat(s, 10)
	fmt.Printf("%s float %v\n", name, f)
}

func (h *Hash) Hash_hgetall(name string) {
	w := &Wrapper{name}
	w.Run(conn.Do(Hset, h.key, "age", 23))
	w.Run(conn.Do(Hset, h.key, "name", "qq"))
	w.Run(conn.Do(Hset, h.key, "phone", 123456))
	ret := w.Run(conn.Do(Hgetall, h.key))

	ss := ret.([]interface{})
	css, _ := redis.Strings(ss, nil)
	cmp, _ := redis.StringMap(ss, nil)
	fmt.Printf("%s strings %v stringMap %v\n", name, css, cmp)
}

func (h *Hash) Hash_hexists(name string) {
	w := &Wrapper{name}
	w.Run(conn.Do(Hset, h.key, "age", 23))
	w.Run(conn.Do(Hexists, h.key, "age"))
}

func (h *Hash) Hash_hlen(name string) {
	w := &Wrapper{name}
	w.Run(conn.Do(Hset, h.key, "age", 23))
	w.Run(conn.Do(Hlen, h.key))

	w.Run(conn.Do(Hset, h.key, "name", "qq"))
	w.Run(conn.Do(Hlen, h.key))
}

func (h *Hash) Hash_hmset(name string) {
	w := &Wrapper{name}

	//key的数量少
	w.Run(conn.Do(Hmset, h.key, "age", 23, "name", "qq", "phone", 123456))
	ret := w.Run(conn.Do(Hgetall, h.key))

	ss := ret.([]interface{})
	css, _ := redis.Strings(ss, nil)
	cmp, _ := redis.StringMap(ss, nil)
	fmt.Printf("%s strings %v stringMap %v\n", name, css, cmp)

	conn.Do(Del, h.key)

	//key的数量多
	args := []interface{}{}
	args = append(args, []interface{}{h.key, "age", 23, "name", "qq", "phone", 123456}...)
	w.Run(conn.Do(Hmset, args...))
	ret = w.Run(conn.Do(Hgetall, h.key))

	ss = ret.([]interface{})
	css, _ = redis.Strings(ss, nil)
	cmp, _ = redis.StringMap(ss, nil)
	fmt.Printf("%s strings %v stringMap %v\n", name, css, cmp)
}

func (h *Hash) Hash_hmget(name string) {
	w := &Wrapper{name}

	w.Run(conn.Do(Hmset, h.key, "age", 23, "name", "qq", "phone", 123456))

	args := []interface{}{}
	args = append(args, []interface{}{h.key, "age", "name", "phone"}...)
	ret := w.Run(conn.Do(Hmget, args...))

	ss := ret.([]interface{})
	css, _ := redis.Strings(ss, nil)
	cmp, _ := redis.StringMap(ss, nil) //参数不是偶数则无法转换完成
	fmt.Printf("%s strings %v stringMap %v\n", name, css, cmp)
}

func (h *Hash) Hash_hkeys(name string) {
	w := &Wrapper{name}
	w.Run(conn.Do(Hmset, h.key, "age", 23, "name", "qq", "phone", 123456))
	ret := w.Run(conn.Do(Hkeys, h.key))

	ss := ret.([]interface{})
	css, _ := redis.Strings(ss, nil)
	cmp, _ := redis.StringMap(ss, nil) //参数不是偶数则无法转换完成
	fmt.Printf("%s strings %v stringMap %v\n", name, css, cmp)
}

func (h *Hash) Hash_hvals(name string) {
	w := &Wrapper{name}
	w.Run(conn.Do(Hmset, h.key, "age", 23, "name", "qq", "phone", 123456))
	ret := w.Run(conn.Do(Hvals, h.key))

	ss := ret.([]interface{})
	css, _ := redis.Strings(ss, nil)
	cmp, _ := redis.StringMap(ss, nil) //参数不是偶数则无法转换完成
	fmt.Printf("%s strings %v stringMap %v\n", name, css, cmp)
}

func (h *Hash) Hash_hincrby(name string) {
	w := &Wrapper{name}
	w.Run(conn.Do(Hmset, h.key, "age", 23, "name", "qq", "phone", 123456))
	ret := w.Run(conn.Do(Hincrby, h.key, "age", 5))
	age, _ := redis.Int(ret, nil)

	ret = w.Run(conn.Do(Hincrby, h.key, "age", -15))
	aage, _ := redis.Int(ret, nil)

	fmt.Printf("%s pos %v neg %v\n", name, age, aage)
}

func (h *Hash) Hash_hincrbyfloat(name string) {
	w := &Wrapper{name}
	w.Run(conn.Do(Hmset, h.key, "age", 23, "name", "qq", "phone", 123456))
	ret := w.Run(conn.Do(Hincrbyfloat, h.key, "age", 0.5))
	age, _ := redis.Float64(ret, nil)

	fmt.Printf("%s float %v\n", name, age)
}

func (h *Hash) Hash_hdel(name string) {
	w := &Wrapper{name}
	w.Run(conn.Do(Hmset, h.key, "age", 23, "name", "qq", "phone", 123456, "address", "tokyo"))
	w.Run(conn.Do(Hdel, h.key, "age", "phone"))
	ret1 := w.Run(conn.Do(Hkeys, h.key))
	keys, _ := redis.Strings(ret1, nil)
	fmt.Printf("%s left keys %v\n", name, keys)

}
