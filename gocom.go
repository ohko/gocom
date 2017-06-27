package gocom

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
)

// Max 返回较大的值
func Max(x, y interface{}) interface{} {
	switch x.(type) {
	case int:
		if x.(int) > y.(int) {
			return x
		}
		return y
	case int32:
		if x.(int32) > y.(int32) {
			return x
		}
		return y
	case int64:
		if x.(int64) > y.(int64) {
			return x
		}
		return y
	case float32:
		if x.(float32) > y.(float32) {
			return x
		}
		return y
	case float64:
		if x.(float64) > y.(float64) {
			return x
		}
		return y
	default:
		panic("type error")
	}
}

// Min 返回较小的值
func Min(x, y interface{}) interface{} {
	switch x.(type) {
	case int:
		if x.(int) < y.(int) {
			return x
		}
		return y
	case int32:
		if x.(int32) < y.(int32) {
			return x
		}
		return y
	case int64:
		if x.(int64) < y.(int64) {
			return x
		}
		return y
	case float32:
		if x.(float32) < y.(float32) {
			return x
		}
		return y
	case float64:
		if x.(float64) < y.(float64) {
			return x
		}
		return y
	default:
		panic("type error")
	}
}

// Ternary 三目运算
func Ternary(b bool, x, y interface{}) interface{} {
	if b {
		return x
	}
	return y
}

// Type 获取对象的类型
func Type(o interface{}) string {
	switch o.(type) {
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case uint:
		return "uint"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "interface{}"
	}
}

// IP2Int ...
func IP2Int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// Int2IP ...
func Int2IP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

// InArrayInt ...
func InArrayInt(a []int, b int) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

// InArrayString ...
func InArrayString(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

// MakeGUID 生成唯一的GUID
func MakeGUID() string {
	b := make([]byte, 16)
	io.ReadFull(crand.Reader, b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[8:10], b[6:8], b[4:6], b[10:])
}

// Bmp1px ...
func Bmp1px(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "image/x-ms-bmp")
	w.Header().Set("Content-Length", "58")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Accept-Ranges", "bytes")
	w.Write([]byte{0x42, 0x4d, 0x3a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})
}
