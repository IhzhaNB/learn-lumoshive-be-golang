# HTTP Test

**Testing HTTP** di Go biasanya dilakukan dengan menggunakan paket **net/http/httptest** dan **net/http**.
Paket ini memungkinkan kita untuk membuat server HTTP dan request palsu untuk menguji handler HTTP kita secara langsung tanpa memerlukan server yang sebenarnya.

### **httptest.NewRequest**

adalah fungsi dari paket **net/http/httptest** yang digunakan untuk membuat **request HTTP palsu (mock request)** untuk keperluan testing.

```go
func NewRequest(method, url string, body io.Reader) *http.Request
```

### **httptest.NewRecorder**

adalah fungsi dari paket **net/http/httptest** yang digunakan untuk membuat objek **http.ResponseWriter** yang disebut **httptest.ResponseRecorder**, yang berfungsi sebagai **writer** untuk merekam hasil **response** dari handler HTTP.

```go
func NewRecorder() *httptest.ResponseRecorder
```

## Implement Code:

1. `main.go`

```go
package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
```

2. `main_test.go`

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Buat request HTTP palsu
	req := httptest.NewRequest("GET", "http://example.com/", nil)

	// Buat recorder untuk menangkap response
	w := httptest.NewRecorder()

	// Panggil handler dengan request dan recorder
	HelloHandler(w, req)

	// Periksa status code response
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// Periksa body response
	expectedBody := "Hello, World!\n"
	body := w.Body.String()
	if body != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, body)
	}
}
```

lalu jalankan unit test

```bash
go test -v
```

## Implemen Code pada studi kasus lain:

1. `main.go`

```go
package main

import (
	"fmt"
	"net/http"
)

// implement other case
const (
	username = "admin"
	password = "password"
)

type Response struct {
	Message string `json:"message"`
}

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || pass != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var res Response
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/post", BasicAuthMiddleware(http.HandlerFunc(PostHandler)))
	http.ListenAndServe(":8080", mux)
}
```

2. `main_test.go`

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostHandler(t *testing.T) {
	// JSON body untuk request
	requestBody := Response{Message: "Hello, World!"}
	body, err := json.Marshal(requestBody)
	assert.NoError(t, err, "Failed to marshal request body")

	// Buat request HTTP palsu dengan basic auth
	req := httptest.NewRequest("POST", "http://example.com/post", bytes.NewReader(body))
	req.SetBasicAuth("admin", "password")

	// Buat recorder untuk menangkap response
	w := httptest.NewRecorder()

	// Panggil handler dengan request dan recorder
	handler := BasicAuthMiddleware(http.HandlerFunc(PostHandler))
	handler.ServeHTTP(w, req)

	// Periksa status code response
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode, "expected status code 200 OK")

	// Periksa body response
	var responseBody Response
	err = json.NewDecoder(w.Body).Decode(&responseBody)
	assert.NoError(t, err, "Failed to decode response body")
	assert.Equal(t, requestBody.Message, responseBody.Message, "expected body message to match request body message")
}

func TestPostHandlerUnauthorized(t *testing.T) {
	// JSON body untuk request
	requestBody := Response{Message: "Hello, World!"}
	body, err := json.Marshal(requestBody)
	assert.NoError(t, err, "Failed to marshal request body")

	// Buat request HTTP palsu tanpa basic auth
	req := httptest.NewRequest("POST", "http://example.com/post", bytes.NewReader(body))

	// Buat recorder untuk menangkap response
	w := httptest.NewRecorder()

	// Panggil handler dengan request dan recorder
	handler := BasicAuthMiddleware(http.HandlerFunc(PostHandler))
	handler.ServeHTTP(w, req)

	// Periksa status code response
	res := w.Result()
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "expected status code 401 Unauthorized")
}
```

lalu jalankan unit test

```bash
go test -v
```
