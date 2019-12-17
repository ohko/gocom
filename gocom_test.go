package gocom

import (
	"log"
	"net"
	"testing"
)

func TestIP2Int(t *testing.T) {
	if 0x7F000001 != IP2Int(net.ParseIP("127.0.0.1")) {
		t.Fail()
	}
}

func TestInt2IP(t *testing.T) {
	if "127.0.0.1" != Int2IP(0x7F000001).String() {
		t.Fail()
	}
}
func TestInArray(t *testing.T) {
	{
		a := []string{"a", "b"}
		b := "b"
		c := "c"
		if !InArray(a, b) {
			t.Fail()
		}
		if InArray(a, c) {
			t.Fail()
		}
	}

	{
		a := []int{1, 2}
		b := 2
		c := 3
		if !InArray(a, b) {
			t.Fail()
		}
		if InArray(a, c) {
			t.Fail()
		}
	}

	{
		a := []float64{1.1, 2.2}
		b := 2.2
		c := 3.3
		if !InArray(a, b) {
			t.Fail()
		}
		if InArray(a, c) {
			t.Fail()
		}
	}
}

func TestMaxMin(t *testing.T) {
	if MaxMin([]int{1, 2, 3}, func(a, b interface{}) interface{} {
		return Ternary(a.(int) > b.(int), a, b)
	}).(int) != 3 {
		t.Fail()
	}

	if MaxMin([]float64{1.2, 4.3, 2.3, 3.4}, func(a, b interface{}) interface{} {
		return Ternary(a.(float64) > b.(float64), a, b)
	}).(float64) != 4.3 {
		t.Fail()
	}

	if MaxMin([]int{1, 2, 3}, func(a, b interface{}) interface{} {
		return Ternary(a.(int) < b.(int), a, b)
	}).(int) != 1 {
		t.Fail()
	}

	if MaxMin([]float64{1.2, 4.3, 2.3, 3.4}, func(a, b interface{}) interface{} {
		return Ternary(a.(float64) < b.(float64), a, b)
	}).(float64) != 1.2 {
		t.Fail()
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

func TestRandIntn(t *testing.T) {
	log.Println(RandIntn(1, 100))
	count := 1000
	for i := 0; i < count; i++ {
		if RandIntn(0, 0) != 0 {
			t.Fail()
		}
	}
	for i := 0; i < count; i++ {
		if RandIntn(0, 1) < 0 || RandIntn(0, 1) > 1 {
			t.Fail()
		}
	}
	for i := 0; i < count; i++ {
		if RandIntn(0, 2) < 0 || RandIntn(0, 2) > 2 {
			t.Fail()
		}
	}
	for i := 0; i < count; i++ {
		if RandIntn(1, 2) < 1 || RandIntn(1, 2) > 2 {
			t.Fail()
		}
	}
	for i := 0; i < count; i++ {
		if RandIntn(1, 3) < 1 || RandIntn(1, 3) > 3 {
			t.Fail()
		}
	}

	c := 10000
	m := make(map[int]int, c)
	for i := 0; i < c; i++ {
		n := RandIntn(0, 10)
		if _, ok := m[n]; !ok {
			m[n] = 0
		}
		m[n]++
	}
	for i := 0; i < 10; i++ {
		if v, ok := m[i]; !ok || v == 0 {
			t.Fail()
		}
	}
}
