package merkle

import (
	"crypto/sha256"
)

type Node interface {
	Hash() [32]byte
}

type Leaf struct {
	Digest [32]byte // Exported field
}

type Internal struct {
	Digest      [32]byte // Exported field
	Left, Right Node     // Already exported
}

type Tree struct {
	Root Node // Already exported
}

func (l *Leaf) Hash() [32]byte {
	return l.Digest
}

func (n *Internal) Hash() [32]byte {
	return n.Digest
}

// Assume we're hashing transactions using sha256
func NewLeaf(data []byte) *Leaf {
	digest := sha256.Sum256(data)
	return &Leaf{Digest: digest}
}

// hashNode computes the SHA-256 hash of two child node hashes.
func hashNode(left, right [32]byte) [32]byte {
	// concatenate both digests and hash
	combined := append(left[:], right[:]...)
	sum := sha256.Sum256(combined)
	return sum
}

func BuildTree(data [][]byte) *Tree {
	var nodes []Node
	for _, d := range data {
		nodes = append(nodes, NewLeaf(d))
	}

	// build upper levels until root
	for len(nodes) > 1 {
		// if odd number of nodes, duplicate last
		if len(nodes)%2 == 1 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
		var nextLevel []Node
		for i := 0; i < len(nodes); i += 2 {
			left, right := nodes[i], nodes[i+1]
			digest := hashNode(left.Hash(), right.Hash())
			nextLevel = append(nextLevel, &Internal{
				Digest: digest,
				Left:   left,
				Right:  right,
			})
		}
		nodes = nextLevel
	}

	return &Tree{Root: nodes[0]}
}

/*
Okay, I have the goal of getting to where I fully understand merkle trees but let's work backwards on concepts to establish the first principles I can work with.

Ask me a short question on merkle trees then I'll respond with my best guess. Based on that establish the underlying principle I'm missing then we'll iterate with another short question until we get to a place where I understand the concept we end on and we can work our way back
*/
