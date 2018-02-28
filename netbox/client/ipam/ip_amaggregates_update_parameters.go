// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// NewIPAMAggregatesUpdateParams creates a new IPAMAggregatesUpdateParams object
// with the default values initialized.
func NewIPAMAggregatesUpdateParams() *IPAMAggregatesUpdateParams {
	var ()
	return &IPAMAggregatesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMAggregatesUpdateParamsWithTimeout creates a new IPAMAggregatesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMAggregatesUpdateParamsWithTimeout(timeout time.Duration) *IPAMAggregatesUpdateParams {
	var ()
	return &IPAMAggregatesUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMAggregatesUpdateParamsWithContext creates a new IPAMAggregatesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMAggregatesUpdateParamsWithContext(ctx context.Context) *IPAMAggregatesUpdateParams {
	var ()
	return &IPAMAggregatesUpdateParams{

		Context: ctx,
	}
}

// NewIPAMAggregatesUpdateParamsWithHTTPClient creates a new IPAMAggregatesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMAggregatesUpdateParamsWithHTTPClient(client *http.Client) *IPAMAggregatesUpdateParams {
	var ()
	return &IPAMAggregatesUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMAggregatesUpdateParams contains all the parameters to send to the API endpoint
for the ipam aggregates update operation typically these are written to a http.Request
*/
type IPAMAggregatesUpdateParams struct {

	/*Data*/
	Data *models.WritableAggregate
	/*ID
	  A unique integer value identifying this aggregate.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) WithTimeout(timeout time.Duration) *IPAMAggregatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) WithContext(ctx context.Context) *IPAMAggregatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) WithHTTPClient(client *http.Client) *IPAMAggregatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) WithData(data *models.WritableAggregate) *IPAMAggregatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) SetData(data *models.WritableAggregate) {
	o.Data = data
}

// WithID adds the id to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) WithID(id int64) *IPAMAggregatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam aggregates update params
func (o *IPAMAggregatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMAggregatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}