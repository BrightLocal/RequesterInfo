package ip

import (
	"net"
	"net/http"
)

func GetIp(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return ip
	}
	return ""
}
