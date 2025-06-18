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

func CreateSimpleBinaryTree(transactions []Transaction, head *Node) (*Node, *Node) {
	listOfNodes := []*Node{}

	for t := 0; t < len(transactions)-1; t += 2 {
		curTransaction := transactions[t]
		nextTransaction := transactions[t+1]
		leftTransaction := Transaction{
			id:        curTransaction.id,
			recipient: curTransaction.recipient,
			payer:     curTransaction.payer,
			amount:    curTransaction.amount,
		}
		rightTransaction := Transaction{
			id:        nextTransaction.id,
			recipient: nextTransaction.recipient,
			payer:     nextTransaction.payer,
			amount:    nextTransaction.amount,
		}
		leftNode := Node{Hash: hashTransaction(leftTransaction)}
		rightNode := Node{Hash: hashTransaction(rightTransaction)}
		internalNode := Node{
			Hash:  hashPair(leftTransaction, rightTransaction),
			Left:  &leftNode,
			Right: &rightNode,
		}
		listOfNodes = append(listOfNodes, &internalNode)
	}

	for _, n := range listOfNodes {
		fmt.Printf("%x ======== %x\n", n.Left.Hash, n.Right.Hash)
	}
	fmt.Println("Good stuff we have len of ", len(listOfNodes))
	if len(listOfNodes)%2 != 0 {
		fmt.Println("odd len ", len(listOfNodes))
	}

	return nil, nil
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
