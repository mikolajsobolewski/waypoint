// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint/pkg/client/gen/models"
)

// WaypointExpediteStatusReport2Reader is a Reader for the WaypointExpediteStatusReport2 structure.
type WaypointExpediteStatusReport2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointExpediteStatusReport2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointExpediteStatusReport2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointExpediteStatusReport2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointExpediteStatusReport2OK creates a WaypointExpediteStatusReport2OK with default headers values
func NewWaypointExpediteStatusReport2OK() *WaypointExpediteStatusReport2OK {
	return &WaypointExpediteStatusReport2OK{}
}

/*
WaypointExpediteStatusReport2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointExpediteStatusReport2OK struct {
	Payload *models.HashicorpWaypointExpediteStatusReportResponse
}

// IsSuccess returns true when this waypoint expedite status report2 o k response has a 2xx status code
func (o *WaypointExpediteStatusReport2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint expedite status report2 o k response has a 3xx status code
func (o *WaypointExpediteStatusReport2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint expedite status report2 o k response has a 4xx status code
func (o *WaypointExpediteStatusReport2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint expedite status report2 o k response has a 5xx status code
func (o *WaypointExpediteStatusReport2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint expedite status report2 o k response a status code equal to that given
func (o *WaypointExpediteStatusReport2OK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointExpediteStatusReport2OK) Error() string {
	return fmt.Sprintf("[PUT /release/{release.id}/status-report][%d] waypointExpediteStatusReport2OK  %+v", 200, o.Payload)
}

func (o *WaypointExpediteStatusReport2OK) String() string {
	return fmt.Sprintf("[PUT /release/{release.id}/status-report][%d] waypointExpediteStatusReport2OK  %+v", 200, o.Payload)
}

func (o *WaypointExpediteStatusReport2OK) GetPayload() *models.HashicorpWaypointExpediteStatusReportResponse {
	return o.Payload
}

func (o *WaypointExpediteStatusReport2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointExpediteStatusReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointExpediteStatusReport2Default creates a WaypointExpediteStatusReport2Default with default headers values
func NewWaypointExpediteStatusReport2Default(code int) *WaypointExpediteStatusReport2Default {
	return &WaypointExpediteStatusReport2Default{
		_statusCode: code,
	}
}

/*
WaypointExpediteStatusReport2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointExpediteStatusReport2Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint expedite status report2 default response
func (o *WaypointExpediteStatusReport2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint expedite status report2 default response has a 2xx status code
func (o *WaypointExpediteStatusReport2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint expedite status report2 default response has a 3xx status code
func (o *WaypointExpediteStatusReport2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint expedite status report2 default response has a 4xx status code
func (o *WaypointExpediteStatusReport2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint expedite status report2 default response has a 5xx status code
func (o *WaypointExpediteStatusReport2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint expedite status report2 default response a status code equal to that given
func (o *WaypointExpediteStatusReport2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointExpediteStatusReport2Default) Error() string {
	return fmt.Sprintf("[PUT /release/{release.id}/status-report][%d] Waypoint_ExpediteStatusReport2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointExpediteStatusReport2Default) String() string {
	return fmt.Sprintf("[PUT /release/{release.id}/status-report][%d] Waypoint_ExpediteStatusReport2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointExpediteStatusReport2Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointExpediteStatusReport2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}