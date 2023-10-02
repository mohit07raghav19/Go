package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMultiplyMatrix(t *testing.T) {
	// Test case with compatible NULL matrices
	t.Run("NULL Matrices", func(t *testing.T) {
		mat1 := [][]int{{}}
		mat2 := [][]int{{}}

		result, err := matrix.MultiplyMatrices(mat1, mat2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expected := [][]int{{}}
		if !matrix.MatrixEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	// Test case with compatible matrices
	t.Run("Compatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2, 3}, {4, 5, 6}}
		mat2 := [][]int{{7, 8}, {9, 10}, {11, 12}}

		result, err := matrix.MultiplyMatrices(mat1, mat2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expected := [][]int{{58, 64}, {139, 154}}
		if !matrix.MatrixEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	
}

func TestMultiplyIncompatibleMatrix(t *testing.T) {
	// Test case with incompatible matrices
	t.Run("Incompatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2, 3}, {4, 5, 6}}
		mat2 := [][]int{{7, 8}, {9, 10}}

		_, err := matrix.MultiplyMatrices(mat1, mat2)
		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})

	t.Run("Incompatible Matrices", func(t *testing.T) {
		mat1 := [][]int{{1, 2}}
		mat2 := [][]int{{}}

		_, err := matrix.MultiplyMatrices(mat1, mat2)
		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})
}