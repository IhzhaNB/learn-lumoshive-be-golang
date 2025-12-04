# Sub & Table Benchmark

## Sub Benchmark

**Sub-benchmark** di Go memungkinkan kita untuk menjalankan benchmark dengan parameter yang berbeda di dalam satu fungsi benchmark, menggunakan metode **b.Run**.

### Contoh code:

`main_test.go`

```go
package main

import (
	"testing"
)

// BenchmarkSum benchmarks the Sum function with various slice sizes
func BenchmarkSum(b *testing.B) {
	b.Run("SmallSlice", func(b *testing.B) {
		numbers := []int{1, 2, 3, 4, 5}
		for i := 0; i < b.N; i++ {
			Sum(numbers)
		}
	})

	b.Run("MediumSlice", func(b *testing.B) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i := 0; i < b.N; i++ {
			Sum(numbers)
		}
	})

	b.Run("LargeSlice", func(b *testing.B) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		for i := 0; i < b.N; i++ {
			Sum(numbers)
		}
	})
}
```

ini hampir sama dengan testing. bedanya dgn test:

```bash
got test -bench=.
```

## Table Benchmark

**Table Benchmark** adalah teknik dimana kita mendefinisikan berbagai skenario benchmark dalam sebuah table yg berisi parameter yg berbeda.

### Contoh Code:

`main_test.go`

```go
package main

import (
	"testing"
)

func BenchmarkSum(b *testing.B) {
    // Tabel benchmark dengan berbagai ukuran slice
    tests := []struct {
        name    string
        numbers []int
    }{
        {"SmallSlice", []int{1, 2, 3, 4, 5}},
        {"MediumSlice", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
        {"LargeSlice", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
    }

    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                Sum(tt.numbers)
            }
        })
    }
}
```

lalu jalankan:

```bash
got test -bench=.
```
