## 📝 Quiz Benchmarking & HTTP Test (Lengkap dengan Sumber)

### Pertanyaan 1

**Manakah pernyataan berikut yang benar mengenai `httptest.NewRecorder()` di Golang?**

**Jawaban:**

Pernyataan yang benar mengenai `httptest.NewRecorder()` di **Golang** adalah:

**Fungsi ini digunakan untuk membuat objek `http.ResponseRecorder` untuk merekam hasil response dari handler HTTP.**

**Sumber**:
pada file [**`http-test/README.md`**](../2.benchmarking-httptest/4.http-test/README.md)

---

### Pertanyaan 2

**Bagaimana cara mengukur waktu eksekusi suatu fungsi di Golang menggunakan benchmarking?**

**Jawaban:**

Cara yang **benar** dan **standar** untuk mengukur waktu eksekusi suatu fungsi di Golang menggunakan benchmarking adalah:

**Menggunakan `b.N` untuk mengukur jumlah iterasi dan waktu eksekusi**

**Sumber**:
pada file [**`benchmark/README.md`**](../2.benchmarking-httptest/2.benchmark/README.md)

---

### Pertanyaan 3

**Dalam benchmarking di Golang, apa tujuan dari menggunakan `b.N`?**

**Jawaban:**

Tujuan utama dari menggunakan **`b.N`** dalam _benchmarking_ di Golang adalah:

**Untuk mendefinisikan jumlah iterasi yang dilakukan dalam benchmark**

**Sumber**:
pada file [**`benchmark/README.md`**](../2.benchmarking-httptest/2.benchmark/README.md)

---

### Pertanyaan 4

**Manakah pernyataan berikut yang benar mengenai Sub-Benchmark di Golang?**

**Jawaban:**

Pernyataan yang benar mengenai **Sub-Benchmark** di Golang adalah:

**Sub-benchmark memungkinkan pengujian beberapa parameter dalam satu fungsi benchmark menggunakan metode `b.Run()`**

**Sumber**:
pada file [**`sub-table-benchmark/README.md`**](../2.benchmarking-httptest/3.sub-table-benchmark/README.md)

---

### Pertanyaan 5

**Untuk menguji performa dan mendapatkan hasil waktu rata-rata dari benchmark, developer ingin menjalankan benchmark menggunakan profiling CPU dan memori. Perintah mana yang tepat?**

**Jawaban:**

Perintah yang tepat untuk menjalankan _benchmark_ sambil mengaktifkan pengukuran memori (_memory allocation_) per operasi dan mendapatkan hasil waktu rata-rata adalah:

$$\text{go test -bench . -benchmem}$$

**Sumber**:
pada file [**`benchmark/README.md`**](../2.benchmarking-httptest/2.benchmark/README.md)
