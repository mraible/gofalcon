// Code generated by go-swagger; DO NOT EDIT.

package identity_entities

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

// NewAPIIdpEntitiesExplorerAPIFetchEntitiesParams creates a new APIIdpEntitiesExplorerAPIFetchEntitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIIdpEntitiesExplorerAPIFetchEntitiesParams() *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	return &APIIdpEntitiesExplorerAPIFetchEntitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIIdpEntitiesExplorerAPIFetchEntitiesParamsWithTimeout creates a new APIIdpEntitiesExplorerAPIFetchEntitiesParams object
// with the ability to set a timeout on a request.
func NewAPIIdpEntitiesExplorerAPIFetchEntitiesParamsWithTimeout(timeout time.Duration) *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	return &APIIdpEntitiesExplorerAPIFetchEntitiesParams{
		timeout: timeout,
	}
}

// NewAPIIdpEntitiesExplorerAPIFetchEntitiesParamsWithContext creates a new APIIdpEntitiesExplorerAPIFetchEntitiesParams object
// with the ability to set a context for a request.
func NewAPIIdpEntitiesExplorerAPIFetchEntitiesParamsWithContext(ctx context.Context) *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	return &APIIdpEntitiesExplorerAPIFetchEntitiesParams{
		Context: ctx,
	}
}

// NewAPIIdpEntitiesExplorerAPIFetchEntitiesParamsWithHTTPClient creates a new APIIdpEntitiesExplorerAPIFetchEntitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIIdpEntitiesExplorerAPIFetchEntitiesParamsWithHTTPClient(client *http.Client) *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	return &APIIdpEntitiesExplorerAPIFetchEntitiesParams{
		HTTPClient: client,
	}
}

/*
APIIdpEntitiesExplorerAPIFetchEntitiesParams contains all the parameters to send to the API endpoint

	for the api idp entities explorer api fetch entities operation.

	Typically these are written to a http.Request.
*/
type APIIdpEntitiesExplorerAPIFetchEntitiesParams struct {

	// Body.
	Body *models.MsaIdsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the api idp entities explorer api fetch entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) WithDefaults() *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the api idp entities explorer api fetch entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the api idp entities explorer api fetch entities params
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) WithTimeout(timeout time.Duration) *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api idp entities explorer api fetch entities params
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api idp entities explorer api fetch entities params
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) WithContext(ctx context.Context) *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api idp entities explorer api fetch entities params
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api idp entities explorer api fetch entities params
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) WithHTTPClient(client *http.Client) *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api idp entities explorer api fetch entities params
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the api idp entities explorer api fetch entities params
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) WithBody(body *models.MsaIdsRequest) *APIIdpEntitiesExplorerAPIFetchEntitiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the api idp entities explorer api fetch entities params
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) SetBody(body *models.MsaIdsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *APIIdpEntitiesExplorerAPIFetchEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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