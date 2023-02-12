package port

import (
	"fmt"
	"net"
	"strconv"
	"sync"

	address "github.com/da-moon/northern-labs-interview/sdk/api/address"
	proto "github.com/da-moon/northern-labs-interview/sdk/api/proto"
	stacktrace "github.com/palantir/stacktrace"
)

const (
	minPort = 2048
	maxPort = 16384
)

// Port represents a port range
type Port struct {
	mutex   sync.Mutex
	minPort int
	maxPort int
	curPort int
}

// New returns a new Port struct
func New(opts ...Option) *Port {
	result := &Port{
		minPort: minPort,
		maxPort: maxPort,
	}
	for _, opt := range opts {
		opt(result)
	}
	result.curPort = result.minPort
	return result
}

// TCP returns a random TCP address
// with an open port in the range [p.minPort, p.maxPort]
func (p *Port) TCP() (*address.Address, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.port(proto.TCP)
}

// UDP returns a random UDP address
// with an open port in the range [p.minPort, p.maxPort]
func (p *Port) UDP() (*address.Address, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.port(proto.UDP)
}
func (p *Port) port(protocol proto.Proto) (*address.Address, error) {
	if protocol == proto.Unknown {
		err := stacktrace.NewError("unknown protocol")
		return nil, err
	}
	_, err := net.ResolveIPAddr("ip", "localhost")
	if err != nil {
		err = stacktrace.Propagate(err, "failed to resolve ip addr")
		return nil, err
	}
	result := &address.Address{}
	addrStr := fmt.Sprintf("localhost:%d", p.curPort)
	if protocol == proto.TCP {
		result, err = address.TCPAddress(addrStr)
		if err != nil {
			err = stacktrace.Propagate(err, "failed to resolve ip addr")
			return nil, err
		}
	} else {
		result, err = address.UDPAddress(addrStr)
		if err != nil {
			err = stacktrace.Propagate(err, "failed to resolve ip addr")
			return nil, err
		}
	}
	for !result.IsOpen() {
		p.curPort++
		if p.curPort > p.maxPort {
			err = stacktrace.NewError("all available ports to test have been exhausted")
			return nil, err
		}
		result.SetPort(strconv.Itoa(p.curPort)) //nolint:gosec // we do not care about the error as we just want to increase current port counter
	}
	return result, nil
}
