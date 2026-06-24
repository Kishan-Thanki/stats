package stats_test

import (
	"fmt"
	"github.com/kishan-thanki/stats"
)

func ExampleAverage() {
	intSlice := []int{1, 2, 3}
	res_int, _ := stats.Average(intSlice)
	fmt.Printf("int:       %v\n", res_int)
	int8Slice := []int8{1, 2, 3}
	res_int8, _ := stats.Average(int8Slice)
	fmt.Printf("int8:      %v\n", res_int8)
	int16Slice := []int16{1, 2, 3}
	res_int16, _ := stats.Average(int16Slice)
	fmt.Printf("int16:     %v\n", res_int16)
	int32Slice := []int32{1, 2, 3}
	res_int32, _ := stats.Average(int32Slice)
	fmt.Printf("int32:     %v\n", res_int32)
	int64Slice := []int64{1, 2, 3}
	res_int64, _ := stats.Average(int64Slice)
	fmt.Printf("int64:     %v\n", res_int64)
	uintSlice := []uint{1, 2, 3}
	res_uint, _ := stats.Average(uintSlice)
	fmt.Printf("uint:      %v\n", res_uint)
	uint8Slice := []uint8{1, 2, 3}
	res_uint8, _ := stats.Average(uint8Slice)
	fmt.Printf("uint8:     %v\n", res_uint8)
	uint16Slice := []uint16{1, 2, 3}
	res_uint16, _ := stats.Average(uint16Slice)
	fmt.Printf("uint16:    %v\n", res_uint16)
	uint32Slice := []uint32{1, 2, 3}
	res_uint32, _ := stats.Average(uint32Slice)
	fmt.Printf("uint32:    %v\n", res_uint32)
	uint64Slice := []uint64{1, 2, 3}
	res_uint64, _ := stats.Average(uint64Slice)
	fmt.Printf("uint64:    %v\n", res_uint64)
	uintptrSlice := []uintptr{1, 2, 3}
	res_uintptr, _ := stats.Average(uintptrSlice)
	fmt.Printf("uintptr:   %v\n", res_uintptr)
	float32Slice := []float32{1.5, 2.5, 3.5}
	res_float32, _ := stats.Average(float32Slice)
	fmt.Printf("float32:   %v\n", res_float32)
	float64Slice := []float64{1.5, 2.5, 3.5}
	res_float64, _ := stats.Average(float64Slice)
	fmt.Printf("float64:   %v\n", res_float64)
	// Output:
	// int:       2
	// int8:      2
	// int16:     2
	// int32:     2
	// int64:     2
	// uint:      2
	// uint8:     2
	// uint16:    2
	// uint32:    2
	// uint64:    2
	// uintptr:   2
	// float32:   2.5
	// float64:   2.5
}

func ExampleMedian() {
	intSlice := []int{3, 1, 2}
	res_int, _ := stats.Median(intSlice)
	fmt.Printf("int:       %v\n", res_int)
	int8Slice := []int8{3, 1, 2}
	res_int8, _ := stats.Median(int8Slice)
	fmt.Printf("int8:      %v\n", res_int8)
	int16Slice := []int16{3, 1, 2}
	res_int16, _ := stats.Median(int16Slice)
	fmt.Printf("int16:     %v\n", res_int16)
	int32Slice := []int32{3, 1, 2}
	res_int32, _ := stats.Median(int32Slice)
	fmt.Printf("int32:     %v\n", res_int32)
	int64Slice := []int64{3, 1, 2}
	res_int64, _ := stats.Median(int64Slice)
	fmt.Printf("int64:     %v\n", res_int64)
	uintSlice := []uint{3, 1, 2}
	res_uint, _ := stats.Median(uintSlice)
	fmt.Printf("uint:      %v\n", res_uint)
	uint8Slice := []uint8{3, 1, 2}
	res_uint8, _ := stats.Median(uint8Slice)
	fmt.Printf("uint8:     %v\n", res_uint8)
	uint16Slice := []uint16{3, 1, 2}
	res_uint16, _ := stats.Median(uint16Slice)
	fmt.Printf("uint16:    %v\n", res_uint16)
	uint32Slice := []uint32{3, 1, 2}
	res_uint32, _ := stats.Median(uint32Slice)
	fmt.Printf("uint32:    %v\n", res_uint32)
	uint64Slice := []uint64{3, 1, 2}
	res_uint64, _ := stats.Median(uint64Slice)
	fmt.Printf("uint64:    %v\n", res_uint64)
	uintptrSlice := []uintptr{3, 1, 2}
	res_uintptr, _ := stats.Median(uintptrSlice)
	fmt.Printf("uintptr:   %v\n", res_uintptr)
	float32Slice := []float32{1.5, 2.5, 3.5}
	res_float32, _ := stats.Median(float32Slice)
	fmt.Printf("float32:   %v\n", res_float32)
	float64Slice := []float64{1.5, 2.5, 3.5}
	res_float64, _ := stats.Median(float64Slice)
	fmt.Printf("float64:   %v\n", res_float64)
	// Output:
	// int:       2
	// int8:      2
	// int16:     2
	// int32:     2
	// int64:     2
	// uint:      2
	// uint8:     2
	// uint16:    2
	// uint32:    2
	// uint64:    2
	// uintptr:   2
	// float32:   2.5
	// float64:   2.5
}

func ExampleMode() {
	intSlice := []int{1, 2, 1}
	res_int, _ := stats.Mode(intSlice)
	fmt.Printf("int:       %v\n", res_int)
	int8Slice := []int8{1, 2, 1}
	res_int8, _ := stats.Mode(int8Slice)
	fmt.Printf("int8:      %v\n", res_int8)
	int16Slice := []int16{1, 2, 1}
	res_int16, _ := stats.Mode(int16Slice)
	fmt.Printf("int16:     %v\n", res_int16)
	int32Slice := []int32{1, 2, 1}
	res_int32, _ := stats.Mode(int32Slice)
	fmt.Printf("int32:     %v\n", res_int32)
	int64Slice := []int64{1, 2, 1}
	res_int64, _ := stats.Mode(int64Slice)
	fmt.Printf("int64:     %v\n", res_int64)
	uintSlice := []uint{1, 2, 1}
	res_uint, _ := stats.Mode(uintSlice)
	fmt.Printf("uint:      %v\n", res_uint)
	uint8Slice := []uint8{1, 2, 1}
	res_uint8, _ := stats.Mode(uint8Slice)
	fmt.Printf("uint8:     %v\n", res_uint8)
	uint16Slice := []uint16{1, 2, 1}
	res_uint16, _ := stats.Mode(uint16Slice)
	fmt.Printf("uint16:    %v\n", res_uint16)
	uint32Slice := []uint32{1, 2, 1}
	res_uint32, _ := stats.Mode(uint32Slice)
	fmt.Printf("uint32:    %v\n", res_uint32)
	uint64Slice := []uint64{1, 2, 1}
	res_uint64, _ := stats.Mode(uint64Slice)
	fmt.Printf("uint64:    %v\n", res_uint64)
	uintptrSlice := []uintptr{1, 2, 1}
	res_uintptr, _ := stats.Mode(uintptrSlice)
	fmt.Printf("uintptr:   %v\n", res_uintptr)
	float32Slice := []float32{1.5, 2.5, 3.5}
	res_float32, _ := stats.Mode(float32Slice)
	fmt.Printf("float32:   %v\n", res_float32)
	float64Slice := []float64{1.5, 2.5, 3.5}
	res_float64, _ := stats.Mode(float64Slice)
	fmt.Printf("float64:   %v\n", res_float64)
	stringSlice := []string{"apple", "banana", "apple"}
	res_string, _ := stats.Mode(stringSlice)
	fmt.Printf("string:    %v\n", res_string)
	// Output:
	// int:       1
	// int8:      1
	// int16:     1
	// int32:     1
	// int64:     1
	// uint:      1
	// uint8:     1
	// uint16:    1
	// uint32:    1
	// uint64:    1
	// uintptr:   1
	// float32:   1.5
	// float64:   1.5
	// string:    apple
}

func ExampleSumSlice() {
	intSlice := []int{1, 2, 3}
	res_int, _ := stats.SumSlice(intSlice)
	fmt.Printf("int:       %v\n", res_int)
	int8Slice := []int8{1, 2, 3}
	res_int8, _ := stats.SumSlice(int8Slice)
	fmt.Printf("int8:      %v\n", res_int8)
	int16Slice := []int16{1, 2, 3}
	res_int16, _ := stats.SumSlice(int16Slice)
	fmt.Printf("int16:     %v\n", res_int16)
	int32Slice := []int32{1, 2, 3}
	res_int32, _ := stats.SumSlice(int32Slice)
	fmt.Printf("int32:     %v\n", res_int32)
	int64Slice := []int64{1, 2, 3}
	res_int64, _ := stats.SumSlice(int64Slice)
	fmt.Printf("int64:     %v\n", res_int64)
	uintSlice := []uint{1, 2, 3}
	res_uint, _ := stats.SumSlice(uintSlice)
	fmt.Printf("uint:      %v\n", res_uint)
	uint8Slice := []uint8{1, 2, 3}
	res_uint8, _ := stats.SumSlice(uint8Slice)
	fmt.Printf("uint8:     %v\n", res_uint8)
	uint16Slice := []uint16{1, 2, 3}
	res_uint16, _ := stats.SumSlice(uint16Slice)
	fmt.Printf("uint16:    %v\n", res_uint16)
	uint32Slice := []uint32{1, 2, 3}
	res_uint32, _ := stats.SumSlice(uint32Slice)
	fmt.Printf("uint32:    %v\n", res_uint32)
	uint64Slice := []uint64{1, 2, 3}
	res_uint64, _ := stats.SumSlice(uint64Slice)
	fmt.Printf("uint64:    %v\n", res_uint64)
	uintptrSlice := []uintptr{1, 2, 3}
	res_uintptr, _ := stats.SumSlice(uintptrSlice)
	fmt.Printf("uintptr:   %v\n", res_uintptr)
	float32Slice := []float32{1.5, 2.5, 3.5}
	res_float32, _ := stats.SumSlice(float32Slice)
	fmt.Printf("float32:   %v\n", res_float32)
	float64Slice := []float64{1.5, 2.5, 3.5}
	res_float64, _ := stats.SumSlice(float64Slice)
	fmt.Printf("float64:   %v\n", res_float64)
	// Output:
	// int:       6
	// int8:      6
	// int16:     6
	// int32:     6
	// int64:     6
	// uint:      6
	// uint8:     6
	// uint16:    6
	// uint32:    6
	// uint64:    6
	// uintptr:   6
	// float32:   7.5
	// float64:   7.5
}

func ExampleMinSlice() {
	intSlice := []int{3, 1, 2}
	res_int, _ := stats.MinSlice(intSlice)
	fmt.Printf("int:       %v\n", res_int)
	int8Slice := []int8{3, 1, 2}
	res_int8, _ := stats.MinSlice(int8Slice)
	fmt.Printf("int8:      %v\n", res_int8)
	int16Slice := []int16{3, 1, 2}
	res_int16, _ := stats.MinSlice(int16Slice)
	fmt.Printf("int16:     %v\n", res_int16)
	int32Slice := []int32{3, 1, 2}
	res_int32, _ := stats.MinSlice(int32Slice)
	fmt.Printf("int32:     %v\n", res_int32)
	int64Slice := []int64{3, 1, 2}
	res_int64, _ := stats.MinSlice(int64Slice)
	fmt.Printf("int64:     %v\n", res_int64)
	uintSlice := []uint{3, 1, 2}
	res_uint, _ := stats.MinSlice(uintSlice)
	fmt.Printf("uint:      %v\n", res_uint)
	uint8Slice := []uint8{3, 1, 2}
	res_uint8, _ := stats.MinSlice(uint8Slice)
	fmt.Printf("uint8:     %v\n", res_uint8)
	uint16Slice := []uint16{3, 1, 2}
	res_uint16, _ := stats.MinSlice(uint16Slice)
	fmt.Printf("uint16:    %v\n", res_uint16)
	uint32Slice := []uint32{3, 1, 2}
	res_uint32, _ := stats.MinSlice(uint32Slice)
	fmt.Printf("uint32:    %v\n", res_uint32)
	uint64Slice := []uint64{3, 1, 2}
	res_uint64, _ := stats.MinSlice(uint64Slice)
	fmt.Printf("uint64:    %v\n", res_uint64)
	uintptrSlice := []uintptr{3, 1, 2}
	res_uintptr, _ := stats.MinSlice(uintptrSlice)
	fmt.Printf("uintptr:   %v\n", res_uintptr)
	float32Slice := []float32{1.5, 2.5, 3.5}
	res_float32, _ := stats.MinSlice(float32Slice)
	fmt.Printf("float32:   %v\n", res_float32)
	float64Slice := []float64{1.5, 2.5, 3.5}
	res_float64, _ := stats.MinSlice(float64Slice)
	fmt.Printf("float64:   %v\n", res_float64)
	// Output:
	// int:       1
	// int8:      1
	// int16:     1
	// int32:     1
	// int64:     1
	// uint:      1
	// uint8:     1
	// uint16:    1
	// uint32:    1
	// uint64:    1
	// uintptr:   1
	// float32:   1.5
	// float64:   1.5
}

func ExampleMaxSlice() {
	intSlice := []int{3, 1, 2}
	res_int, _ := stats.MaxSlice(intSlice)
	fmt.Printf("int:       %v\n", res_int)
	int8Slice := []int8{3, 1, 2}
	res_int8, _ := stats.MaxSlice(int8Slice)
	fmt.Printf("int8:      %v\n", res_int8)
	int16Slice := []int16{3, 1, 2}
	res_int16, _ := stats.MaxSlice(int16Slice)
	fmt.Printf("int16:     %v\n", res_int16)
	int32Slice := []int32{3, 1, 2}
	res_int32, _ := stats.MaxSlice(int32Slice)
	fmt.Printf("int32:     %v\n", res_int32)
	int64Slice := []int64{3, 1, 2}
	res_int64, _ := stats.MaxSlice(int64Slice)
	fmt.Printf("int64:     %v\n", res_int64)
	uintSlice := []uint{3, 1, 2}
	res_uint, _ := stats.MaxSlice(uintSlice)
	fmt.Printf("uint:      %v\n", res_uint)
	uint8Slice := []uint8{3, 1, 2}
	res_uint8, _ := stats.MaxSlice(uint8Slice)
	fmt.Printf("uint8:     %v\n", res_uint8)
	uint16Slice := []uint16{3, 1, 2}
	res_uint16, _ := stats.MaxSlice(uint16Slice)
	fmt.Printf("uint16:    %v\n", res_uint16)
	uint32Slice := []uint32{3, 1, 2}
	res_uint32, _ := stats.MaxSlice(uint32Slice)
	fmt.Printf("uint32:    %v\n", res_uint32)
	uint64Slice := []uint64{3, 1, 2}
	res_uint64, _ := stats.MaxSlice(uint64Slice)
	fmt.Printf("uint64:    %v\n", res_uint64)
	uintptrSlice := []uintptr{3, 1, 2}
	res_uintptr, _ := stats.MaxSlice(uintptrSlice)
	fmt.Printf("uintptr:   %v\n", res_uintptr)
	float32Slice := []float32{1.5, 2.5, 3.5}
	res_float32, _ := stats.MaxSlice(float32Slice)
	fmt.Printf("float32:   %v\n", res_float32)
	float64Slice := []float64{1.5, 2.5, 3.5}
	res_float64, _ := stats.MaxSlice(float64Slice)
	fmt.Printf("float64:   %v\n", res_float64)
	// Output:
	// int:       3
	// int8:      3
	// int16:     3
	// int32:     3
	// int64:     3
	// uint:      3
	// uint8:     3
	// uint16:    3
	// uint32:    3
	// uint64:    3
	// uintptr:   3
	// float32:   3.5
	// float64:   3.5
}
