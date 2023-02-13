package liveness

// Status is an enum representing accepted status values for
// this request handler
type Status int

const (
	// OK represents the OK (healthy) response
	OK Status = iota
	// Unknown is the 'catch all' option
	Unknown
)

// String function returns a string representation of
// Status enum
func (s Status) String() string {
	Status := [...]string{"OK", "unknown"}
	if len(Status) < int(s) {
		return ""
	}
	return Status[s]
}
