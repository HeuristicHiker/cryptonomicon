package tree

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Node struct {
	Left       *Node
	Right      *Node
	Hash       [32]byte
	TDetails   *Transaction
	IsInternal bool
}

type Transaction struct {
	Id        string
	Payer     string
	Recipient string
	Amount    float64
}

// TODO what happens when odd
// what happens when even

func CreateMerkleTree(transactions []Transaction) (*Node, error) {
	if len(transactions) == 0 {
		fmt.Println("Issue with transactions")
		return nil, fmt.Errorf("issue with transactions")
	}

	nodes := CreateLeaves(transactions)

	for len(nodes) > 1 {
		fmt.Println()
		var nextLevel []*Node

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			// supposons que le nœud est impair et que nous devons le dupliquer
			right := nodes[i]

			if i+1 < len(nodes) {
				// Pair exists because it's even
				right = nodes[i+1]
			}

			parentHash := hashPair(left.Hash, right.Hash)

			// Create a combined transaction for internal nodes
			internalTDetail := Transaction{
				Id:        left.TDetails.Id + right.TDetails.Id,
				Payer:     "INTERNAL",
				Recipient: "INTERNAL",
				Amount:    left.TDetails.Amount + right.TDetails.Amount,
			}

			parent := &Node{
				Hash:       parentHash,
				Left:       left,
				Right:      right,
				TDetails:   &internalTDetail,
				IsInternal: true,
			}
			fmt.Printf("[ %s&%s ]-", parent.Left.TDetails.Id, parent.Right.TDetails.Id)

			nextLevel = append(nextLevel, parent)
		}
		nodes = nextLevel
	}

	fmt.Printf("\n[ %s ]", nodes[0].TDetails.Id)

	return nodes[0], nil
}

func hashTransaction(t Transaction) [32]byte {
	tString := fmt.Sprintf("%s:%s:%s:%.2f", t.Id, t.Payer, t.Recipient, t.Amount)
	data := []byte(tString)
	return sha256.Sum256(data)
}

func hashPair(leftHash, rightHash [32]byte) [32]byte {
	// adding a prefix to identify as internal hash for easier verif
	// no collisions with leaves :D
	prefix := []byte{0x01}
	combined := append(prefix, append(leftHash[:], rightHash[:]...)...)
	return sha256.Sum256(combined)
}

func CreateLeaves(transactions []Transaction) []*Node {
	nodes := make([]*Node, len(transactions))
	for i, tx := range transactions {
		nodes[i] = &Node{
			Hash:       hashTransaction(tx),
			TDetails:   &tx,
			IsInternal: false,
		}
	}

	return nodes
}

func PrintLevels(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	level := 1
	totalLeaves := 10

	for len(queue) > 0 {
		levelSize := len(queue)
		fmt.Printf("\nLevel %d and %d: \n", level, levelSize)
		padLeft := (totalLeaves - levelSize) / 2
		padRight := totalLeaves - levelSize - padLeft
		maxSpace := totalLeaves + 2

		for range padLeft {
			filler := strings.Repeat("-", maxSpace-2)
			fmt.Printf("[--%s]", filler)
		}

		for i := 0; i < levelSize; i++ {

			node := queue[0]
			queue = queue[1:] // dequeue

			fLeft := (maxSpace - len(node.TDetails.Id)) / 2
			fRight := maxSpace - len(node.TDetails.Id) - fLeft
			if fLeft < 0 {
				fLeft = 0
			}
			if fRight < 0 {
				fRight = 0
			}
			fillerLeft := strings.Repeat("-", fLeft)
			fillerRight := strings.Repeat("-", fRight)
			if len(node.TDetails.Id) >= maxSpace {
				fLeft = (maxSpace*2 - len(node.TDetails.Id)) / 2
				fRight = maxSpace*2 - len(node.TDetails.Id) - fLeft
				fillerLeft = strings.Repeat("-", fLeft+1)
				fillerRight = strings.Repeat("-", fRight+1)

				fmt.Printf("[%s%s%s]", fillerLeft, node.TDetails.Id, fillerRight)
				padRight--
			} else {
				fmt.Printf("[%s%s%s]", fillerLeft, node.TDetails.Id, fillerRight)
			}

			// Add children to queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			// fmt.Println() // newline after each level
		}

		for range padRight {
			filler := strings.Repeat("-", maxSpace-2)
			fmt.Printf("[%s--]", filler)
		}
		level++
		fmt.Println() // newline after each level
	}
}

// CreateSampleLedger returns up to 'length' transactions from a sample ledger.
// If length is greater than the number of available transactions, all are returned.
func CreateSampleLedger(length int) []Transaction {
	ledger := []Transaction{
		{
			Id:        "0",
			Payer:     "Bob",
			Recipient: "Alice",
			Amount:    5,
		},
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
	}
	if length <= 0 {
		return []Transaction{}
	}
	if length > len(ledger) {
		length = len(ledger)
	}
	return ledger[:length]
}
