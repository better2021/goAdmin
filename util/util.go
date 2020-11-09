package util

import (
	"io/ioutil"
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
	result := make([]byte,n)

	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}


//	获取客户ip地址
func GetClientIp() string {
	addrs,_ := net.InterfaceAddrs()

	for _,addres := range addrs{
		// 检查ip地址判断是否回环地址
		if ipnet,ok := addres.(*net.IPNet);ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
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

// 获取客户ip
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
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