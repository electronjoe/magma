// Code generated by go-swagger; DO NOT EDIT.

package upgrades

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

// NewPostNetworksNetworkIDTiersTierIDImagesParams creates a new PostNetworksNetworkIDTiersTierIDImagesParams object
// with the default values initialized.
func NewPostNetworksNetworkIDTiersTierIDImagesParams() *PostNetworksNetworkIDTiersTierIDImagesParams {
	var ()
	return &PostNetworksNetworkIDTiersTierIDImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDTiersTierIDImagesParamsWithTimeout creates a new PostNetworksNetworkIDTiersTierIDImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNetworksNetworkIDTiersTierIDImagesParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDTiersTierIDImagesParams {
	var ()
	return &PostNetworksNetworkIDTiersTierIDImagesParams{

		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDTiersTierIDImagesParamsWithContext creates a new PostNetworksNetworkIDTiersTierIDImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNetworksNetworkIDTiersTierIDImagesParamsWithContext(ctx context.Context) *PostNetworksNetworkIDTiersTierIDImagesParams {
	var ()
	return &PostNetworksNetworkIDTiersTierIDImagesParams{

		Context: ctx,
	}
}

// NewPostNetworksNetworkIDTiersTierIDImagesParamsWithHTTPClient creates a new PostNetworksNetworkIDTiersTierIDImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNetworksNetworkIDTiersTierIDImagesParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDTiersTierIDImagesParams {
	var ()
	return &PostNetworksNetworkIDTiersTierIDImagesParams{
		HTTPClient: client,
	}
}

/*PostNetworksNetworkIDTiersTierIDImagesParams contains all the parameters to send to the API endpoint
for the post networks network ID tiers tier ID images operation typically these are written to a http.Request
*/
type PostNetworksNetworkIDTiersTierIDImagesParams struct {

	/*Image
	  New image for the tier

	*/
	Image *models.TierImage
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*TierID
	  Tier ID

	*/
	TierID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDTiersTierIDImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) WithContext(ctx context.Context) *PostNetworksNetworkIDTiersTierIDImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDTiersTierIDImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImage adds the image to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) WithImage(image *models.TierImage) *PostNetworksNetworkIDTiersTierIDImagesParams {
	o.SetImage(image)
	return o
}

// SetImage adds the image to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) SetImage(image *models.TierImage) {
	o.Image = image
}

// WithNetworkID adds the networkID to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) WithNetworkID(networkID string) *PostNetworksNetworkIDTiersTierIDImagesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTierID adds the tierID to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) WithTierID(tierID string) *PostNetworksNetworkIDTiersTierIDImagesParams {
	o.SetTierID(tierID)
	return o
}

// SetTierID adds the tierId to the post networks network ID tiers tier ID images params
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) SetTierID(tierID string) {
	o.TierID = tierID
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDTiersTierIDImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Image != nil {
		if err := r.SetBodyParam(o.Image); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param tier_id
	if err := r.SetPathParam("tier_id", o.TierID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
