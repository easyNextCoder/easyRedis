package set

import (
	"log"
	"testing"
)

func Test_set(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Set_sadd"},
		{name: "Set_sismember"},
		{name: "Set_srandmember"},
		{name: "Set_smembers"},
		{name: "Set_sdiff"},
		{name: "Set_sinter"},
		{name: "Set_sunion"},
		{name: "Set_spop"},
		{name: "Set_srem"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := set.funcMap[tt.name]
			if f != nil {
				f()
			} else {
				log.Panicf("can't find %s func, funcMap %+v", tt.name, set.funcMap)
			}
		})
	}
}
