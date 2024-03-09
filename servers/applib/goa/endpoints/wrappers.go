package goaendpoints

import goa "goa.design/goa/v3/pkg"

type Wrappers []func(goa.Endpoint) goa.Endpoint

func (s Wrappers) Wrap(ep goa.Endpoint) goa.Endpoint {
	r := ep
	for _, w := range s {
		r = w(r)
	}
	return r
}
