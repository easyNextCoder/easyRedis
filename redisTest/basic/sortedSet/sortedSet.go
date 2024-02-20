package sortedSet

import (
	. "easyRedis/redisCom"
	"easyRedis/redisPool"
	"easyRedis/util"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

var conn = redisPool.GetRedis()

type SortedSet struct {
	funcMap map[string]func()
	key     string
}

func (h *SortedSet) GetFuncMap() map[string]func() {
	return h.funcMap
}

func (h *SortedSet) GetKey() string {
	return h.key
}

var sortedSet = &SortedSet{funcMap: make(map[string]func()), key: "user_sortedset"}

func init() {
	BindFuncMap(sortedSet, conn)
}

func (l *SortedSet) SortedSet_zadd(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Zadd, l.key, 1, "val1", 3, "val2", 2, -100))
	ret := w.Run(conn.Do(Zrange, l.key, 0, 3))
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

}

func (l *SortedSet) SortedSet_zrange(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Zadd, l.key, 1, "val1", 3, "val2", 2, -100, 4, 25.6))
	ret := w.Run(conn.Do(Zrange, l.key, 0, 2))
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s a strings %s\n", name, ss)

	ret = w.Run(conn.Do(Zrange, l.key, 0, 3))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s b strings %s\n", name, ss)

	ret = w.Run(conn.Do(Zrange, l.key, -1, 5)) //-1表示倒数第一个成员，-2表示
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s c strings %s\n", name, ss)

	ret = w.Run(conn.Do(Zrange, l.key, -1, 5, "withscores")) //同时也返回对应的分数
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s e strings %s\n", name, ss)
}

func (l *SortedSet) SortedSet_zrangebylex(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Zadd, l.key, -2, "a", -1, "aa", 0, "abc", 0, "apple", 0, "b", 0, "c", 0, "d", 0, "d1", 0, "dd", 0, "dobble", 0, "z", 0, "z1", 1, "val1", 3, "val2", 2, -100, 4, 25.6))
	ret := w.Run(conn.Do(Zrangebylex, l.key, "-", "+")) //- +表示得分最小值和最大值
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s a strings %s\n", name, ss)

	ret = w.Run(conn.Do(Zrangebylex, l.key, "-", "+", "limit", 0, 3)) //获取分页数据，limit offset count
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s a strings %s\n", name, ss)

	ret = w.Run(conn.Do(Zrangebylex, l.key, "-", ""))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s a strings %s\n", name, ss)
}

func sortedSetq() {
	conn := redisPool.GetRedis()
	defer conn.Close()

	conn.Do(Del, "set1")
	conn.Do(Zadd, "set1", "2", "xyk", "5", "xyq", "3", "xyl", "1", "xyz", "7", "xym", "10", "xyu")
	val, _ := redis.Int(conn.Do(Zcard, "set1"))
	val1, _ := redis.Strings(conn.Do(Zrange, "set1", 0, -1, "withscores")) //放入的时候会自动排序,withscores可选
	fmt.Println("zadd zcard", val, val1)

	val2, _ := redis.Strings(conn.Do(Zrevrange, "set1", 0, -1))
	fmt.Println("zrevrange", val2)

	val3, _ := redis.Strings(conn.Do(Zrangebyscore, "set1", 1, 3))
	fmt.Println("zrangebyscore", val3)

	val4, _ := redis.Strings(conn.Do(Zrevrangebyscore, "set1", 3, 1))
	fmt.Println("zrevrangebyscore", val4)

	val5, _ := redis.Int(conn.Do(Zscore, "set1", "xyk"))
	fmt.Println("zscore", val5)

	val6, _ := redis.Int(conn.Do(Zrank, "set1", "xyl")) //不支持同时获取多个值
	val7, _ := redis.Int(conn.Do(Zrank, "set1", "xyz"))
	fmt.Println("zrank", val6, val7)

	val8, _ := redis.Int(conn.Do(Zrevrank, "set1", "xyz"))
	fmt.Println("zrevrand", val8)

	val9, _ := redis.Int(conn.Do(Zcount, "set1", 3, 7))
	fmt.Println("zscore", val9)

	conn.Do(Del, "set2")
	conn.Do(Del, "set3")
	conn.Do(Zadd, "set2", "2", "xyk", "5", "xyq", "3", "xyl", "1", "xyz", "7", "xym", "15", "xyu", "100", "xyy")
	conn.Do(Zadd, "set3", "2", "xyk", "5", "xyq", "3", "xyl", "1", "xyz", "7", "xym", "15", "xyu", "100", "xyp")

	conn.Do(Del, "des_set")
	val10, _ := redis.Int(conn.Do(Zinterstore, "des_set", 2, "set2", "set3", "weights", 0.5, 0.5)) //2代表连个两个集合的交集，weights是可选的，得到dest_set中元素的分数
	fmt.Println("zinterstore", val10)
	val11, _ := redis.Strings(conn.Do(Zrange, "des_set", 0, -1, "withscores"))
	fmt.Println("zinterstore", val11)

	conn.Do(Del, "des_set2")
	val12, _ := redis.Int(conn.Do(Zunionstore, "des_set2", 2, "set2", "set3", "weights", 1, 1)) //跟zinterstore的效果是一样的
	fmt.Println("zunionstore", val12)
	val13, _ := redis.Strings(conn.Do(Zrange, "des_set2", 0, -1, "withscores"))
	fmt.Println("zunionstore", val13)

	conn.Do(Del, "set3")
	conn.Do(Zadd, "set3", "2", "xyk", "5", "xyq", "3", "xyl")
	val14, _ := redis.Float64(conn.Do(Zincrby, "set3", -1, "xyk")) //注意几个参数的位置
	val15, _ := redis.String(conn.Do(Zscore, "set3", "xyk"))
	fmt.Println("zincrby", val14, val15)

	conn.Do(Del, "set3")
	conn.Do(Zadd, "set3", "2", "a", "5", "b", "3", "c", 4, "d", 6, "a")
	val16, _ := redis.Strings(conn.Do(Zrange, "set3", 0, -1, "withscores"))
	val17, _ := redis.Bool(conn.Do(Zrem, "set3", "a"))
	val18, _ := redis.Strings(conn.Do(Zrange, "set3", 0, -1))
	fmt.Println("zrem", val16, val17, val18)

	conn.Do(Del, "set3")
	conn.Do(Zadd, "set3", "2", "a", "5", "b", "3", "c", 4, "d", 6, "a")
	val19, _ := redis.Strings(conn.Do(Zrange, "set3", 0, -1, "withscores"))
	val20, _ := redis.Bool(conn.Do(Zremrangebyrank, "set3", 0, 2)) //下标从0开始，[0,2]
	val21, _ := redis.Strings(conn.Do(Zrange, "set3", 0, -1))
	fmt.Println("zremrangerank", val19, val20, val21)

	conn.Do(Del, "set3")
	conn.Do(Zadd, "set3", "2", "a", "5", "b", "3", "c", 4, "d", 6, "a")
	val22, _ := redis.Strings(conn.Do(Zrange, "set3", 0, -1, "withscores"))
	val23, _ := redis.Bool(conn.Do(Zremrangebyscore, "set3", 4, 6)) //下标从0开始，[0,2]
	val24, _ := redis.Strings(conn.Do(Zrange, "set3", 0, -1))
	fmt.Println("zremrangescore", val22, val23, val24)

	conn.Do(Del, "set4")
	conn.Do(Zadd, "set4", "2", "a", 2, "aa", "3", "b", "4", "c", 5, "d", 6, "e", 6, "f", 6, "g")

	val25, _ := redis.Strings(conn.Do(Zrange, "set4", 0, -1, "withscores"))
	val26, _ := redis.Bool(conn.Do(Zpopmax, "set4", 2)) //会先pop出同分最大的，之后找最大的，直到为空
	val27, _ := redis.Strings(conn.Do(Zrange, "set4", 0, -1))
	val28, _ := redis.Bool(conn.Do(Zpopmin, "set4", 20))
	val29, _ := redis.Strings(conn.Do(Zrange, "set4", 0, -1))
	fmt.Println("zpopmax zpopmin", val25, val26, val27, val28, val29)

	//zscan
	conn.Do("flushall")
	for i := 0; i < 1000; i++ {
		str := util.RandString(80)
		k1, _, k2, _ := fmt.Sprintf("key%02d%s", i, str), "val"+strconv.Itoa(i), fmt.Sprintf("name%02d", i), "val"+strconv.Itoa(i)
		conn.Do("zadd", "rootKey", i, k1) //这里分数的位置是在前面
		conn.Do("zadd", "rootKey", i+2, k2)
	}
	//var pos int = 0
	//var res []string
	//xscan_(conn, Zscan, "rootKey", &pos, 10, 10, "key*", &res)
	//sort.StringSlice(res).Sort()
	//fmt.Println("zscan res:", len(res), res)

}
