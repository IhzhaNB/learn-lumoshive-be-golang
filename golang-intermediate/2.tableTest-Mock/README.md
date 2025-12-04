# Table Test & Mock

## Table Test

**Table test** adalah teknik dalam pengujian unit di mana kita dapat mendefinisikan **beberapa kasus uji dalam sebuah tabel** _(array slice)_ dan mengiterasi melalui tabel tersebut untuk menjalankan pengujian.

Keuntungan dari penggunaan **table test** ini adalah untuk **mengurangi pengulangan kode** saat menulis function test yang berbeda-beda untuk setiap kasus uji.

### Contoh code

`main.go`

```go
package main

// Add menjumlahkan dua bilangan
func Add(a, b int) int {
	return a + b
}

```

`main_test.go`

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	// Definisikan tabel test dengan berbagai kasus uji
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive numbers", 1, 2, 3},
		{"Negative numbers", -1, -1, -2},
		{"Mix of positive and negative", 1, -1, 0},
		{"Zero", 0, 0, 0},
	}

	// Iterasi melalui setiap kasus uji dalam tabel test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

```

---

## Mock

**Mock** adalah teknik dalam pengujian unit yang digunakan untuk menggantikan bagian dari sistem yang sedang diuji dengan **objek tiruan** (_mock_) yang meniru perilaku objek asli.

Teknik ini berguna ketika bagian dari sistem yang sedang diuji memiliki **dependensi eksternal** seperti:

- Memanggil **API _third party service_**.
- Interaksi dengan **database**.
- Operasi yang memerlukan banyak sumber daya.

### Contoh Implementasi

Pertama-tama buatlah proyek:

```bash
go mod init nama_proyek
```

lalu kita install library testify

```bash
go get github.com/stretchr/testify
```

setelah itu kita buat folder seperti ini:

```
nama-proyek/
├── model/
├── repository/
└── service/
```

lalu buat code:

#### 1. `model/user.go`

```go
package model

// User represents a user entity
type User struct {
	ID   int
	Name string
}
```

#### 2. `repository/user_repository.go`

```go
package repository

import (
	"nama_proyek/model"
	"database/sql"
	"fmt"
)

// UserRepository defines the methods for interacting with the user repository
type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

// GetUserByID is a mock implementation of UserRepository.GetUserByID
func (s *userRepository) GetUserByID(id int) (*model.User, error) {
	row := s.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
```

**Penjelasan Unit Test:**
Ketika kita menjalankan Test, kita tidak perlu menjalankan test untuk database nya, jadinya kita hanya menjalankan test ke tiruan database nya _(mock)_

#### 3. `service/user_service.go`

```go
package service

import (
	"nama_proyek/model"
	"nama_proyek/repository"
)

// UserService provides methods to interact with users
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser returns a user by ID
func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.repo.GetUserByID(id)
}
```

#### 4. `repository/user_repository_mock.go`

Tahap berikutnya kita buat mock sebagai pengganti object database

```go
package repository

import (
	"nama_proyek/model"

	"github.com/stretchr/testify/mock"
)

// UserRepositoryMock is a mock implementation of UserRepository using testify/mock
type UserRepositoryMock struct {
	mock.Mock
}

// GetUserByID is a mock implementation of UserRepository.GetUserByID
func (m *UserRepositoryMock) GetUserByID(id int) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}
```

#### 5. `service/user_service_test.go`

```go
package service

import (
	"nama_proyek/model"
	"nama_proyek/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUser(t *testing.T) {
	t.Run("user_found", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(repository.UserRepositoryMock)
		userService := NewUserService(mockRepo)

		// Define the user to be returned by the mock
		user := &model.User{ID: 1, Name: "Alice"}

		// Setup expectations
		mockRepo.On("GetUserByID", 1).Return(user, nil)

		// Call the method
		result, err := userService.GetUser(1)

		// Assert expectations
		assert.NoError(t, err)
		assert.Equal(t, user, result)

		// Verify that the expectations were met
		mockRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(repository.UserRepositoryMock)
		userService := NewUserService(mockRepo)

		// Setup expectations
		mockRepo.On("GetUserByID", 999).Return(nil, nil)

		// Call the method
		result, err := userService.GetUser(999)

		// Assert expectations
		assert.NoError(t, err)
		assert.Nil(t, result)

		// Verify that the expectations were met
		mockRepo.AssertExpectations(t)
	})
}
```

#### Berikut contoh folder lengkapnya:

```
nama-proyek/
├── model/
|   └── model.go
├── repository/
|   ├── user_repository.go
|	└── user_repository_mock.go
└── service/
    ├── user_service.go
	└── user_service_test.go
```

Setelah itu kita jalankan unit test nya dengan

### Penjelasan Code

pada `user_repository.go` kita tidak perlu menjalankan test bagian **DB** nya

```bash
go test ./...
```

---

Terkadang dalam kasus tertentu, kita ingin membuat _unit test_ untuk **repository**, namun kita tahu bahwa objek **repository** memiliki properti **DB (database)**. Oleh karena itu, ketika kita ingin melakukan _testing_ terhadap _repository_, kita perlu membuat _database_ terlebih dahulu.

Untuk mengatasi hal ini, kita dapat membuat _mock database_ menggunakan **`sqlmock`**, sehingga kita tidak perlu membuat _database_ nyata untuk _unit test_. Untuk menggunakan `sqlmock`, kita bisa memanfaatkan _library_ **`go-sqlmock`**.

```bash
go get github.com/DATA-DOG/go-sqlmock
```

tambahkan code pada file `repository/user_repository_test.go`

```go
package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUserByID(t *testing.T) {
	// Membuat mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %v", err)
	}
	defer db.Close()

	// Menetapkan ekspektasi untuk query SELECT
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "John Doe")
	mock.ExpectQuery("SELECT id, name FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	// initialization user repo
	repo := NewUserRepository(db)

	// Memanggil fungsi yang akan diuji
	user, err := repo.GetUserByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verifikasi hasil
	if user.ID != 1 || user.Name != "John Doe" {
		t.Errorf("unexpected user: got %v", user)
	}

	// Verifikasi bahwa semua ekspektasi terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	// Membuat mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %v", err)
	}
	defer db.Close()

	// Menetapkan ekspektasi untuk query SELECT yang tidak mengembalikan hasil
	mock.ExpectQuery("SELECT id, name FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	// initialization user repo
	repo := NewUserRepository(db)

	// Memanggil fungsi yang akan diuji
	_, err = repo.GetUserByID(1)
	if err == nil || err.Error() != "user not found" {
		t.Errorf("expected error 'user not found', got %v", err)
	}

	// Verifikasi bahwa semua ekspektasi terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}
```

lalu kita jalankan lagi Test nya

dan contoh folder lengkapnya:

```
nama-proyek/
├── model/
|   └── model.go
├── repository/
|   ├── user_repository.go
|   ├── user_repository_mock.go
|	└── user_repository_test.go
└── service/
    ├── user_service.go
	└── user_service_test.go
```
