// Package bitmask offers a method of turning an array of integers (between 0
// and 64) into a single unsigned integer representing their bitmask. It also
// provides a reverse operation.
package bitmask

import "fmt"

// To turns an array of unsigned integers (between 0 and 64) into a single
// unsigned integer representing their bitmask. Repeated integers have no
// further effect.
func To(bound uint64, nums []uint64) (uint64, error) {
	if bound > 64 {
		return 0, fmt.Errorf("bound (%d) > 64", bound)
	}
	var result uint64 = 0
	for _, num := range nums {
		if num > (bound - 1) {
			return 0, fmt.Errorf("num (%d) > bound (%d)", num, bound)
		}
		result = result | (1 << num)
	}
	return result, nil
}

// From turns a single unsigned integer (a bitmask) into an array of unsigned
// integers between 0 and 64.
func From(bound, mask uint64) ([]uint64, error) {
	if bound > 64 {
		return nil, fmt.Errorf("bound (%d) > 64", bound)
	}
	result := make([]uint64, 0)
	var i uint64 = 0
	for ; i < bound; i++ {
		if (mask & (1 << i)) != 0 {
			result = append(result, i)
		}
	}
	return result, nil
}
