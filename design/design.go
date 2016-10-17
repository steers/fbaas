package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("fbaas", func() {
	Title("FizzBuzz As A Service")
	Description("Practice writing APIs")
	Version("1.0")
	Host("localhost:8080")
	Scheme("http")
	BasePath("/v1")

	Contact(func() {
		Name("Alex Steers")
		Email("alex@steers.rocks")
		URL("https://github.com/steers/rest-fizzbuzz")
	})
	
	License(func() {
		Name("MIT")
		URL("https://github.com/steers/rest-fizzbuzz/blob/master/LICENSE")
	})
})

var FizzBuzz = MediaType("application/vnd.fizzbuzz+json", func() {
	Description("Output of FizzBuzz for one input")
	Attributes(func() {
		Attribute("in", Integer, func() {
			Example(15)
		})
		Attribute("out", String, func() {
			Example("FizzBuzz")
		})
	})

	View("default", func() {
		Attribute("in")
		Attribute("out")
	})
})

var _ = Resource("fizzbuzz", func() {
	Action("single", func() {
		Routing(
			GET("/get/:input"),
		)
		Description("Retrieve FizzBuzz result for one input")
		Params(func() {
			Param("input", Integer, "FizzBuzz input")
		})

		Response(OK, FizzBuzz)
		Response(BadRequest, ErrorMedia)
	})

	Action("multi", func() {
		Routing(
			GET("/between/:begin/:end"),
		)
		Description("Retrieve FizzBuzz result for two inputs")
		Params(func() {
			Param("begin", Integer, "Start of integer series")
			Param("end", Integer, "End of integer series")
		})

		Response(OK, CollectionOf(FizzBuzz))
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/ui", "public/html/index.html")
})

var _ = Resource("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "public/js")
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "public/swagger/swagger.json")
})
