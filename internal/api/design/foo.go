package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("foo", func() {

	HTTP(func() {
		Path("/foo")
	})

	Method("Get", func() {

		Payload(payloadType)
		Result(resultType)

		HTTP(func() {

			GET("/")

			Response(StatusOK)
		})
	})
})

var payloadType = Type("BaseType", func() {

	Attribute("items", ArrayOf(String), func() {
		Enum("a", "b")
	})

	Required("items")
})

var resultType = ResultType("ResultType", func() {
	Attributes(func(){
		Attribute("fromProfileUUID")
		Attribute("fromProfileID")
	})
})