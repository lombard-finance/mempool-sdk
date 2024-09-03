package client

import (
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

func (cli *Client) PostTransaction(rawTransaction string) (*transaction.PostTransaction200Response, error) {
	response, err := cli.post("/tx", strings.NewReader(rawTransaction))
	if err != nil {
		return nil, errors.Wrap(err, "request PostTransaction")
	}

	decoded, err := decodeJSONResponse[transaction.PostTransaction200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode PostTransaction response")
	}

	cli.logger.WithField("txid", decoded.Txid).Debug("posted transaction")

	return &decoded, nil
}
