package gocom

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		x interface{}
		y interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "max(1,2)=>2", args: args{x: 1, y: 2}, want: 2},
		{name: "max(1.1,1.2)=>1.2", args: args{x: 1.1, y: 1.2}, want: 1.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMin(t *testing.T) {
	type args struct {
		x interface{}
		y interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "min(1,2)=>1", args: args{x: 1, y: 2}, want: 1},
		{name: "min(1.1,1.2)=>1.1", args: args{x: 1.1, y: 1.2}, want: 1.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
