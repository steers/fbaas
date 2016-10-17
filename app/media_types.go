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

package app

// Output of FizzBuzz for one input (default view)
//
// Identifier: application/vnd.fizzbuzz+json; view=default
type Fizzbuzz struct {
	In  *int    `form:"in,omitempty" json:"in,omitempty" xml:"in,omitempty"`
	Out *string `form:"out,omitempty" json:"out,omitempty" xml:"out,omitempty"`
}

// FizzbuzzCollection is the media type for an array of Fizzbuzz (default view)
//
// Identifier: application/vnd.fizzbuzz+json; type=collection; view=default
type FizzbuzzCollection []*Fizzbuzz
