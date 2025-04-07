package transaction

import (
	"github.com/lombard-finance/mempool-sdk/api"
)

// GetTransaction200Response
type GetTransaction200Response struct {
	Txid     string      `json:"txid"`
	Version  int         `json:"version"`
	Locktime int         `json:"locktime"`
	Size     int         `json:"size"`
	Weight   int         `json:"weight"`
	Fee      int         `json:"fee"`
	Status   api.Status  `json:"status"`
	Vout     []*api.Vout `json:"vout"`
	Vin      []*api.Vin  `json:"vin"`
}
