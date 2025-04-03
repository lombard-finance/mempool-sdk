package address

import "github.com/lombard-finance/mempool-sdk/api"

// GetAddressTransactions200Response
type GetAddressTransactions200Response []struct {
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

//func (o GetAddressTransactions200Response) MarshalJSON() ([]byte, error) {
//	toSerialize := map[string]interface{}{}
//	toSerialize["key"] = o.Key
//	return json.Marshal(toSerialize)
//}
