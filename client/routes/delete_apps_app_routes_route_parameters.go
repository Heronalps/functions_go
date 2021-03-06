package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteAppsAppRoutesRouteParams creates a new DeleteAppsAppRoutesRouteParams object
// with the default values initialized.
func NewDeleteAppsAppRoutesRouteParams() *DeleteAppsAppRoutesRouteParams {
	var ()
	return &DeleteAppsAppRoutesRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppsAppRoutesRouteParamsWithTimeout creates a new DeleteAppsAppRoutesRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAppsAppRoutesRouteParamsWithTimeout(timeout time.Duration) *DeleteAppsAppRoutesRouteParams {
	var ()
	return &DeleteAppsAppRoutesRouteParams{

		timeout: timeout,
	}
}

// NewDeleteAppsAppRoutesRouteParamsWithContext creates a new DeleteAppsAppRoutesRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAppsAppRoutesRouteParamsWithContext(ctx context.Context) *DeleteAppsAppRoutesRouteParams {
	var ()
	return &DeleteAppsAppRoutesRouteParams{

		Context: ctx,
	}
}

/*DeleteAppsAppRoutesRouteParams contains all the parameters to send to the API endpoint
for the delete apps app routes route operation typically these are written to a http.Request
*/
type DeleteAppsAppRoutesRouteParams struct {

	/*App
	  Name of app for this set of routes.

	*/
	App string
	/*Route
	  Route name

	*/
	Route string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete apps app routes route params
func (o *DeleteAppsAppRoutesRouteParams) WithTimeout(timeout time.Duration) *DeleteAppsAppRoutesRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete apps app routes route params
func (o *DeleteAppsAppRoutesRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete apps app routes route params
func (o *DeleteAppsAppRoutesRouteParams) WithContext(ctx context.Context) *DeleteAppsAppRoutesRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete apps app routes route params
func (o *DeleteAppsAppRoutesRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithApp adds the app to the delete apps app routes route params
func (o *DeleteAppsAppRoutesRouteParams) WithApp(app string) *DeleteAppsAppRoutesRouteParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the delete apps app routes route params
func (o *DeleteAppsAppRoutesRouteParams) SetApp(app string) {
	o.App = app
}

// WithRoute adds the route to the delete apps app routes route params
func (o *DeleteAppsAppRoutesRouteParams) WithRoute(route string) *DeleteAppsAppRoutesRouteParams {
	o.SetRoute(route)
	return o
}

// SetRoute adds the route to the delete apps app routes route params
func (o *DeleteAppsAppRoutesRouteParams) SetRoute(route string) {
	o.Route = route
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppsAppRoutesRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app
	if err := r.SetPathParam("app", o.App); err != nil {
		return err
	}

	// path param route
	if err := r.SetPathParam("route", o.Route); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
