package urandom

import (
	"fmt"

	"github.com/palantir/stacktrace"
)

const uuidLen = 16

// UUID generates a random UUID according to RFC 4122
func UUID() (string, error) {
	uuid, err := Bytes(uuidLen)
	if err != nil {
		err = stacktrace.Propagate(err, "could not generate uuid with length '%v'", uuidLen)
		return "", err
	}
	if len(uuid) != uuidLen {
		err := stacktrace.NewError("could not generate uuid due to wrong length byte slice (%d)", len(uuid))
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
