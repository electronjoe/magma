// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParams creates a new PostNetworksNetworkIDPrometheusAlertReceiverRouteParams object
// with the default values initialized.
func NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParams() *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	var ()
	return &PostNetworksNetworkIDPrometheusAlertReceiverRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithTimeout creates a new PostNetworksNetworkIDPrometheusAlertReceiverRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	var ()
	return &PostNetworksNetworkIDPrometheusAlertReceiverRouteParams{

		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithContext creates a new PostNetworksNetworkIDPrometheusAlertReceiverRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithContext(ctx context.Context) *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	var ()
	return &PostNetworksNetworkIDPrometheusAlertReceiverRouteParams{

		Context: ctx,
	}
}

// NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithHTTPClient creates a new PostNetworksNetworkIDPrometheusAlertReceiverRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	var ()
	return &PostNetworksNetworkIDPrometheusAlertReceiverRouteParams{
		HTTPClient: client,
	}
}

/*PostNetworksNetworkIDPrometheusAlertReceiverRouteParams contains all the parameters to send to the API endpoint
for the post networks network ID prometheus alert receiver route operation typically these are written to a http.Request
*/
type PostNetworksNetworkIDPrometheusAlertReceiverRouteParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Route
	  Alert routing tree to be used

	*/
	Route *models.AlertRoutingTree

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithContext(ctx context.Context) *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithNetworkID(networkID string) *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRoute adds the route to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithRoute(route *models.AlertRoutingTree) *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetRoute(route)
	return o
}

// SetRoute adds the route to the post networks network ID prometheus alert receiver route params
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetRoute(route *models.AlertRoutingTree) {
	o.Route = route
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Route != nil {
		if err := r.SetBodyParam(o.Route); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
