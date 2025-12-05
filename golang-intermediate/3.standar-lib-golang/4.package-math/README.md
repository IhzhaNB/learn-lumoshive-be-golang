# Package Math

**Package math** adalah package bawaan di Go yang menyediakan konstanta dan fungsi matematika dasar. Package ini sangat berguna untuk melakukan berbagai perhitungan numerik.

## Constant

Pada package math menyediakan beberapa konstanta matematika penting yang sering digunakan dalam perhitungan numerik dan ilmiah. Semua konstanta bertipe float64. Berikut contohnya:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // constant in package math
    fmt.Println("Pi:", math.Pi)
    fmt.Println("Euler (e):", math.E)
}
```

output:

```bash
Pi: 3.141592653589793
Euler (e): 2.718281828459045
```

## Basic Math (`.Sqrt()` and `.Abs()`)

Package math di Go menyediakan fungsi-fungsi matematika dasar, termasuk untuk menghitung akar kuadrat dan nilai absolut. Berikut contohnya:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// basic math
	a := 9.0
	b := -3.5
	fmt.Println("Akar kuadrat dari", a, "adalah", math.Sqrt(a)) // akar kuadrat
	fmt.Println("Nilai absolut dari", b, "adalah", math.Abs(b)) // absolut
}
```

output:

```bash
Akar kuadrat dari 9 adalah 3
Nilai absolut dari -3.5 adalah 3.5
```

## Rounding

Package math menyediakan beberapa fungsi untuk melakukan pembulatan angka float64 menjadi nilai terdekat menurut aturan tertentu. Berikut contohnya:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// rounding
	x := 3.7
	fmt.Println("Floor dari", x, "adalah", math.Floor(x)) // Membulatkan kebawah
	fmt.Println("Ceil dari", x, "adalah", math.Ceil(x)) // Membulatkan keatas
	fmt.Println("Pembulatan dari", x, "adalah", math.Round(x)) // Membulatkan sesuai dibelakang koma, yang lebih dari 5 maka keatas, begitupun sebaliknya
}
```

output:

```bash
Floor dari 3.7 adalah 3
Ceil dari 3.7 adalah 4
Pembulatan dari 3.7 adalah 4
```

## `.Min()` & `.Max()`

Fungsi Min dan Max di package math digunakan untuk membandingkan dua angka dan mengembalikan nilai terkecil (Min) atau nilai terbesar (Max) dari keduanya. Keduanya bekerja dengan tipe data float64. Berikut contohnya:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// min and maks
	c := 7.3
	d := 2.9
	fmt.Println("Nilai maksimum antara", c, "dan", d, "adalah", math.Max(c, d))
	fmt.Println("Nilai minimum antara", c, "dan", d, "adalah", math.Min(c, d))
}
```

output:

```bash
Nilai maksimum antara 7.3 dan 2.9 adalah 7.3
Nilai minimum antara 7.3 dan 2.9 adalah 2.9
```

## Pangkat, akar dan modulus

Package math di Golang menyediakan berbagai fungsi matematika, termasuk untuk **pangkat (eksponen)**, **akar kuadrat**, dan **modulus (sisa bagi)**. Berikut contohnya:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// power, square root and modulus
	fmt.Println("2 pangkat 3:", math.Pow(2, 3))
	fmt.Println("Akar kubik dari 27:", math.Cbrt(27))
	fmt.Println("Modulus dari 10.5 dibagi 3:", math.Mod(10.5, 3))
}
```

output:

```bash
2 pangkat 3: 8
Akar kubik dari 27: 3
Modulus dari 10.5 dibagi 3: 1.5
```

Untuk lebih detail mengenai package math dapat mengunjungi link berikut: [https://pkg.go.dev/math](https://pkg.go.dev/math)
