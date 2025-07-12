package main

import (
	"cryptonomicon/tree"
	"fmt"
)

func main() {
	fmt.Println("=== Testing with 20 transactions from JSON ===")
	ledger := tree.CreateSampleLedger(9)
	fmt.Printf("Loaded %d transactions\n", len(ledger))

	trie, err := tree.CreateMerkleTree(ledger)
	if err != nil {
		fmt.Println(err)
		return
	}

	tree.PrintTreeWithLines(trie)

	fmt.Println("\n--- SECURITY FIX VERIFICATION ---")
	fmt.Println("Testing with 9 transactions to see if the vulnerability is fixed...")

	ledger9 := tree.CreateSampleLedger(50)
	trie9, err := tree.CreateMerkleTree(ledger9)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nBefore modification:")
	tree.PrintTreeWithLines(trie9)

	tree.VerifyMerkleProof(trie, ledger9[40])

	// fmt.Println("\nModifying one instance of the last transaction...")
	// tree.ModifyLeafWithID(trie9, "TX009", "HACKED")

	// fmt.Println("\nAfter modification - checking if only ONE instance changed:")
	// tree.PrintTreeWithLines(trie9)

	// fmt.Println("\n=== Demonstrating Flexible Transaction Loading ===")
	// for _, count := range []int{5, 15, 50, 100} {
	// 	ledgerTest := tree.CreateSampleLedger(count)
	// 	fmt.Printf("Requested %d transactions, got %d\n", count, len(ledgerTest))
	// }
}
