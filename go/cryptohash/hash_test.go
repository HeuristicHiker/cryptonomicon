package cryptohash

import (
	"bytes"
	"testing"
)

func TestFindColission(t *testing.T) {
	a, b, err := FindColission(1000000)

	// test determinism
	if !bytes.Equal(a, b) {
		t.Error("a is not equal to b but should be ya goof")
	}

	// test pre-fix collision
	if err != nil {
		t.Error("Expected FindColission(16) not fuck up but we here now bois", err)
	}

	// TODO use separate encodings and compare maybe? Probably need to think of a good way to test a "bad" encoding that would have conflicts in tests
	// Future Connor - trouver un mauvais algorithme d'encodage

}
