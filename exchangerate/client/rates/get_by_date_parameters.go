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

// NewGetByDateParams creates a new GetByDateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetByDateParams() *GetByDateParams {
	return &GetByDateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetByDateParamsWithTimeout creates a new GetByDateParams object
// with the ability to set a timeout on a request.
func NewGetByDateParamsWithTimeout(timeout time.Duration) *GetByDateParams {
	return &GetByDateParams{
		timeout: timeout,
	}
}

// NewGetByDateParamsWithContext creates a new GetByDateParams object
// with the ability to set a context for a request.
func NewGetByDateParamsWithContext(ctx context.Context) *GetByDateParams {
	return &GetByDateParams{
		Context: ctx,
	}
}

// NewGetByDateParamsWithHTTPClient creates a new GetByDateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetByDateParamsWithHTTPClient(client *http.Client) *GetByDateParams {
	return &GetByDateParams{
		HTTPClient: client,
	}
}

/* GetByDateParams contains all the parameters to send to the API endpoint
   for the get by date operation.

   Typically these are written to a http.Request.
*/
type GetByDateParams struct {

	/* AccessKey.

	   Access API key
	*/
	AccessKey string

	/* Base.

	   The base currency

	   Default: "EUR"
	*/
	Base *string

	/* Date.

	   The given date

	   Format: date
	*/
	Date strfmt.Date

	/* Symbols.

	   The exchange rates symbols returned
	*/
	Symbols []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get by date params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetByDateParams) WithDefaults() *GetByDateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get by date params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetByDateParams) SetDefaults() {
	var (
		baseDefault = string("EUR")
	)

	val := GetByDateParams{
		Base: &baseDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get by date params
func (o *GetByDateParams) WithTimeout(timeout time.Duration) *GetByDateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by date params
func (o *GetByDateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by date params
func (o *GetByDateParams) WithContext(ctx context.Context) *GetByDateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by date params
func (o *GetByDateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by date params
func (o *GetByDateParams) WithHTTPClient(client *http.Client) *GetByDateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by date params
func (o *GetByDateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKey adds the accessKey to the get by date params
func (o *GetByDateParams) WithAccessKey(accessKey string) *GetByDateParams {
	o.SetAccessKey(accessKey)
	return o
}

// SetAccessKey adds the accessKey to the get by date params
func (o *GetByDateParams) SetAccessKey(accessKey string) {
	o.AccessKey = accessKey
}

// WithBase adds the base to the get by date params
func (o *GetByDateParams) WithBase(base *string) *GetByDateParams {
	o.SetBase(base)
	return o
}

// SetBase adds the base to the get by date params
func (o *GetByDateParams) SetBase(base *string) {
	o.Base = base
}

// WithDate adds the date to the get by date params
func (o *GetByDateParams) WithDate(date strfmt.Date) *GetByDateParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the get by date params
func (o *GetByDateParams) SetDate(date strfmt.Date) {
	o.Date = date
}

// WithSymbols adds the symbols to the get by date params
func (o *GetByDateParams) WithSymbols(symbols []string) *GetByDateParams {
	o.SetSymbols(symbols)
	return o
}

// SetSymbols adds the symbols to the get by date params
func (o *GetByDateParams) SetSymbols(symbols []string) {
	o.Symbols = symbols
}

// WriteToRequest writes these params to a swagger request
func (o *GetByDateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param access_key
	qrAccessKey := o.AccessKey
	qAccessKey := qrAccessKey
	if qAccessKey != "" {

		if err := r.SetQueryParam("access_key", qAccessKey); err != nil {
			return err
		}
	}

	if o.Base != nil {

		// query param base
		var qrBase string

		if o.Base != nil {
			qrBase = *o.Base
		}
		qBase := qrBase
		if qBase != "" {

			if err := r.SetQueryParam("base", qBase); err != nil {
				return err
			}
		}
	}

	// path param date
	if err := r.SetPathParam("date", o.Date.String()); err != nil {
		return err
	}

	if o.Symbols != nil {

		// binding items for symbols
		joinedSymbols := o.bindParamSymbols(reg)

		// query array param symbols
		if err := r.SetQueryParam("symbols", joinedSymbols...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetByDate binds the parameter symbols
func (o *GetByDateParams) bindParamSymbols(formats strfmt.Registry) []string {
	symbolsIR := o.Symbols

	var symbolsIC []string
	for _, symbolsIIR := range symbolsIR { // explode []string

		symbolsIIV := symbolsIIR // string as string
		symbolsIC = append(symbolsIC, symbolsIIV)
	}

	// items.CollectionFormat: "csv"
	symbolsIS := swag.JoinByFormat(symbolsIC, "csv")

	return symbolsIS
}
