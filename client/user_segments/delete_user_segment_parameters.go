// Code generated by go-swagger; DO NOT EDIT.

package user_segments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteUserSegmentParams creates a new DeleteUserSegmentParams object
// with the default values initialized.
func NewDeleteUserSegmentParams() *DeleteUserSegmentParams {
	var ()
	return &DeleteUserSegmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserSegmentParamsWithTimeout creates a new DeleteUserSegmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserSegmentParamsWithTimeout(timeout time.Duration) *DeleteUserSegmentParams {
	var ()
	return &DeleteUserSegmentParams{

		timeout: timeout,
	}
}

// NewDeleteUserSegmentParamsWithContext creates a new DeleteUserSegmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserSegmentParamsWithContext(ctx context.Context) *DeleteUserSegmentParams {
	var ()
	return &DeleteUserSegmentParams{

		Context: ctx,
	}
}

// NewDeleteUserSegmentParamsWithHTTPClient creates a new DeleteUserSegmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserSegmentParamsWithHTTPClient(client *http.Client) *DeleteUserSegmentParams {
	var ()
	return &DeleteUserSegmentParams{
		HTTPClient: client,
	}
}

/*DeleteUserSegmentParams contains all the parameters to send to the API endpoint
for the delete user segment operation typically these are written to a http.Request
*/
type DeleteUserSegmentParams struct {

	/*EnvironmentKey
	  The environment key, used to tie together flag configuration and users under one environment so they can be managed together.

	*/
	EnvironmentKey string
	/*ProjectKey
	  The project key, used to tie the flags together under one project so they can be managed together.

	*/
	ProjectKey string
	/*UserSegmentKey
	  The user segment's key. The key identifies the user segment in your code.

	*/
	UserSegmentKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user segment params
func (o *DeleteUserSegmentParams) WithTimeout(timeout time.Duration) *DeleteUserSegmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user segment params
func (o *DeleteUserSegmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user segment params
func (o *DeleteUserSegmentParams) WithContext(ctx context.Context) *DeleteUserSegmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user segment params
func (o *DeleteUserSegmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user segment params
func (o *DeleteUserSegmentParams) WithHTTPClient(client *http.Client) *DeleteUserSegmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user segment params
func (o *DeleteUserSegmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentKey adds the environmentKey to the delete user segment params
func (o *DeleteUserSegmentParams) WithEnvironmentKey(environmentKey string) *DeleteUserSegmentParams {
	o.SetEnvironmentKey(environmentKey)
	return o
}

// SetEnvironmentKey adds the environmentKey to the delete user segment params
func (o *DeleteUserSegmentParams) SetEnvironmentKey(environmentKey string) {
	o.EnvironmentKey = environmentKey
}

// WithProjectKey adds the projectKey to the delete user segment params
func (o *DeleteUserSegmentParams) WithProjectKey(projectKey string) *DeleteUserSegmentParams {
	o.SetProjectKey(projectKey)
	return o
}

// SetProjectKey adds the projectKey to the delete user segment params
func (o *DeleteUserSegmentParams) SetProjectKey(projectKey string) {
	o.ProjectKey = projectKey
}

// WithUserSegmentKey adds the userSegmentKey to the delete user segment params
func (o *DeleteUserSegmentParams) WithUserSegmentKey(userSegmentKey string) *DeleteUserSegmentParams {
	o.SetUserSegmentKey(userSegmentKey)
	return o
}

// SetUserSegmentKey adds the userSegmentKey to the delete user segment params
func (o *DeleteUserSegmentParams) SetUserSegmentKey(userSegmentKey string) {
	o.UserSegmentKey = userSegmentKey
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserSegmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentKey
	if err := r.SetPathParam("environmentKey", o.EnvironmentKey); err != nil {
		return err
	}

	// path param projectKey
	if err := r.SetPathParam("projectKey", o.ProjectKey); err != nil {
		return err
	}

	// path param userSegmentKey
	if err := r.SetPathParam("userSegmentKey", o.UserSegmentKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}