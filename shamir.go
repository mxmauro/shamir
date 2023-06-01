package shamir

import (
	"crypto/rand"
	"errors"
)

// -----------------------------------------------------------------------------

const generateOffsetListSwapRounds = 100
const generateOffsetListSwapCount = 100

// -----------------------------------------------------------------------------

// Split divides a secret into several parts to be shared
func Split(secret []byte, numParts, threshold int) ([][]byte, error) {
	// Check parameters
	if threshold < 2 || threshold > 255 {
		return nil, errors.New("threshold value must be between 2 and 255")
	}
	if numParts < threshold || numParts > 255 {
		return nil, errors.New("number of parts must be between threshold value and 255")
	}
	secretLen := len(secret)
	if secretLen == 0 {
		return nil, errors.New("no secret was provided")
	}

	// Generate random list of offsets
	offsetList, err := generateOffsetList(byte(numParts))
	if err != nil {
		return nil, err
	}

	// Prepare each part and initialize the offset
	secretParts := make([][]byte, numParts)
	for idx := range secretParts {
		secretParts[idx] = make([]byte, secretLen+1)
		secretParts[idx][secretLen] = offsetList[idx] + 1
	}

	// Generate a new polynomial for each secret byte
	for idx, val := range secret {
		p := newRandomPolynomial(val, byte(threshold)-1)
		if p == nil {
			return nil, errors.New("unable to generate polynomial")
		}
		for i := 0; i < numParts; i++ {
			secretParts[i][idx] = evaluatePolynomial(p, offsetList[i]+1)
		}
	}

	// Done
	return secretParts, nil
}

// Combine recreates a secret if enough parts are provided
func Combine(secretParts [][]byte) ([]byte, error) {
	// Check parameters
	secretPartsLen := len(secretParts)
	if secretPartsLen < 2 {
		return nil, errors.New("insufficient number of secret parts")
	}
	firstPartsLen := len(secretParts[0])
	if firstPartsLen < 2 {
		return nil, errors.New("secret part size is too small")
	}
	for i := 1; i < len(secretParts); i++ {
		if len(secretParts[i]) != firstPartsLen {
			return nil, errors.New("secret parts must have the same length")
		}
	}
	for i := 0; i < secretPartsLen-1; i++ {
		for j := i + 1; j < secretPartsLen; j++ {
			if secretParts[i][firstPartsLen-1] == secretParts[j][firstPartsLen-1] {
				return nil, errors.New("two or more secret parts are the same")
			}
		}
	}

	// Prepare to save combined secret
	secret := make([]byte, firstPartsLen-1)

	// Calculate each byte
	points := make([]point, secretPartsLen)
	for i := range secret {
		for idx, secretPart := range secretParts {
			points[idx] = point{
				x: secretPart[firstPartsLen-1],
				y: secretPart[i],
			}
		}
		secret[i] = interpolateAt0(points)
	}

	// Done
	return secret, nil
}

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
