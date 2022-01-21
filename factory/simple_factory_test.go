package factory

import (
	"reflect"
	"testing"
)

func TestNewConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IConfigParser
	}{
		{
			name: "JSON Parser",
			args: args{t: "JSON"},
			want: JSONConfigParser{},
		},
		{
			name: "YAML Parser",
			args: args{t: "YAML"},
			want: YAMLConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
