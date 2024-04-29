// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package sync

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// CRMQuotationLinesEndpoint is responsible for communicating with
// the CRMQuotationLines endpoint of the Sync service.
type CRMQuotationLinesEndpoint service

// CRMQuotationLines:
// Service: Sync
// Entity: CRMQuotationLines
// URL: /api/v1/{division}/sync/CRM/QuotationLines
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncCRMQuotationLines
type CRMQuotationLines struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp:
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountDiscount:
	AmountDiscount *float64 `json:"AmountDiscount,omitempty"`

	// AmountDiscountExclVat:
	AmountDiscountExclVat *float64 `json:"AmountDiscountExclVat,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// CostCenter:
	CostCenter *string `json:"CostCenter,omitempty"`

	// CostCenterDescription:
	CostCenterDescription *string `json:"CostCenterDescription,omitempty"`

	// CostUnit:
	CostUnit *string `json:"CostUnit,omitempty"`

	// CostUnitDescription:
	CostUnitDescription *string `json:"CostUnitDescription,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// CustomerItemCode:
	CustomerItemCode *string `json:"CustomerItemCode,omitempty"`

	// CustomField:
	CustomField *string `json:"CustomField,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Item:
	Item *types.GUID `json:"Item,omitempty"`

	// ItemDescription:
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// LineNumber:
	LineNumber *int `json:"LineNumber,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// NetPrice:
	NetPrice *float64 `json:"NetPrice,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// Optional:
	Optional *bool `json:"Optional,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// QuotationID:
	QuotationID *types.GUID `json:"QuotationID,omitempty"`

	// QuotationNumber:
	QuotationNumber *int `json:"QuotationNumber,omitempty"`

	// UnitCode:
	UnitCode *string `json:"UnitCode,omitempty"`

	// UnitDescription:
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// UnitPrice:
	UnitPrice *float64 `json:"UnitPrice,omitempty"`

	// VATAmountFC:
	VATAmountFC *float64 `json:"VATAmountFC,omitempty"`

	// VATCode:
	VATCode *string `json:"VATCode,omitempty"`

	// VATDescription:
	VATDescription *string `json:"VATDescription,omitempty"`

	// VATPercentage:
	VATPercentage *float64 `json:"VATPercentage,omitempty"`
}

func (e *CRMQuotationLines) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *CRMQuotationLinesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "CRM/QuotationLines", method)
}

// List the CRMQuotationLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CRMQuotationLinesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*CRMQuotationLines, error) {
	var entities []*CRMQuotationLines
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/CRM/QuotationLines", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the CRMQuotationLines entitiy in the provided division.
func (s *CRMQuotationLinesEndpoint) Get(ctx context.Context, division int, id *int64) (*CRMQuotationLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/CRM/QuotationLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &CRMQuotationLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
