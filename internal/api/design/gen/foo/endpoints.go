// Code generated by goa v3.5.3, DO NOT EDIT.
//
// foo endpoints
//
// Command:
// $ goa gen
// github.com/kmacmcfarlane/goa-reference-bug-repro/internal/api/design

package foo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "foo" service endpoints.
type Endpoints struct {
	Get goa.Endpoint
}

// NewEndpoints wraps the methods of the "foo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Get: NewGetEndpoint(s),
	}
}

// Use applies the given middleware to all the "foo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Get = m(e.Get)
}

// NewGetEndpoint returns an endpoint function that calls the method "Get" of
// service "foo".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ChildType)
		res, err := s.Get(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResulttype(res, "default")
		return vres, nil
	}
}
