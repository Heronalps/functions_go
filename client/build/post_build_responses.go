package build

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cmdhema/functions_go/models"
)

// PostBuildReader is a Reader for the PostBuild structure.
type PostBuildReader struct {
	formats strfmt.Registry
}

type PostBuildOK struct {
	Payload *models.BuildWrapper
}

type PostBuildDefault struct {
	_statusCode int

	Payload *models.Error
}


/*PostAppsBadRequest handles this case with default header values.

Parameters are missing or invalid.
*/
type PostBuildBadRequest struct {
	Payload *models.Error
}

// ReadResponse reads a server response into the received o.
func (o *PostBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	switch response.Code() {

	case 200:
		result := NewPostBuildOK()
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
		result := NewPostBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

func (o *PostBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *PostBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsOK creates a PostAppsOK with default headers values
func NewPostBuildOK() *PostBuildOK {
	return &PostBuildOK{}
}

// NewPostAppsDefault creates a PostAppsDefault with default headers values
func NewPostBuildDefault(code int) *PostBuildDefault {
	return &PostBuildDefault{
		_statusCode: code,
	}
}

// Code gets the status code for the post apps default response
func (o *PostBuildDefault) Code() int {
	return o._statusCode
}

func (o *PostBuildDefault) Error() string {
	return fmt.Sprintf("[POST /apps][%d] PostBuild default  %+v", o._statusCode, o.Payload)
}