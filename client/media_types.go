//************************************************************************//
// API "fbaas": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/steers/rest-fizzbuzz/design
// --out=$(GOPATH)/src/github.com/steers/rest-fizzbuzz
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Output of FizzBuzz for one input (default view)
//
// Identifier: application/vnd.fizzbuzz+json; view=default
type Fizzbuzz struct {
	In  *int    `form:"in,omitempty" json:"in,omitempty" xml:"in,omitempty"`
	Out *string `form:"out,omitempty" json:"out,omitempty" xml:"out,omitempty"`
}

// DecodeFizzbuzz decodes the Fizzbuzz instance encoded in resp body.
func (c *Client) DecodeFizzbuzz(resp *http.Response) (*Fizzbuzz, error) {
	var decoded Fizzbuzz
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// FizzbuzzCollection is the media type for an array of Fizzbuzz (default view)
//
// Identifier: application/vnd.fizzbuzz+json; type=collection; view=default
type FizzbuzzCollection []*Fizzbuzz

// DecodeFizzbuzzCollection decodes the FizzbuzzCollection instance encoded in resp body.
func (c *Client) DecodeFizzbuzzCollection(resp *http.Response) (FizzbuzzCollection, error) {
	var decoded FizzbuzzCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
