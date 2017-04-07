package inject

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/cmdhema/functions_go/client/build"
)

// New creates a new build API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for build API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

func (a *Client) PostInject(params *build.PostBuildParams) (*PostInjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = build.NewPostBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:		"PostInject",
		Method:		"POST",
		PathPattern:	"/inject",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostInjectReader{formats: a.formats},
		Context:            params.Context,
	})

	if err != nil {
		return nil, err
	}

	return result.(*PostInjectOK), nil
}