// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDGatewaysGatewayIDDescriptionReader is a Reader for the GetNetworksNetworkIDGatewaysGatewayIDDescription structure.
type GetNetworksNetworkIDGatewaysGatewayIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionOK creates a GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionOK() *GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK {
	return &GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK{}
}

/*GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK handles this case with default header values.

The description of the gateway
*/
type GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK struct {
	Payload models.GatewayDescription
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/description][%d] getNetworksNetworkIdGatewaysGatewayIdDescriptionOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK) GetPayload() models.GatewayDescription {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault creates a GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault(code int) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault {
	return &GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID gateways gateway ID description default response
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/description][%d] GetNetworksNetworkIDGatewaysGatewayIDDescription default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
