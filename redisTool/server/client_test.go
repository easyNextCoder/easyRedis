package server

import "testing"

func Test_grabRedBag(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "运行抢红包"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grabRedBag()
		})
	}
}
