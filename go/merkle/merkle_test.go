package merkle

import (
	"crypto/sha256"
	"cryptonomicon/fancy"
	"testing"
)

func TestNewLeaf(t *testing.T) {
	fancy.PrintHeader("Checking if leaves are set all good n stuff")

	t.Run("basic leaf creation", func(t *testing.T) {
		rawData := "Transaction 1"
		data := []byte(rawData)
		testLeaf := NewLeaf(data)

		if len(testLeaf.Hash()) != 32 {
			fancy.PrintError("Leaf should be len 32 * 8 bytes = 256")
			t.Errorf("Expected hash length 32, got %d", len(testLeaf.Hash()))
		}
	})

	t.Run("hash consistency", func(t *testing.T) {
		data := []byte("test data")
		leaf1 := NewLeaf(data)
		leaf2 := NewLeaf(data)

		if leaf1.Hash() != leaf2.Hash() {
			fancy.PrintError("Same data should produce same hash")
			t.Error("Same data should produce same hash")
		}
		fancy.PrintSuccess("Hash consistency verified")
	})

	t.Run("different data produces different hash", func(t *testing.T) {
		leaf1 := NewLeaf([]byte("data1"))
		leaf2 := NewLeaf([]byte("data2"))

		if leaf1.Hash() == leaf2.Hash() {
			fancy.PrintError("Different data should produce different hashes")
			t.Error("Different data should produce different hashes")
		}
		fancy.PrintSuccess("Different hashes for different data")
	})

	t.Run("empty data", func(t *testing.T) {
		leaf := NewLeaf([]byte{})

		if len(leaf.Hash()) != 32 {
			fancy.PrintError("Empty data should still produce 32-byte hash")
			t.Error("Empty data should still produce 32-byte hash")
		}
		fancy.PrintSuccess("Empty data handled correctly")
	})

	t.Run("hash matches expected SHA256", func(t *testing.T) {
		data := []byte("test")
		leaf := NewLeaf(data)
		expected := sha256.Sum256(data)

		if leaf.Hash() != expected {
			fancy.PrintError("Hash should match SHA256 of input data")
			t.Error("Hash should match SHA256 of input data")
		}
		fancy.PrintSuccess("SHA256 hash verified")
	})

	fancy.PrintGreenGiant()

	fancy.PrintHeader("This leaf should be invalid")

	t.Run("hash matches expected SHA256", func(t *testing.T) {
		data := []byte("test")
		data2 := []byte("test2")
		leaf := NewLeaf(data)
		expected := sha256.Sum256(data2)
		if leaf.Hash() == expected {
			t.Error("You failed at failing to create a differnet SHA256. The proability of that means you need to buy a lotery ticket... RIGHT NOW")
			fancy.PrintError("Hash should match SHA256 of input data")
		}
		fancy.PrintSuccess("Invalid SHA256 hash shows as invalid")
		fancy.PrintFireGiant()
	})

}

func TestBuildTree(t *testing.T) {
	var firstLedger = []string{
		"T1",
		"T2",
		"T3",
		"T4",
		"T5",
		"T6",
		"T7",
		"T8",
		"T9",
		"T10",
	}

	// fancy.PrintHeader("Building merkle tree based on ledger")
	// fmt.Println(sampleLedger)

	// Build merkle tree
	// BuildTree(sampleLedger)
	var secondLedger = []string{
		"T1",
		"T2",
		"T3",
		"T4",
		"T5",
		"T6",
		"T7",
		"T10",
		"T9",
		"T10",
	}
	fancy.PrintHeader("Expect the following to be \033[32mValid\033 ledger:")
	CompareTree(firstLedger, firstLedger)
	fancy.PrintHeader("Expect the following to be \033[31mInvalid\033 ledger:")
	CompareTree(firstLedger, secondLedger)
	fancy.PrintSuccess("You FOOL obviously that the transactions are fraudulent do you even SHA my guy???")

	// PrintMerkleCompute(sampleLedger, tree)

}
