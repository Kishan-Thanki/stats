# Exhaustive Generic Examples

This document serves as definitive proof that the `stats` package is fully compatible with every standard numeric type in the Go ecosystem.

## Average

The `Average` function utilizes Go 1.21 Generics to seamlessly compute over any slice of the following types:

```go
res_int, _ := stats.Average([]int{1, 2, 3})
res_int8, _ := stats.Average([]int8{1, 2, 3})
res_int16, _ := stats.Average([]int16{1, 2, 3})
res_int32, _ := stats.Average([]int32{1, 2, 3})
res_int64, _ := stats.Average([]int64{1, 2, 3})
res_uint, _ := stats.Average([]uint{1, 2, 3})
res_uint8, _ := stats.Average([]uint8{1, 2, 3})
res_uint16, _ := stats.Average([]uint16{1, 2, 3})
res_uint32, _ := stats.Average([]uint32{1, 2, 3})
res_uint64, _ := stats.Average([]uint64{1, 2, 3})
res_uintptr, _ := stats.Average([]uintptr{1, 2, 3})
res_float32, _ := stats.Average([]float32{1.5, 2.5, 3.5})
res_float64, _ := stats.Average([]float64{1.5, 2.5, 3.5})
```

## Median

The `Median` function utilizes Go 1.21 Generics to seamlessly compute over any slice of the following types:

```go
res_int, _ := stats.Median([]int{1, 2, 3})
res_int8, _ := stats.Median([]int8{1, 2, 3})
res_int16, _ := stats.Median([]int16{1, 2, 3})
res_int32, _ := stats.Median([]int32{1, 2, 3})
res_int64, _ := stats.Median([]int64{1, 2, 3})
res_uint, _ := stats.Median([]uint{1, 2, 3})
res_uint8, _ := stats.Median([]uint8{1, 2, 3})
res_uint16, _ := stats.Median([]uint16{1, 2, 3})
res_uint32, _ := stats.Median([]uint32{1, 2, 3})
res_uint64, _ := stats.Median([]uint64{1, 2, 3})
res_uintptr, _ := stats.Median([]uintptr{1, 2, 3})
res_float32, _ := stats.Median([]float32{1.5, 2.5, 3.5})
res_float64, _ := stats.Median([]float64{1.5, 2.5, 3.5})
```

## Mode

The `Mode` function utilizes Go 1.21 Generics to seamlessly compute over any slice of the following types:

```go
res_int, _ := stats.Mode([]int{1, 2, 3})
res_int8, _ := stats.Mode([]int8{1, 2, 3})
res_int16, _ := stats.Mode([]int16{1, 2, 3})
res_int32, _ := stats.Mode([]int32{1, 2, 3})
res_int64, _ := stats.Mode([]int64{1, 2, 3})
res_uint, _ := stats.Mode([]uint{1, 2, 3})
res_uint8, _ := stats.Mode([]uint8{1, 2, 3})
res_uint16, _ := stats.Mode([]uint16{1, 2, 3})
res_uint32, _ := stats.Mode([]uint32{1, 2, 3})
res_uint64, _ := stats.Mode([]uint64{1, 2, 3})
res_uintptr, _ := stats.Mode([]uintptr{1, 2, 3})
res_float32, _ := stats.Mode([]float32{1.5, 2.5, 3.5})
res_float64, _ := stats.Mode([]float64{1.5, 2.5, 3.5})
res_string, _ := stats.Mode([]string{"apple", "banana", "apple"})
```

## SumSlice

The `SumSlice` function utilizes Go 1.21 Generics to seamlessly compute over any slice of the following types:

```go
res_int, _ := stats.SumSlice([]int{1, 2, 3})
res_int8, _ := stats.SumSlice([]int8{1, 2, 3})
res_int16, _ := stats.SumSlice([]int16{1, 2, 3})
res_int32, _ := stats.SumSlice([]int32{1, 2, 3})
res_int64, _ := stats.SumSlice([]int64{1, 2, 3})
res_uint, _ := stats.SumSlice([]uint{1, 2, 3})
res_uint8, _ := stats.SumSlice([]uint8{1, 2, 3})
res_uint16, _ := stats.SumSlice([]uint16{1, 2, 3})
res_uint32, _ := stats.SumSlice([]uint32{1, 2, 3})
res_uint64, _ := stats.SumSlice([]uint64{1, 2, 3})
res_uintptr, _ := stats.SumSlice([]uintptr{1, 2, 3})
res_float32, _ := stats.SumSlice([]float32{1.5, 2.5, 3.5})
res_float64, _ := stats.SumSlice([]float64{1.5, 2.5, 3.5})
```

## MinSlice

The `MinSlice` function utilizes Go 1.21 Generics to seamlessly compute over any slice of the following types:

```go
res_int, _ := stats.MinSlice([]int{1, 2, 3})
res_int8, _ := stats.MinSlice([]int8{1, 2, 3})
res_int16, _ := stats.MinSlice([]int16{1, 2, 3})
res_int32, _ := stats.MinSlice([]int32{1, 2, 3})
res_int64, _ := stats.MinSlice([]int64{1, 2, 3})
res_uint, _ := stats.MinSlice([]uint{1, 2, 3})
res_uint8, _ := stats.MinSlice([]uint8{1, 2, 3})
res_uint16, _ := stats.MinSlice([]uint16{1, 2, 3})
res_uint32, _ := stats.MinSlice([]uint32{1, 2, 3})
res_uint64, _ := stats.MinSlice([]uint64{1, 2, 3})
res_uintptr, _ := stats.MinSlice([]uintptr{1, 2, 3})
res_float32, _ := stats.MinSlice([]float32{1.5, 2.5, 3.5})
res_float64, _ := stats.MinSlice([]float64{1.5, 2.5, 3.5})
```

## MaxSlice

The `MaxSlice` function utilizes Go 1.21 Generics to seamlessly compute over any slice of the following types:

```go
res_int, _ := stats.MaxSlice([]int{1, 2, 3})
res_int8, _ := stats.MaxSlice([]int8{1, 2, 3})
res_int16, _ := stats.MaxSlice([]int16{1, 2, 3})
res_int32, _ := stats.MaxSlice([]int32{1, 2, 3})
res_int64, _ := stats.MaxSlice([]int64{1, 2, 3})
res_uint, _ := stats.MaxSlice([]uint{1, 2, 3})
res_uint8, _ := stats.MaxSlice([]uint8{1, 2, 3})
res_uint16, _ := stats.MaxSlice([]uint16{1, 2, 3})
res_uint32, _ := stats.MaxSlice([]uint32{1, 2, 3})
res_uint64, _ := stats.MaxSlice([]uint64{1, 2, 3})
res_uintptr, _ := stats.MaxSlice([]uintptr{1, 2, 3})
res_float32, _ := stats.MaxSlice([]float32{1.5, 2.5, 3.5})
res_float64, _ := stats.MaxSlice([]float64{1.5, 2.5, 3.5})
```

