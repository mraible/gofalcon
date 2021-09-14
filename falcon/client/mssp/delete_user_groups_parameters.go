// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// NewDeleteUserGroupsParams creates a new DeleteUserGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserGroupsParams() *DeleteUserGroupsParams {
	return &DeleteUserGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserGroupsParamsWithTimeout creates a new DeleteUserGroupsParams object
// with the ability to set a timeout on a request.
func NewDeleteUserGroupsParamsWithTimeout(timeout time.Duration) *DeleteUserGroupsParams {
	return &DeleteUserGroupsParams{
		timeout: timeout,
	}
}

// NewDeleteUserGroupsParamsWithContext creates a new DeleteUserGroupsParams object
// with the ability to set a context for a request.
func NewDeleteUserGroupsParamsWithContext(ctx context.Context) *DeleteUserGroupsParams {
	return &DeleteUserGroupsParams{
		Context: ctx,
	}
}

// NewDeleteUserGroupsParamsWithHTTPClient creates a new DeleteUserGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserGroupsParamsWithHTTPClient(client *http.Client) *DeleteUserGroupsParams {
	return &DeleteUserGroupsParams{
		HTTPClient: client,
	}
}

/* DeleteUserGroupsParams contains all the parameters to send to the API endpoint
   for the delete user groups operation.

   Typically these are written to a http.Request.
*/
type DeleteUserGroupsParams struct {

	/* UserGroupIds.

	   User group IDs to delete
	*/
	UserGroupIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserGroupsParams) WithDefaults() *DeleteUserGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user groups params
func (o *DeleteUserGroupsParams) WithTimeout(timeout time.Duration) *DeleteUserGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user groups params
func (o *DeleteUserGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user groups params
func (o *DeleteUserGroupsParams) WithContext(ctx context.Context) *DeleteUserGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user groups params
func (o *DeleteUserGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user groups params
func (o *DeleteUserGroupsParams) WithHTTPClient(client *http.Client) *DeleteUserGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user groups params
func (o *DeleteUserGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserGroupIds adds the userGroupIds to the delete user groups params
func (o *DeleteUserGroupsParams) WithUserGroupIds(userGroupIds []string) *DeleteUserGroupsParams {
	o.SetUserGroupIds(userGroupIds)
	return o
}

// SetUserGroupIds adds the userGroupIds to the delete user groups params
func (o *DeleteUserGroupsParams) SetUserGroupIds(userGroupIds []string) {
	o.UserGroupIds = userGroupIds
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserGroupIds != nil {

		// binding items for user_group_ids
		joinedUserGroupIds := o.bindParamUserGroupIds(reg)

		// query array param user_group_ids
		if err := r.SetQueryParam("user_group_ids", joinedUserGroupIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteUserGroups binds the parameter user_group_ids
func (o *DeleteUserGroupsParams) bindParamUserGroupIds(formats strfmt.Registry) []string {
	userGroupIdsIR := o.UserGroupIds

	var userGroupIdsIC []string
	for _, userGroupIdsIIR := range userGroupIdsIR { // explode []string

		userGroupIdsIIV := userGroupIdsIIR // string as string
		userGroupIdsIC = append(userGroupIdsIC, userGroupIdsIIV)
	}

	// items.CollectionFormat: "multi"
	userGroupIdsIS := swag.JoinByFormat(userGroupIdsIC, "multi")

	return userGroupIdsIS
}
