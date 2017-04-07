package inject

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cmdhema/functions_go/models"
)

// PostInjectReader is a Reader for the PostInject structure.
type PostInjectReader struct {
	formats strfmt.Registry
}

type PostInjectOK struct {
	Payload *models.BuildWrapper
}

type PostInjectDefault struct {
	_statusCode int

	Payload *models.Error
}


/*PostAppsBadRequest handles this case with default header values.

Parameters are missing or invalid.
*/
type PostInjectBadRequest struct {
	Payload *models.Error
}

// ReadResponse reads a server response into the received o.
func (o *PostInjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	switch response.Code() {

	case 200:
		result := NewPostInjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

		//case 400:
		//	result := NewPostAppsBadRequest()
		//	if err := result.readResponse(response, consumer, o.formats); err != nil {
		//		return nil, err
		//	}
		//	return nil, result
		//
		//case 409:
		//	result := NewPostAppsConflict()
		//	if err := result.readResponse(response, consumer, o.formats); err != nil {
		//		return nil, err
		//	}
		//	return nil, result

	default:
		result := NewPostInjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

func (o *PostInjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *PostInjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsOK creates a PostAppsOK with default headers values
func NewPostInjectOK() *PostInjectOK {
	return &PostInjectOK{}
}

// NewPostAppsDefault creates a PostAppsDefault with default headers values
func NewPostInjectDefault(code int) *PostInjectDefault {
	return &PostInjectDefault{
		_statusCode: code,
	}
}

// Code gets the status code for the post apps default response
func (o *PostInjectDefault) Code() int {
	return o._statusCode
}

func (o *PostInjectDefault) Error() string {
	return fmt.Sprintf("[POST /apps][%d] PostInject default  %+v", o._statusCode, o.Payload)
}
