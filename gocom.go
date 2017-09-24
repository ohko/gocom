package gocom

import (
	"bytes"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

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

// RandIntn 返回随机数：x <= n <= y
func RandIntn(x, y int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn((y-x)+1) + x
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

// BasicAuth ...
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
	if strings.Contains(logFileName, "_%s") {
		logFileName = fmt.Sprintf(logFileName, time.Now().Format("2006-01-02"))
	}
	f, e := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if e != nil {
		return nil, e
	}
	return log.New(f, "", log.Ldate|log.Ltime), nil
}

// OutJSON ...
func OutJSON(w io.Writer, r *http.Request, no int, msg interface{}) error {
	js, err := jsoniter.MarshalToString(map[string]interface{}{
		"no":  no,
		"msg": msg,
	})
	if err != nil {
		return err
	}

	r.ParseForm()
	callback := r.FormValue("callback")
	if callback != "" {
		fmt.Fprint(w, callback+"(")
		fmt.Fprint(w, js)
		fmt.Fprintln(w, ")")
	} else {
		fmt.Fprintln(w, js)
	}
	return nil
}
