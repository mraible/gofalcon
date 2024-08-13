// Code generated by go-swagger; DO NOT EDIT.

package a_s_p_m

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewGetServiceViolationTypesParams creates a new GetServiceViolationTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServiceViolationTypesParams() *GetServiceViolationTypesParams {
	return &GetServiceViolationTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceViolationTypesParamsWithTimeout creates a new GetServiceViolationTypesParams object
// with the ability to set a timeout on a request.
func NewGetServiceViolationTypesParamsWithTimeout(timeout time.Duration) *GetServiceViolationTypesParams {
	return &GetServiceViolationTypesParams{
		timeout: timeout,
	}
}

// NewGetServiceViolationTypesParamsWithContext creates a new GetServiceViolationTypesParams object
// with the ability to set a context for a request.
func NewGetServiceViolationTypesParamsWithContext(ctx context.Context) *GetServiceViolationTypesParams {
	return &GetServiceViolationTypesParams{
		Context: ctx,
	}
}

// NewGetServiceViolationTypesParamsWithHTTPClient creates a new GetServiceViolationTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServiceViolationTypesParamsWithHTTPClient(client *http.Client) *GetServiceViolationTypesParams {
	return &GetServiceViolationTypesParams{
		HTTPClient: client,
	}
}

/*
GetServiceViolationTypesParams contains all the parameters to send to the API endpoint

	for the get service violation types operation.

	Typically these are written to a http.Request.
*/
type GetServiceViolationTypesParams struct {

	// Body.
	Body *models.TypesGenericUserFacingRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get service violation types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceViolationTypesParams) WithDefaults() *GetServiceViolationTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get service violation types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceViolationTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get service violation types params
func (o *GetServiceViolationTypesParams) WithTimeout(timeout time.Duration) *GetServiceViolationTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service violation types params
func (o *GetServiceViolationTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service violation types params
func (o *GetServiceViolationTypesParams) WithContext(ctx context.Context) *GetServiceViolationTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service violation types params
func (o *GetServiceViolationTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service violation types params
func (o *GetServiceViolationTypesParams) WithHTTPClient(client *http.Client) *GetServiceViolationTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service violation types params
func (o *GetServiceViolationTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get service violation types params
func (o *GetServiceViolationTypesParams) WithBody(body *models.TypesGenericUserFacingRequest) *GetServiceViolationTypesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get service violation types params
func (o *GetServiceViolationTypesParams) SetBody(body *models.TypesGenericUserFacingRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceViolationTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
