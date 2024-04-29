// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesorder

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// SalesOrderOrderChargeLinesEndpoint is responsible for communicating with
// the SalesOrderOrderChargeLines endpoint of the SalesOrder service.
type SalesOrderOrderChargeLinesEndpoint service

// SalesOrderOrderChargeLines:
// Service: SalesOrder
// Entity: SalesOrderOrderChargeLines
// URL: /api/v1/{division}/salesorder/SalesOrderOrderChargeLines
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderSalesOrderOrderChargeLines
type SalesOrderOrderChargeLines struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Line ID of shipping method or order charges
	ID *types.GUID `json:"ID,omitempty"`

	// AmountDC: Amount excluded VAT in reporting currency for shipping cost or order charges
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFCExclVAT: Amount excluded VAT in trading currency for shipping cost or order charges
	AmountFCExclVAT *float64 `json:"AmountFCExclVAT,omitempty"`

	// AmountFCInclVAT: Amount included VAT in trading currency for shipping cost or order charges
	AmountFCInclVAT *float64 `json:"AmountFCInclVAT,omitempty"`

	// AmountVATFC: VAT amount in trading currency for shipping cost or order charges
	AmountVATFC *float64 `json:"AmountVATFC,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// IsShippingCost: Indicates whether the order charge line is shipping cost
	IsShippingCost *bool `json:"IsShippingCost,omitempty"`

	// LineNumber: Line number of shipping cost and order charges
	LineNumber *int `json:"LineNumber,omitempty"`

	// OrderCharge: ID of order charges is mandatory for order charge. However, it is not required for shipping cost
	OrderCharge *types.GUID `json:"OrderCharge,omitempty"`

	// OrderChargeCode: Code of shipping method or order charges
	OrderChargeCode *string `json:"OrderChargeCode,omitempty"`

	// OrderChargeDescription: Description from shipping method or order charges master
	OrderChargeDescription *string `json:"OrderChargeDescription,omitempty"`

	// OrderChargesLineDescription: Line description of shipping cost or order charges (only available in WD Premium packages)
	OrderChargesLineDescription *string `json:"OrderChargesLineDescription,omitempty"`

	// OrderID: The OrderID identifies the sales order. All the lines of a sales order have the same OrderID
	OrderID *types.GUID `json:"OrderID,omitempty"`

	// VATCode: VAT code that is used for shipping cost or order charges
	VATCode *string `json:"VATCode,omitempty"`

	// VATDescription: VAT description for shipping cost or order charges
	VATDescription *string `json:"VATDescription,omitempty"`

	// VATPercentage: The vat percentage of the VAT code. This is the percentage at the moment the invoice is created. It&#39;s also used for the default calculation of VAT amounts and VAT base amounts
	VATPercentage *float64 `json:"VATPercentage,omitempty"`
}

func (e *SalesOrderOrderChargeLines) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SalesOrderOrderChargeLinesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "salesorder/SalesOrderOrderChargeLines", method)
}

// List the SalesOrderOrderChargeLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesOrderOrderChargeLinesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SalesOrderOrderChargeLines, error) {
	var entities []*SalesOrderOrderChargeLines
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/SalesOrderOrderChargeLines", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SalesOrderOrderChargeLines entitiy in the provided division.
func (s *SalesOrderOrderChargeLinesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SalesOrderOrderChargeLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/SalesOrderOrderChargeLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SalesOrderOrderChargeLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty SalesOrderOrderChargeLines entity
func (s *SalesOrderOrderChargeLinesEndpoint) New() *SalesOrderOrderChargeLines {
	return &SalesOrderOrderChargeLines{}
}

// Create the SalesOrderOrderChargeLines entity in the provided division.
func (s *SalesOrderOrderChargeLinesEndpoint) Create(ctx context.Context, division int, entity *SalesOrderOrderChargeLines) (*SalesOrderOrderChargeLines, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/SalesOrderOrderChargeLines", division) // #nosec
	e := &SalesOrderOrderChargeLines{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the SalesOrderOrderChargeLines entity in the provided division.
func (s *SalesOrderOrderChargeLinesEndpoint) Update(ctx context.Context, division int, entity *SalesOrderOrderChargeLines) (*SalesOrderOrderChargeLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/SalesOrderOrderChargeLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &SalesOrderOrderChargeLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the SalesOrderOrderChargeLines entity in the provided division.
func (s *SalesOrderOrderChargeLinesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/SalesOrderOrderChargeLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return err
	}

	_, r, requestError := s.client.NewRequestAndDo(ctx, "DELETE", u.String(), nil, nil)
	if requestError != nil {
		return requestError
	}

	if r.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(r.Body) // #nosec
		return fmt.Errorf("Failed with status %v and body %v", r.StatusCode, body)
	}

	return nil
}
