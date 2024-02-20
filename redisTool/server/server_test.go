package server

import "testing"

func Test_serve1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "run server1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serve1()
		})
	}
}

func Test_serve2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "run server2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serve2()
		})
	}
}
