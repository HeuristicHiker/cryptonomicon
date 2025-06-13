package cryptohash

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

// Future Connor, refactorisez la sortie de test sophistiquée dans son propre package

// ANSI color codes for fancy output
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Bold    = "\033[1m"
	Blink   = "\033[5m"
)

func printHeader(title string) {
	fmt.Printf("\n%s%s╔══════════════════════════════════════════════════════════════════════╗%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s║  🧪 %-60s  ║%s\n", Bold, Cyan, title, Reset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════════════════════════════╝%s\n\n", Bold, Cyan, Reset)
}

func printProgress(message string) {
	symbols := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	fmt.Printf("%s%s", Yellow, message)
	for i := 0; i < 10; i++ {
		fmt.Printf("\r%s%s %s", Yellow, symbols[i%len(symbols)], message)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r%s✅ %s%s\n", Green, message, Reset)
}

func printSuccess(message string) {
	fmt.Printf("%s%s🎉 SUCCESS: %s%s\n", Bold, Green, message, Reset)
}

func printError(message string) {
	fmt.Printf("%s%s💥 ERROR: %s%s\n", Bold, Red, message, Reset)
}

func TestFindColission(t *testing.T) {
	printHeader("🔐 CRYPTOGRAPHIC COLLISION FINDER TEST")

	printProgress("Searching for hash collisions in the digital wilderness...")
	a, b, err := FindColission(1000000)

	// test determinism
	fmt.Printf("%s%s🔍 Testing collision validity...%s\n", Bold, Blue, Reset)
	if !bytes.Equal(a, b) {
		printError("Collision check failed - values don't match!")
		t.Error("a is not equal to b but should be ya goof")
		return
	}
	printSuccess("Hash collision successfully found! 🎯")

	// test pre-fix collision
	fmt.Printf("%s%s⚡ Verifying error handling...%s\n", Bold, Blue, Reset)
	if err != nil {
		printError(fmt.Sprintf("FindColission returned unexpected error: %v", err))
		t.Errorf("Expected FindColission(16) not mess up but we here now bois %e", err)
		return
	}
	printSuccess("Error handling verification passed! ✨")

	fmt.Printf("\n%s%s🏆 COLLISION TEST COMPLETE - ALL SYSTEMS NOMINAL! 🏆%s\n", Bold, Magenta, Reset)
	fmt.Printf("%s%s════════════════════════════════════════════════════════%s\n\n", Bold, Magenta, Reset)

	// TODO use separate encodings and compare maybe? Probably need to think of a good way to test a "bad" encoding that would have conflicts in tests
	// Future Connor - trouver un mauvais algorithme d'encodage
}

func TestBirthdayParadoxProof(t *testing.T) {
	printHeader("🎂 BIRTHDAY PARADOX PROBABILITY SIMULATOR")

	groupSize := 23      // 23
	simulations := 10000 // 10000

	fmt.Printf("%s%s📊 Configuration:%s\n", Bold, Yellow, Reset)
	fmt.Printf("   👥 Group Size: %s%d people%s\n", Cyan, groupSize, Reset)
	fmt.Printf("   🔄 Simulations: %s%d runs%s\n", Cyan, simulations, Reset)
	fmt.Printf("\n")

	printProgress("Running Monte Carlo simulations across parallel universes...")
	probability, err := BirthdayParadoxProof(groupSize, simulations)

	fmt.Printf("%s%s🧮 Analyzing probability distribution...%s\n", Bold, Blue, Reset)
	if err != nil {
		printError(fmt.Sprintf("Probability calculation failed: %v", err))
		t.Errorf("The probability ist broken %e", err)
		return
	}
	printSuccess("Probability calculation completed successfully!")

	fmt.Printf("%s%s📈 Validating results against theoretical expectations...%s\n", Bold, Blue, Reset)
	if probability < 0.45 || probability > 0.55 {
		printError(fmt.Sprintf("Probability %.2f%% is outside expected range for group size %d", probability*100, groupSize))
		t.Errorf("Probability %.2f%% is outside expected range for group size %d", probability*100, groupSize)
		return
	}
	printSuccess("Results within expected theoretical bounds!")

	// Fancy probability display
	fmt.Printf("\n%s%s┌─────────────────────────────────────────────────────────────────────┐%s\n", Bold, Green, Reset)
	fmt.Printf("%s%s│  🎯 FINAL RESULT: %.2f%% probability of collision! 🎯         │%s\n", Bold, Green, probability*100, Reset)
	fmt.Printf("%s%s└─────────────────────────────────────────────────────────────────────┘%s\n", Bold, Green, Reset)

	// Create a simple ASCII bar chart
	barLength := int(probability * 50)
	fmt.Printf("\n%s%s📊 Probability Visualization:%s\n", Bold, Yellow, Reset)
	fmt.Printf("0%%  ")
	for i := 0; i < 50; i++ {
		if i < barLength {
			fmt.Printf("%s█%s", Green, Reset)
		} else {
			fmt.Printf("░")
		}
	}
	fmt.Printf("  100%%\n")

	fmt.Printf("\n%s%s🌟 BIRTHDAY PARADOX TEST COMPLETE - MATHEMATICS ROCKS! 🌟%s\n", Bold, Magenta, Reset)
	fmt.Printf("%s%s═══════════════════════════════════════════════════════════%s\n\n", Bold, Magenta, Reset)
}
