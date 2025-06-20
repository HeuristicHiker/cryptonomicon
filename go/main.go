package main

import (
	"cryptonomicon/tree"
	"fmt"
)

func main() {
	// ledger := merklev2.CreateSampleLedger()
	// ledgerByteSlice := merklev2.LedgerToByteSlices(ledger)
	// merklev2.NewMerkleTree(ledger.Transactions)
	ledger := tree.CreateSampleLedger()
	// head := tree.Node{}

	trie := tree.CreateSimpleBinaryTree(ledger)

	fmt.Println(trie)

}

// return head
// head.left and head.right = last 2 hashes
