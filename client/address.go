package client

import (
	"fmt"
	"github.com/lombard-finance/mempool-sdk/api/address"
	"github.com/pkg/errors"
	"net/url"
)

func (cli *Client) GetAddress(addr string) (address.GetAddress200Response, error) {
	response, err := cli.get(fmt.Sprintf("/address/%s", url.PathEscape(addr)))
	if err != nil {
		return address.GetAddress200Response{}, errors.Wrap(err, "request GetAddress") // Return an empty struct
	}
	decoded, err := decodeJSONResponse[address.GetAddress200Response](response)
	if err != nil {
		return address.GetAddress200Response{}, errors.Wrap(err, "decode GetAddress response") // Return an empty struct
	}
	return decoded, nil
}

// GetAddressTransactions Get transaction history for the specified address/scripthash, sorted with newest first. Returns up to 50 mempool transactions plus the first 25 confirmed transactions. You can request more confirmed transactions using an after_txid query parameter.
func (cli *Client) GetAddressTransactions(addr string) (address.GetAddressTransactions200Response, error) {
	response, err := cli.get(fmt.Sprintf("/address/%s/txs", url.PathEscape(addr)))
	if err != nil {
		return nil, errors.Wrap(err, "request GetAddressTransactions")
	}
	decoded, err := decodeJSONResponse[address.GetAddressTransactions200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode GetAddressTransactions response")
	}
	cli.logger.WithField("address", addr).WithField("transactions", decoded).Debug("fetched transactions")
	return decoded, nil
}

// GetAddressUTXOs Get the list of unspent transaction outputs associated with the address/scripthash
func (cli *Client) GetAddressUTXOs(addr string) (address.GetAddressUtxos200Response, error) {
	response, err := cli.get(fmt.Sprintf("/address/%s/utxo", url.PathEscape(addr)))
	if err != nil {
		return nil, errors.Wrap(err, "request GetAddressUTXOs")
	}
	decoded, err := decodeJSONResponse[address.GetAddressUtxos200Response](response)
	return decoded, nil
}
