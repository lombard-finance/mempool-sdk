package client

import (
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/lombard-finance/mempool-sdk/api/fee"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestService_GetRecommendedFees(t *testing.T) {
	logger := logrus.New()

	base, err := url.Parse("https://mempool.space/signet/api")
	require.NoError(t, err)

	timeout := time.Minute

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: timeout,
			}).DialContext,
			MaxIdleConns:        32,
			MaxConnsPerHost:     32,
			MaxIdleConnsPerHost: 32,
			IdleConnTimeout:     600 * time.Second,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       timeout,
	}

	type fields struct {
		logger  *logrus.Entry
		base    *url.URL
		client  *http.Client
		timeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		want    *fee.GetRecommendedFees200Response
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				logger:  logger.WithField("test", "test"),
				base:    base,
				client:  client,
				timeout: timeout,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &Client{
				logger:  tt.fields.logger,
				base:    tt.fields.base,
				client:  tt.fields.client,
				timeout: tt.fields.timeout,
			}
			got, err := cli.GetRecommendedFees()
			require.NoError(t, err)
			t.Log(got)

			// check that all fees are not nil and are not 0
			require.True(t, got.MinimumFee > 0)
			require.True(t, got.EconomyFee > 0)
			require.True(t, got.HourFee > 0)
			require.True(t, got.HalfHourFee > 0)
			require.True(t, got.FastestFee > 0)

			// compare fastest and minimum fee
			require.True(t, got.FastestFee >= got.MinimumFee)

		})
	}
}
