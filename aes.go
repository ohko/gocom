package gocom

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	pos := length - unpadding
	if pos < 0 || pos >= len(origData) {
		return origData
	}
	return origData[:pos]
}

// AESCBCEncrypt aes cbc encrypt
// key length: 16/24/32
func AESCBCEncrypt(key, data []byte) []byte {
	out := make([]byte, len(data))
	copy(out, data)
	out = pkcs5Padding(out, aes.BlockSize)
	c, _ := aes.NewCipher(key)
	cipher.NewCBCEncrypter(c, key[:aes.BlockSize]).
		CryptBlocks(out, out)
	return out
}

// AESCBCDecrypt aes cbc decrypt
// key length: 16/24/32
func AESCBCDecrypt(key, data []byte) []byte {
	out := make([]byte, len(data))
	copy(out, data)
	c, _ := aes.NewCipher(key)
	cipher.NewCBCDecrypter(c, key[:aes.BlockSize]).
		CryptBlocks(out, out)
	return pkcs5UnPadding(out)
}

// AESEncode aes encrypt/decrypt
// key length: 16/24/32
// iv length: 16
func AESEncode(key, iv, data []byte) []byte {
	block, _ := aes.NewCipher(key)
	buf := make([]byte, len(data))

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(buf, data)
	return buf
}
