package merklev2

import (
	"crypto/sha256"
	"fmt"
	"math"
	"strings"
)

type Node struct {
	Left        *Node
	Right       *Node
	Hash        [32]byte
	transaction Transaction
	AfectedBy   string
}

type Transaction struct {
	id        string
	payer     string
	recipient string
	amount    float64
}

type Ledger struct {
	Name         string
	Transactions []Transaction
}

type Relationships struct {
	Leaves   []*Leaf
	Parents  []*Internal
	Siblings []*Node
	Children []*Node
}

type Hashable interface {
	Hash() [32]byte
}

type Leaf struct {
	Transaction Transaction
	Digest      [32]byte
}

type Internal struct {
	Relationships Relationships
	Digest        [32]byte
	Transactions  []Transaction
}

type MerkleTree struct {
	Internals []*Internal
	HashSet   [][]*Node
	LeafNodes []*Leaf
}

//		Ledger		//

func LedgerToByteSlices(ledger Ledger) [][]byte {
	var data [][]byte
	for _, tx := range ledger.Transactions {
		// Convert transaction to byte slice (simple string representation)
		txString := fmt.Sprintf("%s:%s:%s:%.2f", tx.id, tx.payer, tx.recipient, tx.amount)
		data = append(data, []byte(txString))
	}
	return data
}

func CreateSampleLedger() Ledger {
	ledger := Ledger{
		Name: "First Ledger",
		Transactions: []Transaction{
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
		},
	}

	return ledger
}

//		Hashing		//

func hashData(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func hashPair(left, right [32]byte) [32]byte {
	combined := append(left[:], right[:]...)
	return sha256.Sum256(combined)
}

//		Tree Creation		//

func createLeafNodes(ledger Ledger) []*Node {
	leaves := make([]*Node, len(ledger.Transactions))

	for idx, t := range ledger.Transactions {
		tString := fmt.Sprintf("%s:%s:%s:%.2f", t.id, t.payer, t.recipient, t.amount)
		data := []byte(tString)
		leaves[idx] = &Node{
			Hash:        hashData(data),
			transaction: t,
			AfectedBy:   t.id,
		}
	}
	return leaves
}

func buildInternals(tree *MerkleTree) *MerkleTree {
	var currentLevel []*Node

	// I think I only need to do this once
	// check that it's not happening over and over
	for _, leaf := range tree.LeafNodes {
		currentLevel = append(currentLevel, &Node{
			Hash:        leaf.Digest,
			transaction: leaf.Transaction,
			AfectedBy:   leaf.Transaction.id,
		})
	}

	level := 0

	// looking at internals now
	for len(currentLevel) > 1 {
		var nextLevel []*Node

		// get dem binary pairs
		for i := 0; i < len(currentLevel); i += 2 {
			left := currentLevel[i]
			var right *Node

			// Creating a duplicate node feels inefficent but for now...
			if i+1 < len(currentLevel) {
				right = currentLevel[i+1]
			} else {
				// we've been duped!
				right = currentLevel[i]
			}

			combinedHash := hashPair(left.Hash, right.Hash)
			combinedTransaction := Transaction{
				id:        left.transaction.id + "+" + right.transaction.id,
				payer:     left.transaction.payer + "+" + right.transaction.payer, // yes, this is stupid but whatever I mean technically I could parse out payers involved idk
				recipient: left.transaction.recipient + "+" + right.transaction.recipient,
				amount:    left.transaction.amount + right.transaction.amount, // maybe we use this for valid transaction amounts accross levels in a closed system
			}

			parent := &Node{
				Left:        left,
				Right:       right,
				Hash:        combinedHash,
				transaction: combinedTransaction,
				AfectedBy:   left.AfectedBy + right.AfectedBy, // s'il vous plaît seigneur, dites-moi que cela est vrai pour les validations aux transactions
			}

			nextLevel = append(nextLevel, parent)
		}

		if level < len(tree.HashSet) {
			tree.HashSet[level] = currentLevel
		}

		fmt.Printf("Level %d: %d nodes -> %d nodes\n", level, len(currentLevel), len(nextLevel))

		currentLevel = nextLevel
		level++
	}

	if len(currentLevel) == 1 {
		fmt.Printf("✓ C'est le hachage racine au niveau %d\n", level)
		if level < len(tree.HashSet) {
			tree.HashSet[level] = currentLevel
		}
	}

	return tree
}

func NewMerkleTree(transactions []Transaction) *MerkleTree {

	estimatedNumberOfLevels := math.Log2(float64(len(transactions)))
	intEstimate := int(estimatedNumberOfLevels)
	if float64(intEstimate) < estimatedNumberOfLevels {
		intEstimate++
	}
	// If the length of nodes is longer than this then I'm not following the expected
	// math of the number of levels being log(n) where n is leaves
	// ... I think - future Connor can let me know if this is dumb
	var initialLeaves []*Leaf
	for _, t := range transactions {
		// convert transaction details into a byte slice... in a messy way that could 100% be better
		data := append([]byte(t.id), []byte(t.payer)...)
		data = append(data, []byte(t.recipient)...)
		data = append(data, []byte(fmt.Sprintf("%.2f", t.amount))...)
		newDigest := sha256.Sum256(data)
		newLeaf := Leaf{
			Transaction: t,
			Digest:      newDigest,
		}
		initialLeaves = append(initialLeaves, &newLeaf)
	}

	// okay so outer slice should be hard set to estimated number of levels while the inner slice will be dynamic
	// If it exceeds this number then I'm mathing wrong
	tree := MerkleTree{
		Internals: []*Internal{},
		LeafNodes: initialLeaves,
		HashSet:   make([][]*Node, intEstimate),
	}

	// passed by reference so shouldn't need to pull out return values
	buildInternals(&tree)

	if len(tree.HashSet) == intEstimate {
		fmt.Println("Expected levels, wow, you must have mathedor ")
	}
	return &tree
}

func buildMerkleTree(nodes []*Node, totalLeaves int, internals []*Node) (*Node, []*Node) {
	// Okay we got our root bois
	if len(nodes) == 1 {
		PrintLeaves(nodes, totalLeaves)
		return nodes[0], internals
	}

	var nextLevel []*Node
	for i := 0; i < len(nodes); i += 2 {
		var left = nodes[i]
		var right *Node
		if i+1 < len(nodes) {
			right = nodes[i+1]
		} else {
			// Duplicate last node if odd number of nodes
			right = nodes[i]
		}
		newTransaction := Transaction{
			id:        left.transaction.id + right.transaction.id,
			payer:     left.transaction.payer + right.transaction.payer,
			recipient: left.transaction.recipient + right.transaction.recipient,
			amount:    left.transaction.amount + right.transaction.amount,
		}
		newAffectedBy := filterDuplicateLedgerIds(left.AfectedBy, right.AfectedBy)
		parent := &Node{
			Left:        left,
			Right:       right,
			Hash:        hashPair(left.Hash, right.Hash),
			transaction: newTransaction,
			AfectedBy:   newAffectedBy,
		}
		nextLevel = append(nextLevel, parent)
		internals = append(internals, nextLevel...)
	}
	PrintLeaves(nodes, totalLeaves)
	return buildMerkleTree(nextLevel, totalLeaves, internals)
}

func filterDuplicateLedgerIds(left, right string) string {
	seen := make(map[rune]bool)

	for _, s := range left {
		if !seen[s] {
			seen[s] = true
		}
	}

	for _, s := range right {
		if !seen[s] {
			seen[s] = true
		}
	}

	str := ""
	for s := range seen {
		str += string(s)
	}

	return str
}
func BuildMerkleRoot(ledger Ledger) (*Node, []*Node) {
	leaves := createLeafNodes(ledger)
	totalLeaves := len(leaves)
	var internals []*Node

	return buildMerkleTree(leaves, totalLeaves, internals)
}

//			Print stuff			//

func PrintLedger(ledger Ledger) {
	fmt.Println("Looking at leger: ", ledger.Name)
	fmt.Printf("|")
	for _, t := range ledger.Transactions {
		fmt.Printf(" %s |", t.id)
	}
	fmt.Printf("\n")
}

func PrintLeaves(leaves []*Node, totalLeaves int) {
	padLeft := (totalLeaves - len(leaves)) / 2
	padRight := totalLeaves - len(leaves) - padLeft
	maxSpace := totalLeaves*2 + 2

	for range padLeft {
		filler := strings.Repeat("-", maxSpace-2)
		fmt.Printf("[-x%s]", filler)
	}

	for _, l := range leaves {
		// This is actual trash water but I want things to look good when visualizing the problem
		// ... judge thou as ye may
		fLeft := (maxSpace - len(l.transaction.id)) / 2
		fRight := maxSpace - len(l.transaction.id) - fLeft
		if fLeft < 0 {
			fLeft = 0
		}
		if fRight < 0 {
			fRight = 0
		}
		fillerLeft := strings.Repeat("-", fLeft)
		fillerRight := strings.Repeat("-", fRight)
		if len(l.transaction.id) >= maxSpace {
			fLeft = (maxSpace*2 - len(l.transaction.id)) / 2
			fRight = maxSpace*2 - len(l.transaction.id) - fLeft
			fillerLeft = strings.Repeat("-", fLeft+1)
			fillerRight = strings.Repeat("-", fRight+1)

			fmt.Printf("[%s%s%s]", fillerLeft, l.transaction.id, fillerRight)
			padRight--
		} else {
			fmt.Printf("[%s%s%s]", fillerLeft, l.transaction.id, fillerRight)
		}
	}
	for range padRight {
		filler := strings.Repeat("-", maxSpace-2)
		fmt.Printf("[%sx-]", filler)
	}
	fmt.Println("")
}
