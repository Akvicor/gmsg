package gclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client  *http.Client
	retries int
	timeout time.Duration
}

type Option func(*Client)

func WithRetries(retries int) Option {
	return func(c *Client) {
		c.retries = retries
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

func NewClient(opts ...Option) *Client {
	c := &Client{
		retries: 3,
		timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.client = &http.Client{
		Timeout: c.timeout,
	}

	return c
}

func (c *Client) request(ctx context.Context, method, url string, body any, headers map[string]string) ([]byte, error) {
	var requestBody []byte
	var err error

	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	attempt := 0
	for attempt <= c.retries {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			response, err := c.doRequest(ctx, method, url, requestBody, headers)
			if err == nil {
				return response, nil
			}

			attempt++
			if attempt > c.retries {
				return nil, err
			}

			// Exponential backoff: 1s, 2s, 4s...
			backoff := time.Duration(1<<uint(attempt-1)) * time.Second
			time.Sleep(backoff)
		}
	}

	return nil, errors.New("max retries exceeded")
}

func (c *Client) doRequest(ctx context.Context, method, url string, body []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set(HTTPContentType, HTTPContentTypeJsonUTF8)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New(string(respBody))
	}

	return respBody, nil
}

func (c *Client) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return c.request(ctx, http.MethodGet, url, nil, headers)
}

func (c *Client) Post(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error) {
	return c.request(ctx, http.MethodPost, url, body, headers)
}

func (c *Client) Put(ctx context.Context, url string, body any, headers map[string]string) ([]byte, error) {
	return c.request(ctx, http.MethodPut, url, body, headers)
}

func (c *Client) Delete(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return c.request(ctx, http.MethodDelete, url, nil, headers)
}
