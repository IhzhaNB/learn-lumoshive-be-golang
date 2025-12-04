package main

import (
	"fmt"
	"os"
	"testing"
)

// Setup global state or resources
var globalSetupDone bool

// TestMain adalah fungsi utama untuk pengaturan dan pembersihan sebelum dan setelah semua tes
func TestMain(m *testing.M) {
	// Setup sebelum menjalankan tes
	fmt.Println("Setup: Initializing resources...")
	globalSetupDone = true

	// Jalankan tes
	exitCode := m.Run()

	// Teardown setelah menjalankan tes
	fmt.Println("Teardown: Cleaning up resources...")
	globalSetupDone = false

	// Kembalikan kode keluar tes
	os.Exit(exitCode)
}

// TestExample1 adalah contoh tes yang akan dijalankan
func TestExample1(t *testing.T) {
	if !globalSetupDone {
		t.Fatal("Setup was not done correctly")
	}
	fmt.Println("Running TestExample1")
	// Tambahkan logika tes di sini
}

// TestExample2 adalah contoh tes lainnya
func TestExample2(t *testing.T) {
	if !globalSetupDone {
		t.Fatal("Setup was not done correctly")
	}
	fmt.Println("Running TestExample2")
	// Tambahkan logika tes di sini
}

func TestAdd(t *testing.T) {
	t.Run("Positive numbers", func(t *testing.T) {
		result := Add(1, 2)
		expected := 3
		if result != expected {
			t.Errorf("Add(1, 2) = %d; want %d", result, expected)
		}
	})

	t.Run("Negative numbers", func(t *testing.T) {
		result := Add(-1, -1)
		expected := -2
		if result != expected {
			t.Errorf("Add(-1, -1) = %d; want %d", result, expected)
		}
	})

	t.Run("Mix of positive and negative", func(t *testing.T) {
		result := Add(1, -1)
		expected := 0
		if result != expected {
			t.Errorf("Add(1, -1) = %d; want %d", result, expected)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		result := Add(0, 0)
		expected := 0
		if result != expected {
			t.Errorf("Add(0, 0) = %d; want %d", result, expected)
		}
	})
}
