// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package bulk

import (
	"context"
	"encoding/json"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// CRMQuotationsEndpoint is responsible for communicating with
// the CRMQuotations endpoint of the Bulk service.
type CRMQuotationsEndpoint service

// CRMQuotations:
// Service: Bulk
// Entity: CRMQuotations
// URL: /api/v1/{division}/bulk/CRM/Quotations
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=BulkCRMQuotations
type CRMQuotations struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// QuotationID:
	QuotationID *types.GUID `json:"QuotationID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountDiscount:
	AmountDiscount *float64 `json:"AmountDiscount,omitempty"`

	// AmountDiscountExclVat:
	AmountDiscountExclVat *float64 `json:"AmountDiscountExclVat,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// CloseDate:
	CloseDate *types.Date `json:"CloseDate,omitempty"`

	// ClosingDate:
	ClosingDate *types.Date `json:"ClosingDate,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// DeliveryAccount:
	DeliveryAccount *types.GUID `json:"DeliveryAccount,omitempty"`

	// DeliveryAccountCode:
	DeliveryAccountCode *string `json:"DeliveryAccountCode,omitempty"`

	// DeliveryAccountContact:
	DeliveryAccountContact *types.GUID `json:"DeliveryAccountContact,omitempty"`

	// DeliveryAccountContactFullName:
	DeliveryAccountContactFullName *string `json:"DeliveryAccountContactFullName,omitempty"`

	// DeliveryAccountName:
	DeliveryAccountName *string `json:"DeliveryAccountName,omitempty"`

	// DeliveryAddress:
	DeliveryAddress *types.GUID `json:"DeliveryAddress,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// DueDate:
	DueDate *types.Date `json:"DueDate,omitempty"`

	// IncotermAddress:
	IncotermAddress *string `json:"IncotermAddress,omitempty"`

	// IncotermCode:
	IncotermCode *string `json:"IncotermCode,omitempty"`

	// IncotermVersion:
	IncotermVersion *int `json:"IncotermVersion,omitempty"`

	// InvoiceAccount:
	InvoiceAccount *types.GUID `json:"InvoiceAccount,omitempty"`

	// InvoiceAccountCode:
	InvoiceAccountCode *string `json:"InvoiceAccountCode,omitempty"`

	// InvoiceAccountContact:
	InvoiceAccountContact *types.GUID `json:"InvoiceAccountContact,omitempty"`

	// InvoiceAccountContactFullName:
	InvoiceAccountContactFullName *string `json:"InvoiceAccountContactFullName,omitempty"`

	// InvoiceAccountName:
	InvoiceAccountName *string `json:"InvoiceAccountName,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Opportunity:
	Opportunity *types.GUID `json:"Opportunity,omitempty"`

	// OpportunityName:
	OpportunityName *string `json:"OpportunityName,omitempty"`

	// OrderAccount:
	OrderAccount *types.GUID `json:"OrderAccount,omitempty"`

	// OrderAccountCode:
	OrderAccountCode *string `json:"OrderAccountCode,omitempty"`

	// OrderAccountContact:
	OrderAccountContact *types.GUID `json:"OrderAccountContact,omitempty"`

	// OrderAccountContactFullName:
	OrderAccountContactFullName *string `json:"OrderAccountContactFullName,omitempty"`

	// OrderAccountName:
	OrderAccountName *string `json:"OrderAccountName,omitempty"`

	// PaymentCondition:
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription:
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode:
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// QuotationDate:
	QuotationDate *types.Date `json:"QuotationDate,omitempty"`

	// QuotationLines:
	QuotationLines *json.RawMessage `json:"QuotationLines,omitempty"`

	// QuotationNumber:
	QuotationNumber *int `json:"QuotationNumber,omitempty"`

	// Remarks:
	Remarks *string `json:"Remarks,omitempty"`

	// SalesChannel:
	SalesChannel *types.GUID `json:"SalesChannel,omitempty"`

	// SalesChannelCode:
	SalesChannelCode *string `json:"SalesChannelCode,omitempty"`

	// SalesChannelDescription:
	SalesChannelDescription *string `json:"SalesChannelDescription,omitempty"`

	// SalesPerson:
	SalesPerson *types.GUID `json:"SalesPerson,omitempty"`

	// SalesPersonFullName:
	SalesPersonFullName *string `json:"SalesPersonFullName,omitempty"`

	// SelectionCode:
	SelectionCode *types.GUID `json:"SelectionCode,omitempty"`

	// SelectionCodeCode:
	SelectionCodeCode *string `json:"SelectionCodeCode,omitempty"`

	// SelectionCodeDescription:
	SelectionCodeDescription *string `json:"SelectionCodeDescription,omitempty"`

	// ShippingMethod:
	ShippingMethod *types.GUID `json:"ShippingMethod,omitempty"`

	// ShippingMethodDescription:
	ShippingMethodDescription *string `json:"ShippingMethodDescription,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// StatusDescription:
	StatusDescription *string `json:"StatusDescription,omitempty"`

	// VATAmountFC:
	VATAmountFC *float64 `json:"VATAmountFC,omitempty"`

	// VersionNumber:
	VersionNumber *int `json:"VersionNumber,omitempty"`

	// YourRef:
	YourRef *string `json:"YourRef,omitempty"`
}

func (e *CRMQuotations) GetPrimary() *types.GUID {
	return e.QuotationID
}

func (s *CRMQuotationsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "CRM/Quotations", method)
}

// List the CRMQuotations entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CRMQuotationsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*CRMQuotations, error) {
	var entities []*CRMQuotations
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/CRM/Quotations", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the CRMQuotations entitiy in the provided division.
func (s *CRMQuotationsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*CRMQuotations, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/CRM/Quotations", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &CRMQuotations{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
