// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// DeleteFegNetworkIDFederationReader is a Reader for the DeleteFegNetworkIDFederation structure.
type DeleteFegNetworkIDFederationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFegNetworkIDFederationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFegNetworkIDFederationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteFegNetworkIDFederationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFegNetworkIDFederationNoContent creates a DeleteFegNetworkIDFederationNoContent with default headers values
func NewDeleteFegNetworkIDFederationNoContent() *DeleteFegNetworkIDFederationNoContent {
	return &DeleteFegNetworkIDFederationNoContent{}
}

/*DeleteFegNetworkIDFederationNoContent handles this case with default header values.

Success
*/
type DeleteFegNetworkIDFederationNoContent struct {
}

func (o *DeleteFegNetworkIDFederationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /feg/{network_id}/federation][%d] deleteFegNetworkIdFederationNoContent ", 204)
}

func (o *DeleteFegNetworkIDFederationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFegNetworkIDFederationDefault creates a DeleteFegNetworkIDFederationDefault with default headers values
func NewDeleteFegNetworkIDFederationDefault(code int) *DeleteFegNetworkIDFederationDefault {
	return &DeleteFegNetworkIDFederationDefault{
		_statusCode: code,
	}
}

/*DeleteFegNetworkIDFederationDefault handles this case with default header values.

Unexpected Error
*/
type DeleteFegNetworkIDFederationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete feg network ID federation default response
func (o *DeleteFegNetworkIDFederationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFegNetworkIDFederationDefault) Error() string {
	return fmt.Sprintf("[DELETE /feg/{network_id}/federation][%d] DeleteFegNetworkIDFederation default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFegNetworkIDFederationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFegNetworkIDFederationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
