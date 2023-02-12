package proto

import "strings"

// Proto enum type for network protocols
type Proto int

const (
	// TCP is the TCP protocol
	TCP Proto = iota
	// UDP is the TCP protocol
	UDP
	// Unknown is the 'catch all' option
	Unknown
)

// String returns the string representation of a Network Protocol type
func (s Proto) String() string {
	Proto := [...]string{"tcp", "udp", "unknown"}
	if len(Proto) < int(s) {
		return ""
	}
	return Proto[s]
}

// ToType converts a string to a Network Protocol type
func ToType(in string) Proto {
	in = strings.ToLower(in)
	in = strings.TrimSpace(in)
	switch in {
	case "tcp":
		return TCP
	case "udp":
		return UDP
	default:
		return Unknown
	}
}
