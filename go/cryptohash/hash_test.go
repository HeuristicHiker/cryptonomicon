package cryptohash

import (
	"bytes"
	"fmt"
	"testing"

	testingpkg "cryptonomicon/testing"
)

func TestFindColission(t *testing.T) {
	testingpkg.PrintHeader("🔐 CRYPTOGRAPHIC COLLISION FINDER TEST")

	testingpkg.PrintProgress("Searching for hash collisions in the digital wilderness...")
	a, b, err := FindColission(1000000)

	// test determinism
	fmt.Printf("%s%s🔍 Testing collision validity...%s\n", testingpkg.Bold, testingpkg.Blue, testingpkg.Reset)
	if !bytes.Equal(a, b) {
		testingpkg.PrintError("Collision check failed - values don't match!")
		t.Error("a is not equal to b but should be ya goof")
		return
	}
	testingpkg.PrintSuccess("Hash collision successfully found! 🎯")

	// test pre-fix collision
	fmt.Printf("%s%s⚡ Verifying error handling...%s\n", testingpkg.Bold, testingpkg.Blue, testingpkg.Reset)
	if err != nil {
		testingpkg.PrintError(fmt.Sprintf("FindColission returned unexpected error: %v", err))
		t.Errorf("Expected FindColission(16) not mess up but we here now bois %e", err)
		return
	}
	testingpkg.PrintSuccess("Error handling verification passed! ✨")

	fmt.Printf("\n%s%s🏆 COLLISION TEST COMPLETE - ALL SYSTEMS NOMINAL! 🏆%s\n", testingpkg.Bold, testingpkg.Magenta, testingpkg.Reset)
	fmt.Printf("%s%s════════════════════════════════════════════════════════%s\n\n", testingpkg.Bold, testingpkg.Magenta, testingpkg.Reset)

	// TODO use separate encodings and compare maybe? Probably need to think of a good way to test a "bad" encoding that would have conflicts in tests
	// Future Connor - trouver un mauvais algorithme d'encodage
}

func TestBirthdayParadoxProof(t *testing.T) {
	testingpkg.PrintHeader("🎂 BIRTHDAY PARADOX PROBABILITY SIMULATOR")

	groupSize := 23      // 23
	simulations := 10000 // 10000

	testingpkg.PrintProbabilityConfiguration(groupSize, simulations)

	probability, err := BirthdayParadoxProof(groupSize, simulations)

	if err != nil {
		testingpkg.PrintError(fmt.Sprintf("Probability calculation failed: %v", err))
		t.Errorf("The probability ist broken %e", err)
		return
	}
	testingpkg.PrintSuccess("Probability calculation completed successfully!")

	if probability < 0.45 || probability > 0.55 {
		testingpkg.PrintError(fmt.Sprintf("Probability %.2f%% is outside expected range for group size %d", probability*100, groupSize))
		t.Errorf("Probability %.2f%% is outside expected range for group size %d", probability*100, groupSize)
		return
	}
	// Create a simple ASCII bar chart
	testingpkg.PrintBarChart(probability)

	testingpkg.PrintProbabilityResult(probability)

}
