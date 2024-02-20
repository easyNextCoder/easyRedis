package list

import (
	. "easyRedis/redisCom"
	"easyRedis/redisPool"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var conn = redisPool.GetRedis()

type List struct {
	funcMap map[string]func()
	key     string
}

func (h *List) GetFuncMap() map[string]func() {
	return h.funcMap
}

func (h *List) GetKey() string {
	return h.key
}

var list = &List{funcMap: make(map[string]func()), key: "user_list"}

func init() {
	BindFuncMap(list, conn)
}

func (l *List) List_rpush(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Rpush, l.key, "val1", "val2", 1, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)
}

func (l *List) List_lpush(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Lpush, l.key, "val1", "val2", 1, 2, 2, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)
}

func (l *List) List_linsert(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Rpush, l.key, "val1", "val2", 1, 2, 2, 2, 3, 99.9, -100))
	w.Run(conn.Do(Linsert, l.key, AFTER, "val2", "val3"))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	w.Run(conn.Do(Linsert, l.key, BEFORE, 99.9, 99.8))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	w.Run(conn.Do(Linsert, l.key, BEFORE, "99.8", "99.7"))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	w.Run(conn.Do(Linsert, l.key, BEFORE, "2", -1))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	w.Run(conn.Do(Linsert, l.key, AFTER, 2, "-1"))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

}

func (l *List) List_lrange(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Rpush, l.key, "val1", "val2", 1, 2, 2, 2, 3, 99.9, -100))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lrange, l.key, 0, 8))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lrange, l.key, 0, 100))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lrange, l.key, 0, 0))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lrange, l.key, 1, 1))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lrange, l.key, 1, 0))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss) //空
}

func (l *List) List_lpop(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Rpush, l.key, "val1", "val2", 1, 2, 2, 2, 3, 99.9, -100))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	var err error
	var val string
	for err == nil {
		val, err = redis.String(conn.Do(Lpop, l.key))
		if err != nil {
			fmt.Printf("%s lpop err %s\n", name, err)
			break
		}
		fmt.Printf("%s lpop val %s\n", name, val)

		val, err = redis.String(conn.Do(Rpop, l.key))
		if err != nil {
			fmt.Printf("%s rpop err %s\n", name, err)
			break
		}
		fmt.Printf("%s rpop val %s\n", name, val)
	}

}

func (l *List) List_rpop(name string) {
	l.List_lpop(name)
}

func (l *List) List_lrem(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Rpush, l.key, "val1", "val2", 1, 2, 2, 2, 2, 9, 9, 2, 2, 3, 99.9, -100, "1"))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lrem, l.key, 2, 2))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s %s %s %d %d ,strings %s\n", name, Lrem, l.key, 2, 2, ss)

	ret = w.Run(conn.Do(Lrem, l.key, 1, "2"))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s %s %s %d \"%s\" ,strings %s\n", name, Lrem, l.key, 1, "2", ss)

	ret = w.Run(conn.Do(Lrem, l.key, -1, 2))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s %s %s %d %d ,strings %s\n", name, Lrem, l.key, -1, 2, ss)

	ret = w.Run(conn.Do(Lrem, l.key, 0, 9))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s %s %s %d %d ,strings %s\n", name, Lrem, l.key, 0, 9, ss)

	ret = w.Run(conn.Do(Lrem, l.key, 0, 1)) //会把 1 和 "1" 全部删除掉
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s %s %s %d %d ,strings %s\n", name, Lrem, l.key, 0, 1, ss)

}

func (l *List) List_lindex(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Lpush, l.key, "val1", "val2", 1, 2, 2, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1)) //-1表示倒数第一个
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lindex, l.key, 0))
	ssnew, _ := redis.String(ret, nil)
	fmt.Printf("%s strings %s\n", name, ssnew)

}

func (l *List) List_lset(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Rpush, l.key, "val1", "val2", 1, 2, 2, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1))
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lset, l.key, 0, "val0"))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Lset, l.key, -1, "val_1"))
	ret = w.Run(conn.Do(Lrange, l.key, 0, -1))
	ss, _ = redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	//panic
	//ret = w.Run(conn.Do(Lset, l.key, 100, "val100"))
	//ret = w.Run(conn.Do(Lrange, l.key, 0, -1))
	//ss, _ = redis.Strings(ret, nil)
	//fmt.Printf("%s strings %s\n", name, ss)

}

func (l *List) List_llen(name string) {

	w := &Wrapper{name}

	w.Run(conn.Do(Rpush, l.key, "val1", "val2", 1, 2, 2, 2, 3, 99.9, -1))
	ret := w.Run(conn.Do(Lrange, l.key, 0, -1))
	ss, _ := redis.Strings(ret, nil)
	fmt.Printf("%s strings %s\n", name, ss)

	ret = w.Run(conn.Do(Llen, l.key))
	ssn, _ := redis.String(ret, nil)
	fmt.Printf("%s strings %s\n", name, ssn)

}
