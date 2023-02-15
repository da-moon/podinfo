package value

import (
	"flag"
	"fmt"
	"strconv"
)

var _ flag.Value = &Int{}

// Int provides a flag value that's aware if it has been set.
type Int struct {
	v *int
}

// Merge will overlay this value if it has been set.
func (u *Int) Merge(onto *int) {
	if u.v != nil {
		*onto = *(u.v)
	}
}

// Set implements the flag.Value interface.
func (u *Int) Set(v string) error {
	parsed, err := strconv.ParseInt(v, 0, 64)
	u.RawSet(int(parsed))
	return err
}

// RawSet sets the underlying value directly
func (u *Int) RawSet(v int) {
	if u.v == nil {
		u.v = new(int)
	}
	*(u.v) = v
}

// Get returns the actual underlying value
func (u *Int) Get() int {
	var result int
	if u.v != nil {
		result = *(u.v)
	}
	return result
}

// String implements the flag.Value interface.
func (u *Int) String() string {
	return fmt.Sprintf("%v", u.Get())
}
