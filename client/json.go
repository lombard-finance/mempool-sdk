package client

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

func decodeJSONResponse[T any](body io.Reader) (T, error) {
	var res T

	if body == nil {
		return res, errors.New("no body to read")
	}

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return res, errors.Wrap(err, "parse response")
	}

	return res, nil
}
