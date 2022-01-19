package design

//go:generate goa gen github.com/kmacmcfarlane/goa-repro/internal/api/design

import (
	. "goa.design/goa/v3/dsl"
)

var applicationName = "endorsements-api"

var _ = API("API", func() {
	Title("API")
	Description("HTTP API")
	Version("1.0.0")

	Server(applicationName, func() {
		Description("Server that hosts the API")
		Services("foo")
		Host(applicationName, func() {
			URI("http://localhost/")
		})
	})
})

var FormValidationError = Type("FormValidationError", func() {
	Description("Form field errors")

	Attribute("fields", ArrayOf(FormFieldError), "form fields that contain errors")

	Required("fields")
})

var FormFieldError = Type("FormFieldError", func() {
	Attribute("fieldName", String, "Field name that has an error", func() {
		Example("email")
	})
	Attribute("errorCode", String, "Code of the error", func() {
		Example("INVALID_EMAIL")
	})
	Required("fieldName")
	Required("errorCode")
})

var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint`)
})