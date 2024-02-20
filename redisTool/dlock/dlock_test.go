package dlock

import "testing"

func Test_runRedisDLock(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试自己写的分布式锁"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runRedisDLock()
		})
	}
}
