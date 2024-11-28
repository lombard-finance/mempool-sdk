package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/lombard-finance/mempool-sdk/api/transaction"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (cli *Client) GetTransaction(txid string) (*transaction.GetTransaction200Response, error) {
	response, err := cli.get(fmt.Sprintf("/tx/%s", url.PathEscape(txid)))
	if err != nil {
		return nil, errors.Wrap(err, "request GetTransaction")
	}

	decoded, err := decodeJSONResponse[transaction.GetTransaction200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode GetTransaction response")
	}

	cli.logger.WithFields(logrus.Fields{
		"txid":        txid,
		"transaction": decoded,
	}).Debug("fetched transaction")

	return &decoded, nil
}

func (cli *Client) PostTransaction(rawTransaction string) (string, error) {
	response, err := cli.post("/tx", strings.NewReader(rawTransaction))
	if err != nil {
		return "", errors.Wrap(err, "request PostTransaction")
	}

	// Use json.Decoder directly to decode the response
	var decoded interface{}
	if err := json.NewDecoder(response).Decode(&decoded); err != nil {
		return "", errors.Wrap(err, "decode PostTransaction response")
	}

	// Convert the decoded JSON to a string representation
	encoded, err := json.Marshal(decoded)
	if err != nil {
		return "", errors.Wrap(err, "encode PostTransaction response to string")
	}

	cli.logger.WithField("txid", string(encoded)).Debug("posted transaction")
	return string(encoded), nil
}
