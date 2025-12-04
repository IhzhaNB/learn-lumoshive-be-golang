# Benchmark

**Benchmarking** mengacu pada proses pengujian dan evaluasi kinerja berbagai bagian dari aplikasi atau sistem. Di Go, **benchmark** digunakan untuk mengukur waktu eksekusi dari kode untuk mengetahui seberapa efisien fungsi atau bagian tertentu dari aplikasi.

Di Go, _benchmarking_ dilakukan menggunakan **`package testing`**. _Benchmark_ di Go mirip dengan _unit test_ tetapi bertujuan untuk mengukur **waktu eksekusi**, bukan memverifikasi hasil.

## Aturan Pembuatan Benchmark pada Golang

- **Definisikan Fungsi Benchmark:** Fungsi _benchmark_ di Go menggunakan nama fungsi yang dimulai dengan **`Benchmark`** diikuti dengan nama yang mendeskripsikan apa yang diuji. Misalnya, `BenchmarkFunctionName`.
- **Gunakan Parameter `b *testing.B`:** Parameter ini adalah objek _benchmark_ yang menyediakan metode untuk menjalankan _benchmark_ berulang kali.
- **Gunakan `b.N`:** Untuk mengukur kinerja, _benchmark_ menjalankan kode dalam _loop_ sebanyak **`b.N`** kali. Go akan menentukan nilai `b.N` secara otomatis untuk mendapatkan hasil yang dapat diandalkan.

## Implement:

1. `main.go`

```go
package main

import "fmt"

// Sum calculates the sum of elements in a slice
func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(nums))
}
```

2. `main_test.go` untuk melakukan test pada function:

```go
package main

import (
	"testing"
)

// BenchmarkSum benchmarks the Sum function
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
```

## commonds benchmark golang

### run all benchmark

```bash
go test -bench=.
```

### run specifically name benchmark

```bash
go test -bench=BenchmarkFuncName
```

### run benchmark show profile CPU and memory (Jalankan benchmark dengan profil CPU dan memori:)

```bash
go test -bench=. -cpuprofile=cpu.pprof -memprofile=mem.pprof
```

example output bechmark in golang

```bash
BenchmarkSum-8    10000000        0.0750 ns/op
``

detail information
-BenchmarkSum-8: Nama benchmark dan jumlah core CPU yang digunakan.
-10000000: Jumlah iterasi yang dilakukan.
-0.0750 ns/op: Waktu rata-rata per iterasi dalam nanodetik.
```
