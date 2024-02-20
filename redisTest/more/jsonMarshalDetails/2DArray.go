package jsonMarshalDetails

import (
	"easyRedis/redisPool"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type AItem struct {
	Aid int   `json:"aid"`
	L   int   `json:"l"`
	V   int64 `json:"v"`
}

type two struct {
	Arr [][]int  `xorm:"arr" json:"arr"`
	As  []*AItem `json:"as"`
}

func twoDArrJsonMarshal() {
	conn := redisPool.GetRedis()
	defer conn.Close()

	var t two
	t.As = []*AItem{&AItem{
		Aid: 1,
		L:   2,
		V:   3,
	}}

	t.Arr = [][]int{{1, 0, 0}, {2, 1, 2}}
	by, err := json.Marshal(&t)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	conn.Do("hset", "achv", "9214", by)
	by1, err1 := redis.Bytes(conn.Do("hget", "achv", "9214"))
	if err1 != nil {
		fmt.Println("err1", err1)
		return
	}

	var t1 two

	err2 := json.Unmarshal(by1, &t1)
	if err2 != nil {
		fmt.Println("err1", err2)
		return
	}

	fmt.Println("res t1 is:", t1.Arr, t.As, t.As[0])

}
