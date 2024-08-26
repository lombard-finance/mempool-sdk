package client

import (
	"github.com/lombard-finance/mempool-sdk/api/address"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestService_GetAddressTransactions(t *testing.T) {
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
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *address.GetAddressTransactions200Response
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
			args:    args{addr: "tb1qjqyhgv3yn357r0fzv98eatx2fnyanate8dw575"},
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
			got, err := cli.GetAddressTransactions(tt.args.addr)
			require.NoError(t, err)
			t.Log(got)

			for i, s2 := range got {
				t.Log(i, s2)
			}

			t.Logf("%v", got[0].Vout[0].ScriptpubkeyAddress)
		})
	}
}
