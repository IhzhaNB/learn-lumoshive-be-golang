# Package Strings

Package **strings** dalam Golang menyediakan berbagai fungsi untuk memanipulasi string.

## `strings.Contains()`

Digunakan untuk mengecek apakah sebuah string mengandung substring tertentu. Berikut
contohnya:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, Lumoshive Academy!"
    fmt.Println(strings.Contains(str, "Lumoshive")) // true
    fmt.Println(strings.Contains(str, "agency")) // false
    fmt.Println(strings.Contains("abcdef", "cd")) // true
}
```

output:

```bash
go run .

``
true
false
true
```

## `strings.ToUpper()` dan `strings.ToLower()`

Digunakan untuk mengonversi string menjadi huruf besar atau huruf kecil. Berikut contoh kodenya:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "lumoshive"
    fmt.Println(strings.ToUpper(str))
    fmt.Println(strings.ToLower("ACADEMY"))
}
```

output:

```bash
LUMOSHIVE
academy
```

## `strings.Trim()` dan `strings.TrimSpace()`

Digunakan untuk menghapus karakter tertentu diawal dan diakhir string. Berikut contoh kodenya:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Trim("***Hello, Lumoshive Academy!***", "*"))
    fmt.Println(strings.TrimSpace(" Hello, Lumoshive Academy! "))
    fmt.Println(strings.Trim("--Hello, Lumoshive Academy!--", "-"))
}
```

output:

```bash
Hello, Lumoshive Academy!
Hello, Lumoshive Academy!
Hello, Lumoshive Academy!
```

## `strings.Split()`

Digunakan untuk memisahkan string berdasarkan delimiter tertentu. Berikut contoh kodenya

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fruits := "apel,pisang,mangga"
    words := strings.Split(fruits, ",")
    fmt.Println(words)
    fmt.Println(strings.Split("lumoshive|academy|happy coding", "|"))
}
```

output:

```bash
[apel pisang mangga]
[lumshive academy happy coding]
```

## strings.Join()

Digunakan untuk menggabungkan elemen slice menjadi string dengan separator tertentu. Berikut contoh kodenya:

```go
package main

mport (
    "fmt"
    "strings"
)

func main() {
    words := []string{"Hello", "Lumoshive", "Academy"}
    fmt.Println(strings.Join(words, "-"))
    fmt.Println(strings.Join([]string{"go", "is", "fun"}, " "))
}
```

output:

```bash
Hello-Lumoshive-Academy
go is fun
```

## `strings.ReplaceAll()`

Digunakan mengganti semua kemunculan substring tertentu dalam string.. Berikut contoh kodenya:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, world"
    fmt.Println(strings.ReplaceAll(str, "world", "Lumoshive"))
    fmt.Println(strings.ReplaceAll("2025-02-11", "-", "/"))
}
```

output:

```bash
Hello, Lumoshive
2025/02/11
```

## `strings.HasPrefix()` dan `strings.HasSuffix()`

Digunakan memeriksa apakah string memiliki awalan atau akhiran tertentu. Berikut contoh kodenya:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.HasPrefix("Lumoshive is great", "Lumoshive")) // memeriksa dari awalan string
    fmt.Println(strings.HasSuffix("data.txt", ".txt")) // memeriksa dari akhiran string
    fmt.Println(strings.HasPrefix("www.lumoshive.com", "http"))
}
```

output:

```bash
true
true
false
```
