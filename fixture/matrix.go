package fixture

import (
	"math/rand"
)

// MakeSymMatrix generates a random, symmetric matrix.
func MakeSymMatrix(m uint32) []float64 {
	M := make([]float64, m*m)

	for i := uint32(0); i < m; i++ {
		M[i*m+i] = rand.Float64()

		for j := i + 1; j < m; j++ {
			M[i*m+j] = rand.Float64()
			M[j*m+i] = M[i*m+j]
		}
	}

	return M
}
