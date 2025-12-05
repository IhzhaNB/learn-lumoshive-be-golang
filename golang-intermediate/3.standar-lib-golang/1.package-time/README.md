# Package Time

Paket time dalam bahasa Go (Golang) menyediakan fungsionalitas untuk mengelola waktu dan tanggal. Paket ini sangat penting untuk berbagai keperluan mulai dari pengukuran waktu eksekusi hingga penjadwalan tugas. Berikut adalah beberapa method dan fungsi utama yang ada dalam paket time:

### Method dan fungsi utama yg ada dalam paket time:

- `Now`
- `Parse`
- `Format`
- `Sleep`
- `Ticker`
- `After`
- `AfterFunc`
- `NewTimer`

## `time.Now()`

Fungsi time.Now digunakan untuk mendapatkan waktu saat ini sesuai dengan zona waktu lokasi sistem tempat program berjalan.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println("Current Time:", now)
}
```

```bash
go run .

``
Current Time: 2024-07-08 14:41:14.5000209 +07 m=+0.002691901
```

### Method Paket `time` yang Terkait `Time.now`:

| Method           | Deskripsi                                        | Contoh Penggunaan  |
| :--------------- | :----------------------------------------------- | :----------------- |
| **Year()**       | Mengembalikan tahun dari objek Time.             | `now.Year()`       |
| **Month()**      | Mengembalikan bulan dari objek Time.             | `now.Month()`      |
| **Day()**        | Mengembalikan hari dari objek Time.              | `now.Day()`        |
| **Hour()**       | Mengembalikan jam dari objek Time.               | `now.Hour()`       |
| **Minute()**     | Mengembalikan menit dari objek Time.             | `now.Minute()`     |
| **Second()**     | Mengembalikan detik dari objek Time.             | `now.Second()`     |
| **Nanosecond()** | Mengembalikan nanodetik dari objek Time.         | `now.Nanosecond()` |
| **Location()**   | Mengembalikan zona waktu dari objek Time.        | `now.Location()`   |
| **Weekday()**    | Mengembalikan hari dalam minggu dari objek Time. | `now.Weekday()`    |
| **Local()**      | Mengkonversi waktu Time ke waktu lokal.          | `now.Local()`      |

### Contoh penerapannya:

`main.go`

```go
package main

import (
    "fmt"
    "time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Println("Current time:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Nanosecond:", now.Nanosecond())
	fmt.Println("Location:", now.Location())

    formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted time:", formattedTime)
}
```

output:

```bash
Current time: 2024-07-08 14:43:55.9698547 +0700 +07 m=+0.002668001
Year: 2024
Month: July
Day: 8
Hour: 14
Minute: 43
Second: 55
Nanosecond: 969854700
Location: Local
Formatted time: 2024-07-08 14:43:55
```

## `time.Parse()`

**`time.Parse`** adalah fungsi dalam paket **time** di Golang yang digunakan untuk **mengkonversi _string_ yang berisi representasi waktu ke dalam objek Time**.

### Contoh penerapannya:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    layout := "2006-01-02 15:04:05"
	str := "2023-06-23 14:30:00"
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time:", parsedTime)
}
```

output:

```bash
Parsed time: 2023-06-23 14:30:00 +0000 UTC
```

## `time.Format()`

**time.Format** adalah method dalam paket time di Golang yang digunakan untuk mengkonversi objek Time ke dalam _string_ sesuai dengan _layout_ yang ditentukan.

### Contoh penerapannya:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    // Format tanggal dan waktu lengkap
	layout1 := "2006-01-02 15:04:05"
	formattedTime1 := now.Format(layout1)
	fmt.Println("Formatted time 1:", formattedTime1)

	// Format hanya tanggal
	layout2 := "02-Jan-2006"
	formattedTime2 := now.Format(layout2)
	fmt.Println("Formatted time 2:", formattedTime2)

	// Format dengan nama hari dan bulan
	layout3 := "Monday, 02-Jan-2006 03:04:05 PM"
	formattedTime3 := now.Format(layout3)
	fmt.Println("Formatted time 3:", formattedTime3)

	// Format dengan zona waktu
	layout4 := "2006-01-02 15:04:05 MST"
	formattedTime4 := now.Format(layout4)
	fmt.Println("Formatted time 4:", formattedTime4)
}
```

output:

```bash
Formatted time 1: 2024-07-08 14:47:49
Formatted time 2: 08-Jul-2024
Formatted time 3: Monday, 08-Jul-2024 02:47:49 PM
Formatted time 4: 2024-07-08 14:47:49 +07
```

## `time.Sleep()`

**time.Sleep** adalah fungsi dalam paket time di Golang yang digunakan untuk menangguhkan eksekusi saat ini selama durasi tertentu.

### Contoh penerapannya:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // start time sleep
	fmt.Println("Start")

	// Menunda eksekusi selama 2 detik
	time.Sleep(2 * time.Second)

	fmt.Println("2 seconds later")

	// Menunda eksekusi selama 500 milidetik (0.5 detik)
	time.Sleep(500 * time.Millisecond)

	fmt.Println("0.5 seconds later")
}
```

output:

```bash
Start
2 seconds later
0.5 seconds later
```

## `time.Ticker()`

**time.Ticker** adalah tipe dalam paket time di Golang yang digunakan untuk menghasilkan peristiwa secara berkala pada interval waktu tertentu. time.Ticker mengirim nilai waktu (time.Time) secara berulang ke channel yang bisa diterima dalam goroutine lain.

### Contoh penerapannya:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(6 * time.Second)
	fmt.Println("Main function exiting")
}
```

output:

```bash
Tick at 2024-07-08 14:51:27.9941987 +0700 +07 m=+1.003502701
Tick at 2024-07-08 14:51:29.0052751 +0700 +07 m=+2.014579101
Tick at 2024-07-08 14:51:29.9965241 +0700 +07 m=+3.005828101
Tick at 2024-07-08 14:51:30.9946199 +0700 +07 m=+4.003923901
Main function exiting
```

## `time.After()`

**time.After** adalah fungsi dalam paket time di Golang yang digunakan untuk membuat channel yang akan mengirimkan nilai waktu tunggal setelah durasi tertentu berlalu.

### Contoh penerapannya:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Start")

	// Membuat channel yang mengirimkan waktu tunggal setelah 2 detik
	timeout := time.After(2 * time.Second)

	// Menunggu nilai dari channel timeout atau done
	select {
	case <-timeout:
		fmt.Println("Timeout occurred")
	case <-time.After(3 * time.Second):
		fmt.Println("Operation took too long")
	}

	fmt.Println("End")
}
```

output:

```bash
Start
Timeout occurred
End
```

## `time.AfterFunc()`

**time.AfterFunc** adalah fungsi dalam paket **time** di **Golang** yang digunakan untuk menjadwalkan eksekusi fungsi tertentu setelah durasi tertentu berlalu. Fungsi ini mirip dengan **time.After**, namun **time.AfterFunc** memungkinkan Anda untuk menjalankan sebuah fungsi **callback** (atau fungsi anonim) secara **asinkron** setelah jangka waktu tertentu.

### Contoh penerapannya:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Start")

	// Menjadwalkan eksekusi fungsi callback setelah 2 detik
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Callback executed after 2 seconds")
	})

	// Menunggu agar callback dieksekusi sebelum keluar
	time.Sleep(3 * time.Second)
	fmt.Println("End")
}
```

output:

```bash
Start
Callback executed after 2 seconds
End
```

## time.NewTimer()

**time.NewTimer** adalah fungsi dalam paket **time** di **Golang** yang digunakan untuk membuat objek **Timer** yang akan mengirimkan nilai waktu tunggal ke dalam channel setelah durasi tertentu berlalu. Fungsi ini berguna untuk menunda eksekusi goroutine untuk jangka waktu tertentu dan kemudian mengirimkan waktu tersebut melalui channel.

### Contoh penerapannya:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Start")

	// Membuat timer yang mengirimkan waktu setelah 2 detik
	timer := time.NewTimer(2 * time.Second)

	// Menunggu nilai dari channel timer.C
	<-timer.C
	fmt.Println("Timer expired")

	fmt.Println("End")
}
```

output:

```bash
Start
Timer expired
End
```
