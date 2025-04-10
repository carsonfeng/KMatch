package hungarian_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	hungarian "github.com/carsonfeng/KMatch"
)

var testsMax = []struct {
	name   string
	m      [][]float64
	result map[int]map[int]float64
}{
	{"Default", [][]float64{
		{6, 2, 3, 4, 5},
		{3, 8, 2, 8, 1},
		{9.1, 9, 5, 4, 2},
		{6, 7, 3, 4, 3},
		{1, 2, 6, 4, 9},
	}, map[int]map[int]float64{
		0: {2: 3},
		1: {3: 8},
		2: {0: 9.1},
		3: {1: 7},
		4: {4: 9},
	}},

	{"KMatch Test Case 1", [][]float64{
		{2, 100},
		{100, hungarian.NegativeInf},
	}, map[int]map[int]float64{
		0: {1: 100},
		1: {0: 100},
	}},

	{"KMatch Test Case 2", [][]float64{
		{-0.1, -0.1},
		{-0.1, 0},
	}, map[int]map[int]float64{
		0: {0: -0.1},
		1: {1: 0},
	}},

	{"KMatch Test Case 3", [][]float64{
		{3.16, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{2.58, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{2.35, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{1.76, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{7.9, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
	}, map[int]map[int]float64{
		4: {0: 7.9},
	}},

	{"KMatch Test Case 4", [][]float64{
		{3.16, 4.16, 5.16},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{hungarian.NegativeInf, hungarian.NegativeInf, 6.16},
	}, map[int]map[int]float64{
		0: {1: 4.16},
		2: {2: 6.16},
	}},

	{"KMatch Test Case 5 - Rows < Columns", [][]float64{
		{5.8, 6.9, 8.3, 7.2},
	}, map[int]map[int]float64{
		0: {2: 8.3},
	}},

	{"KMatch Test Case 5.2 - Rows < Columns", [][]float64{
		{5.8, 6.9, 7.2, 8.3},
	}, map[int]map[int]float64{
		0: {3: 8.3},
	}},

	{"KMatch Test Case 5.3 - Rows < Columns", [][]float64{
		{8.3, 5.8, 6.9, 7.2},
	}, map[int]map[int]float64{
		0: {0: 8.3},
	}},

	{"KMatch Test Case 6 - Rows > Columns", [][]float64{
		{5.8},
		{6.9},
		{8.3},
		{7.2},
	}, map[int]map[int]float64{
		2: {0: 8.3},
	}},

	{"KMatch Test Case 6.1 - Rows > Columns", [][]float64{
		{5.8},
		{6.9},
		{7.2},
		{8.3},
	}, map[int]map[int]float64{
		3: {0: 8.3},
	}},

	{"KMatch Test Case 6.2 - Rows > Columns", [][]float64{
		{8.3},
		{5.8},
		{6.9},
		{7.2},
	}, map[int]map[int]float64{
		0: {0: 8.3},
	}},

	{"KMatch Test Case 7", [][]float64{
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, 3.16},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, 2.58},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, 2.35},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, 1.76},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, 7.9},
	}, map[int]map[int]float64{
		4: {4: 7.9},
	}},

	{"KMatch Test Case 8", [][]float64{
		{20.888, 100, hungarian.NegativeInf, hungarian.NegativeInf},
		{100, 20, hungarian.NegativeInf, hungarian.NegativeInf},
		{hungarian.NegativeInf, hungarian.NegativeInf, 20.888, 21},
		{hungarian.NegativeInf, hungarian.NegativeInf, 21, 20.888},
	}, map[int]map[int]float64{
		0: {1: 100},
		1: {0: 100},
		2: {3: 21},
		3: {2: 21},
	}},

	{"KMatch Test Case 9 - Matrix with Weight 0 (but Connectable)", [][]float64{
		{0, 1, 2},
		{1, 0, 3},
		{2, 3, 0},
	}, map[int]map[int]float64{
		0: {0: 0},
		1: {2: 3},
		2: {1: 3},
	}},
	// 10 New Test Cases
	{"KMatch Test Case 10", [][]float64{
		{3, 1},
		{1, 4},
	}, map[int]map[int]float64{
		0: {0: 3},
		1: {1: 4},
	}},

	{"KMatch Test Case 11", [][]float64{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}, map[int]map[int]float64{
		0: {2: 3},
		1: {0: 3},
		2: {1: 3},
	}},

	{"KMatch Test Case 12", [][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}, map[int]map[int]float64{
		0: {0: 1},
		1: {1: 1},
		2: {2: 1},
	}},

	{"KMatch Test Case 13", [][]float64{
		{0, hungarian.NegativeInf, hungarian.NegativeInf},
		{hungarian.NegativeInf, 0, hungarian.NegativeInf},
		{hungarian.NegativeInf, hungarian.NegativeInf, 100},
	}, map[int]map[int]float64{
		0: {0: 0},
		1: {1: 0},
		2: {2: 100},
	}},

	{"KMatch Test Case 14", [][]float64{
		{0.08, 0.1},
		{0.13, 0.14},
	}, map[int]map[int]float64{
		0: {1: 0.1},
		1: {0: 0.13},
	}},

	{"KMatch Test Case 15", [][]float64{
		{0.15, 0.1125},
	}, map[int]map[int]float64{
		0: {0: 0.15},
	}},

	{"KMatch Test Case 16", [][]float64{
		{0.1125, 0.15},
	}, map[int]map[int]float64{
		0: {1: 0.15},
	}},

	{"KMatch Test Case 17", [][]float64{
		{2.64, 3.168},
		{3.2, 3.84},
		{2.4, 2.88},
	}, map[int]map[int]float64{
		0: {0: 2.64},
		1: {1: 3.84},
	}},

	{"KMatch Test Case 18 - No Edges, Should Not Crash", [][]float64{
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
		{hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf, hungarian.NegativeInf},
	}, map[int]map[int]float64{}},

	{"KMatch Test Case 19 - Empty Matrix, Should Not Crash", [][]float64{}, map[int]map[int]float64{}},

	{"KMatch Test Case 20 - Max Weight Priority", [][]float64{
		{1, 100},
		{hungarian.NegativeInf, 2},
		{hungarian.NegativeInf, 3},
	}, map[int]map[int]float64{
		0: {1: 100},
	}},
}

func sumWeight(m map[int]map[int]float64) float64 {
	total := 0.0
	for _, colMap := range m {
		for _, val := range colMap {
			total += val
		}
	}
	return total
}

func TestSolveMax(t *testing.T) {
	for i, value := range testsMax {
		calculate := hungarian.SolveMax(value.m)
		t.Logf("Case %d, Test Name: %s, Match Count: %d", i, value.name, len(calculate))
		t.Logf("Case %d SolveMax Result: %v", i, calculate)

		expectedSum := sumWeight(value.result)
		actualSum := sumWeight(calculate)

		// 🎯 Match count can be compared, but more importantly check the total weight
		if math.Abs(expectedSum-actualSum) > 1e-6 {
			t.Errorf("❌ Case %d FAILED: Total weight mismatch, expected = %.6f, got = %.6f", i, expectedSum, actualSum)
		} else {
			t.Logf("✅ Case %d passed → Total weight OK (%.6f)\n", i, actualSum)
		}
	}
}

func BenchmarkSolveMax8x8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hungarian.SolveMax([][]float64{
			{6, 2, 3, 4, 5, 11, 3, 8},
			{3, 8, 2, 8, 1, 12, 5, 4},
			{7, 9, 5, 10, 2, 11, 6, 8},
			{6, 7, 3, 4, 3, 5, 5, 3},
			{1, 2, 6, 13, 9, 11, 3, 6},
			{6, 2, 3, 4, 5, 11, 3, 8},
			{4, 6, 8, 9, 7, 1, 5, 3},
			{9, 1, 2, 5, 2, 7, 3, 8},
		})
	}
}

func BenchmarkSolveMax10x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hungarian.SolveMax([][]float64{
			{6, 2, 3, 4, 5, 11, 3, 8, 15, 18},
			{3, 8, 2, 12, 33, 8, 1, 12, 5, 4},
			{7, 9, 5, 11, 10, 2, 22, 11, 6, 8},
			{6, 7, 3, 4, 32, 3, 5, 5, 23, 3},
			{1, 2, 21, 6, 13, 9, 11, 3, 18, 6},
			{6, 2, 17, 3, 4, 41, 5, 11, 3, 8},
			{4, 6, 13, 8, 9, 7, 27, 1, 5, 3},
			{9, 1, 2, 16, 5, 2, 7, 31, 3, 8},
			{7, 1, 13, 8, 9, 4, 27, 6, 5, 3},
			{9, 2, 6, 16, 5, 1, 7, 31, 3, 8},
		})
	}
}

func BenchmarkSolveMax12x12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hungarian.SolveMax([][]float64{
			{6, 2, 72, 3, 4, 5, 11, 3, 19, 8, 15, 18},
			{3, 8, 2, 18, 12, 33, 8, 1, 34, 12, 5, 4},
			{7, 9, 5, 11, 10, 51, 2, 22, 11, 6, 15, 8},
			{6, 7, 3, 4, 32, 3, 5, 9, 5, 16, 23, 3},
			{1, 12, 2, 21, 6, 13, 9, 11, 17, 3, 18, 6},
			{6, 2, 16, 37, 17, 3, 4, 41, 5, 11, 3, 8},
			{4, 15, 6, 13, 8, 9, 7, 19, 27, 1, 5, 3},
			{9, 1, 73, 39, 2, 16, 5, 2, 7, 31, 3, 8},
			{6, 2, 72, 3, 4, 5, 11, 3, 19, 8, 15, 18},
			{3, 8, 2, 18, 12, 33, 8, 1, 34, 12, 5, 4},
			{7, 9, 5, 11, 10, 51, 2, 22, 11, 6, 15, 8},
			{6, 7, 3, 4, 32, 3, 5, 9, 5, 16, 23, 3},
		})
	}
}

func generateRandomMatrix(size int) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
		for j := range matrix[i] {
			r := rand.Float64()
			matrix[i][j] = r * 100
			if r > 0.5 {
				matrix[i][j] = -matrix[i][j]
			}
			if r < 0.1 {
				matrix[i][j] = hungarian.NegativeInf
			}
		}
	}
	return matrix
}

func BenchmarkSolveMax100x100(b *testing.B) {
	matrix := generateRandomMatrix(100)
	for i := 0; i < b.N; i++ {
		hungarian.SolveMax(matrix)
	}
}

func BenchmarkSolveMax200x200(b *testing.B) {
	matrix := generateRandomMatrix(200)
	for i := 0; i < b.N; i++ {
		hungarian.SolveMax(matrix)
	}
}

func BenchmarkSolveMax300x300(b *testing.B) {
	matrix := generateRandomMatrix(300)
	for i := 0; i < b.N; i++ {
		hungarian.SolveMax(matrix)
	}
}

func BenchmarkSolveMax500x500(b *testing.B) {
	matrix := generateRandomMatrix(500)
	for i := 0; i < b.N; i++ {
		hungarian.SolveMax(matrix)
	}
}

//func BenchmarkSolveMax1000x1000(b *testing.B) {
//	matrix := generateRandomMatrix(1000)
//	for i := 0; i < b.N; i++ {
//		hungarian.SolveMax(matrix)
//	}
//}
//
//func BenchmarkSolveMax10000x10000(b *testing.B) {
//	matrix := generateRandomMatrix(10000)
//	for i := 0; i < b.N; i++ {
//		hungarian.SolveMax(matrix)
//	}
//}
