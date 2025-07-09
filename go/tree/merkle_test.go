package tree

import "testing"

func TestHashPair(t *testing.T) {
	tests := []struct {
		left, right [32]byte
		want        [32]byte
	}{
		// Add test cases here
		// What should the expected hash output be for these inputs?
		// Try to compute the expected value using the same hash function as in hashPair.
		// For now, let's use a placeholder value (all zeros) and challenge ourselves to replace it with the real hash.
		{left: [32]byte{1}, right: [32]byte{2}, want: [32]byte{0x01}},
	}
	for _, tc := range tests {
		got := hashPair(tc.left, tc.right)
		if got != tc.want {
			t.Errorf("hashPair(%v, %v) = %v, want %v", tc.left, tc.right, got, tc.want)
		}
	}
}
