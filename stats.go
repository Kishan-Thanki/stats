package stats

import (
	"cmp"
	"errors"

	"slices"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type IntegerOrFloat interface {
	Integer | Float
}

// Average returns the arithmetic mean of a slice of integers or floats.
// It returns an error if the slice is empty.
func Average[T IntegerOrFloat](slice []T) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot calculate average of an empty slice")
	}
	var sum float64
	for _, v := range slice {
		sum += float64(v)
	}
	return sum / float64(len(slice)), nil
}

// Median returns the median of a slice of integers or floats.
// It clones and sorts the slice internally to avoid modifying the input.
// It returns an error if the slice is empty.
func Median[T IntegerOrFloat](slice []T) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot calculate median of an empty slice")
	}
	cloned := slices.Clone(slice)
	slices.Sort(cloned)

	n := len(cloned)
	if n%2 == 1 {
		return float64(cloned[n/2]), nil
	}
	return (float64(cloned[n/2-1]) + float64(cloned[n/2])) / 2.0, nil
}

// Mode returns the value that appears most frequently in the slice.
// If multiple values have the same maximum frequency, it returns the one that appeared first.
// It returns an error if the slice is empty.
func Mode[T cmp.Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("cannot calculate mode of an empty slice")
	}
	counts := make(map[T]int)
	for _, v := range slice {
		counts[v]++
	}

	var mode T
	maxCount := 0
	for _, v := range slice {
		if counts[v] > maxCount {
			maxCount = counts[v]
			mode = v
		}
	}
	return mode, nil
}

// SumSlice returns the sum of a slice of integers or floats.
// It returns an error if the slice is empty, or if the sum overflows the bounds of the type T (for integer types).
func SumSlice[T IntegerOrFloat](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("cannot calculate sum of an empty slice")
	}

	var sum T
	var zero T
	for _, v := range slice {
		prev := sum
		sum += v

		if v > zero && sum < prev {
			return zero, errors.New("sum overflows type")
		}
		if v < zero && sum > prev {
			return zero, errors.New("sum underflows type")
		}
	}
	return sum, nil
}

// MinSlice returns the minimum value in a slice of ordered values.
// It returns an error if the slice is empty.
func MinSlice[T cmp.Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("cannot calculate minimum of an empty slice")
	}
	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// MaxSlice returns the maximum value in a slice of ordered values.
// It returns an error if the slice is empty.
func MaxSlice[T cmp.Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("cannot calculate maximum of an empty slice")
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}
