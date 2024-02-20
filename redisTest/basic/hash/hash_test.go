package hash

import (
	"log"
	"testing"
)

func Test_hash(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Hash_hset"},
		{name: "Hash_hget"},
		{name: "Hash_hgetall"},
		{name: "Hash_hexists"},
		{name: "Hash_hlen"},
		{name: "Hash_hmset"},
		{name: "Hash_hmget"},
		{name: "Hash_hkeys"},
		{name: "Hash_hvals"},
		{name: "Hash_hincrby"},
		{name: "Hash_hincrbyfloat"},
		{name: "Hash_hdel"},
		{name: "Hash_llen"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := hash.funcMap[tt.name]
			if f != nil {
				f()
			} else {
				log.Panicf("can't find %s func, funcMap %+v", tt.name, hash.funcMap)
			}
		})
	}
}
