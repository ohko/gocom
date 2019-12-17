package gocom

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

// MustInit some settings that are usually used
// - flag.Parse()
// - runtime.GOMAXPROCS(runtime.NumCPU)
// - show file's line in log
func MustInit() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Flags() | log.Lshortfile)
}

// Ternary /'tɜːnərɪ/ ternary operator
func Ternary(b bool, x, y interface{}) interface{} {
	if b {
		return x
	}
	return y
}

// IP2Int net.IP => uint32
// net.ParseIP("127.0.0.1") => 0x7F000001
func IP2Int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// Int2IP uint32 => net.IP(IPv4)
// 0x7F000001 => net.ParseIP("127.0.0.1")
func Int2IP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

// InArray check if b if in a
// type a must be []string/[]int/[]float64
func InArray(a interface{}, b interface{}) bool {
	switch a.(type) {
	case []string:
		for _, v := range a.([]string) {
			if v == b {
				return true
			}
		}
	case []int:
		for _, v := range a.([]int) {
			if v == b {
				return true
			}
		}
	case []float64:
		for _, v := range a.([]float64) {
			if v == b {
				return true
			}
		}
	default:
		panic("unkonw type")
	}
	return false
}

// MaxMin return max/min in arr
// arr type must be []int/[]float64
func MaxMin(arr interface{}, compare func(a, b interface{}) interface{}) interface{} {
	var max interface{}

	switch arr.(type) {
	case []int:
		max = arr.([]int)[0]
		for _, v := range arr.([]int) {
			max = compare(max, v)
		}
	case []float64:
		max = arr.([]float64)[0]
		for _, v := range arr.([]float64) {
			max = compare(max, v)
		}
	default:
		panic("unkonw type")
	}

	return max
}

// MakeGUID make GUID
func MakeGUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[8:10], b[6:8], b[4:6], b[10:])
}

// RandIntn return min <= x <= max
func RandIntn(min, max int) int {
	if min == 0 && max == 0 {
		return 0
	}
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max+1-min) + min
}

// Bmp1px http write 1px bmp picture
func Bmp1px(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "image/x-ms-bmp")
	w.Header().Set("Content-Length", "58")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Accept-Ranges", "bytes")
	w.Write([]byte{
		0x42, 0x4d, 0x3a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x28, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0x20, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
	})
}

// BasicAuth http basic auth
func BasicAuth(f http.HandlerFunc, user, pass string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
			}
		}()

		basicAuthPrefix := "Basic "

		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, basicAuthPrefix) {
			payload, err := base64.StdEncoding.DecodeString(
				auth[len(basicAuthPrefix):],
			)
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 &&
					bytes.Equal(pair[0], []byte(user)) &&
					bytes.Equal(pair[1], []byte(pass)) {
					f(w, r)
					return
				}
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// NewLogger ...
func NewLogger(logFileName string) (*log.Logger, error) {
	f, e := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if e != nil {
		return nil, e
	}
	return log.New(f, "", log.Ldate|log.Ltime), nil
}
