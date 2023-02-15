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

// WaypointListPipelineRunsReader is a Reader for the WaypointListPipelineRuns structure.
type WaypointListPipelineRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListPipelineRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListPipelineRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListPipelineRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListPipelineRunsOK creates a WaypointListPipelineRunsOK with default headers values
func NewWaypointListPipelineRunsOK() *WaypointListPipelineRunsOK {
	return &WaypointListPipelineRunsOK{}
}

/*
WaypointListPipelineRunsOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListPipelineRunsOK struct {
	Payload *models.HashicorpWaypointListPipelineRunsResponse
}

// IsSuccess returns true when this waypoint list pipeline runs o k response has a 2xx status code
func (o *WaypointListPipelineRunsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list pipeline runs o k response has a 3xx status code
func (o *WaypointListPipelineRunsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list pipeline runs o k response has a 4xx status code
func (o *WaypointListPipelineRunsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list pipeline runs o k response has a 5xx status code
func (o *WaypointListPipelineRunsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list pipeline runs o k response a status code equal to that given
func (o *WaypointListPipelineRunsOK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointListPipelineRunsOK) Error() string {
	return fmt.Sprintf("[GET /project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/runs][%d] waypointListPipelineRunsOK  %+v", 200, o.Payload)
}

func (o *WaypointListPipelineRunsOK) String() string {
	return fmt.Sprintf("[GET /project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/runs][%d] waypointListPipelineRunsOK  %+v", 200, o.Payload)
}

func (o *WaypointListPipelineRunsOK) GetPayload() *models.HashicorpWaypointListPipelineRunsResponse {
	return o.Payload
}

func (o *WaypointListPipelineRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListPipelineRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListPipelineRunsDefault creates a WaypointListPipelineRunsDefault with default headers values
func NewWaypointListPipelineRunsDefault(code int) *WaypointListPipelineRunsDefault {
	return &WaypointListPipelineRunsDefault{
		_statusCode: code,
	}
}

/*
WaypointListPipelineRunsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListPipelineRunsDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list pipeline runs default response
func (o *WaypointListPipelineRunsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint list pipeline runs default response has a 2xx status code
func (o *WaypointListPipelineRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list pipeline runs default response has a 3xx status code
func (o *WaypointListPipelineRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list pipeline runs default response has a 4xx status code
func (o *WaypointListPipelineRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list pipeline runs default response has a 5xx status code
func (o *WaypointListPipelineRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list pipeline runs default response a status code equal to that given
func (o *WaypointListPipelineRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointListPipelineRunsDefault) Error() string {
	return fmt.Sprintf("[GET /project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/runs][%d] Waypoint_ListPipelineRuns default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListPipelineRunsDefault) String() string {
	return fmt.Sprintf("[GET /project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/runs][%d] Waypoint_ListPipelineRuns default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListPipelineRunsDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListPipelineRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}