// Code generated by goa v3.5.3, DO NOT EDIT.
//
// foo HTTP server types
//
// Command:
// $ goa gen
// github.com/kmacmcfarlane/goa-reference-bug-repro/internal/api/design

package server

import (
	foo "github.com/kmacmcfarlane/goa-reference-bug-repro/internal/api/design/gen/foo"
	fooviews "github.com/kmacmcfarlane/goa-reference-bug-repro/internal/api/design/gen/foo/views"
	goa "goa.design/goa/v3/pkg"
)

// GetRequestBody is the type of the "foo" service "Get" endpoint HTTP request
// body.
type GetRequestBody struct {
	// profile uuid of the endorsement sender
	FromProfileUUID *string `json:"fromProfileUuid"`
	// profile id of the endorsement sender
	FromProfileID *int    `json:"fromProfileId"`
	ValidateMe    *string `form:"validateMe,omitempty" json:"validateMe,omitempty" xml:"validateMe,omitempty"`
}

// GetResponseBody is the type of the "foo" service "Get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// profile uuid of the endorsement sender
	FromProfileUUID string `json:"fromProfileUuid"`
	// profile id of the endorsement sender
	FromProfileID int `json:"fromProfileId"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "Get" endpoint of the "foo" service.
func NewGetResponseBody(res *fooviews.ResulttypeView) *GetResponseBody {
	body := &GetResponseBody{
		FromProfileUUID: *res.FromProfileUUID,
		FromProfileID:   *res.FromProfileID,
	}
	return body
}

// NewGetChildType builds a foo service Get endpoint payload.
func NewGetChildType(body *GetRequestBody) *foo.ChildType {
	v := &foo.ChildType{
		FromProfileUUID: *body.FromProfileUUID,
		FromProfileID:   *body.FromProfileID,
		ValidateMe:      body.ValidateMe,
	}

	return v
}

// ValidateGetRequestBody runs the validations defined on GetRequestBody
func ValidateGetRequestBody(body *GetRequestBody) (err error) {
	if body.FromProfileUUID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fromProfileUUID", "body"))
	}
	if body.FromProfileID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fromProfileID", "body"))
	}
	if body.FromProfileUUID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.fromProfileUUID", *body.FromProfileUUID, goa.FormatUUID))
	}
	if body.FromProfileID != nil {
		if *body.FromProfileID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.fromProfileID", *body.FromProfileID, 1, true))
		}
	}
	if body.ValidateMe != nil {
		if !(*body.ValidateMe == "one" || *body.ValidateMe == "two") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.validateMe", *body.ValidateMe, []interface{}{"one", "two"}))
		}
	}
	return
}
