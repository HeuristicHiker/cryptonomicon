package main

import (
	"fmt"

	"cryptonomicon/merkle"
)

func main() {
	data := []string{
		"transaction1",
		"transaction2",
		"transaction3",
		"transaction4",
	}

	// Build merkle tree
	tree := merkle.BuildTree(merkle.ConvertToBytes(data))

	// Print the tree in visual format using the printer package
	merkle.PrintTree(tree)

	// Optionally print with full hash details
	// fmt.Println("\n")
	merkle.PrintTreeDetailed(tree)

	fmt.Printf("\nRoot Hash: %x\n", tree.Root.Hash())
}

// Helper function to convert strings to byte slices
