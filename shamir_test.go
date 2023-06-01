package shamir

import (
	"bytes"
	"crypto/rand"
	"testing"
)

// -----------------------------------------------------------------------------

const shamirTestsIterations = 10000

func TestShamir(t *testing.T) {
	t.Logf("Executing %v tests", shamirTestsIterations)
	for counter := 1; counter <= shamirTestsIterations; counter++ {
		// Create a random secret
		secretLen := 16 + int(getRandomByte())%32
		secret := make([]byte, secretLen)
		_, _ = rand.Read(secret)

		// Split it
		partsCount := 3 + int(getRandomByte())%6
		threshold := 2 + int(getRandomByte())%(partsCount-2)
		secretParts, err := Split(secret, partsCount, threshold)
		if err != nil {
			t.Fatal("unable to split secret:", err)
		}

		// Take a random number of parts (minimum threshold)
		selectedPartsCount := threshold + int(getRandomByte())%(partsCount-threshold)
		selectedParts := make([][]byte, selectedPartsCount)
		selectedPartsCheck := make(map[int]struct{})
		for idx := 0; idx < selectedPartsCount; idx++ {
			var r int

			for {
				r = int(getRandomByte()) % partsCount
				if _, ok := selectedPartsCheck[r]; !ok {
					break
				}
			}

			selectedPartsCheck[r] = struct{}{}
			selectedParts[idx] = secretParts[r]
		}

		// Rebuild secret
		var recoveredSecret []byte
		recoveredSecret, err = Combine(selectedParts)
		if err != nil {
			t.Fatal("unable to recombine secret:", err)
		}

		if !bytes.Equal(secret, recoveredSecret) {
			t.Fatal("secrets and recovered secret mismatch")
		}
	}
}

// -----------------------------------------------------------------------------

func getRandomByte() byte {
	var tempLen [1]byte

	_, _ = rand.Read(tempLen[:])
	return tempLen[0]
}
