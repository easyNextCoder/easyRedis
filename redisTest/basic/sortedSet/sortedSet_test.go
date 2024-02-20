package sortedSet

import (
	"log"
	"testing"
)

func Test_sortedSet(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "SortedSet_zadd"},
		{name: "SortedSet_zrange"},
		{name: "SortedSet_zrangebylex"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := sortedSet.funcMap[tt.name]
			if f != nil {
				f()
			} else {
				log.Panicf("can't find %s func, funcMap %+v", tt.name, sortedSet.funcMap)
			}
		})
	}
}
