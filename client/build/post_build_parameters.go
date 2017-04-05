package build

import (
	//"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	//cr "github.com/go-openapi/runtime/client"
	"github.com/cmdhema/functions_go/models"
	strfmt "github.com/go-openapi/strfmt"
)

type PostBuildParams struct {

	/*App
	  name of the function.

	*/
	Body *models.BuildWrapper
	Context context.Context
}

func NewPostBuildParams() *PostBuildParams {
	var ()

	return &PostBuildParams{

	}
}

// WithContext adds the context to the post apps params
func (o *PostBuildParams) WithContext(ctx context.Context) *PostBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post apps params
func (o *PostBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post apps params
func (o *PostBuildParams) WithBody(body *models.BuildWrapper) *PostBuildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post apps params
func (o *PostBuildParams) SetBody(body *models.BuildWrapper) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.BuildWrapper)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
