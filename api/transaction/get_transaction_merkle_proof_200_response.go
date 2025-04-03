package transaction

// GetTransactionMerkleProof200Response
type GetTransactionMerkleProof200Response struct {
	BlockHeight uint64   `json:"block_height"`
	Merkle      []string `json:"merkle"`
	Pos         uint64   `json:"pos"`
}
