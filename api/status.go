package api

type Status struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight uint64 `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int    `json:"block_time"`
}
