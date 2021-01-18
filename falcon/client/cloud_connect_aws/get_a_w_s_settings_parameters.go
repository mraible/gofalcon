// Code generated by go-swagger; DO NOT EDIT.

package cloud_connect_aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAWSSettingsParams creates a new GetAWSSettingsParams object
// with the default values initialized.
func NewGetAWSSettingsParams() *GetAWSSettingsParams {

	return &GetAWSSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAWSSettingsParamsWithTimeout creates a new GetAWSSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAWSSettingsParamsWithTimeout(timeout time.Duration) *GetAWSSettingsParams {

	return &GetAWSSettingsParams{

		timeout: timeout,
	}
}

// NewGetAWSSettingsParamsWithContext creates a new GetAWSSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAWSSettingsParamsWithContext(ctx context.Context) *GetAWSSettingsParams {

	return &GetAWSSettingsParams{

		Context: ctx,
	}
}

// NewGetAWSSettingsParamsWithHTTPClient creates a new GetAWSSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAWSSettingsParamsWithHTTPClient(client *http.Client) *GetAWSSettingsParams {

	return &GetAWSSettingsParams{
		HTTPClient: client,
	}
}

/*GetAWSSettingsParams contains all the parameters to send to the API endpoint
for the get a w s settings operation typically these are written to a http.Request
*/
type GetAWSSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get a w s settings params
func (o *GetAWSSettingsParams) WithTimeout(timeout time.Duration) *GetAWSSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a w s settings params
func (o *GetAWSSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a w s settings params
func (o *GetAWSSettingsParams) WithContext(ctx context.Context) *GetAWSSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a w s settings params
func (o *GetAWSSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a w s settings params
func (o *GetAWSSettingsParams) WithHTTPClient(client *http.Client) *GetAWSSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a w s settings params
func (o *GetAWSSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAWSSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}