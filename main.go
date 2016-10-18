//go:generate goagen bootstrap -d github.com/steers/rest-fizzbuzz/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/steers/fbaas/app"
)

func main() {
	// Create service
	service := goa.New("fbaas")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "fizzbuzz" controller
	c := NewFizzbuzzController(service)
	app.MountFizzbuzzController(service, c)
	// Mount "js" controller
	c2 := NewJsController(service)
	app.MountJsController(service, c2)
	// Mount "public" controller
	c3 := NewPublicController(service)
	app.MountPublicController(service, c3)
	// Mount "swagger" controller
	c4 := NewSwaggerController(service)
	app.MountSwaggerController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
