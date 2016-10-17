package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// MultiFizzbuzzPath computes a request path to the multi action of fizzbuzz.
func MultiFizzbuzzPath(begin int, end int) string {
	return fmt.Sprintf("/v1/between/%v/%v", begin, end)
}

// Retrieve FizzBuzz result for two inputs
func (c *Client) MultiFizzbuzz(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewMultiFizzbuzzRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMultiFizzbuzzRequest create the request corresponding to the multi action endpoint of the fizzbuzz resource.
func (c *Client) NewMultiFizzbuzzRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SingleFizzbuzzPath computes a request path to the single action of fizzbuzz.
func SingleFizzbuzzPath(input int) string {
	return fmt.Sprintf("/v1/get/%v", input)
}

// Retrieve FizzBuzz result for one input
func (c *Client) SingleFizzbuzz(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSingleFizzbuzzRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSingleFizzbuzzRequest create the request corresponding to the single action endpoint of the fizzbuzz resource.
func (c *Client) NewSingleFizzbuzzRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
