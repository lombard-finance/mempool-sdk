package client

import (
	"github.com/lombard-finance/mempool-sdk/api/fee"
	"github.com/pkg/errors"
)

func (cli *Client) GetRecommendedFees() (*fee.GetRecommendedFees200Response, error) {
	response, err := cli.get("/v1/fees/recommended")
	if err != nil {
		return nil, errors.Wrap(err, "request GetRecommendedFees")
	}

	decoded, err := decodeJSONResponse[fee.GetRecommendedFees200Response](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode GetRecommendedFees response")
	}

	cli.logger.WithField("fee", decoded).Debug("fetched recommended fees")

	return &decoded, nil
}
