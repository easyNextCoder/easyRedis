package set

import (
	cmd "easyRedis/redisCom"
	"easyRedis/redisPool"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var conn = redisPool.GetRedis()

type Set struct {
	funcMap map[string]func()
	key     string
}

func (h *Set) GetFuncMap() map[string]func() {
	return h.funcMap
}

func (h *Set) GetKey() string {
	return h.key
}

var set = &Set{funcMap: make(map[string]func()), key: "user_set"}

func init() {
	cmd.BindFuncMap(set, conn)
}

func setq() {
	//sscan
	//conn.Do("flushall")
	//for i := 0; i < 1000; i++ {
	//	str := util.RandString(80)
	//	k1, _, k2, _ := fmt.Sprintf("key%02d%s", i, str), "val"+strconv.Itoa(i), fmt.Sprintf("name%02d", i), "val"+strconv.Itoa(i)
	//	conn.Do("sadd", "rootKey", k1)
	//	conn.Do("sadd", "rootKey", k2)
	//}
	//var pos int = 0
	//var res []string
	//xscan_(conn, Sscan, "rootKey", &pos, 10, 10, "name*", &res)
	//sort.StringSlice(res).Sort()
	//fmt.Println("sscan res:", len(res), res)

}

func (l *Set) Set_sadd(name string) {

	w := &cmd.Wrapper{name}

	w.Run(conn.Do(cmd.Sadd, l.key, "val1", "val2", 1, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(cmd.Scard, l.key))
	i, _ := redis.Int(ret, nil)
	fmt.Printf("%s strings %d\n", name, i)
}

func (l *Set) Set_sismember(name string) {

	w := &cmd.Wrapper{name}

	w.Run(conn.Do(cmd.Sadd, l.key, "val1", "val2", 1, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(cmd.Sismember, l.key, "99.9"))
	i, _ := redis.Int(ret, nil)
	fmt.Printf("%s strings %d\n", name, i)
}

func (l *Set) Set_srandmember(name string) {

	w := &cmd.Wrapper{name}

	w.Run(conn.Do(cmd.Sadd, l.key, "val1", "val2", 1, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(cmd.Srandmember, l.key, "2"))
	s, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, s)
}

func (l *Set) Set_smembers(name string) {

	w := &cmd.Wrapper{name}

	w.Run(conn.Do(cmd.Sadd, l.key, "val1", "val2", 1, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(cmd.Smembers, l.key))
	s, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, s)
}

func (l *Set) Set_sdiff(name string) {

	w := &cmd.Wrapper{name}

	key0 := l.key
	key1 := l.key + "1"
	key2 := l.key + "2"

	w.Run(conn.Do(cmd.Sadd, key0, "val1", "val2", 1, 2, 3, 99.9, -1))
	w.Run(conn.Do(cmd.Sadd, key1, "val1", "val2", -1, 100))
	w.Run(conn.Do(cmd.Sadd, key2, 2, 3))

	ret := w.Run(conn.Do(cmd.Sdiff, key0, key1, key2))
	s, _ := redis.Strings(ret, nil) //results: [1,99.9]
	fmt.Printf("%s strings %s\n", name, s)

	ret = w.Run(conn.Do(cmd.Smembers, key0))
	s, _ = redis.Strings(ret, nil)
	fmt.Printf("%s key0 now strings %s\n", name, s)

	ret = w.Run(conn.Do(cmd.Smembers, key1))
	s, _ = redis.Strings(ret, nil)
	fmt.Printf("%s key1 now strings %s\n", name, s)

	ret = w.Run(conn.Do(cmd.Smembers, key2))
	s, _ = redis.Strings(ret, nil)
	fmt.Printf("%s key2 now strings %s\n", name, s)

	conn.Do(cmd.Del, key0, key1, key2)
}

func (l *Set) Set_sinter(name string) {

	w := &cmd.Wrapper{name}

	key0 := l.key
	key1 := l.key + "1"
	key2 := l.key + "2"

	w.Run(conn.Do(cmd.Sadd, key0, "val1", "val2", 1, 2, 3, 99.9, -1))
	w.Run(conn.Do(cmd.Sadd, key1, "val1", "val2", -1, 100))
	w.Run(conn.Do(cmd.Sadd, key2, "val1", -1))

	ret := w.Run(conn.Do(cmd.Sinter, key0, key1, key2))
	s, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, s)

	conn.Do(cmd.Del, key0, key1, key2)
}

func (l *Set) Set_sunion(name string) {

	w := &cmd.Wrapper{name}

	key0 := l.key
	key1 := l.key + "1"
	key2 := l.key + "2"

	w.Run(conn.Do(cmd.Sadd, key0, "val1", "val2", 1, 2, 3, 99.9, -1))
	w.Run(conn.Do(cmd.Sadd, key1, "val1", "val2", -1, 100))
	w.Run(conn.Do(cmd.Sadd, key2, "val1", -1))

	ret := w.Run(conn.Do(cmd.Sunion, key0, key1, key2))
	s, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, s)

	conn.Do(cmd.Del, key0, key1, key2)
}

func (l *Set) Set_spop(name string) {

	w := &cmd.Wrapper{name}

	w.Run(conn.Do(cmd.Sadd, l.key, "val1", "val2", 1, 2, 3, 99.9, -1))

	ret := w.Run(conn.Do(cmd.Spop, l.key))
	s, _ := redis.String(ret, nil)
	fmt.Printf("%s strings %s\n", name, s)

}

func (l *Set) Set_srem(name string) {

	w := &cmd.Wrapper{name}

	w.Run(conn.Do(cmd.Sadd, l.key, "val1", "val2", 1, 2, 3, 99.9, -1))

	ret := w.Run(conn.Do(cmd.Srem, l.key, 1, "2", "-1", 3))
	ret = w.Run(conn.Do(cmd.Smembers, l.key))
	s, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, s)

}
