package build

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
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

func (a *Client) PostBuild(params *PostBuildParams) (*PostBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:		"PostBuild",
		Method:		"POST",
		PathPattern:	"/build",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostBuildReader{formats: a.formats},
		Context:            params.Context,
	})

	if err != nil {
		return nil, err
	}

	return result.(*PostBuildOK), nil
}