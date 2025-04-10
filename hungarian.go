package hungarian

import (
	"math"
	"sort"
)

var NegativeInf = math.Inf(-1)

// SolveMax is the unified entry point that automatically determines if perfect matching is possible
// and chooses KM algorithm or falls back to greedy matching
func SolveMax(cost [][]float64) map[int]map[int]float64 {
	if canPerfectMatch(cost) {
		return kmSolve(cost)
	}
	return greedyGlobalMax(cost)
}

// ==== KM algorithm implementation for maximum weight perfect matching ====

func kmSolve(cost [][]float64) map[int]map[int]float64 {
	matrix := padding(cost)
	n := len(matrix)

	xy := make([]int, n)
	yx := make([]int, n)
	for i := range xy {
		xy[i] = -1
		yx[i] = -1
	}

	lx := make([]float64, n)
	ly := make([]float64, n)
	for i := range lx {
		lx[i] = NegativeInf
		for j := 0; j < n; j++ {
			if matrix[i][j] > lx[i] && !math.IsInf(matrix[i][j], -1) {
				lx[i] = matrix[i][j]
			}
		}
	}

	for root := 0; root < n; root++ {
		slack := make([]float64, n)
		slackx := make([]int, n)
		prev := make([]int, n)
		S := make([]bool, n)
		T := make([]bool, n)
		queue := []int{root}
		S[root] = true

		for i := range slack {
			slack[i] = math.Inf(1)
			slackx[i] = -1
			prev[i] = -1
		}

		var ty = -1
		found := false

		for !found {
			for len(queue) > 0 && !found {
				x := queue[0]
				queue = queue[1:]
				for y := 0; y < n; y++ {
					if T[y] || math.IsInf(matrix[x][y], -1) {
						continue
					}
					delta := lx[x] + ly[y] - matrix[x][y]
					if math.Abs(delta) < 1e-9 {
						prev[y] = x
						if yx[y] == -1 {
							ty = y
							found = true
							break
						}
						T[y] = true
						queue = append(queue, yx[y])
						S[yx[y]] = true
					} else if slack[y] > delta {
						slack[y] = delta
						slackx[y] = x
					}
				}
			}

			if found {
				break
			}

			delta := math.Inf(1)
			for y := 0; y < n; y++ {
				if !T[y] && slack[y] < delta {
					delta = slack[y]
				}
			}

			// Infinite loop protection
			if math.IsInf(delta, 1) {
				goto BUILD_RESULT
			}

			for i := 0; i < n; i++ {
				if S[i] {
					lx[i] -= delta
				}
			}
			for j := 0; j < n; j++ {
				if T[j] {
					ly[j] += delta
				} else {
					slack[j] -= delta
				}
			}

			for y := 0; y < n; y++ {
				if !T[y] && slack[y] < 1e-9 {
					prev[y] = slackx[y]
					if yx[y] == -1 {
						ty = y
						found = true
						break
					}
					T[y] = true
					queue = append(queue, yx[y])
					S[yx[y]] = true
				}
			}
		}

		if ty < 0 {
			continue
		}

		// Augmenting path backtracking
		for ty >= 0 {
			x := prev[ty]
			tmp := xy[x]
			xy[x] = ty
			yx[ty] = x
			ty = tmp
		}
	}

BUILD_RESULT:
	result := make(map[int]map[int]float64)
	for x := 0; x < len(cost); x++ {
		y := xy[x]
		if y >= 0 && y < len(cost[x]) && !math.IsInf(cost[x][y], -1) {
			if result[x] == nil {
				result[x] = make(map[int]float64)
			}
			result[x][y] = cost[x][y]
		}
	}
	return result
}

// ==== 👇 Greedy strategy: Global maximum weight edge priority matching (for non-perfect graphs) ====

type edge struct {
	i, j   int
	weight float64
}

func greedyGlobalMax(cost [][]float64) map[int]map[int]float64 {
	edges := []edge{}
	for i := range cost {
		for j := 0; j < len(cost[i]); j++ {
			if !math.IsInf(cost[i][j], -1) {
				edges = append(edges, edge{i, j, cost[i][j]})
			}
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].weight > edges[b].weight
	})

	usedRow := make(map[int]bool)
	usedCol := make(map[int]bool)
	result := make(map[int]map[int]float64)

	for _, e := range edges {
		if usedRow[e.i] || usedCol[e.j] {
			continue
		}
		if result[e.i] == nil {
			result[e.i] = make(map[int]float64)
		}
		result[e.i][e.j] = e.weight
		usedRow[e.i] = true
		usedCol[e.j] = true
	}

	return result
}

// ==== Simple usable perfect matching determination ====

func canPerfectMatch(cost [][]float64) bool {
	rows := len(cost)
	cols := 0
	for _, row := range cost {
		if len(row) > cols {
			cols = len(row)
		}
	}
	n := rows
	if cols > n {
		n = cols
	}

	// Count how many rows and how many cols have at least one valid (non -Inf) edge
	rowCount := 0
	colHasEdge := make(map[int]bool)
	for i := 0; i < len(cost); i++ {
		hasEdge := false
		for j := 0; j < len(cost[i]); j++ {
			if !math.IsInf(cost[i][j], -1) {
				hasEdge = true
				colHasEdge[j] = true
			}
		}
		if hasEdge {
			rowCount++
		}
	}
	// If there aren't enough usable rows and cols, can't form perfect match
	return rowCount >= n && len(colHasEdge) >= n
}

// ==== Automatic conversion to square matrix (negative infinity padding) ====

func padding(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	maxCol := 0
	for _, row := range matrix {
		if len(row) > maxCol {
			maxCol = len(row)
		}
	}
	n := rows
	if maxCol > n {
		n = maxCol
	}

	out := make([][]float64, n)
	for i := 0; i < n; i++ {
		out[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i < len(matrix) && j < len(matrix[i]) {
				out[i][j] = matrix[i][j]
			} else {
				out[i][j] = NegativeInf
			}
		}
	}
	return out
}

// NewBiGraph creates a new bipartite graph
func NewBiGraph(row, col int) (G [][]float64) {
	G = make([][]float64, row)
	for i := range G {
		G[i] = make([]float64, col)
	}
	return
}
