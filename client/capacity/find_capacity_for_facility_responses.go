// Code generated by go-swagger; DO NOT EDIT.

package capacity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindCapacityForFacilityReader is a Reader for the FindCapacityForFacility structure.
type FindCapacityForFacilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindCapacityForFacilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindCapacityForFacilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindCapacityForFacilityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindCapacityForFacilityOK creates a FindCapacityForFacilityOK with default headers values
func NewFindCapacityForFacilityOK() *FindCapacityForFacilityOK {
	return &FindCapacityForFacilityOK{}
}

/* FindCapacityForFacilityOK describes a response with status code 200, with default header values.

ok
*/
type FindCapacityForFacilityOK struct {
	Payload *types.CapacityList
}

func (o *FindCapacityForFacilityOK) Error() string {
	return fmt.Sprintf("[GET /capacity][%d] findCapacityForFacilityOK  %+v", 200, o.Payload)
}
func (o *FindCapacityForFacilityOK) GetPayload() *types.CapacityList {
	return o.Payload
}

func (o *FindCapacityForFacilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.CapacityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindCapacityForFacilityUnauthorized creates a FindCapacityForFacilityUnauthorized with default headers values
func NewFindCapacityForFacilityUnauthorized() *FindCapacityForFacilityUnauthorized {
	return &FindCapacityForFacilityUnauthorized{}
}

/* FindCapacityForFacilityUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindCapacityForFacilityUnauthorized struct {
}

func (o *FindCapacityForFacilityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /capacity][%d] findCapacityForFacilityUnauthorized ", 401)
}

func (o *FindCapacityForFacilityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
