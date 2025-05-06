// See the LICENSE file for license details.

package shamir

import (
	"crypto/rand"
)

// -----------------------------------------------------------------------------

const (
	generateOffsetListSwapRounds = 100
	generateOffsetListSwapCount  = 100
)

// -----------------------------------------------------------------------------

func generateOffsetList(size byte) ([]byte, error) {
	var swap [2 * generateOffsetListSwapCount]byte

	offsetList := make([]byte, size)
	for i := byte(0); i < size; i++ {
		offsetList[i] = i
	}

	for r := 0; r < generateOffsetListSwapRounds; r++ {
		_, err := rand.Read(swap[:])
		if err != nil {
			return nil, err
		}
		for i := 0; i < 2*generateOffsetListSwapCount; i += 2 {
			x := swap[i] % size
			y := swap[i+1] % size
			offsetList[x], offsetList[y] = offsetList[y], offsetList[x]
		}
	}

	// Done
	return offsetList, nil
}
