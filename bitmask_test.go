package bitmask

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	nums := []uint64{0, 2, 6, 23}
	var bound uint64 = 24
	mask, err := To(bound, nums)
	if err != nil {
		t.Fatalf("To(bound, nums) err = %v, expected nil", err)
	}
	masked, err := From(bound, mask)
	if err != nil {
		t.Fatalf("From(bound, mask) err = %v, expected nil", err)
	}
	if len(nums) != len(masked) {
		msg := "Length mismatch: input %d vs. output %d"
		t.Fatalf(msg, len(nums), len(masked))
	}
	// This test suggests that input is unique and sorted. Since we control
	// the input this is OK to assume.
	for i, num := range nums {
		if num != masked[i] {
			msg := "Mismatch at position %d: input %d vs. output %d"
			t.Errorf(msg, i, num, masked[i])
		}
	}
}

func TestToOutOfBound(t *testing.T) {
	nums := []uint64{3}
	var bound uint64 = 2
	_, err := To(bound, nums)
	expErr := "num (3) > bound (2)"
	if err == nil || err.Error() != expErr {
		t.Fatalf("To(bound, nums) err = %v, expected %q", err, expErr)
	}
}

func TestToBoundOver64(t *testing.T) {
	nums := []uint64{3}
	var bound uint64 = 65
	_, err := To(bound, nums)
	expErr := "bound (65) > 64"
	if err == nil || err.Error() != expErr {
		t.Fatalf("To(bound, nums) err = %v, expected %q", err, expErr)
	}
}

func TestFromBoundOver64(t *testing.T) {
	var mask uint64 = 0
	var bound uint64 = 65
	_, err := From(bound, mask)
	expErr := "bound (65) > 64"
	if err == nil || err.Error() != expErr {
		t.Fatalf("From(bound, mask) err = %v, expected %q", err, expErr)
	}
}
