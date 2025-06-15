# What is a merkle tree?
## What it is
Binary tree of cryptographic hashes where:
- Leef Nodes = Hashes of raw data blocks
- Internal nodes = hashes of the concatenation of their two child nodes
- Merkle root = single hash at top - succinctly commits to entire dataset

## Why it matters
- Tamper evidence - any change to any leaf propagates all the way up changing root
- Efficient proofs - prove inclusion of single leaf with only O(log n) hashes ("merkle proof") instead of revealing entire dataset

# Core concepts and props
