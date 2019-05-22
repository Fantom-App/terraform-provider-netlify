// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/open-api/go/models"
)

// NewCreateHookBySiteIDParams creates a new CreateHookBySiteIDParams object
// with the default values initialized.
func NewCreateHookBySiteIDParams() *CreateHookBySiteIDParams {
	var ()
	return &CreateHookBySiteIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHookBySiteIDParamsWithTimeout creates a new CreateHookBySiteIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateHookBySiteIDParamsWithTimeout(timeout time.Duration) *CreateHookBySiteIDParams {
	var ()
	return &CreateHookBySiteIDParams{

		timeout: timeout,
	}
}

// NewCreateHookBySiteIDParamsWithContext creates a new CreateHookBySiteIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateHookBySiteIDParamsWithContext(ctx context.Context) *CreateHookBySiteIDParams {
	var ()
	return &CreateHookBySiteIDParams{

		Context: ctx,
	}
}

// NewCreateHookBySiteIDParamsWithHTTPClient creates a new CreateHookBySiteIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateHookBySiteIDParamsWithHTTPClient(client *http.Client) *CreateHookBySiteIDParams {
	var ()
	return &CreateHookBySiteIDParams{
		HTTPClient: client,
	}
}

/*CreateHookBySiteIDParams contains all the parameters to send to the API endpoint
for the create hook by site Id operation typically these are written to a http.Request
*/
type CreateHookBySiteIDParams struct {

	/*Hook*/
	Hook *models.Hook
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create hook by site Id params
func (o *CreateHookBySiteIDParams) WithTimeout(timeout time.Duration) *CreateHookBySiteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hook by site Id params
func (o *CreateHookBySiteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hook by site Id params
func (o *CreateHookBySiteIDParams) WithContext(ctx context.Context) *CreateHookBySiteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hook by site Id params
func (o *CreateHookBySiteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hook by site Id params
func (o *CreateHookBySiteIDParams) WithHTTPClient(client *http.Client) *CreateHookBySiteIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hook by site Id params
func (o *CreateHookBySiteIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHook adds the hook to the create hook by site Id params
func (o *CreateHookBySiteIDParams) WithHook(hook *models.Hook) *CreateHookBySiteIDParams {
	o.SetHook(hook)
	return o
}

// SetHook adds the hook to the create hook by site Id params
func (o *CreateHookBySiteIDParams) SetHook(hook *models.Hook) {
	o.Hook = hook
}

// WithSiteID adds the siteID to the create hook by site Id params
func (o *CreateHookBySiteIDParams) WithSiteID(siteID string) *CreateHookBySiteIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the create hook by site Id params
func (o *CreateHookBySiteIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHookBySiteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Hook != nil {
		if err := r.SetBodyParam(o.Hook); err != nil {
			return err
		}
	}

	// query param site_id
	qrSiteID := o.SiteID
	qSiteID := qrSiteID
	if qSiteID != "" {
		if err := r.SetQueryParam("site_id", qSiteID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
