// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutWifiNetworkIDNameReader is a Reader for the PutWifiNetworkIDName structure.
type PutWifiNetworkIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWifiNetworkIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutWifiNetworkIDNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutWifiNetworkIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutWifiNetworkIDNameNoContent creates a PutWifiNetworkIDNameNoContent with default headers values
func NewPutWifiNetworkIDNameNoContent() *PutWifiNetworkIDNameNoContent {
	return &PutWifiNetworkIDNameNoContent{}
}

/*PutWifiNetworkIDNameNoContent handles this case with default header values.

Success
*/
type PutWifiNetworkIDNameNoContent struct {
}

func (o *PutWifiNetworkIDNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/name][%d] putWifiNetworkIdNameNoContent ", 204)
}

func (o *PutWifiNetworkIDNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWifiNetworkIDNameDefault creates a PutWifiNetworkIDNameDefault with default headers values
func NewPutWifiNetworkIDNameDefault(code int) *PutWifiNetworkIDNameDefault {
	return &PutWifiNetworkIDNameDefault{
		_statusCode: code,
	}
}

/*PutWifiNetworkIDNameDefault handles this case with default header values.

Unexpected Error
*/
type PutWifiNetworkIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put wifi network ID name default response
func (o *PutWifiNetworkIDNameDefault) Code() int {
	return o._statusCode
}

func (o *PutWifiNetworkIDNameDefault) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/name][%d] PutWifiNetworkIDName default  %+v", o._statusCode, o.Payload)
}

func (o *PutWifiNetworkIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutWifiNetworkIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
