// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// NewUpdateCSPMScanScheduleParams creates a new UpdateCSPMScanScheduleParams object
// with the default values initialized.
func NewUpdateCSPMScanScheduleParams() *UpdateCSPMScanScheduleParams {
	var ()
	return &UpdateCSPMScanScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCSPMScanScheduleParamsWithTimeout creates a new UpdateCSPMScanScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCSPMScanScheduleParamsWithTimeout(timeout time.Duration) *UpdateCSPMScanScheduleParams {
	var ()
	return &UpdateCSPMScanScheduleParams{

		timeout: timeout,
	}
}

// NewUpdateCSPMScanScheduleParamsWithContext creates a new UpdateCSPMScanScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCSPMScanScheduleParamsWithContext(ctx context.Context) *UpdateCSPMScanScheduleParams {
	var ()
	return &UpdateCSPMScanScheduleParams{

		Context: ctx,
	}
}

// NewUpdateCSPMScanScheduleParamsWithHTTPClient creates a new UpdateCSPMScanScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCSPMScanScheduleParamsWithHTTPClient(client *http.Client) *UpdateCSPMScanScheduleParams {
	var ()
	return &UpdateCSPMScanScheduleParams{
		HTTPClient: client,
	}
}

/*UpdateCSPMScanScheduleParams contains all the parameters to send to the API endpoint
for the update c s p m scan schedule operation typically these are written to a http.Request
*/
type UpdateCSPMScanScheduleParams struct {

	/*Body*/
	Body *models.RegistrationScanScheduleUpdateRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update c s p m scan schedule params
func (o *UpdateCSPMScanScheduleParams) WithTimeout(timeout time.Duration) *UpdateCSPMScanScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update c s p m scan schedule params
func (o *UpdateCSPMScanScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update c s p m scan schedule params
func (o *UpdateCSPMScanScheduleParams) WithContext(ctx context.Context) *UpdateCSPMScanScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update c s p m scan schedule params
func (o *UpdateCSPMScanScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update c s p m scan schedule params
func (o *UpdateCSPMScanScheduleParams) WithHTTPClient(client *http.Client) *UpdateCSPMScanScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update c s p m scan schedule params
func (o *UpdateCSPMScanScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update c s p m scan schedule params
func (o *UpdateCSPMScanScheduleParams) WithBody(body *models.RegistrationScanScheduleUpdateRequestV1) *UpdateCSPMScanScheduleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update c s p m scan schedule params
func (o *UpdateCSPMScanScheduleParams) SetBody(body *models.RegistrationScanScheduleUpdateRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCSPMScanScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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