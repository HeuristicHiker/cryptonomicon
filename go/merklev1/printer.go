package merkle

import (
	"fmt"
	"strings"
)

// PrintTree prints the merkle tree in a visual format to the terminal
func PrintTree(t *Tree) {
	if t.Root == nil {
		fmt.Println("Empty tree")
		return
	}

	fmt.Println("Merkle Tree Structure:")
	fmt.Println("=" + strings.Repeat("=", 60))
	printNode(t.Root, "", true, 0)
	fmt.Println("=" + strings.Repeat("=", 60))
}

// printNode recursively prints nodes with tree-like formatting
func printNode(node Node, prefix string, isLast bool, level int) {
	if node == nil {
		return
	}

	// Choose the appropriate connector
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	// Print current node
	hash := node.Hash()
	nodeType := "LEAF"
	if internal, ok := node.(*Internal); ok {
		nodeType = "INTERNAL"
		fmt.Printf("%s%s%s [Level %d]: %x\n", prefix, connector, nodeType, level, hash[:4])

		// Prepare prefix for children
		childPrefix := prefix
		if isLast {
			childPrefix += "    "
		} else {
			childPrefix += "│   "
		}

		// Print children (right first, then left for better visual layout)
		if internal.Right != nil && internal.Left != nil {
			printNode(internal.Right, childPrefix, false, level+1)
			printNode(internal.Left, childPrefix, true, level+1)
		}
	} else {
		fmt.Printf("%s%s%s [Level %d]: %x\n", prefix, connector, nodeType, level, hash[:4])
	}
}

// PrintTreeDetailed prints the tree with full hash values
func PrintTreeDetailed(t *Tree) {
	if t.Root == nil {
		fmt.Println("Empty tree")
		return
	}

	fmt.Println("Detailed Merkle Tree Structure:")
	fmt.Println("=" + strings.Repeat("=", 80))
	printNodeDetailed(t.Root, "", true, 0)
	fmt.Println("=" + strings.Repeat("=", 80))
}

// printNodeDetailed prints nodes with full hash values
func printNodeDetailed(node Node, prefix string, isLast bool, level int) {
	if node == nil {
		return
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	hash := node.Hash()
	nodeType := "LEAF"
	if internal, ok := node.(*Internal); ok {
		nodeType = "INTERNAL"
		fmt.Printf("%s%s%s [L%d]: %x\n", prefix, connector, nodeType, level, hash)

		childPrefix := prefix
		if isLast {
			childPrefix += "    "
		} else {
			childPrefix += "│   "
		}

		if internal.Right != nil && internal.Left != nil {
			printNodeDetailed(internal.Right, childPrefix, false, level+1)
			printNodeDetailed(internal.Left, childPrefix, true, level+1)
		}
	} else {
		fmt.Printf("%s%s%s [L%d]: %x\n", prefix, connector, nodeType, level, hash)
	}
}
