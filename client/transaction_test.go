package client

import (
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/lombard-finance/mempool-sdk/api/transaction"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestService_GetTransaction(t *testing.T) {
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
		txid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *transaction.GetTransaction200Response
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
			args:    args{txid: "6b58397274d9fcce5a8f3918e2c5fe778968d6f60d697dde0fe35c6a6ef94f5b"},
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
			got, err := cli.GetTransaction(tt.args.txid)
			require.NoError(t, err)
			t.Log(got)

			// check that txid matches and check status fields
			require.Equal(t, tt.args.txid, got.Txid)
			require.True(t, got.Status.Confirmed)
			require.Equal(t, uint64(211313), got.Status.BlockHeight)
			require.Equal(t, "0000002d608dad0ba5ecf5b388f4fafac879fccaa67c3b233e1f5000ff686932", got.Status.BlockHash)
			require.Equal(t, 1725201424, got.Status.BlockTime)
			require.Equal(t, 1, len(got.Vout))
			require.Equal(t, 1, len(got.Vin))

		})
	}
}

func TestService_PostTransaction(t *testing.T) {
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
		rawTransaction string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful spent tx",
			fields: fields{
				logger:  logger.WithField("test", "test"),
				base:    base,
				client:  client,
				timeout: timeout,
			},
			args:    args{rawTransaction: "0200000001fd5b5fcd1cb066c27cfc9fda5428b9be850b81ac440ea51f1ddba2f987189ac1010000008a4730440220686a40e9d2dbffeab4ca1ff66341d06a17806767f12a1fc4f55740a7af24c6b5022049dd3c9a85ac6c51fecd5f4baff7782a518781bbdd94453c8383755e24ba755c01410436d554adf4a3eb03a317c77aa4020a7bba62999df633bba0ea8f83f48b9e01b0861d3b3c796840f982ee6b14c3c4b7ad04fcfcc3774f81bff9aaf52a15751fedfdffffff02416c00000000000017a914bc791b2afdfe1e1b5650864a9297b20d74c61f4787d71d0000000000001976a9140a59837ccd4df25adc31cdad39be6a8d97557ed688ac00000000"},
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
			err := cli.PostTransaction(tt.args.rawTransaction)
			require.Error(t, err)
			require.Contains(t, err.Error(), "bad-txns-inputs-missingorspent")

		})
	}
}
