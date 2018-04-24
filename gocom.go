package gocom

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// MustInit 通常都会使用的设置
func MustInit() {
	flag.Parse()
	log.SetFlags(log.Flags() | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Max 返回较大的值
func Max(x ...interface{}) (interface{}, error) {
	if len(x) == 0 {
		return nil, errors.New("size error")
	}

	switch x[0].(type) {
	case int:
		max := x[0].(int)
		for _, v := range x {
			if v.(int) > max {
				max = v.(int)
			}
		}
		return max, nil
	case int32:
		max := x[0].(int32)
		for _, v := range x {
			if v.(int32) > max {
				max = v.(int32)
			}
		}
		return max, nil
	case int64:
		max := x[0].(int64)
		for _, v := range x {
			if v.(int64) > max {
				max = v.(int64)
			}
		}
		return max, nil
	case float32:
		max := x[0].(float32)
		for _, v := range x {
			if v.(float32) > max {
				max = v.(float32)
			}
		}
		return max, nil
	case float64:
		max := x[0].(float64)
		for _, v := range x {
			if v.(float64) > max {
				max = v.(float64)
			}
		}
		return max, nil
	default:
		return nil, errors.New("type error")
	}
}

// Min 返回较小的值
func Min(x ...interface{}) (interface{}, error) {
	if len(x) == 0 {
		return nil, errors.New("size error")
	}

	switch x[0].(type) {
	case int:
		min := x[0].(int)
		for _, v := range x {
			if v.(int) < min {
				min = v.(int)
			}
		}
		return min, nil
	case int32:
		min := x[0].(int32)
		for _, v := range x {
			if v.(int32) < min {
				min = v.(int32)
			}
		}
		return min, nil
	case int64:
		min := x[0].(int64)
		for _, v := range x {
			if v.(int64) < min {
				min = v.(int64)
			}
		}
		return min, nil
	case float32:
		min := x[0].(float32)
		for _, v := range x {
			if v.(float32) < min {
				min = v.(float32)
			}
		}
		return min, nil
	case float64:
		min := x[0].(float64)
		for _, v := range x {
			if v.(float64) < min {
				min = v.(float64)
			}
		}
		return min, nil
	default:
		return nil, errors.New("type error")
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

// InArrayString ...
func InArrayString(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
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

// InArrayFloat ...
func InArrayFloat(a []float64, b float64) bool {
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
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[8:10], b[6:8], b[4:6], b[10:])
}

// RandIntn 返回随机数：x <= n <= y
func RandIntn(x, y int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn((y-x)+1) + x
}

// Bmp1px 1像素bmp图片
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

// BasicAuth 简单验证
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
func OutJSON(w http.ResponseWriter, r *http.Request, no int, msg interface{}, gz bool) error {
	w.Header().Set("Content-Type", "application/json")
	var ww io.Writer
	if gz {
		w.Header().Set("Content-Encoding", "gzip")
		g := gzip.NewWriter(w)
		ww = g
		defer g.Close()
	} else {
		ww = w
	}

	bs, err := json.Marshal(map[string]interface{}{
		"no":  no,
		"msg": msg,
	})
	if err != nil {
		return err
	}

	r.ParseForm()
	callback := r.FormValue("callback")
	if callback != "" {
		fmt.Fprint(ww, callback+"(")
		fmt.Fprint(ww, string(bs))
		fmt.Fprintln(ww, ")")
	} else {
		fmt.Fprintln(ww, string(bs))
	}
	return nil
}

// Md5 ...
func Md5(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

// MakeSSLFile 生成域名证书
// openssl genrsa -des3 -passout pass:123456 -out ssl.pass.key 2048
// openssl rsa -passin pass:123456 -in ssl.pass.key -out ssl.key
// rm -rf ssl.pass.key
// openssl req -new -subj "/C=US/ST=Mars/L=iTranswarp/O=iTranswarp/OU=iTranswarp/CN=*.com" -key ssl.key -out ssl.csr
// openssl x509 -req -days 365 -in ssl.csr -signkey ssl.key -out ssl.crt
// rm -rf ssl.csr
//
// 校验
// openssl x509 -noout -modulus -in ssl.crt | openssl md5
// openssl rsa -noout -modulus -in ssl.key | openssl md5
func MakeSSLFile(origin string) {
	exec.Command(`openssl`, `genrsa`, `-des3`, `-passout`, `pass:123456`, `-out`, origin+`.pass.key`, `2048`).CombinedOutput()
	exec.Command("openssl", `rsa`, `-passin`, `pass:123456`, `-in`, origin+`.pass.key`, `-out`, origin+`.key`).CombinedOutput()
	exec.Command("rm", origin+`.pass.key`).CombinedOutput()
	exec.Command("openssl", `req`, `-new`, `-subj`, `/C=US/ST=Mars/L=iTranswarp/O=iTranswarp/OU=iTranswarp/CN=`+origin, `-key`, origin+`.key`, `-out`, origin+`.csr`).CombinedOutput()
	exec.Command("openssl", `x509`, `-req`, `-days`, `365`, `-in`, origin+`.csr`, `-signkey`, origin+`.key`, `-out`, origin+`.crt`).CombinedOutput()
	exec.Command("rm", origin+`.csr`).CombinedOutput()
}

// MakeTLSFile 生成TLS双向认证证书
// # 1.创建根证书密钥文件(自己做CA)root.key：
// openssl genrsa -des3 -passout pass:123 -out ssl/root.key 2048
// # 2.创建根证书的申请文件root.csr：
// openssl req -passin pass:123 -new -subj "/C=CN/ST=Shanghai/L=Shanghai/O=MyCompany/OU=MyCompany/CN=localhost/emailAddress=hk@cdeyun.com" -key ssl/root.key -out ssl/root.csr
// # 3.创建根证书root.crt：
// openssl x509 -passin pass:123 -req -days 3650 -sha256 -extensions v3_ca -signkey ssl/root.key -in ssl/root.csr -out ssl/root.crt
// rm -rf ssl/root.csr
//
// # 1.创建客户端证书私钥
// openssl genrsa -des3 -passout pass:456 -out ssl/ssl.key 2048
// # 2.去除key口令
// openssl rsa -passin pass:456 -in ssl/ssl.key -out ssl/ssl.key
// # 3.创建客户端证书申请文件ssl.csr
// openssl req -new -subj "/C=CN/ST=Shanghai/L=Shanghai/O=MyCompany/OU=MyCompany/CN=localhost/emailAddress=hk@cdeyun.com" -key ssl/ssl.key -out ssl/ssl.csr
// # 4.创建客户端证书文件ssl.crt
// openssl x509 -passin pass:123 -req -days 365 -sha256 -extensions v3_req -CA ssl/root.crt -CAkey ssl/root.key -CAcreateserial -in ssl/ssl.csr -out ssl/ssl.crt
// rm -rf ssl/ssl.csr
// rm -rf ssl/root.srl
// # 5.将客户端证书文件ssl.crt和客户端证书密钥文件ssl.key合并成客户端证书安装包ssl.pfx
// openssl pkcs12 -export -passout pass:789 -in ssl/ssl.crt -inkey ssl/ssl.key -out ssl/ssl.pfx
func MakeTLSFile(passRoot, passKey, passPfx, domain, email string) bool {
	domain = "ssl/" + domain
	os.Mkdir("ssl", 0755)
	// 1.1.创建根证书密钥文件(自己做CA)root.key：
	exec.Command(`openssl`, `genrsa`, `-des3`, `-passout`, `pass:`+passRoot, `-out`, domain+`.root.key`, `2048`).CombinedOutput()

	// 1.2.创建根证书的申请文件root.csr：
	exec.Command(`openssl`, `req`, `-passin`, `pass:`+passRoot, `-new`, `-subj`, `/C=CN/ST=Shanghai/L=Shanghai/O=MyCompany/OU=MyCompany/CN=`+domain+`/emailAddress=`+email, `-key`, domain+`.root.key`, `-out`, domain+`.root.csr`).CombinedOutput()

	// 1.3.创建根证书root.crt：
	exec.Command(`openssl`, `x509`, `-passin`, `pass:`+passRoot, `-req`, `-days`, `3650`, `-sha256`, `-extensions`, `v3_ca`, `-signkey`, domain+`.root.key`, `-in`, domain+`.root.csr`, `-out`, domain+`.root.crt`).CombinedOutput()
	exec.Command(`rm`, domain+`.root.csr`).CombinedOutput()

	// 2.1.创建客户端证书私钥
	exec.Command(`openssl`, `genrsa`, `-des3`, `-passout`, `pass:`+passKey, `-out`, domain+`.ssl.key`, `2048`).CombinedOutput()

	// 2.2.去除key口令
	exec.Command(`openssl`, `rsa`, `-passin`, `pass:`+passKey, `-in`, domain+`.ssl.key`, `-out`, domain+`.ssl.key`).CombinedOutput()

	// 2.3.创建客户端证书申请文件ssl.csr
	exec.Command(`openssl`, `req`, `-new`, `-subj`, `/C=US/ST=Mars/L=iTranswarp/O=iTranswarp/OU=iTranswarp/CN=`+domain+`/emailAddress=`+email, `-key`, domain+`.ssl.key`, `-out`, domain+`.ssl.csr`).CombinedOutput()

	// 2.4.创建客户端证书文件ssl.crt
	exec.Command(`openssl`, `x509`, `-passin`, `pass:`+passRoot, `-req`, `-days`, `365`, `-sha256`, `-extensions`, `v3_req`, `-CA`, domain+`.root.crt`, `-CAkey`, domain+`.root.key`, `-CAcreateserial`, `-in`, domain+`.ssl.csr`, `-out`, domain+`.ssl.crt`).CombinedOutput()
	exec.Command(`rm`, domain+`.ssl.csr`).CombinedOutput()

	// 2.5.将客户端证书文件ssl.crt和客户端证书密钥文件ssl.key合并成客户端证书安装包ssl.pfx
	exec.Command(`openssl`, `pkcs12`, `-export`, `-passout`, `pass:`+passPfx, `-in`, domain+`.ssl.crt`, `-inkey`, domain+`.ssl.key`, `-out`, domain+`.ssl.pfx`).CombinedOutput()
	exec.Command(`rm`, domain+`.srl`).CombinedOutput()

	// 3.校验
	bs1, _ := exec.Command(`openssl`, `x509`, `-noout`, `-modulus`, `-in`, domain+`.ssl.crt`).CombinedOutput()
	bs2, _ := exec.Command(`openssl`, `rsa`, `-noout`, `-modulus`, `-in`, domain+`.ssl.key`).CombinedOutput()
	return string(bs1) == string(bs2)
}
