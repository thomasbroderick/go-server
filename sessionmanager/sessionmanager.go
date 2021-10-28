package sessionmanager

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func CreateNewSession() string {
	// Create a byte slice
	token := make([]byte, 2)
	// Seed the rand function and generate random bytes for token
	rand.Seed(time.Now().UnixNano())
	rand.Read(token)
	// Hash the created token and then encode it into a hex string
	hash := md5.Sum(token)
	pass := hex.EncodeToString(hash[:])
	return pass
}
