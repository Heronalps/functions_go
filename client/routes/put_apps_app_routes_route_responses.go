package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/iron-io/functions_go/models"
)

// PutAppsAppRoutesRouteReader is a Reader for the PutAppsAppRoutesRoute structure.
type PutAppsAppRoutesRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAppsAppRoutesRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutAppsAppRoutesRouteCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutAppsAppRoutesRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutAppsAppRoutesRouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutAppsAppRoutesRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutAppsAppRoutesRouteCreated creates a PutAppsAppRoutesRouteCreated with default headers values
func NewPutAppsAppRoutesRouteCreated() *PutAppsAppRoutesRouteCreated {
	return &PutAppsAppRoutesRouteCreated{}
}

/*PutAppsAppRoutesRouteCreated handles this case with default header values.

Route updated
*/
type PutAppsAppRoutesRouteCreated struct {
	Payload *models.RouteWrapper
}

func (o *PutAppsAppRoutesRouteCreated) Error() string {
	return fmt.Sprintf("[PUT /apps/{app}/routes/{route}][%d] putAppsAppRoutesRouteCreated  %+v", 201, o.Payload)
}

func (o *PutAppsAppRoutesRouteCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAppsAppRoutesRouteBadRequest creates a PutAppsAppRoutesRouteBadRequest with default headers values
func NewPutAppsAppRoutesRouteBadRequest() *PutAppsAppRoutesRouteBadRequest {
	return &PutAppsAppRoutesRouteBadRequest{}
}

/*PutAppsAppRoutesRouteBadRequest handles this case with default header values.

One or more of the routes were invalid due to parameters being missing or invalid.
*/
type PutAppsAppRoutesRouteBadRequest struct {
	Payload *models.Error
}

func (o *PutAppsAppRoutesRouteBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{app}/routes/{route}][%d] putAppsAppRoutesRouteBadRequest  %+v", 400, o.Payload)
}

func (o *PutAppsAppRoutesRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAppsAppRoutesRouteInternalServerError creates a PutAppsAppRoutesRouteInternalServerError with default headers values
func NewPutAppsAppRoutesRouteInternalServerError() *PutAppsAppRoutesRouteInternalServerError {
	return &PutAppsAppRoutesRouteInternalServerError{}
}

/*PutAppsAppRoutesRouteInternalServerError handles this case with default header values.

Could not accept routes due to internal error.
*/
type PutAppsAppRoutesRouteInternalServerError struct {
	Payload *models.Error
}

func (o *PutAppsAppRoutesRouteInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /apps/{app}/routes/{route}][%d] putAppsAppRoutesRouteInternalServerError  %+v", 500, o.Payload)
}

func (o *PutAppsAppRoutesRouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAppsAppRoutesRouteDefault creates a PutAppsAppRoutesRouteDefault with default headers values
func NewPutAppsAppRoutesRouteDefault(code int) *PutAppsAppRoutesRouteDefault {
	return &PutAppsAppRoutesRouteDefault{
		_statusCode: code,
	}
}

/*PutAppsAppRoutesRouteDefault handles this case with default header values.

Unexpected error
*/
type PutAppsAppRoutesRouteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put apps app routes route default response
func (o *PutAppsAppRoutesRouteDefault) Code() int {
	return o._statusCode
}

func (o *PutAppsAppRoutesRouteDefault) Error() string {
	return fmt.Sprintf("[PUT /apps/{app}/routes/{route}][%d] PutAppsAppRoutesRoute default  %+v", o._statusCode, o.Payload)
}

func (o *PutAppsAppRoutesRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
