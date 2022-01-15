package strategy

import (
	"testing"
)

func TestNewCache(t *testing.T) {
	type args struct {
		e EvictionAlgo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testFIFO",
			args: args{FIFO{}},
		},
		{
			name: "testLRU",
			args: args{LRU{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCache(tt.args.e)
			got.EvictionAlgo.Evict()
		})
	}
}
