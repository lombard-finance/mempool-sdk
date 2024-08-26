package address

import "github.com/lombard-finance/mempool-sdk/api"

// GetAddressUtxos200Response List of UTXOs
type GetAddressUtxos200Response []struct {
	Txid   string     `json:"txid"`
	Vout   uint32     `json:"vout"`
	Value  uint64     `json:"value"`
	Status api.Status `json:"status"`
}
