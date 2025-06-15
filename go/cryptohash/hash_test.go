package cryptohash

import (
	"bytes"
	"fmt"
	"testing"

	"cryptonomicon/fancy"
)

func TestFindColission(t *testing.T) {
	fancy.PrintHeader("🔐 CRYPTOGRAPHIC COLLISION FINDER TEST")

	fancy.PrintProgress("Searching for hash collisions in the digital wilderness...")
	a, b, err := FindColission(1000000)

	// test determinism
	fmt.Printf("🔍 Testing collision validity...\n")
	if !bytes.Equal(a, b) {
		fancy.PrintError("Collision check failed - values don't match!")
		t.Error("a is not equal to b but should be ya goof")
		return
	}
	fancy.PrintSuccess("Hash collision successfully found! 🎯")

	// test pre-fix collision
	fmt.Printf("⚡ Verifying error handling...\n")
	if err != nil {
		fancy.PrintError(fmt.Sprintf("FindColission returned unexpected error: %v", err))
		t.Errorf("Expected FindColission(16) not mess up but we here now bois %e", err)
		return
	}
	fancy.PrintSuccess("Error handling verification passed! ✨")

	fmt.Printf("\n🏆 COLLISION TEST COMPLETE - ALL SYSTEMS NOMINAL! 🏆\n")
	fmt.Printf("════════════════════════════════════════════════════════\n\n")

	// TODO use separate encodings and compare maybe? Probably need to think of a good way to test a "bad" encoding that would have conflicts in tests
	// Future Connor - trouver un mauvais algorithme d'encodage
}

func TestBirthdayParadoxProof(t *testing.T) {
	fancy.PrintHeader("🎂 BIRTHDAY PARADOX PROBABILITY SIMULATOR")

	groupSize := 23      // 23
	simulations := 10000 // 10000

	fancy.PrintProbabilityConfiguration(groupSize, simulations)

	probability, err := BirthdayParadoxProof(groupSize, simulations)

	if err != nil {
		fancy.PrintError(fmt.Sprintf("Probability calculation failed: %v", err))
		t.Errorf("The probability ist broken %e", err)
		return
	}
	fancy.PrintSuccess("Probability calculation completed successfully!")

	if probability < 0.45 || probability > 0.55 {
		fancy.PrintError(fmt.Sprintf("Probability %.2f%% is outside expected range for group size %d", probability*100, groupSize))
		t.Errorf("Probability %.2f%% is outside expected range for group size %d", probability*100, groupSize)
		return
	}
	// Create a simple ASCII bar chart
	fancy.PrintBarChart(probability)

	fancy.PrintProbabilityResult(probability)

}
