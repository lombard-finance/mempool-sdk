package client

import (
	"fmt"
	"net/url"

	"github.com/lombard-finance/mempool-sdk/api/utxo"
	"github.com/pkg/errors"
)

// IsUTXOSpent checks if a specific UTXO is spent
func (cli *Client) IsUTXOSpent(txid string, vout uint32) (bool, error) {
	response, err := cli.get(fmt.Sprintf("/tx/%s/outspend/%d", url.PathEscape(txid), vout))
	if err != nil {
		return false, errors.Wrap(err, "request IsUTXOSpent")
	}

	decoded, err := decodeJSONResponse[utxo.GetTxOutspendStatus200Response](response)
	if err != nil {
		return false, errors.Wrap(err, "decode IsUTXOSpent response")
	}

	cli.logger.WithField("txid", txid).WithField("vout", vout).WithField("spent", decoded.Spent).Debug("checked UTXO spent status")
	return decoded.Spent, nil
}
