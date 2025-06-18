package main

import "cryptonomicon/tree"

func main() {
	// ledger := merklev2.CreateSampleLedger()
	// ledgerByteSlice := merklev2.LedgerToByteSlices(ledger)
	// merklev2.NewMerkleTree(ledger.Transactions)

	ledger := tree.CreateSampleLedger()
	head := tree.Node{}

	tree.CreateSimpleBinaryTree(ledger, &head)

}

// return head
// head.left and head.right = last 2 hashes
