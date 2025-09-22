package tochka

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/AA122AA/tochka-api-client/tochka/dto"
)

type AcquiringClient struct {
	client   *Client
	basePath string
}

func (c *AcquiringClient) GetPayments(ctx context.Context) (*dto.GetPaymentsResponse, error) {
	u, err := c.client.buildURL(c.basePath, "/payments")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("customerCode", "1234567ab")

	u.RawQuery = params.Encode()

	req, err := c.client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	log.Printf("req = %+v", req)

	respBody, bodyCloser, err := c.client.do(req)
	if err != nil {
		return nil, err
	}
	defer bodyCloser()

	var data dto.GetPaymentsResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
