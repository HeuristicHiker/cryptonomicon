# Moving comments to notes

## NewMerkleTree
See the [NewMerkleTree](./new_merkle.go#239)
// First, we want to setup the leaves

// Track level
//	// Number of nodes
// Track relationships
//	// Which are parent/sib/children
//	// Slice of nodes
//

// init leaf nodes
//	// stores hash of data
//	// hash datablocks
// more than 1 node?
//	// pair nodes
//	// track relationsh
//	// create level
//	// Duplicate if odd number
//		// l node + r node
// final
//	// node is merkle root
//		// is tree
//		// has
//			// root hash - has value
//			// merkle proofs (inclusion proofs)
//				// verify data
//				// path of hashes from leaf to root
//					// only uses subset
//			// hash function used
//			//


// Basically, I want to round this up
// Yes, I do realize creating a var from converting a float into an int
// And then literally typing into float64 on the next line and comparins
// SEEMS stupid
// BUT if you compare as ints then it rounds and loses the point
// BUT BUT you still want an int since you won't have x.5 levels you'll only have x levels
// I'm mainly writing this for future Connor who may still think it's dumb after reading this
// Je me mets au défi 🤺