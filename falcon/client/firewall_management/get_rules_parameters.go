// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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
	"github.com/go-openapi/swag"
)

// NewGetRulesParams creates a new GetRulesParams object
// with the default values initialized.
func NewGetRulesParams() *GetRulesParams {
	var ()
	return &GetRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRulesParamsWithTimeout creates a new GetRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRulesParamsWithTimeout(timeout time.Duration) *GetRulesParams {
	var ()
	return &GetRulesParams{

		timeout: timeout,
	}
}

// NewGetRulesParamsWithContext creates a new GetRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRulesParamsWithContext(ctx context.Context) *GetRulesParams {
	var ()
	return &GetRulesParams{

		Context: ctx,
	}
}

// NewGetRulesParamsWithHTTPClient creates a new GetRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRulesParamsWithHTTPClient(client *http.Client) *GetRulesParams {
	var ()
	return &GetRulesParams{
		HTTPClient: client,
	}
}

/*GetRulesParams contains all the parameters to send to the API endpoint
for the get rules operation typically these are written to a http.Request
*/
type GetRulesParams struct {

	/*Ids
	  The rules to retrieve, identified by ID

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rules params
func (o *GetRulesParams) WithTimeout(timeout time.Duration) *GetRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rules params
func (o *GetRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rules params
func (o *GetRulesParams) WithContext(ctx context.Context) *GetRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rules params
func (o *GetRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rules params
func (o *GetRulesParams) WithHTTPClient(client *http.Client) *GetRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rules params
func (o *GetRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get rules params
func (o *GetRulesParams) WithIds(ids []string) *GetRulesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get rules params
func (o *GetRulesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "multi")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}