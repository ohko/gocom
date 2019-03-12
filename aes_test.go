package gocom

import (
	"bytes"
	"testing"
)

func TestAESCBC(t *testing.T) {
	keys := [][]byte{
		[]byte("1234567812345678"),
		[]byte("123456781234567812345678"),
		[]byte("12345678123456781234567812345678"),
	}
	msg := "hello"

	for _, key := range keys {
		en := AESCBCEncrypt(key, []byte(msg))
		de := AESCBCDecrypt(key, en)
		if bytes.Compare([]byte(msg), de) != 0 {
			t.Fail()
		}
	}
}

func TestAESEncode(t *testing.T) {
	keys := [][]byte{
		[]byte("1234567812345678"),
		[]byte("123456781234567812345678"),
		[]byte("12345678123456781234567812345678"),
	}
	iv := []byte("1234567812345678")
	msg := "hello"

	for _, key := range keys {
		en := AESEncode(key, iv, []byte(msg))
		de := AESEncode(key, iv, en)
		if bytes.Compare([]byte(msg), de) != 0 {
			t.Fail()
		}
	}
}
