// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RafaySystems/rcloud-base/components/common/api/def/clients/scheduler/models"
)

// ClusterRegisterClusterReader is a Reader for the ClusterRegisterCluster structure.
type ClusterRegisterClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterRegisterClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterRegisterClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewClusterRegisterClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewClusterRegisterClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewClusterRegisterClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterRegisterClusterOK creates a ClusterRegisterClusterOK with default headers values
func NewClusterRegisterClusterOK() *ClusterRegisterClusterOK {
	return &ClusterRegisterClusterOK{}
}

/* ClusterRegisterClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterRegisterClusterOK struct {
	Payload *models.RPCRegisterClusterResponse
}

func (o *ClusterRegisterClusterOK) Error() string {
	return fmt.Sprintf("[POST /infra/v3/scheduler/cluster/register][%d] clusterRegisterClusterOK  %+v", 200, o.Payload)
}
func (o *ClusterRegisterClusterOK) GetPayload() *models.RPCRegisterClusterResponse {
	return o.Payload
}

func (o *ClusterRegisterClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCRegisterClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterRegisterClusterForbidden creates a ClusterRegisterClusterForbidden with default headers values
func NewClusterRegisterClusterForbidden() *ClusterRegisterClusterForbidden {
	return &ClusterRegisterClusterForbidden{}
}

/* ClusterRegisterClusterForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ClusterRegisterClusterForbidden struct {
	Payload interface{}
}

func (o *ClusterRegisterClusterForbidden) Error() string {
	return fmt.Sprintf("[POST /infra/v3/scheduler/cluster/register][%d] clusterRegisterClusterForbidden  %+v", 403, o.Payload)
}
func (o *ClusterRegisterClusterForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ClusterRegisterClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterRegisterClusterNotFound creates a ClusterRegisterClusterNotFound with default headers values
func NewClusterRegisterClusterNotFound() *ClusterRegisterClusterNotFound {
	return &ClusterRegisterClusterNotFound{}
}

/* ClusterRegisterClusterNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ClusterRegisterClusterNotFound struct {
	Payload string
}

func (o *ClusterRegisterClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /infra/v3/scheduler/cluster/register][%d] clusterRegisterClusterNotFound  %+v", 404, o.Payload)
}
func (o *ClusterRegisterClusterNotFound) GetPayload() string {
	return o.Payload
}

func (o *ClusterRegisterClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterRegisterClusterDefault creates a ClusterRegisterClusterDefault with default headers values
func NewClusterRegisterClusterDefault(code int) *ClusterRegisterClusterDefault {
	return &ClusterRegisterClusterDefault{
		_statusCode: code,
	}
}

/* ClusterRegisterClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterRegisterClusterDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the cluster register cluster default response
func (o *ClusterRegisterClusterDefault) Code() int {
	return o._statusCode
}

func (o *ClusterRegisterClusterDefault) Error() string {
	return fmt.Sprintf("[POST /infra/v3/scheduler/cluster/register][%d] Cluster_RegisterCluster default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterRegisterClusterDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ClusterRegisterClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}