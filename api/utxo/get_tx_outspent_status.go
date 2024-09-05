package utxo

import "github.com/lombard-finance/mempool-sdk/api"

// GetTxOutspendStatus200Response represents the response for the transaction outspend status
type GetTxOutspendStatus200Response struct {
	Spent  bool       `json:"spent"`
	Txid   string     `json:"txid"`
	Vin    int        `json:"vin"`
	Status api.Status `json:"status"`
}
