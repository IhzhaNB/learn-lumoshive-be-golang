# Package Regexp

Package regexp digunakan untuk melakukan pencocokan pola berbasis regular expression (regex). Ini sangat berguna untuk validasi format string, pencarian, penggantian, dan ekstraksi data dari teks.

## `regexp.Compile()`

Digunakan untuk mengompilasi regular expression menjadi objek regexp dan objek error. Function `regexp.Compile()` sangat cocok digunakan ketika belum yakin dengan format penulisan regular expression yang benar. Berikut contohnya:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile("[a-z]+")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	fmt.Println("Regex compiled successfully:", re.String())
	// do something to the string
}
```

#### `regexp.Compile("[a-z]+")`

Membuat objek regex yang mencocokkan satu atau lebih huruf kecil (a sampai z). Fungsi ini mengembalikan objek regex (re) dan error (err) jika ada kesalahan.

output:

```bash
Regex compiled successfully: [a-z]+
```

## `regexp.MustCompile()`

Digunakan untuk digunakan untuk mengompilasi regular expression menjadi sebuah objek `regexp.Regexp`. Dengan catatan pola yang digunakan sudah benar, berikut contohnya:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[A-Za-z]+$")
	fmt.Println("Regex compiled successfully:", re.String())
	// do something to the string
}
```

### `regexp.MustCompile("^[A-Za-z]+$")`

Membuat objek regex yang mencocokkan string yang terdiri dari satu atau lebih huruf besar (A-Z) atau huruf kecil (a-z) dari awal (^) sampai akhir ($) string. Fungsi MustCompile akan langsung menghasilkan objek regex dan panic jika pola regex tidak valid.

output:

```bash
Regex compiled successfully: ^[A-Za-z]+$
```

## `MatchString`

Digunakan untuk mengecek apakah string cocok dengan pola regex, berikut contohnya:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[A-Za-z]+$")
	// sample
	fmt.Println(re.MatchString("Lumoshive"))
	fmt.Println(re.MatchString("Lumoshive123"))
}
```

- `re.MatchString("Lumoshive")` Memeriksa apakah string "Lumoshive" hanya berisi huruf dari A sampai Z (besar atau
  kecil). Hasilnya true karena string sesuai pola.
- `re.MatchString("Lumoshive123")`
  Memeriksa apakah string "Lumoshive123" hanya berisi huruf dari A sampai Z (besar atau kecil). Hasilnya false karena ada angka di dalam string yang tidak cocok dengan
  pola.

output:

```bash
true
false
```

## `FindString`

Digunakan untuk mencari dan mengembalikan string pertama yang cocok dengan pola (regex) dalam suatu teks, berikut contohnya:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindString("Kode 123 Golang 456"))
}
```

- `re.FindString("Kode 123 Golang 456")`
  Mencari dan mengembalikan substring pertama dalam string yang cocok dengan pola regex, yaitu angka pertama yang ditemukan yaitu "123".

output:

```bash
123
```

## `FindAllString`

Digunakan untuk mencari semua substring dalam sebuah string yang cocok dengan pola regex tertentu, berikut contohnya:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString("Kode 123 Golang 456", -1))
}
```

- `re.FindAllString("Kode 123 Golang 456", -1)`
  Mencari dan mengembalikan semua substring dalam string yang cocok dengan pola regex. Argumen -1 berarti mencari semua kemunculan yang cocok. Pada contoh ini, hasilnya adalah slice ["123", "456"].

output:

```bash
[123 456]
```

## `ReplaceAllString`

Digunakan untuk mengganti semua bagian dari string yang cocok dengan pola regex, dengan string pengganti tertentu, berikut contohnya:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.ReplaceAllString("Kode 123 Golang 456", "x"))
}
```

- `re.ReplaceAllString("Kode 123 Golang 456", "x")`
  Mengganti semua substring dalam string yang cocok dengan pola regex dengan string "x". Pada contoh ini, angka "123" dan "456" akan diganti menjadi "x" sehingga hasilnya menjadi "Kode x Golang x".

output:

```bash
Kode x Golang x
```

## `Split`

Digunakan untuk memecah string (mirip strings.Split), tetapi menggunakan pola regex sebagai pemisahnya, berikut contohnya:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`,\s`)
	fmt.Println(re.Split("rust, golang, python", -1))
	fmt.Println(re.Split("rust, golang, python", 2))
}
```

- **regexp.MustCompile(\,\s`)**
  Membuat objek regex yang mencocokkan pola koma diikuti oleh satu spasi (", "). Fungsi MustCompile langsung mengembalikan objek regex dan akan panic jika pola regex tidak valid.
- `re.Split("rust, golang, python", -1)`
  Memecah string "rust, golang, python" menjadi slice berdasarkan pola regex. Argumen -1 berarti memecah sebanyak mungkin (semua kemunculan). Hasilnya adalah slice: ["rust", "golang", "python"].
- `re.Split("rust, golang, python", 2)`
  Memecah string yang sama tetapi hanya menjadi maksimal 2 bagian. Hasilnya adalah slice: ["rust", "golang, python"].

output:

```bash
[rust golang python]
[rust golang, python]
```

## Studi Kasus

Untuk mengimplementasikan package regexp, kita akan membuat sistem untuk memvalidasi apakah format nomor telepon sudah sesuai atau belum. Berikut contoh kode dalam bahasa Go:

```go
package main

import (
	"fmt"
	"regexp"
)

func isValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?\d{10,15}$`)
	return re.MatchString(phone)
}
func main() {
	examples := []string{"081234567890", "+6281234567890", "12345", "abcdefg"}
	for _, phone := range examples {
		fmt.Printf("%s valid: %t\n", phone, isValidPhoneNumber(phone))
	}
}
```

- `regexp.MustCompile(^+?\d{10,15}$)`
  Membuat objek regex yang mencocokkan string nomor telepon dengan pola:

  - `^` awal string
  - `\+?` tanda plus opsional di depan
  - `\d{10,15}` 10 sampai 15 digit angka
  - `$` akhir string

  Fungsi MustCompile akan langsung mengembalikan objek regex dan panic jika pola regex tidak valid.

output:

```bash
081234567890 valid: true
+6281234567890 valid: true
12345 valid: false
abcdefg valid: false
```

Pada contoh ini, nomor seperti 081234567890 dan +6281234567890 dikenali sebagai format yang valid, sedangkan 12345 dan abcdefg dianggap tidak valid. Proses validasi ini dilakukan dengan mencocokkan input terhadap pola regex tertentu
