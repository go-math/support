// Package fixture provides a number of functions for managing fixture data in
// tests.
package fixture

import (
	"math/rand"
)

// MakeSymMatrix returns a random symmetric matrix.
func MakeSymMatrix(n uint) []float64 {
	M := make([]float64, n*n)

	for i := uint(0); i < n; i++ {
		M[i*n+i] = rand.Float64()

		for j := i + 1; j < n; j++ {
			M[i*n+j] = rand.Float64()
			M[j*n+i] = M[i*n+j]
		}
	}

	return M
}
