package main

import (
	merklev2 "cryptonomicon/merklev2"
)

func main() {
	ledger := merklev2.CreateSampleLedger()
	// ledgerByteSlice := merklev2.LedgerToByteSlices(ledger)
	merklev2.NewMerkleTree(ledger.Transactions)

}
