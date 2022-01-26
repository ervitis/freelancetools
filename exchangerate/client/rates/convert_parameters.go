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
	"github.com/go-openapi/swag"
)

// NewConvertParams creates a new ConvertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConvertParams() *ConvertParams {
	return &ConvertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConvertParamsWithTimeout creates a new ConvertParams object
// with the ability to set a timeout on a request.
func NewConvertParamsWithTimeout(timeout time.Duration) *ConvertParams {
	return &ConvertParams{
		timeout: timeout,
	}
}

// NewConvertParamsWithContext creates a new ConvertParams object
// with the ability to set a context for a request.
func NewConvertParamsWithContext(ctx context.Context) *ConvertParams {
	return &ConvertParams{
		Context: ctx,
	}
}

// NewConvertParamsWithHTTPClient creates a new ConvertParams object
// with the ability to set a custom HTTPClient for a request.
func NewConvertParamsWithHTTPClient(client *http.Client) *ConvertParams {
	return &ConvertParams{
		HTTPClient: client,
	}
}

/* ConvertParams contains all the parameters to send to the API endpoint
   for the convert operation.

   Typically these are written to a http.Request.
*/
type ConvertParams struct {

	/* Apikey.

	   Access API key
	*/
	Apikey string

	/* BaseCurrency.

	   Amount
	*/
	BaseCurrency *float64

	/* DateFrom.

	   The currency from conversion
	*/
	DateFrom *string

	/* DateTo.

	   The currency to convert
	*/
	DateTo *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the convert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConvertParams) WithDefaults() *ConvertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the convert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConvertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the convert params
func (o *ConvertParams) WithTimeout(timeout time.Duration) *ConvertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the convert params
func (o *ConvertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the convert params
func (o *ConvertParams) WithContext(ctx context.Context) *ConvertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the convert params
func (o *ConvertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the convert params
func (o *ConvertParams) WithHTTPClient(client *http.Client) *ConvertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the convert params
func (o *ConvertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApikey adds the apikey to the convert params
func (o *ConvertParams) WithApikey(apikey string) *ConvertParams {
	o.SetApikey(apikey)
	return o
}

// SetApikey adds the apikey to the convert params
func (o *ConvertParams) SetApikey(apikey string) {
	o.Apikey = apikey
}

// WithBaseCurrency adds the baseCurrency to the convert params
func (o *ConvertParams) WithBaseCurrency(baseCurrency *float64) *ConvertParams {
	o.SetBaseCurrency(baseCurrency)
	return o
}

// SetBaseCurrency adds the baseCurrency to the convert params
func (o *ConvertParams) SetBaseCurrency(baseCurrency *float64) {
	o.BaseCurrency = baseCurrency
}

// WithDateFrom adds the dateFrom to the convert params
func (o *ConvertParams) WithDateFrom(dateFrom *string) *ConvertParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the convert params
func (o *ConvertParams) SetDateFrom(dateFrom *string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the convert params
func (o *ConvertParams) WithDateTo(dateTo *string) *ConvertParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the convert params
func (o *ConvertParams) SetDateTo(dateTo *string) {
	o.DateTo = dateTo
}

// WriteToRequest writes these params to a swagger request
func (o *ConvertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrBaseCurrency float64

		if o.BaseCurrency != nil {
			qrBaseCurrency = *o.BaseCurrency
		}
		qBaseCurrency := swag.FormatFloat64(qrBaseCurrency)
		if qBaseCurrency != "" {

			if err := r.SetQueryParam("base_currency", qBaseCurrency); err != nil {
				return err
			}
		}
	}

	if o.DateFrom != nil {

		// query param date_from
		var qrDateFrom string

		if o.DateFrom != nil {
			qrDateFrom = *o.DateFrom
		}
		qDateFrom := qrDateFrom
		if qDateFrom != "" {

			if err := r.SetQueryParam("date_from", qDateFrom); err != nil {
				return err
			}
		}
	}

	if o.DateTo != nil {

		// query param date_to
		var qrDateTo string

		if o.DateTo != nil {
			qrDateTo = *o.DateTo
		}
		qDateTo := qrDateTo
		if qDateTo != "" {

			if err := r.SetQueryParam("date_to", qDateTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
