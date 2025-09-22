package tochka

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL    string
	apiVersion string
	apiKey     string
	httpClient *http.Client

	Acquiring *AcquiringClient
}

func NewClient(baseURL, apiVersion, apiKey string) *Client {
	c := &Client{
		baseURL:    baseURL,    // https://enter.tochka.com/sandbox/v2
		apiVersion: apiVersion, // v1.0
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}

	c.Acquiring = &AcquiringClient{
		client:   c,
		basePath: fmt.Sprintf("/acquiring/%s", c.apiVersion),
	}

	return c
}

func (c *Client) buildURL(basePath, path string) (*url.URL, error) {
	URL, err := url.Parse(c.baseURL + basePath + path)
	if err != nil {
		return nil, err
	}

	return URL, nil
}

func (c *Client) newRequest(ctx context.Context, method, url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("newRequest: error while creating request - %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) do(req *http.Request) ([]byte, func() error, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("do: error while doing request - %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("do: error while reading body - %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("do: unexpected status code: %d, status %v, body: %v", resp.StatusCode, resp.Status, string(respBody))
	}

	return respBody, resp.Body.Close, nil
}
