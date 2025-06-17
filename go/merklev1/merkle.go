package merkle

import (
	"crypto/sha256"
	"fmt"
)

type Node interface {
	Hash() [32]byte
}

type Leaf struct {
	Digest       [32]byte
	asByteStream []byte
	plainText    string
}

type Internal struct {
	Digest      [32]byte
	Left, Right Node
}

type Tree struct {
	Root Node
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
	return &Leaf{
		Digest:       digest,
		asByteStream: data,
		plainText:    string(data),
	}
}

// hashNode computes the SHA-256 hash of two child node hashes.
func hashNode(left, right [32]byte) [32]byte {
	// concatenate both digests and hash
	combined := append(left[:], right[:]...)
	sum := sha256.Sum256(combined)
	return sum
}

func BuildTree(ledger []string) *Tree {
	// I want to just pass in a ledger homie

	data := ConvertToBytes(ledger)

	fmt.Printf("Wow my guy looks like you have %d transactions. Better verify them bois\n", len(ledger))
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

	// fmt.Println("Root: ", nodes[0])

	return &Tree{Root: nodes[0]}
}

func CompareTree(ledger1, ledger2 []string) (bool, error) {
	data := ConvertToBytes(ledger1)
	data2 := ConvertToBytes(ledger2)
	fmt.Println("Comparing: ")
	fmt.Println(ledger1)
	fmt.Println("to")
	fmt.Println(ledger2)

	// Build trees for both ledgers
	var nodes1, nodes2 []Node
	for _, d := range data {
		nodes1 = append(nodes1, NewLeaf(d))
	}
	for _, d := range data2 {
		nodes2 = append(nodes2, NewLeaf(d))
	}

	// Compare and print leaf level
	fmt.Printf("Comparing %d vs %d transactions\n", len(nodes1), len(nodes2))
	printComparison(nodes1, nodes2)

	// Build upper levels and compare at each level
	for len(nodes1) > 1 || len(nodes2) > 1 {
		// Handle tree 1
		if len(nodes1) > 1 {
			if len(nodes1)%2 == 1 {
				nodes1 = append(nodes1, nodes1[len(nodes1)-1])
			}
			var nextLevel1 []Node
			for i := 0; i < len(nodes1); i += 2 {
				left, right := nodes1[i], nodes1[i+1]
				digest := hashNode(left.Hash(), right.Hash())
				nextLevel1 = append(nextLevel1, &Internal{
					Digest: digest,
					Left:   left,
					Right:  right,
				})
			}
			nodes1 = nextLevel1
		}

		// Handle tree 2
		if len(nodes2) > 1 {
			if len(nodes2)%2 == 1 {
				nodes2 = append(nodes2, nodes2[len(nodes2)-1])
			}
			var nextLevel2 []Node
			for i := 0; i < len(nodes2); i += 2 {
				left, right := nodes2[i], nodes2[i+1]
				digest := hashNode(left.Hash(), right.Hash())
				nextLevel2 = append(nextLevel2, &Internal{
					Digest: digest,
					Left:   left,
					Right:  right,
				})
			}
			nodes2 = nextLevel2
		}

		printComparison(nodes1, nodes2)
	}

	fmt.Println("\nComparison complete")

	return true, nil
}

// Helper function to print colored comparison
func printComparison(nodes1, nodes2 []Node) {
	maxLen := len(nodes1)
	padLeft := (10 - maxLen) / 2
	padRight := 10 - maxLen - padLeft

	if len(nodes2) > maxLen {
		maxLen = len(nodes2)
	}

	for range padLeft {
		fmt.Printf("[-x-----------]")
	}

	for i := 0; i < maxLen; i++ {
		var hash1, hash2 [32]byte
		var hasNode1, hasNode2 bool

		if i < len(nodes1) {
			hash1 = nodes1[i].Hash()
			hasNode1 = true
		}
		if i < len(nodes2) {
			hash2 = nodes2[i].Hash()
			hasNode2 = true
		}

		if hasNode1 && hasNode2 && hash1 == hash2 {
			fmt.Printf("\033[32m[%p]\033[0m", nodes1[i]) // Green for matching
		} else {
			fmt.Printf("\033[31m[%p]\033[0m", nodes1[i]) // Red for different

		}
	}
	for range padRight {
		fmt.Printf("[-x-----------]")
	}

	fmt.Printf("\n")
}

func PrintLineOfNodes(nodes []Node) {
	nodeLen := len(nodes)
	padLeft := (10 - nodeLen) / 2
	padRight := 10 - nodeLen - padLeft
	for range padLeft {
		fmt.Printf("[-x-----------]")
	}
	for _, node := range nodes {
		fmt.Printf("[%p]", node)
	}
	for range padRight {
		fmt.Printf("[-x-----------]")
	}
	fmt.Printf("\n")
}

// Utils
func ConvertToBytes(data []string) [][]byte {
	result := make([][]byte, len(data))
	for i, s := range data {
		result[i] = []byte(s)
	}
	return result
}

func PrintMerkleCompute(ledger []string, t *Tree) {
	asBytes := ConvertToBytes(ledger)
	fmt.Println("----")
	for _, b := range asBytes {
		fmt.Printf("| %8s %d|", b, len(b))
		for _, p := range b {
			fmt.Printf("%8b |", p)
		}
		fmt.Printf("\n")
	}
}

func NewTreeBuild(ledger []string) {
	// hash all transactions
	data := ConvertToBytes(ledger)

	// set the nodes
	var nodes []Node
	for _, d := range data {
		newLeaf := NewLeaf(d)
		// fmt.Printf("Transaction %s -> Hash: %x\n", newLeaf.plainText, newLeaf.Digest)
		nodes = append(nodes, newLeaf)
	}

	printableNodes := make(map[int][]Node)
	iterations := 0

	printNodeLine(nodes, 10)
	for len(nodes) > 1 {
		// TODO - balance the binary tree - feels meh
		if len(nodes)%2 == 1 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
		var nextLevel []Node
		iterations++

		for i := 0; i < len(nodes); i += 2 {
			left, right := nodes[i], nodes[i+1]
			digest := hashNode(left.Hash(), right.Hash())
			fmt.Printf("%x", digest)
			newInternal := &Internal{
				Digest: digest,
				Left:   left,
				Right:  right,
			}
			nextLevel = append(nextLevel, newInternal)
			printableNodes[iterations] = append(printableNodes[i], newInternal)
			printNodeLine(nextLevel, 10)
			// shortenedHash := nextLevel[0].Hash()
			// fmt.Printf("\033[31m[%x]\033[0m", shortenedHash[28:])
			// fmt.Printf("\033[32m[%x]\033[0m", shortenedHash[28:])
		}
		nodes = nextLevel
		fmt.Println("asdjkfnasklfdm ", iterations)
		for _, j := range printableNodes {
			printNodeLine(j, 10)
		}
		fmt.Println("------")
	}

	// printNodeLine(nodes, 10)

	// pair off transactions
}

func printNodeLine(nodes []Node, maxLen int) {
	lenToPad := len(nodes)
	padLeft := (maxLen - lenToPad) / 2
	padRight := maxLen - lenToPad - padLeft

	for idx := range padLeft {
		if idx == 0 {
			fmt.Printf("[-x------]")
		} else {
			fmt.Printf("[--------]")
		}
	}
	for _, n := range nodes {
		shortenedHash := n.Hash()
		fmt.Printf("[%x]", shortenedHash[28:])
	}
	for idx := range padRight {
		if idx == padRight-1 {
			fmt.Printf("[------x-]")
		} else {
			fmt.Printf("[--------]")
		}
	}
	fmt.Printf("\n")
}
