package main

import (
	"cryptonomicon/tree"
	"fmt"
)

func main() {
	// ledger := merklev2.CreateSampleLedger()
	// ledgerByteSlice := merklev2.LedgerToByteSlices(ledger)
	// merklev2.NewMerkleTree(ledger.Transactions)
	ledger := tree.CreateSampleLedger(11)
	// head := tree.Node{}

	trie, err := tree.CreateMerkleTree(ledger)
	if err != nil {
		fmt.Println(err)
	}

	tree.PrintLevels(trie)

}

// return head
// head.left and head.right = last 2 hashes
