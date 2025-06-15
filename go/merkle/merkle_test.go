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

	fancy.PrintSuccess("Leafs are leafy")

}

func TestBuildTree(t *testing.T) {
	var sampleLedger = []string{
		"Transaction 1",
		"Transaction 2",
		"Transaction 3",
		"Transaction 4",
		"Transaction 5",
		"Transaction 6",
		"Transaction 7",
		"Transaction 8",
		"Transaction 9",
		"Transaction 10",
	}
	fancy.PrintHeader("Verifying a ledger")

}
