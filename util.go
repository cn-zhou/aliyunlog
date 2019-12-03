package aliyunlog

import (
	"net"
	"time"
)

func GetLocalIp() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && ipnet.IP.IsGlobalUnicast() {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetUnixMsToTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000")
}
