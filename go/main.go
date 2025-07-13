package main

import (
	"cryptonomicon/tree"
	"fmt"
)

func main() {

	fmt.Println("=== Merkle Proof Data Flow Demo ===")
	fmt.Println("Demonstrating what each actor knows and how information flows")

	// === SETUP: Create the blockchain scenario ===
	fmt.Println("\n🌐 BLOCKCHAIN NETWORK: Creating Block #12345...")
	allTransactions := tree.CreateSampleLedger(8) // 8 transactions in the block

	// === ALICE (Full Node) - Has complete information ===
	fmt.Println("\n👩‍💻 ALICE (Full Node): Building complete tree...")
	alice, err := tree.SimulateFullNode(allTransactions, 12345)
	if err != nil {
		fmt.Printf("Error creating full node: %v\n", err)
		return
	}

	fmt.Printf("Alice knows:\n")
	fmt.Printf("  • Complete tree with %d transactions\n", len(alice.AllTransactions))
	fmt.Printf("  • Root hash: %x...\n", alice.RootHash[:8])
	fmt.Printf("  • Block number: %d\n", alice.BlockNumber)
	fmt.Printf("  • All transaction details:\n")
	for i, tx := range alice.AllTransactions {
		fmt.Printf("    %d: %s (%s → %s, $%.2f)\n", i, tx.Id, tx.Payer, tx.Recipient, tx.Amount)
	}

	// === BLOCKCHAIN NETWORK - Provides consensus ===
	fmt.Println("\n🔗 BLOCKCHAIN NETWORK: Publishing block header...")
	blockHeader := tree.SimulateBlockchain(alice.RootHash, 12345)
	fmt.Printf("Network consensus provides:\n")
	fmt.Printf("  • Block #%d root hash: %x...\n", blockHeader.BlockNumber, blockHeader.RootHash[:8])
	fmt.Printf("  • Timestamp: %d\n", blockHeader.Timestamp)

	// === BOB (Light Client) - Has minimal information ===
	fmt.Println("\n👨‍💼 BOB (Light Client): I want to verify my transaction...")
	myTransaction := allTransactions[2] // Bob thinks TX003 is his transaction
	bob := tree.SimulateLightClient(blockHeader.RootHash, myTransaction, 12345)

	fmt.Printf("Bob knows:\n")
	fmt.Printf("  • His transaction: %s (%s → %s, $%.2f)\n",
		bob.MyTransaction.Id, bob.MyTransaction.Payer, bob.MyTransaction.Recipient, bob.MyTransaction.Amount)
	fmt.Printf("  • Trusted root hash: %x... (from blockchain)\n", bob.TrustedRootHash[:8])
	fmt.Printf("  • Block number: %d\n", bob.BlockNumber)
	fmt.Printf("  • Bob does NOT know: other transactions, tree structure, sibling hashes\n")

	// === PROOF REQUEST: Bob asks Alice for proof ===
	fmt.Println("\n📤 BOB → ALICE: Requesting proof...")
	proofRequest := tree.CreateProofRequest(bob.MyTransaction, bob.BlockNumber)
	fmt.Printf("Bob's request:\n")
	fmt.Printf("  • Transaction to prove: %s\n", proofRequest.Transaction.Id)
	fmt.Printf("  • Block number: %d\n", proofRequest.BlockNumber)

	// === PROOF RESPONSE: Alice processes the request ===
	fmt.Println("\n📥 ALICE → BOB: Processing request...", proofRequest.Transaction.Id)
	proofResponse := tree.ProcessProofRequest(alice, proofRequest)
	fmt.Printf("Alice's response:\n")
	fmt.Printf("  • Found: %v\n", proofResponse.Found)
	fmt.Printf("  • Message: %s\n", proofResponse.Message)
	fmt.Printf("  • Root hash: %x...\n", proofResponse.RootHash[:8])

	if proofResponse.Found {
		fmt.Printf("  • Proof path: %d elements\n", len(proofResponse.ProofPath))
		for i, element := range proofResponse.ProofPath {
			side := "right"
			if element.IsLeft {
				side = "left"
			}
			fmt.Printf("    %d: %s sibling = %x...\n", i, side, element.Hash[:8])
		}
	}

	// === VERIFICATION: Bob verifies the proof ===
	fmt.Println("\n🔍 BOB: Verifying proof...")
	if proofResponse.Found {
		fmt.Printf("Bob now has:\n")
		fmt.Printf("  • His transaction: %s\n", bob.MyTransaction.Id)
		fmt.Printf("  • Trusted root hash: %x...\n", bob.TrustedRootHash[:8])
		fmt.Printf("  • Proof path from Alice: %d elements\n", len(proofResponse.ProofPath))
		tree.VerifyMerkleProof(proofResponse.ProofPath, alice.RootHash, bob.MyTransaction)

		// TODO: Once VerifyMerkleProof is implemented, Bob would call:
		// isValid := tree.VerifyMerkleProof(bob.TrustedRootHash, bob.MyTransaction, proofResponse.ProofPath)
		// fmt.Printf("  • Verification result: %v\n", isValid)

		fmt.Printf("  • Bob can now verify without trusting Alice!\n")
	} else {
		fmt.Printf("  • Transaction not found in block - Bob's transaction is not included\n")
	}

	// === SUMMARY ===
	fmt.Println("\n📋 SUMMARY:")
	fmt.Printf("• Alice (Full Node): Knows everything, generates proofs\n")
	fmt.Printf("• Bob (Light Client): Knows minimal info, verifies proofs\n")
	fmt.Printf("• Blockchain Network: Provides trusted root hashes\n")
	fmt.Printf("• Security: Bob doesn't need to trust Alice's claims\n")
	fmt.Printf("• Efficiency: Bob only needs log(n) proof elements, not all transactions\n")
}
