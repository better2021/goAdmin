package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

// 导出随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmASDFGHJKLZCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//	获取客户ip地址
func GetClientIp() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addres := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addres.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return "Can not find the client ip address!"
}

//	获取服务端ip
func GetServerIP() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}

// ClientPublicIP 获取公网ip
func ClientPublicIP() string {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// 获取星期
func Getweek() string {
	var datStr string
	day := time.Now().Weekday()
	switch day {
	case 0:
		datStr = "星期天"
	case 1:
		datStr = "星期一"
	case 2:
		datStr = "星期二"
	case 3:
		datStr = "星期三"
	case 4:
		datStr = "星期四"
	case 5:
		datStr = "星期五"
	case 6:
		datStr = "星期六"
	}
	return datStr
}

/**
crypto的加解密
*/
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

// AesEncrypt 加密函数
func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	text := base64.StdEncoding.EncodeToString(plaintext)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err, "err")
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(text))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	text, _ := base64.StdEncoding.DecodeString(string(ciphertext))
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err, "err")
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(text))
	blockMode.CryptBlocks(origData, text)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
