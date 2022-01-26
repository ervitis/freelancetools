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

// NewHistoricalParams creates a new HistoricalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHistoricalParams() *HistoricalParams {
	return &HistoricalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHistoricalParamsWithTimeout creates a new HistoricalParams object
// with the ability to set a timeout on a request.
func NewHistoricalParamsWithTimeout(timeout time.Duration) *HistoricalParams {
	return &HistoricalParams{
		timeout: timeout,
	}
}

// NewHistoricalParamsWithContext creates a new HistoricalParams object
// with the ability to set a context for a request.
func NewHistoricalParamsWithContext(ctx context.Context) *HistoricalParams {
	return &HistoricalParams{
		Context: ctx,
	}
}

// NewHistoricalParamsWithHTTPClient creates a new HistoricalParams object
// with the ability to set a custom HTTPClient for a request.
func NewHistoricalParamsWithHTTPClient(client *http.Client) *HistoricalParams {
	return &HistoricalParams{
		HTTPClient: client,
	}
}

/* HistoricalParams contains all the parameters to send to the API endpoint
   for the historical operation.

   Typically these are written to a http.Request.
*/
type HistoricalParams struct {

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

// WithDefaults hydrates default values in the historical params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HistoricalParams) WithDefaults() *HistoricalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the historical params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HistoricalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the historical params
func (o *HistoricalParams) WithTimeout(timeout time.Duration) *HistoricalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the historical params
func (o *HistoricalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the historical params
func (o *HistoricalParams) WithContext(ctx context.Context) *HistoricalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the historical params
func (o *HistoricalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the historical params
func (o *HistoricalParams) WithHTTPClient(client *http.Client) *HistoricalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the historical params
func (o *HistoricalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApikey adds the apikey to the historical params
func (o *HistoricalParams) WithApikey(apikey string) *HistoricalParams {
	o.SetApikey(apikey)
	return o
}

// SetApikey adds the apikey to the historical params
func (o *HistoricalParams) SetApikey(apikey string) {
	o.Apikey = apikey
}

// WithBaseCurrency adds the baseCurrency to the historical params
func (o *HistoricalParams) WithBaseCurrency(baseCurrency *float64) *HistoricalParams {
	o.SetBaseCurrency(baseCurrency)
	return o
}

// SetBaseCurrency adds the baseCurrency to the historical params
func (o *HistoricalParams) SetBaseCurrency(baseCurrency *float64) {
	o.BaseCurrency = baseCurrency
}

// WithDateFrom adds the dateFrom to the historical params
func (o *HistoricalParams) WithDateFrom(dateFrom *string) *HistoricalParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the historical params
func (o *HistoricalParams) SetDateFrom(dateFrom *string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the historical params
func (o *HistoricalParams) WithDateTo(dateTo *string) *HistoricalParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the historical params
func (o *HistoricalParams) SetDateTo(dateTo *string) {
	o.DateTo = dateTo
}

// WriteToRequest writes these params to a swagger request
func (o *HistoricalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
