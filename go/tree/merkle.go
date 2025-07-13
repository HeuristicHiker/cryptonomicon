package tree

import (
	"crypto/sha256"
	"cryptonomicon/fancy"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func CreateMerkleTree(transactions []Transaction) (*Node, error) {
	if len(transactions) == 0 {
		fmt.Println("Issue with transactions")
		return nil, fmt.Errorf("issue with transactions")
	}

	nodes := CreateLeaves(transactions)

	for len(nodes) > 1 {
		if len(nodes)%2 == 1 {
			// need to create new node for duplication rather than passing shared reference
			lastNode := nodes[len(nodes)-1]
			duplicateNode := &Node{
				Hash:  lastNode.Hash, // Same hash value
				Left:  nil,           // Leaf node
				Right: nil,           // Leaf node
				TDetails: &Transaction{ // NEW transaction instance
					Id:        lastNode.TDetails.Id,
					Payer:     lastNode.TDetails.Payer,
					Recipient: lastNode.TDetails.Recipient,
					Amount:    lastNode.TDetails.Amount,
				},
				IsInternal: false,
			}
			nodes = append(nodes, duplicateNode)
		}
		// fmt.Println()
		var nextLevel []*Node

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i+1]

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
			// fmt.Printf("[ %s & %s ]-", parent.Left.TDetails.Id, parent.Right.TDetails.Id)

			nextLevel = append(nextLevel, parent)
		}
		nodes = nextLevel
	}

	// fmt.Printf("\n[ %s ]\n", nodes[0].TDetails.Id)

	return nodes[0], nil
}

func HashTransaction(t Transaction) [32]byte {
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
			Hash:       HashTransaction(tx),
			TDetails:   &transactions[i],
			IsInternal: false,
		}
	}
	return nodes
}

func PrintLevels(root *Node) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	// First pass: collect all levels to determine max width
	var levels [][]*Node
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]*Node, levelSize)
		copy(currentLevel, queue[:levelSize])
		levels = append(levels, currentLevel)

		var nextLevel []*Node
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		queue = nextLevel
	}

	// Find the maximum width (bottom level)
	maxWidth := len(levels[len(levels)-1])

	// Print each level centered
	for _, level := range levels {
		levelWidth := len(level)
		padding := (maxWidth - levelWidth) * 2 // 2 spaces per node for centering

		// Print leading spaces for centering
		fmt.Print(strings.Repeat(" ", padding))

		// Print nodes at this level
		for _, node := range level {
			if node.TDetails != nil {
				fmt.Printf("[%s] ", node.TDetails.Id)
			} else {
				fmt.Printf("[nil] ")
			}
		}
		fmt.Println()
	}
}

// PrintTree prints a more detailed tree structure showing parent-child relationships.
func PrintTree(root *Node, prefix string, isLast bool) {
	if root == nil {
		return
	}

	// Print current node
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	nodeInfo := "nil"
	if root.TDetails != nil {
		nodeInfo = fmt.Sprintf("ID:%s", root.TDetails.Id)
		if root.IsInternal {
			nodeInfo += " (internal)"
		}
	}

	fmt.Printf("%s%s%s\n", prefix, connector, nodeInfo)

	// Prepare prefix for children
	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	// Print children
	if root.Left != nil || root.Right != nil {
		if root.Right != nil {
			PrintTree(root.Right, childPrefix, root.Left == nil)
		}
		if root.Left != nil {
			PrintTree(root.Left, childPrefix, true)
		}
	}
}

// PrintTreeWithLines prints the tree with ASCII art showing parent-child connections
func PrintTreeWithLines(root *Node) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	fmt.Println("Tree Structure (with parent-child connections):")
	printNodeWithLines(root, "", true, true)
}

// Helper function for recursive tree printing with lines
func printNodeWithLines(node *Node, prefix string, isRoot bool, isLast bool) {
	if node == nil {
		return
	}

	// Print current node
	nodeStr := ""
	if node.TDetails != nil {
		if node.IsInternal {
			nodeStr = fmt.Sprintf("Internal[%s]", node.TDetails.Id)
		} else {
			nodeStr = fmt.Sprintf("Leaf[%s]", node.TDetails.Id)
		}
	} else {
		nodeStr = "nil"
	}

	if isRoot {
		fmt.Printf("ROOT: %s\n", nodeStr)
	} else {
		connector := "├── "
		if isLast {
			connector = "└── "
		}
		fmt.Printf("%s%s%s\n", prefix, connector, nodeStr)
	}

	// Prepare prefix for children
	childPrefix := prefix
	if !isRoot {
		if isLast {
			childPrefix += "    "
		} else {
			childPrefix += "│   "
		}
	}

	// Print children
	hasLeft := node.Left != nil
	hasRight := node.Right != nil

	if hasLeft {
		printNodeWithLines(node.Left, childPrefix, false, !hasRight)
	}
	if hasRight {
		printNodeWithLines(node.Right, childPrefix, false, true)
	}
}

// CreateSampleLedger loads transactions from JSON file and returns up to 'length' transactions.
func CreateSampleLedger(length int) []Transaction {
	// Load transactions from JSON file
	data, err := ioutil.ReadFile("tree/transactions.json")
	if err != nil {
		fmt.Printf("Error reading transactions.json: %v\n", err)
		// Fallback to simple generated transactions
		return generateSimpleTransactions(length)
	}

	var allTransactions []Transaction
	err = json.Unmarshal(data, &allTransactions)
	if err != nil {
		fmt.Printf("Error parsing transactions.json: %v\n", err)
		// Fallback to simple generated transactions
		return generateSimpleTransactions(length)
	}

	// Return requested number of transactions (up to available)
	if length <= 0 {
		return []Transaction{}
	}
	if length > len(allTransactions) {
		length = len(allTransactions)
	}

	return allTransactions[:length]
}

// generateSimpleTransactions creates basic transactions as fallback
func generateSimpleTransactions(length int) []Transaction {
	transactions := make([]Transaction, length)
	for i := 0; i < length; i++ {
		transactions[i] = Transaction{
			Id:        fmt.Sprintf("%d", i),
			Payer:     "payer" + fmt.Sprintf("%d", i),
			Recipient: "recipient" + fmt.Sprintf("%d", i),
			Amount:    float64(i * 10),
		}
	}
	return transactions
}

// ModifyLeafWithID demonstrates the security vulnerability of shared references
func ModifyLeafWithID(root *Node, targetID string, newID string) bool {
	if root == nil {
		return false
	}

	// If this is a leaf node with matching ID
	if root.Left == nil && root.Right == nil && root.TDetails != nil && root.TDetails.Id == targetID {
		fmt.Printf("Found leaf with ID '%s', changing to '%s'\n", targetID, newID)
		root.TDetails.Id = newID
		return true
	}

	// Recursively search children
	return ModifyLeafWithID(root.Left, targetID, newID) || ModifyLeafWithID(root.Right, targetID, newID)
}

// ProofElement represents one step in a Merkle proof
type ProofElement struct {
	Hash   [32]byte
	IsLeft bool // true if this hash should be on the left when combining
}

// === LIGHT CLIENT (Bob) DATA STRUCTURES ===

// LightClientKnowledge represents what Bob (light client) knows
type LightClientKnowledge struct {
	TrustedRootHash [32]byte    // From blockchain consensus
	MyTransaction   Transaction // Transaction Bob wants to verify
	BlockNumber     int         // Which block he thinks it's in
}

// ProofRequest represents what Bob sends to Alice to request a proof
type ProofRequest struct {
	Transaction Transaction // "Please prove this transaction is in the block"
	BlockNumber int         // "I think it's in block #12345"
}

// === FULL NODE (Alice) DATA STRUCTURES ===

// FullNodeKnowledge represents what Alice (full node) knows
type FullNodeKnowledge struct {
	CompleteTree    *Node         // The entire Merkle tree
	AllTransactions []Transaction // All transactions in the block
	RootHash        [32]byte      // Root hash of the tree
	BlockNumber     int           // Which block this represents
}

// ProofResponse represents what Alice sends back to Bob
type ProofResponse struct {
	Found       bool           // Whether the transaction was found
	Transaction Transaction    // Echo back the transaction
	ProofPath   []ProofElement // The sibling hashes needed for verification
	RootHash    [32]byte       // The root hash for verification
	Message     string         // Human-readable status
}

// === NETWORK/BLOCKCHAIN DATA STRUCTURES ===

// BlockHeader represents what the network consensus provides
type BlockHeader struct {
	BlockNumber int      // Block number
	RootHash    [32]byte // Merkle root hash
	Timestamp   int64    // When block was created
	// ... other block metadata
}

// CreateProofRequest simulates what Bob (light client) would send to Alice (full node)
func CreateProofRequest(transaction Transaction, blockNumber int) ProofRequest {
	return ProofRequest{
		Transaction: transaction,
		BlockNumber: blockNumber,
	}
}

// SimulateLightClient creates what Bob knows (minimal information)
func SimulateLightClient(trustedRootHash [32]byte, myTransaction Transaction, blockNumber int) LightClientKnowledge {
	return LightClientKnowledge{
		TrustedRootHash: trustedRootHash,
		MyTransaction:   myTransaction,
		BlockNumber:     blockNumber,
	}
}

// SimulateFullNode creates what Alice knows (complete information)
func SimulateFullNode(transactions []Transaction, blockNumber int) (FullNodeKnowledge, error) {
	tree, err := CreateMerkleTree(transactions)
	if err != nil {
		return FullNodeKnowledge{}, err
	}

	return FullNodeKnowledge{
		CompleteTree:    tree,
		AllTransactions: transactions,
		RootHash:        tree.Hash,
		BlockNumber:     blockNumber,
	}, nil
}

// SimulateBlockchain creates what the network consensus provides
func SimulateBlockchain(rootHash [32]byte, blockNumber int) BlockHeader {
	return BlockHeader{
		BlockNumber: blockNumber,
		RootHash:    rootHash,
		Timestamp:   1234567890, // Mock timestamp
	}
}

// GenerateMerkleProof generates a proof for a given transaction in the tree
func GenerateMerkleProof(root *Node, targetTransaction Transaction) ([]ProofElement, error) {
	// Base case: if this is a leaf node
	if !root.IsInternal {
		if targetTransaction.Id == root.TDetails.Id {
			fancy.PrintGreen("Found transaction: " + root.TDetails.Id)
			// Found the target - return empty proof (no siblings needed at leaf level)
			return []ProofElement{}, nil
		} else {
			fancy.PrintRed("Leaf isn't transaction: " + root.TDetails.Id)
			// Not found - return nil to indicate not found (not an error)
			return nil, nil
		}
	}

	// Internal node: search left subtree first
	if root.Left != nil {
		fancy.PrintCyan("Checking left: " + root.Left.TDetails.Id)
		leftProof, err := GenerateMerkleProof(root.Left, targetTransaction)
		if err != nil {
			return nil, err
		}
		if leftProof != nil {
			// Found in left subtree - collect RIGHT sibling
			// Adding sibling node hash so Bob can verify parent hash! ah, duh
			sibling := ProofElement{
				Hash:   root.Right.Hash,
				IsLeft: false, // Right sibling goes on right when combining
			}
			return append(leftProof, sibling), nil
		}
	}

	// Not found in left, search right subtree
	if root.Right != nil {
		fancy.PrintBlue("Checking right: " + root.Right.TDetails.Id)
		rightProof, err := GenerateMerkleProof(root.Right, targetTransaction)
		if err != nil {
			return nil, err
		}
		if rightProof != nil {
			// Found in right subtree - collect LEFT sibling
			sibling := ProofElement{
				Hash:   root.Left.Hash,
				IsLeft: true, // Left sibling goes on left when combining
			}
			return append(rightProof, sibling), nil
		}
	}

	// Not found in either subtree
	return nil, nil
}

// ProcessProofRequest simulates Alice processing Bob's request
func ProcessProofRequest(fullNode FullNodeKnowledge, request ProofRequest) ProofResponse {
	// Alice tries to generate a proof for Bob's transaction
	proof, err := GenerateMerkleProof(fullNode.CompleteTree, request.Transaction)

	if err != nil {
		return ProofResponse{
			Found:       false,
			Transaction: request.Transaction,
			ProofPath:   nil,
			RootHash:    fullNode.RootHash,
			Message:     fmt.Sprintf("Transaction not found: %v", err),
		}
	}

	return ProofResponse{
		Found:       true,
		Transaction: request.Transaction,
		ProofPath:   proof,
		RootHash:    fullNode.RootHash,
		Message:     "Proof generated successfully",
	}
}

func VerifyMerkleProof(proof []ProofElement, rootHash [32]byte, transaction Transaction) bool {
	/*
		Hash his TX003 transaction
		Combine with Element 0 (right) → gets [TX003TX004] hash
		Combine with Element 1 (left) → gets [TX001TX002TX003TX004] hash
		Combine with Element 2 (right) → gets ROOT hash
		Compare with trusted root hash from blockchain
	*/
	fmt.Println("Bob to prove: ")
	combinedHash := [32]byte{}
	for _, t := range proof {
		fmt.Printf("Hash: %x, IsLeft: %v\n", t.Hash, t.IsLeft)
		if t.IsLeft {
			combinedHash = hashPair(combinedHash, t.Hash)
		} else {
			combinedHash = hashPair(t.Hash, combinedHash)
		}
	}
	validHash := combinedHash == rootHash
	if validHash {
		fancy.PrintGreen("Success!")
	} else {
		fancy.PrintGreen("Success!")
	}
	return validHash
}
