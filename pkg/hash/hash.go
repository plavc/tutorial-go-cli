package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(in string) string {
	sha := sha256.New()
	sha.Write([]byte(in))
	return hex.EncodeToString(sha.Sum(nil))
}
