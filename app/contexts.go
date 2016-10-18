//************************************************************************//
// API "fbaas": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/steers/fbaas/design
// --out=$(GOPATH)/src/github.com/steers/fbaas
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// MultiFizzbuzzContext provides the fizzbuzz multi action context.
type MultiFizzbuzzContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Begin int
	End   int
}

// NewMultiFizzbuzzContext parses the incoming request URL and body, performs validations and creates the
// context used by the fizzbuzz controller multi action.
func NewMultiFizzbuzzContext(ctx context.Context, service *goa.Service) (*MultiFizzbuzzContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := MultiFizzbuzzContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramBegin := req.Params["begin"]
	if len(paramBegin) > 0 {
		rawBegin := paramBegin[0]
		if begin, err2 := strconv.Atoi(rawBegin); err2 == nil {
			rctx.Begin = begin
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("begin", rawBegin, "integer"))
		}
	}
	paramEnd := req.Params["end"]
	if len(paramEnd) > 0 {
		rawEnd := paramEnd[0]
		if end, err2 := strconv.Atoi(rawEnd); err2 == nil {
			rctx.End = end
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("end", rawEnd, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *MultiFizzbuzzContext) OK(r FizzbuzzCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.fizzbuzz+json; type=collection")
	if r == nil {
		r = FizzbuzzCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *MultiFizzbuzzContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// SingleFizzbuzzContext provides the fizzbuzz single action context.
type SingleFizzbuzzContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Input int
}

// NewSingleFizzbuzzContext parses the incoming request URL and body, performs validations and creates the
// context used by the fizzbuzz controller single action.
func NewSingleFizzbuzzContext(ctx context.Context, service *goa.Service) (*SingleFizzbuzzContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := SingleFizzbuzzContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramInput := req.Params["input"]
	if len(paramInput) > 0 {
		rawInput := paramInput[0]
		if input, err2 := strconv.Atoi(rawInput); err2 == nil {
			rctx.Input = input
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("input", rawInput, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SingleFizzbuzzContext) OK(r *Fizzbuzz) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.fizzbuzz+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *SingleFizzbuzzContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}
