// Code generated by go-swagger; DO NOT EDIT.

package rates

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

// NewGetLatestParams creates a new GetLatestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLatestParams() *GetLatestParams {
	return &GetLatestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestParamsWithTimeout creates a new GetLatestParams object
// with the ability to set a timeout on a request.
func NewGetLatestParamsWithTimeout(timeout time.Duration) *GetLatestParams {
	return &GetLatestParams{
		timeout: timeout,
	}
}

// NewGetLatestParamsWithContext creates a new GetLatestParams object
// with the ability to set a context for a request.
func NewGetLatestParamsWithContext(ctx context.Context) *GetLatestParams {
	return &GetLatestParams{
		Context: ctx,
	}
}

// NewGetLatestParamsWithHTTPClient creates a new GetLatestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLatestParamsWithHTTPClient(client *http.Client) *GetLatestParams {
	return &GetLatestParams{
		HTTPClient: client,
	}
}

/* GetLatestParams contains all the parameters to send to the API endpoint
   for the get latest operation.

   Typically these are written to a http.Request.
*/
type GetLatestParams struct {

	/* Apikey.

	   Access API key
	*/
	Apikey string

	/* BaseCurrency.

	   The base currency

	   Default: "EUR"
	*/
	BaseCurrency *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get latest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLatestParams) WithDefaults() *GetLatestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get latest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLatestParams) SetDefaults() {
	var (
		baseCurrencyDefault = string("EUR")
	)

	val := GetLatestParams{
		BaseCurrency: &baseCurrencyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get latest params
func (o *GetLatestParams) WithTimeout(timeout time.Duration) *GetLatestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest params
func (o *GetLatestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest params
func (o *GetLatestParams) WithContext(ctx context.Context) *GetLatestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest params
func (o *GetLatestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest params
func (o *GetLatestParams) WithHTTPClient(client *http.Client) *GetLatestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest params
func (o *GetLatestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApikey adds the apikey to the get latest params
func (o *GetLatestParams) WithApikey(apikey string) *GetLatestParams {
	o.SetApikey(apikey)
	return o
}

// SetApikey adds the apikey to the get latest params
func (o *GetLatestParams) SetApikey(apikey string) {
	o.Apikey = apikey
}

// WithBaseCurrency adds the baseCurrency to the get latest params
func (o *GetLatestParams) WithBaseCurrency(baseCurrency *string) *GetLatestParams {
	o.SetBaseCurrency(baseCurrency)
	return o
}

// SetBaseCurrency adds the baseCurrency to the get latest params
func (o *GetLatestParams) SetBaseCurrency(baseCurrency *string) {
	o.BaseCurrency = baseCurrency
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apikey
	qrApikey := o.Apikey
	qApikey := qrApikey
	if qApikey != "" {

		if err := r.SetQueryParam("apikey", qApikey); err != nil {
			return err
		}
	}

	if o.BaseCurrency != nil {

		// query param base_currency
		var qrBaseCurrency string

		if o.BaseCurrency != nil {
			qrBaseCurrency = *o.BaseCurrency
		}
		qBaseCurrency := qrBaseCurrency
		if qBaseCurrency != "" {

			if err := r.SetQueryParam("base_currency", qBaseCurrency); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
