package cryptox

import "crypto/sha256"

// SHA256 computes sha256 given bytes arrays
func SHA256(in ...[]byte) []byte {
	algo := sha256.New()
	for _, v := range in {
		algo.Write(v)
	}
	return algo.Sum(nil)
}
