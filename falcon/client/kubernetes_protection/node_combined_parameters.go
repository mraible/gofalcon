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
	"github.com/go-openapi/swag"
)

// NewNodeCombinedParams creates a new NodeCombinedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodeCombinedParams() *NodeCombinedParams {
	return &NodeCombinedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodeCombinedParamsWithTimeout creates a new NodeCombinedParams object
// with the ability to set a timeout on a request.
func NewNodeCombinedParamsWithTimeout(timeout time.Duration) *NodeCombinedParams {
	return &NodeCombinedParams{
		timeout: timeout,
	}
}

// NewNodeCombinedParamsWithContext creates a new NodeCombinedParams object
// with the ability to set a context for a request.
func NewNodeCombinedParamsWithContext(ctx context.Context) *NodeCombinedParams {
	return &NodeCombinedParams{
		Context: ctx,
	}
}

// NewNodeCombinedParamsWithHTTPClient creates a new NodeCombinedParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodeCombinedParamsWithHTTPClient(client *http.Client) *NodeCombinedParams {
	return &NodeCombinedParams{
		HTTPClient: client,
	}
}

/*
NodeCombinedParams contains all the parameters to send to the API endpoint

	for the node combined operation.

	Typically these are written to a http.Request.
*/
type NodeCombinedParams struct {

	/* Filter.

	   Search Kubernetes nodes using a query in Falcon Query Language (FQL). Supported filters:  agent_id,agent_type,annotations_list,cid,cloud_account_id,cloud_name,cloud_region,cloud_service,cluster_id,cluster_name,container_count,container_runtime_version,first_seen,image_digest,ipv4,kac_agent_id,last_seen,linux_sensor_coverage,node_name,pod_count,resource_status
	*/
	Filter *string

	/* Limit.

	   The upper-bound on the number of records to retrieve.
	*/
	Limit *int64

	/* Offset.

	   The offset from where to begin.
	*/
	Offset *int64

	/* Sort.

	   Field to sort results by
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the node combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeCombinedParams) WithDefaults() *NodeCombinedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the node combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeCombinedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the node combined params
func (o *NodeCombinedParams) WithTimeout(timeout time.Duration) *NodeCombinedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node combined params
func (o *NodeCombinedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node combined params
func (o *NodeCombinedParams) WithContext(ctx context.Context) *NodeCombinedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node combined params
func (o *NodeCombinedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node combined params
func (o *NodeCombinedParams) WithHTTPClient(client *http.Client) *NodeCombinedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node combined params
func (o *NodeCombinedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the node combined params
func (o *NodeCombinedParams) WithFilter(filter *string) *NodeCombinedParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the node combined params
func (o *NodeCombinedParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the node combined params
func (o *NodeCombinedParams) WithLimit(limit *int64) *NodeCombinedParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the node combined params
func (o *NodeCombinedParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the node combined params
func (o *NodeCombinedParams) WithOffset(offset *int64) *NodeCombinedParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the node combined params
func (o *NodeCombinedParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the node combined params
func (o *NodeCombinedParams) WithSort(sort *string) *NodeCombinedParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the node combined params
func (o *NodeCombinedParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *NodeCombinedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
