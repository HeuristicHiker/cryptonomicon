package cryptohash

import (
	"bytes"
	"fmt"
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
		t.Errorf("Expected FindColission(16) not mess up but we here now bois %e", err)
	}

	// TODO use separate encodings and compare maybe? Probably need to think of a good way to test a "bad" encoding that would have conflicts in tests
	// Future Connor - trouver un mauvais algorithme d'encodage

}

func TestBirthdayParadoxProof(t *testing.T) {
	groupSize := 23      // 23
	simulations := 10000 // 10000

	probability, err := BirthdayParadoxProof(groupSize, simulations)

	if err != nil {
		t.Errorf("The probability ist broken %e", err)
	}

	if probability < 0.45 || probability > 0.55 {
		t.Errorf("Probability %.2f%% is outside expected range for group size %d", probability*100, groupSize)
	}
	fmt.Printf(" The probability is %.2f%% \n", 100*probability)

}
