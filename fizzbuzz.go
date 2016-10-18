package main

import (
	"github.com/goadesign/goa"
	"github.com/steers/fbaas/app"
	"fmt"
)

// FizzbuzzController implements the fizzbuzz resource.
type FizzbuzzController struct {
	*goa.Controller
}

// Here's where all the magic happens!
func fizzbuzz(input int) (output string) {
	var (
		mod3 = input % 3 == 0
		mod5 = input % 5 == 0
	)
	switch {
	case mod3 && mod5:
		output = "FizzBuzz"
	case mod3:
		output = "Fizz"
	case mod5:
		output = "Buzz"
	default:
		output = fmt.Sprintf("%v", input)
	}
	return 
}

// NewFizzbuzzController creates a fizzbuzz controller.
func NewFizzbuzzController(service *goa.Service) *FizzbuzzController {
	return &FizzbuzzController{Controller: service.NewController("FizzbuzzController")}
}

// Multi runs the multi action.
func (c *FizzbuzzController) Multi(ctx *app.MultiFizzbuzzContext) error {
	min, max := ctx.Begin, ctx.End
	if min > max {
		min, max = ctx.End, ctx.Begin
	}

	size := max - min + 1
	if size > 1000 {
		return ctx.BadRequest(goa.ErrBadRequest(fmt.Sprintf("range size (%v) exceeded limit (1000)", size)))
	}

	idx := 0
	res := make(app.FizzbuzzCollection, size)
	for i := min; i <= max; i++ {
		input, output := i, fizzbuzz(i)
		res[idx] = &app.Fizzbuzz{In: &input, Out: &output}
		idx++
	}

	return ctx.OK(res)
}

// Single runs the single action.
func (c *FizzbuzzController) Single(ctx *app.SingleFizzbuzzContext) error {
	output := fizzbuzz(ctx.Input)
	res := &app.Fizzbuzz{In: &ctx.Input, Out: &output}
	return ctx.OK(res)
}
