package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	logger *logrus.Entry

	base    *url.URL
	client  *http.Client
	timeout time.Duration
}

// New creates a new client, connecting with a standard HTTP.
func New(address string, logger *logrus.Entry, timeout time.Duration) (*Client, error) {

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
		Timeout: timeout,
	}
	if !strings.HasPrefix(address, "http") {
		address = fmt.Sprintf("https://%s", address)
	}
	if !strings.HasSuffix(address, "/") {
		address = fmt.Sprintf("%s/", address)
	}
	base, err := url.Parse(address)
	if err != nil {
		return nil, errors.Wrap(err, "invalid URL")
	}
	return &Client{
		logger:  logger.WithField("client", "mempool"),
		base:    base,
		client:  client,
		timeout: timeout,
	}, nil
}

// Stop closes the client, freeing up resources.
func (cli *Client) Stop() {
	cli.client.CloseIdleConnections()
}

func (cli *Client) get(endpoint string) (io.Reader, error) {
	log := cli.logger.WithField("id", fmt.Sprintf("%02x", rand.Int31())).
		WithField("endpoint", endpoint)
	log.Trace("GET request")

	requestEndpoint, err := url.Parse(fmt.Sprintf("%s%s", strings.TrimSuffix(cli.base.String(), "/"), endpoint))
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint")
	}

	opCtx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	req, err := http.NewRequestWithContext(opCtx, http.MethodGet, requestEndpoint.String(), nil)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to create GET request")
	}

	resp, err := cli.client.Do(req)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to call GET endpoint")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// Nothing found.  This is not an error, so we return nil on both counts.
		cancel()
		return nil, nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to read GET response")
	}

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		cancel()
		log.Trace("status_code", resp.StatusCode)
		log.Trace("data", string(data))
		log.Trace("GET failed")
		return nil, errors.Errorf("Method %s, StatusCode: %d, Endpoint: %s, Data: %s", http.MethodGet, resp.StatusCode, endpoint, data)
	}
	cancel()

	log.Trace("response", string(data))
	log.Trace("GET response")

	return bytes.NewReader(data), nil
}

func (cli *Client) post(endpoint string, body io.Reader) (io.Reader, error) {
	return cli.requestWithBody(endpoint, http.MethodPost, body)
}

func (cli *Client) put(endpoint string, body io.Reader) (io.Reader, error) {
	return cli.requestWithBody(endpoint, http.MethodPut, body)
}

func (cli *Client) patch(endpoint string, body io.Reader) (io.Reader, error) {
	return cli.requestWithBody(endpoint, http.MethodPatch, body)
}

func (cli *Client) requestWithBody(endpoint string, method string, body io.Reader) (io.Reader, error) {
	log := cli.logger.WithField("id", fmt.Sprintf("%02x", rand.Int31())).
		WithField("endpoint", endpoint).
		WithField("method", method)

	// build url
	requestEndpoint, err := url.Parse(fmt.Sprintf("%s%s", strings.TrimSuffix(cli.base.String(), "/"), endpoint))
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint")
	}

	// build request
	opCtx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	req, err := http.NewRequestWithContext(opCtx, method, requestEndpoint.String(), body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "create request with context")
	}

	// add headers
	req.Header.Set("Content-type", "application/json")
	//req.Header.Set("Accept", "application/json")

	// do the request
	resp, err := cli.client.Do(req)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "read response")
	}

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		log.Trace("status_code", resp.StatusCode)
		log.Trace("data", string(data))
		log.Tracef("%s failed", method)
		cancel()
		return nil, errors.Errorf("Method %s, StatusCode: %d, Endpoint: %s, Data: %s", method, resp.StatusCode, endpoint, data)
	}
	cancel()

	log.Trace("response", string(data))
	log.Tracef("%s response", method)

	return bytes.NewReader(data), nil
}
