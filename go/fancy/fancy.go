package fancy

import (
	"fmt"
	"time"
)

// ANSI color codes for fancy output
const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
	bold    = "\033[1m"
	blink   = "\033[5m"
)

func PrintHeader(title string) {
	fmt.Printf("\n%s%s╔══════════════════════════════════════════════════════════════════════╗%s\n", bold, cyan, reset)
	fmt.Printf("%s%s║  🧪 %-60s  ║%s\n", bold, cyan, title, reset)
	fmt.Printf("%s%s╚══════════════════════════════════════════════════════════════════════╝%s\n\n", bold, cyan, reset)
}

func PrintProgress(message string) {
	symbols := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	fmt.Printf("%s%s", yellow, message)
	for i := 0; i < 10; i++ {
		fmt.Printf("\r%s%s %s", yellow, symbols[i%len(symbols)], message)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r%s✅ %s%s\n", green, message, reset)
}

func PrintSuccess(message string) {
	fmt.Printf("%s%s🎉 SUCCESS: %s%s\n", bold, green, message, reset)
}

func PrintError(message string) {
	fmt.Printf("%s%s💥 ERROR: %s%s\n", bold, red, message, reset)
}

// For stats
func PrintProbabilityConfiguration(groupSize, simulations int) {
	fmt.Printf("%s%s📊 Configuration:%s\n", bold, yellow, reset)
	fmt.Printf("   👥 Group Size: %s%d people%s\n", cyan, groupSize, reset)
	fmt.Printf("   🔄 Simulations: %s%d runs%s\n", cyan, simulations, reset)
	fmt.Printf("\n")

	PrintProgress("Running simulations across parallel universes...")

	fmt.Printf("%s%s🧮 Analyzing probability distribution...%s\n", bold, blue, reset)
}

func PrintBarChart(probability float64) {
	barLength := int(probability * 50)
	fmt.Printf("\n%s%s📊 Probability Visualization:%s\n", bold, yellow, reset)
	fmt.Printf("%.2f%%  ", probability)
	for i := 0; i < 50; i++ {
		if i < barLength {
			fmt.Printf("%s█%s", green, reset)
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
		bold,
		green,
		reset,
	)
	fmt.Printf("%s%s│  🎯 FINAL RESULT: %.2f%% probability of collision! 🎯         │%s\n",
		bold,
		green,
		probability*100,
		reset,
	)
	fmt.Printf("%s%s└─────────────────────────────────────────────────────────────────────┘%s\n",
		bold,
		green,
		reset,
	)

	fmt.Printf("\n%s%s🌟 BIRTHDAY PARADOX TEST COMPLETE - MATHEMATICS ROCKS! 🌟%s\n", bold, magenta, reset)
	fmt.Printf("%s%s═══════════════════════════════════════════════════════════%s\n\n", bold, magenta, reset)
}

func PrintGreenGiant() {
	fmt.Println("Oh yeah")
	fmt.Println("  🌿🌿🌿🌿🌿🌿🌿 ")
	fmt.Println("🌿🌱🌱🌿🌱🌱🌿🌱🌿")
	fmt.Println("🌿🌱🌱👁️🌱🌱👁️🌱🌿")
	fmt.Println("🌿🌱🌱🌱👃🌱🌱🌱🌿")
	fmt.Println("🌿🌱🌱🌱👄🌱🌱🌱🌿	That's a leaf alright")
	fmt.Println(" 🌿🌱🌱🌱🌱🌱🌱🌿")
	fmt.Println("   🌿🌿🌿🌿🌿🌿")
	fmt.Println("")
}
