package urandom

import (
	"encoding/base64"
	"net"

	"github.com/palantir/stacktrace"
)

// String returns a URL-safe, base64 encoded
// securely generated random string.
func String(s int) (string, error) {

	b, err := Bytes(s)
	if err != nil {
		err = stacktrace.Propagate(err, "could not generate random string")
		return "", err
	}
	result := base64.URLEncoding.EncodeToString(b)
	if len(result) > s {
		result = result[:s]
	}
	return result, nil
}

// IPV4Address returns a valid IPv4 address as string
func IPV4Address() string {
	return ipAddr(net.IPv4len)
}

// IPV6Address returns a valid IPv6 address as net.IP
func IPV6Address() string {
	return ipAddr(net.IPv6len)
}
func ipAddr(length int) string {
	var ip net.IP
	for i := 0; i < length; i++ {
		number := uint8(Int(255))
		ip = append(ip, number)
	}
	return ip.String()
}
