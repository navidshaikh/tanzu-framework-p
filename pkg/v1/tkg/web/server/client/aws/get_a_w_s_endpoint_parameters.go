// Code generated by go-swagger; DO NOT EDIT.

package aws

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
)

// NewGetAWSEndpointParams creates a new GetAWSEndpointParams object
// with the default values initialized.
func NewGetAWSEndpointParams() *GetAWSEndpointParams {

	return &GetAWSEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAWSEndpointParamsWithTimeout creates a new GetAWSEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAWSEndpointParamsWithTimeout(timeout time.Duration) *GetAWSEndpointParams {

	return &GetAWSEndpointParams{

		timeout: timeout,
	}
}

// NewGetAWSEndpointParamsWithContext creates a new GetAWSEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAWSEndpointParamsWithContext(ctx context.Context) *GetAWSEndpointParams {

	return &GetAWSEndpointParams{

		Context: ctx,
	}
}

// NewGetAWSEndpointParamsWithHTTPClient creates a new GetAWSEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAWSEndpointParamsWithHTTPClient(client *http.Client) *GetAWSEndpointParams {

	return &GetAWSEndpointParams{
		HTTPClient: client,
	}
}

/*GetAWSEndpointParams contains all the parameters to send to the API endpoint
for the get a w s endpoint operation typically these are written to a http.Request
*/
type GetAWSEndpointParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get a w s endpoint params
func (o *GetAWSEndpointParams) WithTimeout(timeout time.Duration) *GetAWSEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a w s endpoint params
func (o *GetAWSEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a w s endpoint params
func (o *GetAWSEndpointParams) WithContext(ctx context.Context) *GetAWSEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a w s endpoint params
func (o *GetAWSEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a w s endpoint params
func (o *GetAWSEndpointParams) WithHTTPClient(client *http.Client) *GetAWSEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a w s endpoint params
func (o *GetAWSEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAWSEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
