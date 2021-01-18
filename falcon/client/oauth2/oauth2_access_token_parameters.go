// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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

// NewOauth2AccessTokenParams creates a new Oauth2AccessTokenParams object
// with the default values initialized.
func NewOauth2AccessTokenParams() *Oauth2AccessTokenParams {
	var ()
	return &Oauth2AccessTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOauth2AccessTokenParamsWithTimeout creates a new Oauth2AccessTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOauth2AccessTokenParamsWithTimeout(timeout time.Duration) *Oauth2AccessTokenParams {
	var ()
	return &Oauth2AccessTokenParams{

		timeout: timeout,
	}
}

// NewOauth2AccessTokenParamsWithContext creates a new Oauth2AccessTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewOauth2AccessTokenParamsWithContext(ctx context.Context) *Oauth2AccessTokenParams {
	var ()
	return &Oauth2AccessTokenParams{

		Context: ctx,
	}
}

// NewOauth2AccessTokenParamsWithHTTPClient creates a new Oauth2AccessTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOauth2AccessTokenParamsWithHTTPClient(client *http.Client) *Oauth2AccessTokenParams {
	var ()
	return &Oauth2AccessTokenParams{
		HTTPClient: client,
	}
}

/*Oauth2AccessTokenParams contains all the parameters to send to the API endpoint
for the oauth2 access token operation typically these are written to a http.Request
*/
type Oauth2AccessTokenParams struct {

	/*ClientID
	  The API client ID to authenticate your API requests. For information on generating API clients, see [API documentation inside Falcon](https://falcon.crowdstrike.com/support/documentation/1/crowdstrike-api-introduction-for-developers).

	*/
	ClientID string
	/*ClientSecret
	  The API client secret to authenticate your API requests. For information on generating API clients, see [API documentation inside Falcon](https://falcon.crowdstrike.com/support/documentation/1/crowdstrike-api-introduction-for-developers).

	*/
	ClientSecret string
	/*MemberCid
	  For MSSP Master CIDs, optionally lock the token to act on behalf of this member CID

	*/
	MemberCid *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the oauth2 access token params
func (o *Oauth2AccessTokenParams) WithTimeout(timeout time.Duration) *Oauth2AccessTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oauth2 access token params
func (o *Oauth2AccessTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oauth2 access token params
func (o *Oauth2AccessTokenParams) WithContext(ctx context.Context) *Oauth2AccessTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oauth2 access token params
func (o *Oauth2AccessTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oauth2 access token params
func (o *Oauth2AccessTokenParams) WithHTTPClient(client *http.Client) *Oauth2AccessTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oauth2 access token params
func (o *Oauth2AccessTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the oauth2 access token params
func (o *Oauth2AccessTokenParams) WithClientID(clientID string) *Oauth2AccessTokenParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the oauth2 access token params
func (o *Oauth2AccessTokenParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the oauth2 access token params
func (o *Oauth2AccessTokenParams) WithClientSecret(clientSecret string) *Oauth2AccessTokenParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the oauth2 access token params
func (o *Oauth2AccessTokenParams) SetClientSecret(clientSecret string) {
	o.ClientSecret = clientSecret
}

// WithMemberCid adds the memberCid to the oauth2 access token params
func (o *Oauth2AccessTokenParams) WithMemberCid(memberCid *string) *Oauth2AccessTokenParams {
	o.SetMemberCid(memberCid)
	return o
}

// SetMemberCid adds the memberCid to the oauth2 access token params
func (o *Oauth2AccessTokenParams) SetMemberCid(memberCid *string) {
	o.MemberCid = memberCid
}

// WriteToRequest writes these params to a swagger request
func (o *Oauth2AccessTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param client_id
	frClientID := o.ClientID
	fClientID := frClientID
	if fClientID != "" {
		if err := r.SetFormParam("client_id", fClientID); err != nil {
			return err
		}
	}

	// form param client_secret
	frClientSecret := o.ClientSecret
	fClientSecret := frClientSecret
	if fClientSecret != "" {
		if err := r.SetFormParam("client_secret", fClientSecret); err != nil {
			return err
		}
	}

	if o.MemberCid != nil {

		// form param member_cid
		var frMemberCid string
		if o.MemberCid != nil {
			frMemberCid = *o.MemberCid
		}
		fMemberCid := frMemberCid
		if fMemberCid != "" {
			if err := r.SetFormParam("member_cid", fMemberCid); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}