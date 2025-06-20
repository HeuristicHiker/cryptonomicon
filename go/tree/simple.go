package tree

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Hash  [32]byte
}

type Transaction struct {
	id        string
	payer     string
	recipient string
	amount    float64
}

func CreateSimpleBinaryTree(transactions []Transaction) *Node {
	if len(transactions) == 0 {
		return nil
	}

	// Create leaf nodes for each transaction
	nodes := make([]*Node, len(transactions))
	for i, tx := range transactions {
		nodes[i] = &Node{
			Hash: hashTransaction(tx),
		}
	}

	// Build tree level by level until we have one root node
	for len(nodes) > 1 {
		var nextLevel []*Node

		// Process pairs of nodes
		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				// Pair exists
				leftHash := nodes[i].Hash
				rightHash := nodes[i+1].Hash
				combined := append(leftHash[:], rightHash[:]...)
				parentHash := sha256.Sum256(combined)

				parent := &Node{
					Hash:  parentHash,
					Left:  nodes[i],
					Right: nodes[i+1],
				}
				nextLevel = append(nextLevel, parent)
			} else {
				// Odd node, promote to next level
				nextLevel = append(nextLevel, nodes[i])
			}
		}
		nodes = nextLevel
	}

	return nodes[0]
}

func hashTransaction(t Transaction) [32]byte {
	tString := fmt.Sprintf("%s:%s:%s:%.2f", t.id, t.payer, t.recipient, t.amount)
	data := []byte(tString)
	return sha256.Sum256(data)
}

func hashPair(left, right Transaction) [32]byte {
	leftHash := hashTransaction(left)
	rightHash := hashTransaction(right)
	combined := append(leftHash[:], rightHash[:]...)
	return sha256.Sum256(combined)
}

func PrintTree(*Node) {

}

func CreateSampleLedger() []Transaction {
	ledger := []Transaction{
		{
			id:        "1",
			payer:     "Alice",
			recipient: "Bob",
			amount:    5,
		},
		{
			id:        "2",
			payer:     "Bob",
			recipient: "Charlie",
			amount:    2,
		},
		{
			id:        "3",
			payer:     "Charlie",
			recipient: "Dave",
			amount:    1,
		},
		{
			id:        "4",
			payer:     "Dave",
			recipient: "Eve",
			amount:    0.5,
		},
		{
			id:        "5",
			payer:     "Eve",
			recipient: "Alice",
			amount:    0.5,
		},
		{
			id:        "6",
			payer:     "Eve",
			recipient: "Alice",
			amount:    0.5,
		},
		{
			id:        "7",
			payer:     "Eve",
			recipient: "Alice",
			amount:    0.5,
		},
		{
			id:        "8",
			payer:     "Eve",
			recipient: "Alice",
			amount:    0.5,
		},
		{
			id:        "9",
			payer:     "Eve",
			recipient: "Alice",
			amount:    0.5,
		},
		{
			id:        "10",
			payer:     "Eve",
			recipient: "Alice",
			amount:    0.5,
		},
	}
	return ledger
}
