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

// NewCreateExecutorNodeParams creates a new CreateExecutorNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateExecutorNodeParams() *CreateExecutorNodeParams {
	return &CreateExecutorNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExecutorNodeParamsWithTimeout creates a new CreateExecutorNodeParams object
// with the ability to set a timeout on a request.
func NewCreateExecutorNodeParamsWithTimeout(timeout time.Duration) *CreateExecutorNodeParams {
	return &CreateExecutorNodeParams{
		timeout: timeout,
	}
}

// NewCreateExecutorNodeParamsWithContext creates a new CreateExecutorNodeParams object
// with the ability to set a context for a request.
func NewCreateExecutorNodeParamsWithContext(ctx context.Context) *CreateExecutorNodeParams {
	return &CreateExecutorNodeParams{
		Context: ctx,
	}
}

// NewCreateExecutorNodeParamsWithHTTPClient creates a new CreateExecutorNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateExecutorNodeParamsWithHTTPClient(client *http.Client) *CreateExecutorNodeParams {
	return &CreateExecutorNodeParams{
		HTTPClient: client,
	}
}

/*
CreateExecutorNodeParams contains all the parameters to send to the API endpoint

	for the create executor node operation.

	Typically these are written to a http.Request.
*/
type CreateExecutorNodeParams struct {

	// Body.
	Body *models.TypesExecutorNode

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create executor node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExecutorNodeParams) WithDefaults() *CreateExecutorNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create executor node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExecutorNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create executor node params
func (o *CreateExecutorNodeParams) WithTimeout(timeout time.Duration) *CreateExecutorNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create executor node params
func (o *CreateExecutorNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create executor node params
func (o *CreateExecutorNodeParams) WithContext(ctx context.Context) *CreateExecutorNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create executor node params
func (o *CreateExecutorNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create executor node params
func (o *CreateExecutorNodeParams) WithHTTPClient(client *http.Client) *CreateExecutorNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create executor node params
func (o *CreateExecutorNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create executor node params
func (o *CreateExecutorNodeParams) WithBody(body *models.TypesExecutorNode) *CreateExecutorNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create executor node params
func (o *CreateExecutorNodeParams) SetBody(body *models.TypesExecutorNode) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExecutorNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
