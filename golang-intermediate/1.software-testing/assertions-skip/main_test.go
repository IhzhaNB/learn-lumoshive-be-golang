package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// // exampel not use assertion
// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5
// 	if result != expected {
// 		t.Error("Expected 5, but got", result)
// 	}
// }

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	assert.Equal(t, expected, result, "Add(2, 3) should be 5")
}

// // example assertion with require
// func TestSubtract(t *testing.T) {
// 	result := Subtract(5, 3)
// 	expected := 2
// 	assert.Equal(t, expected, result, "Subtract(5, 3) should be 2")
// }

// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5

// 	// Menggunakan assert untuk assertion yang tidak menghentikan eksekusi
// 	assert.Equal(t, expected, result, "Add(2, 3) should be 5")

// 	// Menggunakan require untuk assertion yang menghentikan eksekusi jika gagal
// 	require.Equal(t, expected, result, "Add(2, 3) should be 5")

// 	// Tes lanjutan yang hanya dijalankan jika assertion require berhasil
// 	result = Add(-1, 1)
// 	expected = 0
// 	assert.Equal(t, expected, result, "Add(-1, 1) should be 0")
// }

// func TestSubtract(t *testing.T) {
// 	result := Subtract(5, 3)
// 	expected := 2

// 	// Menggunakan assert
// 	assert.Equal(t, expected, result, "Subtract(5, 3) should be 2")

// 	// Menggunakan require
// 	require.Equal(t, expected, result, "Subtract(5, 3) should be 2")

// 	// Tes lanjutan
// 	result = Subtract(10, 5)
// 	expected = 5
// 	assert.Equal(t, expected, result, "Subtract(10, 5) should be 5")
// }

// TestIsEvenTrue menguji apakah bilangan genap diidentifikasi dengan benar
func TestIsEvenTrue(t *testing.T) {
	result := IsEven(4)
	expected := true
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// TestIsEvenFalse menguji apakah bilangan ganjil diidentifikasi dengan benar
// Menggunakan t.Skip untuk melewati tes berdasarkan kondisi
func TestIsEvenFalse(t *testing.T) {
	// Kondisi yang memicu skip (misalnya, jika bilangan adalah bilangan genap)
	if true { // Gantilah kondisi ini dengan logika sebenarnya jika diperlukan
		t.Skip("Skipping test for odd numbers")
	}
	result := IsEven(5)
	expected := false
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// TestIsEvenEdge menguji batas kasus (misalnya, bilangan 0)
// Menggunakan t.Skipf untuk melewati tes dengan pesan format
func TestIsEvenEdge(t *testing.T) {
	// Kondisi yang memicu skip
	if true { // Gantilah kondisi ini dengan logika sebenarnya jika diperlukan
		t.Skipf("Skipping edge case test for %d", 0)
	}
	result := IsEven(0)
	expected := true
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
