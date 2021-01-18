// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// NewRTRDeleteSessionParams creates a new RTRDeleteSessionParams object
// with the default values initialized.
func NewRTRDeleteSessionParams() *RTRDeleteSessionParams {
	var ()
	return &RTRDeleteSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRTRDeleteSessionParamsWithTimeout creates a new RTRDeleteSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRTRDeleteSessionParamsWithTimeout(timeout time.Duration) *RTRDeleteSessionParams {
	var ()
	return &RTRDeleteSessionParams{

		timeout: timeout,
	}
}

// NewRTRDeleteSessionParamsWithContext creates a new RTRDeleteSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRTRDeleteSessionParamsWithContext(ctx context.Context) *RTRDeleteSessionParams {
	var ()
	return &RTRDeleteSessionParams{

		Context: ctx,
	}
}

// NewRTRDeleteSessionParamsWithHTTPClient creates a new RTRDeleteSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRTRDeleteSessionParamsWithHTTPClient(client *http.Client) *RTRDeleteSessionParams {
	var ()
	return &RTRDeleteSessionParams{
		HTTPClient: client,
	}
}

/*RTRDeleteSessionParams contains all the parameters to send to the API endpoint
for the r t r delete session operation typically these are written to a http.Request
*/
type RTRDeleteSessionParams struct {

	/*SessionID
	  RTR Session id

	*/
	SessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the r t r delete session params
func (o *RTRDeleteSessionParams) WithTimeout(timeout time.Duration) *RTRDeleteSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r delete session params
func (o *RTRDeleteSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r delete session params
func (o *RTRDeleteSessionParams) WithContext(ctx context.Context) *RTRDeleteSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r delete session params
func (o *RTRDeleteSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r delete session params
func (o *RTRDeleteSessionParams) WithHTTPClient(client *http.Client) *RTRDeleteSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r delete session params
func (o *RTRDeleteSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionID adds the sessionID to the r t r delete session params
func (o *RTRDeleteSessionParams) WithSessionID(sessionID string) *RTRDeleteSessionParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the r t r delete session params
func (o *RTRDeleteSessionParams) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *RTRDeleteSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param session_id
	qrSessionID := o.SessionID
	qSessionID := qrSessionID
	if qSessionID != "" {
		if err := r.SetQueryParam("session_id", qSessionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}