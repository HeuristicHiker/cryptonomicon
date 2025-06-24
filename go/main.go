package main

import (
	"cryptonomicon/tree"
)

func main() {
	// ledger := merklev2.CreateSampleLedger()
	// ledgerByteSlice := merklev2.LedgerToByteSlices(ledger)
	// merklev2.NewMerkleTree(ledger.Transactions)
	ledger := tree.CreateSampleLedger()
	// head := tree.Node{}

	trie := tree.CreateSimpleBinaryTree(ledger)

	tree.PrintFromRoot(trie)

}

// return head
// head.left and head.right = last 2 hashes
