# stats

_The definitive, overflow-safe, generic Slice Mathematics package for Go._

[![Go CI](https://github.com/kishan-thanki/stats/actions/workflows/go.yml/badge.svg)](https://github.com/kishan-thanki/stats/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/kishan-thanki/stats.svg)](https://pkg.go.dev/github.com/kishan-thanki/stats)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kishan-thanki/stats/blob/main/LICENSE)

Calculating the Average, Median, or Mode of an array of integers often requires writing boilerplate `for` loops, casting to `float64`, and handling manual overflow protection.

This package solves this problem by providing a strictly-typed, high-performance, and overflow-safe statistics package explicitly designed for Go slices, utilizing Go 1.21+ Generics.

## Features (The Production Gold)

- **Average**: Safe arithmetic mean for any numeric slice.
- **Median**: Safe median calculation that does not mutate your original slice.
- **Mode**: Frequency-mapped mode calculation for any comparable slice.
- **Aggregations**: Boilerplate-free `SumSlice`, `MinSlice`, and `MaxSlice`.

## Installation

Get the package using `go get`:

```bash
go get github.com/kishan-thanki/stats
```

## Usage

Because the repository _is_ the package, usage is completely flat and idiomatic.

```go
package main

import (
	"fmt"

	"github.com/kishan-thanki/stats"
)

func main() {
	// 1. Float64
	floatSlice := []float64{1.5, 2.5, 3.5}
	avgFloat, _ := stats.Average(floatSlice)
	fmt.Printf("Float64 Average: %.2f\n", avgFloat)

	// 2. Int64
	int64Slice := []int64{100, 200, 300}
	minInt64, _ := stats.MinSlice(int64Slice)
	maxInt64, _ := stats.MaxSlice(int64Slice)
	fmt.Printf("Int64 Min: %d | Max: %d\n", minInt64, maxInt64)

	// 3. Uint8
	uint8Slice := []uint8{255, 128, 64}
	medUint8, _ := stats.Median(uint8Slice)
	fmt.Printf("Uint8 Median: %.1f\n", medUint8)

	// 4. Strings
	stringSlice := []string{"a", "b", "a"}
	modeString, _ := stats.Mode(stringSlice)
	fmt.Printf("String Mode: %s\n", modeString)
}
```

For an exhaustive list of all 14 supported generic types (`int8`, `uint64`, `float32`, etc.), please see the [EXAMPLES.md](./EXAMPLES.md) file.

## API Reference

| Function   | Signature                                               | Notes                                                                            |
| ---------- | ------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Average`  | `Average[T IntegerOrFloat](slice []T) (float64, error)` | Returns the arithmetic mean; error if slice is empty                             |
| `Median`   | `Median[T IntegerOrFloat](slice []T) (float64, error)`  | Returns the median; makes a safe copy (does not mutate); error if slice is empty |
| `Mode`     | `Mode[T cmp.Ordered](slice []T) (T, error)`             | Returns the most frequent value; error if slice is empty                         |
| `SumSlice` | `SumSlice[T IntegerOrFloat](slice []T) (T, error)`      | Returns the sum of slice elements; panics on overflow; error if empty            |
| `MinSlice` | `MinSlice[T cmp.Ordered](slice []T) (T, error)`         | Returns the minimum element of a slice; error if empty                           |
| `MaxSlice` | `MaxSlice[T cmp.Ordered](slice []T) (T, error)`         | Returns the maximum element of a slice; error if empty                           |

## License

This project is licensed under the terms of the MIT License.
For more details, see the [LICENSE](./LICENSE) file.
