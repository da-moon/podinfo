package urandom

import (
	"reflect"
)

// SliceRand assign a random value from a string to the "r" arg
func SliceRand(s, r interface{}) {
	sl := reflect.ValueOf(s).Elem()
	// Handle types or exit if not supported type
	switch s.(type) {
	case *[]string:
		sli := sl.Interface().([]string)
		idx := Int(len(sli))
		*r.(*string) = sli[idx]
	case *[]int:
		sli := sl.Interface().([]int)
		idx := Int(len(sli))
		*r.(*int) = sli[idx]
	case *[]float64:
		sli := sl.Interface().([]float64)
		idx := Int(len(sli))
		*r.(*float64) = sli[idx]
	default:
		return
	}
}

// SliceRandString returns a random value from a string slice
func SliceRandString(s []string) string {
	idx := Int(len(s))
	return s[idx]
}

// SliceRandS is an alias of SliceRandString
func SliceRandS(s []string) string {
	return SliceRandString(s)
}

// SliceRandInt returns a random value from an int slice
func SliceRandInt(s []int) int {
	idx := Int(len(s))
	return s[idx]
}

// SliceRandI is an alias of SliceRandInt
func SliceRandI(s []int) int {
	return SliceRandInt(s)
}

// SliceRandFloat returns a random value from a float64 slice
func SliceRandFloat(s []float64) float64 {
	idx := Int(len(s))
	return s[idx]
}

// SliceRandF is an alias of SliceRandFloat
func SliceRandF(s []float64) float64 {
	return SliceRandFloat(s)
}
