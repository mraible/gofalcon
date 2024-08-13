// Code generated by go-swagger; DO NOT EDIT.

package compliance_assessments

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

// NewExtAggregateFailedRulesCountBySeverityParams creates a new ExtAggregateFailedRulesCountBySeverityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtAggregateFailedRulesCountBySeverityParams() *ExtAggregateFailedRulesCountBySeverityParams {
	return &ExtAggregateFailedRulesCountBySeverityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtAggregateFailedRulesCountBySeverityParamsWithTimeout creates a new ExtAggregateFailedRulesCountBySeverityParams object
// with the ability to set a timeout on a request.
func NewExtAggregateFailedRulesCountBySeverityParamsWithTimeout(timeout time.Duration) *ExtAggregateFailedRulesCountBySeverityParams {
	return &ExtAggregateFailedRulesCountBySeverityParams{
		timeout: timeout,
	}
}

// NewExtAggregateFailedRulesCountBySeverityParamsWithContext creates a new ExtAggregateFailedRulesCountBySeverityParams object
// with the ability to set a context for a request.
func NewExtAggregateFailedRulesCountBySeverityParamsWithContext(ctx context.Context) *ExtAggregateFailedRulesCountBySeverityParams {
	return &ExtAggregateFailedRulesCountBySeverityParams{
		Context: ctx,
	}
}

// NewExtAggregateFailedRulesCountBySeverityParamsWithHTTPClient creates a new ExtAggregateFailedRulesCountBySeverityParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtAggregateFailedRulesCountBySeverityParamsWithHTTPClient(client *http.Client) *ExtAggregateFailedRulesCountBySeverityParams {
	return &ExtAggregateFailedRulesCountBySeverityParams{
		HTTPClient: client,
	}
}

/*
ExtAggregateFailedRulesCountBySeverityParams contains all the parameters to send to the API endpoint

	for the ext aggregate failed rules count by severity operation.

	Typically these are written to a http.Request.
*/
type ExtAggregateFailedRulesCountBySeverityParams struct {

	/* Filter.

	     Filter results using a query in Falcon Query Language (FQL). Supported Filters:
	cloud_info.cluster_name: Kubernetes cluster name
	cloud_info.cloud_provider: Cloud provider
	image_id: Image ID
	compliance_finding.severity: Compliance finding severity; available values: 4, 3, 2, 1 (4: critical, 3: high, 2: medium, 1:low)
	compliance_finding.framework: Compliance finding framework (available values: CIS)
	image_digest: Image digest (sha256 digest)
	image_registry: Image registry
	asset_type: asset type (container, image)
	cid: Customer ID
	compliance_finding.name: Compliance finding Name
	image_repository: Image repository
	cloud_info.cloud_region: Cloud region
	compliance_finding.id: Compliance finding ID
	image_tag: Image tag
	cloud_info.cloud_account_id: Cloud account ID

	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ext aggregate failed rules count by severity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtAggregateFailedRulesCountBySeverityParams) WithDefaults() *ExtAggregateFailedRulesCountBySeverityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ext aggregate failed rules count by severity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtAggregateFailedRulesCountBySeverityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ext aggregate failed rules count by severity params
func (o *ExtAggregateFailedRulesCountBySeverityParams) WithTimeout(timeout time.Duration) *ExtAggregateFailedRulesCountBySeverityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ext aggregate failed rules count by severity params
func (o *ExtAggregateFailedRulesCountBySeverityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ext aggregate failed rules count by severity params
func (o *ExtAggregateFailedRulesCountBySeverityParams) WithContext(ctx context.Context) *ExtAggregateFailedRulesCountBySeverityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ext aggregate failed rules count by severity params
func (o *ExtAggregateFailedRulesCountBySeverityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ext aggregate failed rules count by severity params
func (o *ExtAggregateFailedRulesCountBySeverityParams) WithHTTPClient(client *http.Client) *ExtAggregateFailedRulesCountBySeverityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ext aggregate failed rules count by severity params
func (o *ExtAggregateFailedRulesCountBySeverityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the ext aggregate failed rules count by severity params
func (o *ExtAggregateFailedRulesCountBySeverityParams) WithFilter(filter *string) *ExtAggregateFailedRulesCountBySeverityParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the ext aggregate failed rules count by severity params
func (o *ExtAggregateFailedRulesCountBySeverityParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *ExtAggregateFailedRulesCountBySeverityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
