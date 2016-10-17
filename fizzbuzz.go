package main

import (
	"github.com/goadesign/goa"
	"github.com/steers/rest-fizzbuzz/app"
)

// FizzbuzzController implements the fizzbuzz resource.
type FizzbuzzController struct {
	*goa.Controller
}

// NewFizzbuzzController creates a fizzbuzz controller.
func NewFizzbuzzController(service *goa.Service) *FizzbuzzController {
	return &FizzbuzzController{Controller: service.NewController("FizzbuzzController")}
}

// Multi runs the multi action.
func (c *FizzbuzzController) Multi(ctx *app.MultiFizzbuzzContext) error {
	// FizzbuzzController_Multi: start_implement

	// Put your logic here

	// FizzbuzzController_Multi: end_implement
	res := app.FizzbuzzCollection{}
	return ctx.OK(res)
}

// Single runs the single action.
func (c *FizzbuzzController) Single(ctx *app.SingleFizzbuzzContext) error {
	// FizzbuzzController_Single: start_implement

	// Put your logic here

	// FizzbuzzController_Single: end_implement
	res := &app.Fizzbuzz{}
	return ctx.OK(res)
}
