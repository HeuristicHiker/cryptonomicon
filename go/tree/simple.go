package tree

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Left     *Node
	Right    *Node
	Hash     [32]byte
	TDetails *Transaction
}

type Transaction struct {
	Id        string
	Payer     string
	Recipient string
	Amount    float64
}

func CreateSimpleBinaryTree(transactions []Transaction) *Node {
	if len(transactions) == 0 {
		return nil
	}

	nodes := CreateLeaves(transactions)

	for len(nodes) > 1 {
		var nextLevel []*Node

		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				// Pair exists
				left := nodes[i]
				right := nodes[i+1]
				parentHash := hashPair(nodes[i].Hash, nodes[i+1].Hash)
				internalTDetail := Transaction{
					Id:        left.TDetails.Id,
					Payer:     left.TDetails.Payer,
					Recipient: left.TDetails.Recipient,
					Amount:    left.TDetails.Amount,
				}

				parent := &Node{
					Hash:     parentHash,
					Left:     left,
					Right:    right,
					TDetails: &internalTDetail,
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
	tString := fmt.Sprintf("%s:%s:%s:%.2f", t.Id, t.Payer, t.Recipient, t.Amount)
	data := []byte(tString)
	return sha256.Sum256(data)
}

func hashPair(leftHash, rightHash [32]byte) [32]byte {
	combined := append(leftHash[:], rightHash[:]...)
	return sha256.Sum256(combined)
}

func CreateLeaves(transactions []Transaction) []*Node {
	nodes := make([]*Node, len(transactions))
	for i, tx := range transactions {
		nodes[i] = &Node{
			Hash:     hashTransaction(tx),
			TDetails: &tx,
		}
	}

	return nodes
}

func PrintFromRoot(root *Node) {
	if root.Left == nil && root.Right == nil {
		return
	}
	fmt.Println(root.TDetails)
	// fmt.Println("Looking at root: ", root.TDetails.Id, " ", root.Left.TDetails.Id, " ", root.Right.TDetails.Id)
	if root.Left != nil {
		PrintFromRoot(root.Left)
	} else {
		PrintFromRoot(root.Right)
	}
}

func CreateSampleLedger() []Transaction {
	ledger := []Transaction{
		{
			Id:        "1",
			Payer:     "Alice",
			Recipient: "Bob",
			Amount:    5,
		},
		{
			Id:        "2",
			Payer:     "Bob",
			Recipient: "Charlie",
			Amount:    2,
		},
		{
			Id:        "3",
			Payer:     "Charlie",
			Recipient: "Dave",
			Amount:    1,
		},
		{
			Id:        "4",
			Payer:     "Dave",
			Recipient: "Eve",
			Amount:    0.5,
		},
		{
			Id:        "5",
			Payer:     "Eve",
			Recipient: "Alice",
			Amount:    0.5,
		},
		{
			Id:        "6",
			Payer:     "Eve",
			Recipient: "Alice",
			Amount:    0.5,
		},
		{
			Id:        "7",
			Payer:     "Eve",
			Recipient: "Alice",
			Amount:    0.5,
		},
		{
			Id:        "8",
			Payer:     "Eve",
			Recipient: "Alice",
			Amount:    0.5,
		},
		{
			Id:        "9",
			Payer:     "Eve",
			Recipient: "Alice",
			Amount:    0.5,
		},
		{
			Id:        "10",
			Payer:     "Eve",
			Recipient: "Alice",
			Amount:    0.5,
		},
	}
	return ledger
}
