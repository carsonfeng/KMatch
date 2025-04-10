# KMatch

🎯 **KMatch** is a Go implementation of **maximum weight bipartite matching**, built with:

- ✅ [Kuhn-Munkres (Hungarian) algorithm](https://en.wikipedia.org/wiki/Hungarian_algorithm) for finding **perfect matching with maximum total weight**
- ✅ A smart fallback: **global max-weight greedy matcher** when perfect matching is not possible
- ✅ Support for **non-square**, **sparse**, and **invalid (-Inf)** graphs

---

## 🚀 Features

- ✅ **Perfect Matching** with max weight using KM algorithm
- ❎ If perfect matching fails (e.g., in sparse graphs), auto fallback to best greedy matching
- 🎯 Handles non-square matrices with automatic padding
- ❄️ Supports `math.Inf(-1)` to denote unmatchable edges
- 🧪 Full test coverage including correctness and edge conditions
- 📊 Benchmarks included up to large matrix sizes (500x500)

---

## 🗂 Project Structure

```bash
.
├── cmd/
│   └── main.go          # Example CLI usage
├── hungarian.go         # Core algorithm (KM + fallback)
└── hungarian_test.go    # 22+ test cases with expected total weights
```

---

## 📦 Installation

```bash
go get github.com/carsonfeng/KMatch
```

> or clone manually:

```bash
git clone https://github.com/carsonfeng/KMatch.git
cd KMatch
```

---

## 🧪 Running Tests

```bash
go test -v .
```

### 🧪 Run performance benchmarks:

```bash
go test -bench=. -benchmem
```

---

## 🛠 Usage

Here's a ready-to-run minimal usage example:

```go
package main

import (
	"fmt"
	"github.com/carsonfeng/KMatch/hungarian"
)

func main() {
	matrix := [][]float64{
		{1, 100, 2},
		{50, hungarian.NegativeInf, 10},
		{30, 20, 70},
	}

	result := hungarian.SolveMax(matrix)

	fmt.Println("🔗 Matching result:")
	for i, row := range result {
		for j, val := range row {
			fmt.Printf("Row %d → Col %d = %.2f\n", i, j, val)
		}
	}
}
```

Run with:

```bash
go run cmd/main.go
```

---

## 📚 API Documentation

### Variables

- `hungarian.NegativeInf` - A constant equal to `math.Inf(-1)` for denoting invalid/unmatchable edges

### Functions

- `hungarian.SolveMax(cost [][]float64) map[int]map[int]float64` - Main entry point that automatically chooses between KM algorithm or greedy fallback
- `hungarian.NewBiGraph(row, col int) [][]float64` - Helper to create a new bipartite graph matrix

---

## ✅ Example Outputs

```bash
🔗 Matching result:
Row 0 → Col 1 = 100.00
Row 1 → Col 0 = 50.00
Row 2 → Col 2 = 70.00
```

---

## 🎯 How It Works

KMatch uses a dual-strategy approach to handle both perfect and non-perfect bipartite matching scenarios:

1. **Decision Logic:**
   ```text
   SolveMax(matrix) ➞
   ├── if perfect match possible ➞ KM Algorithm (O(n³))
   └── else ➞ Fallback to Greedy Global Max Edge Matching (O(m log m))
   ```

2. **KM Algorithm (Primary):**
   - Used when a perfect matching is possible
   - Ensures maximum total weight while matching every node
   - Optimal solution with O(n³) time complexity
   - Works with the Kuhn-Munkres approach to find max-weight perfect matching

3. **Greedy Fallback (Secondary):**
   - Used when perfect matching is impossible (e.g., sparse graphs with insufficient valid edges)
   - Sorts all edges by weight and matches greedily from highest to lowest
   - Prioritizes maximum weight edges to ensure best possible total weight
   - Not guaranteed to be optimal, but provides best-effort matching

The algorithm automatically determines if perfect matching is possible by checking if there are enough valid edges connecting distinct rows and columns.

---

## 📜 License

Released under the [MIT License](https://opensource.org/licenses/MIT).

---

## 🤝 Contributing

Pull requests are welcome! Please create issues for any bugs or enhancements.

---

## 🔥 Credits

Built by [@carsonfeng](https://github.com/carsonfeng) for real-world robust bipartite matching in task scheduling, recommender matching, and graph optimization applications.
