package more

import (
	. "easyRedis/redisCom"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

func xscan_(i interface{}, cmd, key string, pos *int, count int, n int, match string, res *[]string) {
	if n <= 0 {
		return
	}
	conn := i.(redis.Conn)
	//第一次扫描
	val11, err := conn.Do(cmd, key, *pos, "match", match, "count", count)
	if err != nil {
		fmt.Println(err)
	}
	//在这里解析结果
	replies := val11.([]interface{})
	repla1 := replies[0].([]byte)
	repla2 := replies[1].([]interface{})
	repla1final, _ := strconv.Atoi(string(repla1))
	repla2final := make([]string, len(repla2))
	for idx, v := range repla2 {
		repla2final[idx] = string(v.([]byte))
		*res = append(*res, repla2final[idx])
	}

	*pos = repla1final
	fmt.Println(repla1final, repla2final)
	xscan_(conn, cmd, key, pos, count, n-1, match, res)
}

func scan_(i interface{}, pos *int, count int, n int, match string, res *[]string) {
	if n <= 0 {
		return
	}
	conn := i.(redis.Conn)
	//第一次扫描
	val11, err := conn.Do(Scan, *pos, "match", match, "count", count)
	if err != nil {
		fmt.Println(err)
	}
	//在这里解析结果
	replies := val11.([]interface{})
	repla1 := replies[0].([]byte)
	repla2 := replies[1].([]interface{})
	repla1final, _ := strconv.Atoi(string(repla1))
	repla2final := make([]string, len(repla2))
	for idx, v := range repla2 {
		repla2final[idx] = string(v.([]byte))
		*res = append(*res, repla2final[idx])
	}

	*pos = repla1final

	fmt.Println(repla1final, repla2final)
	scan_(conn, pos, count, n-1, match, res)

}
