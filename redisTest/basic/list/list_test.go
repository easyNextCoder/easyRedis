package list

import (
	"log"
	"testing"
)

func Test_list(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "List_rpush"},
		{name: "List_lpush"},
		{name: "List_linsert"},
		{name: "List_lrange"},
		{name: "List_lpop"},
		{name: "List_rpop"},
		{name: "List_lrem"},
		{name: "List_lindex"},
		{name: "List_lset"},
		{name: "List_llen"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := list.funcMap[tt.name]
			if f != nil {
				f()
			} else {
				log.Panicf("can't find %s func, funcMap %+v", tt.name, list.funcMap)
			}
		})
	}
}
