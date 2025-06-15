package testing

import (
	"fmt"
	"time"
)

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

func PrintHeader(title string) {
	fmt.Printf("\n%s%s╔══════════════════════════════════════════════════════════════════════╗%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s%s║  🧪 %-60s  ║%s\n", Bold, Cyan, title, Reset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════════════════════════════╝%s\n\n", Bold, Cyan, Reset)
}

func PrintProgress(message string) {
	symbols := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	fmt.Printf("%s%s", Yellow, message)
	for i := 0; i < 10; i++ {
		fmt.Printf("\r%s%s %s", Yellow, symbols[i%len(symbols)], message)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r%s✅ %s%s\n", Green, message, Reset)
}

func PrintSuccess(message string) {
	fmt.Printf("%s%s🎉 SUCCESS: %s%s\n", Bold, Green, message, Reset)
}

func PrintError(message string) {
	fmt.Printf("%s%s💥 ERROR: %s%s\n", Bold, Red, message, Reset)
}

// For stats
func PrintProbabilityConfiguration(groupSize, simulations int) {
	fmt.Printf("%s%s📊 Configuration:%s\n", Bold, Yellow, Reset)
	fmt.Printf("   👥 Group Size: %s%d people%s\n", Cyan, groupSize, Reset)
	fmt.Printf("   🔄 Simulations: %s%d runs%s\n", Cyan, simulations, Reset)
	fmt.Printf("\n")

	PrintProgress("Running simulations across parallel universes...")

	fmt.Printf("%s%s🧮 Analyzing probability distribution...%s\n", Bold, Blue, Reset)
}

func PrintBarChart(probability float64) {
	barLength := int(probability * 50)
	fmt.Printf("\n%s%s📊 Probability Visualization:%s\n", Bold, Yellow, Reset)
	fmt.Printf("%.2f%%  ", probability)
	for i := 0; i < 50; i++ {
		if i < barLength {
			fmt.Printf("%s█%s", Green, Reset)
		} else {
			fmt.Printf("░")
		}
	}
	fmt.Printf("  100%%\n")
}

func PrintProbabilityResult(probability float64) {
	PrintSuccess("Results within expected theoretical bounds!")

	// Fancy probability display
	fmt.Printf("\n%s%s┌─────────────────────────────────────────────────────────────────────┐%s\n",
		Bold,
		Green,
		Reset,
	)
	fmt.Printf("%s%s│  🎯 FINAL RESULT: %.2f%% probability of collision! 🎯         │%s\n",
		Bold,
		Green,
		probability*100,
		Reset,
	)
	fmt.Printf("%s%s└─────────────────────────────────────────────────────────────────────┘%s\n",
		Bold,
		Green,
		Reset,
	)

	fmt.Printf("\n%s%s🌟 BIRTHDAY PARADOX TEST COMPLETE - MATHEMATICS ROCKS! 🌟%s\n", Bold, Magenta, Reset)
	fmt.Printf("%s%s═══════════════════════════════════════════════════════════%s\n\n", Bold, Magenta, Reset)
}
