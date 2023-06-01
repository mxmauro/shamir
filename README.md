# shamir

This library implements the Shamir's Secret Sharing algorithm in Golang

## Usage

```golang
package main

import (
    "bytes"
    "crypto/rand"

    "github.com/mxmauro/shamir"
)

func main() {
    // Create a random secret
    secret := make([]byte, 16)
    _, _ = rand.Read(secret)

    // Split it in five parts. Three parts are required in order to recreate it.
    secretParts, err := shamir.Split(secret, 5, 3)
    if err != nil {
        // unable to split secret
    }

    // In this example we will take three parts at random and unsorted
    selectedParts := make([][]byte, 3)
    selectedParts[0] = secretParts[1]
    selectedParts[1] = secretParts[4]
    selectedParts[2] = secretParts[2]

    // Rebuild secret
    var recoveredSecret []byte
    recoveredSecret, err = shamir.Combine(selectedParts)
	if err != nil {
		// unable to recombine secret
	}

	if !bytes.Equal(secret, recoveredSecret) {
		// secrets and recovered secret mismatch
	}
}
```

## LICENSE

[MIT](/LICENSE)

