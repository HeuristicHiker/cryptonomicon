package main

import (
	"fmt"

	"cryptonomicon/merkle"
	"cryptonomicon/merkle/printer"
)

func main() {
	data := []string{
		"transaction1",
		"transaction2",
		"transaction3",
		"transaction4",
	}

	// Build merkle tree
	tree := merkle.BuildTree(convertToBytes(data))

	// Print the tree in visual format using the printer package
	printer.PrintTree(tree)

	// Optionally print with full hash details
	// fmt.Println("\n")
	printer.PrintTreeDetailed(tree)

	fmt.Printf("\nRoot Hash: %x\n", tree.Root.Hash())
}

// Helper function to convert strings to byte slices
func convertToBytes(data []string) [][]byte {
	result := make([][]byte, len(data))
	for i, s := range data {
		result[i] = []byte(s)
	}
	return result
}
