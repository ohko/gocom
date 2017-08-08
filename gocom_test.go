package gocom

import (
	"fmt"
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

func TestRandIntn(t *testing.T) {
	fmt.Print("0-0:")
	for i := 0; i < 10; i++ {
		fmt.Print(RandIntn(0, 0), ",")
	}
	fmt.Println()
	fmt.Print("0-1:")
	for i := 0; i < 10; i++ {
		fmt.Print(RandIntn(0, 1), ",")
	}
	fmt.Println()
	fmt.Print("0-2:")
	for i := 0; i < 10; i++ {
		fmt.Print(RandIntn(0, 2), ",")
	}
	fmt.Println()
	fmt.Print("1-2:")
	for i := 0; i < 10; i++ {
		fmt.Print(RandIntn(1, 2), ",")
	}
	fmt.Println()
	fmt.Print("1-3:")
	for i := 0; i < 10; i++ {
		fmt.Print(RandIntn(1, 3), ",")
	}
	fmt.Println()
}
