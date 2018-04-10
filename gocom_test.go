package gocom

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	type args struct {
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"1", args{[]interface{}{1}}, 1},
		{"1,2", args{[]interface{}{1, 2}}, 1},
		{"3,1,2", args{[]interface{}{3, 1, 2}}, 1},
		{"3.0,1.0,2.0", args{[]interface{}{3.0, 1.0, 2.0}}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Min(tt.args.x...); !reflect.DeepEqual(got, tt.want) {
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

	c := 100
	m := make(map[int]int, c)
	for i := 0; i < c; i++ {
		n := RandIntn(0, c)
		if _, ok := m[n]; !ok {
			m[n] = 0
		}
		m[n]++
	}
	fmt.Println(m)
}

func TestMax(t *testing.T) {
	type args struct {
		x []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"1", args{[]interface{}{1}}, 1},
		{"1,2", args{[]interface{}{1, 2}}, 2},
		{"3,1,2", args{[]interface{}{3, 1, 2}}, 3},
		{"3.0,1.0,2.0", args{[]interface{}{3.0, 1.0, 2.0}}, 3.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Max(tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeGUID(t *testing.T) {
	c := 1000000
	m := make(map[string]bool, c)
	for i := 0; i < c; i++ {
		n := MakeGUID()
		if _, ok := m[n]; ok {
			t.Error()
		}
		m[n] = true
	}
}

func TestInArray(t *testing.T) {
	type args struct {
		a []interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{a: []interface{}{1, 2, 3}, b: 1}, true},
		{"2", args{a: []interface{}{1.1, 2.2, 3.3}, b: 1.1}, true},
		{"3", args{a: []interface{}{1.1, 2.2, 3.3}, b: 1.2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]byte("hk")}, "ae4171856a75f7b67d51fc0e1f95902e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.data); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}
