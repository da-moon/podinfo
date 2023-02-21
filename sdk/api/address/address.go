package address

import (
	"fmt"
	"net"
	"strconv"
	"sync"

	proto "github.com/da-moon/podinfo/sdk/api/proto"
	stacktrace "github.com/palantir/stacktrace"
)

// Address represents a network address
type Address struct {
	mutex sync.RWMutex
	Host  string `json:"host"`
	Port  string `json:"port"`
	proto proto.Proto
}

// TCPAddress parses the string representation of an address
// and returns a TCP address type
func TCPAddress(s string) (*Address, error) {
	host, port, err := net.SplitHostPort(s)
	if err != nil {
		return nil, err
	}
	_, err = strconv.Atoi(port)
	if err != nil {
		err = stacktrace.Propagate(err, "error parsing TCP address %s", s)
		return nil, err
	}
	result := &Address{
		Host:  host,
		Port:  port,
		proto: proto.TCP,
	}
	return result, nil
}

// UDPAddress parses the string representation of an address
// and returns a UDP address type
func UDPAddress(s string) (*Address, error) {
	host, port, err := net.SplitHostPort(s)
	if err != nil {
		return nil, err
	}
	_, err = strconv.Atoi(port)
	if err != nil {
		err = stacktrace.Propagate(err, "error parsing UDP address %s", s)
		return nil, err
	}
	result := &Address{
		Host:  host,
		Port:  port,
		proto: proto.UDP,
	}
	return result, nil
}

// Parse parses the string representation of an address.
// The address must be in the form "host:port".
// The host may be an Host address or a hostname.
// The port must be an integer.
func (a *Address) Parse() (net.Addr, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return Parse(a.proto, a.Host, a.Port)
}

// Parse parses Protocol and the string representation of an address
// and returns a net.Addr.
func Parse(p proto.Proto, host, port string) (net.Addr, error) {
	if p == proto.TCP {
		return net.ResolveTCPAddr(p.String(), net.JoinHostPort(host, port))
	}
	return net.ResolveUDPAddr(p.String(), net.JoinHostPort(host, port))
}

// ToString returns the string representation of the address
func (a *Address) ToString() string {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

// IsOpen checks if the network address is accessible
func (a *Address) IsOpen() bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	if a.proto == proto.Unknown {
		return false
	}
	ln, err := net.Listen(a.proto.String(), fmt.Sprintf("%s:%s", a.Host, a.Port))
	if ln != nil {
		defer ln.Close()
	}
	return err == nil
}
