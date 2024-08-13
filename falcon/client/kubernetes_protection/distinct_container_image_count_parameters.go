// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// NewDistinctContainerImageCountParams creates a new DistinctContainerImageCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDistinctContainerImageCountParams() *DistinctContainerImageCountParams {
	return &DistinctContainerImageCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDistinctContainerImageCountParamsWithTimeout creates a new DistinctContainerImageCountParams object
// with the ability to set a timeout on a request.
func NewDistinctContainerImageCountParamsWithTimeout(timeout time.Duration) *DistinctContainerImageCountParams {
	return &DistinctContainerImageCountParams{
		timeout: timeout,
	}
}

// NewDistinctContainerImageCountParamsWithContext creates a new DistinctContainerImageCountParams object
// with the ability to set a context for a request.
func NewDistinctContainerImageCountParamsWithContext(ctx context.Context) *DistinctContainerImageCountParams {
	return &DistinctContainerImageCountParams{
		Context: ctx,
	}
}

// NewDistinctContainerImageCountParamsWithHTTPClient creates a new DistinctContainerImageCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewDistinctContainerImageCountParamsWithHTTPClient(client *http.Client) *DistinctContainerImageCountParams {
	return &DistinctContainerImageCountParams{
		HTTPClient: client,
	}
}

/*
DistinctContainerImageCountParams contains all the parameters to send to the API endpoint

	for the distinct container image count operation.

	Typically these are written to a http.Request.
*/
type DistinctContainerImageCountParams struct {

	/* Filter.

	   Search Kubernetes containers using a query in Falcon Query Language (FQL). Supported filters:  agent_id,agent_type,allow_privilege_escalation,cid,cloud_account_id,cloud_name,cloud_region,cloud_service,cluster_id,cluster_name,container_id,container_image_id,container_name,cve_id,detection_name,first_seen,image_detection_count,image_digest,image_has_been_assessed,image_id,image_registry,image_repository,image_tag,image_vulnerability_count,insecure_mount_source,insecure_mount_type,insecure_propagation_mode,interactive_mode,ipv4,ipv6,kac_agent_id,labels,last_seen,namespace,node_name,node_uid,package_name_version,pod_id,pod_name,port,privileged,root_write_access,run_as_root_group,run_as_root_user,running_status
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the distinct container image count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DistinctContainerImageCountParams) WithDefaults() *DistinctContainerImageCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the distinct container image count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DistinctContainerImageCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the distinct container image count params
func (o *DistinctContainerImageCountParams) WithTimeout(timeout time.Duration) *DistinctContainerImageCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distinct container image count params
func (o *DistinctContainerImageCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distinct container image count params
func (o *DistinctContainerImageCountParams) WithContext(ctx context.Context) *DistinctContainerImageCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distinct container image count params
func (o *DistinctContainerImageCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distinct container image count params
func (o *DistinctContainerImageCountParams) WithHTTPClient(client *http.Client) *DistinctContainerImageCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distinct container image count params
func (o *DistinctContainerImageCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the distinct container image count params
func (o *DistinctContainerImageCountParams) WithFilter(filter *string) *DistinctContainerImageCountParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the distinct container image count params
func (o *DistinctContainerImageCountParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *DistinctContainerImageCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
