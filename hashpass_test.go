package sha256password

import (
	"math/rand"
	"testing"
	"time"
)

func randomPass(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
	}
	return string(runes)
}

func TestRandomHash(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 10000; i++ {
		pass := randomPass(rng)
		hashpass, _ := EncryptPassword(pass)
		if !CompWithEncrypted(pass, hashpass) {
			t.Errorf("CheskPass(%q, %q) = false", pass, hashpass)
		}
	}
}
