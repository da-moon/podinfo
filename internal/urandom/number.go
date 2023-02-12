package urandom

import (
	"crypto/rand"
	"math"
	"math/big"
)

// Int64 returns a cryptographically secure random int
func Int64(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		// TODO: gracefully deal with this instead of panicing
		panic(err)
	}
	return nBig.Int64()
}
func Int(max int) int {
	return int(Int64(int64(max)))
}

// Integer returns a random integer value between lb(lowerbound) and ub(upperbound)
func Integer(lb, ub int) int {
	diff := ub - lb
	// [ NOTE ] ensuring lowerbound is actuall less than upperbound
	diff = int(math.Abs(float64(diff)))
	bias := Int(diff)
	return bias + lb
}
