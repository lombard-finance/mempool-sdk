package client

import (
	"github.com/lombard-finance/mempool-sdk/api/blocks"
	"github.com/pkg/errors"
)

// GetBlocksHeight Returns the height of the last block
func (cli *Client) GetBlocksHeight() (blocks.GetBlocksHeight200Response, error) {
	response, err := cli.get("/blocks/tip/height")
	if err != nil {
		return 0, errors.Wrap(err, "request GetBlocksHeight")
	}
	decoded, err := decodeJSONResponse[blocks.GetBlocksHeight200Response](response)
	return decoded, nil
}
