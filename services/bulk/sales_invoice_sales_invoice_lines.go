// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package bulk

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// SalesInvoiceSalesInvoiceLinesEndpoint is responsible for communicating with
// the SalesInvoiceSalesInvoiceLines endpoint of the Bulk service.
type SalesInvoiceSalesInvoiceLinesEndpoint service

// SalesInvoiceSalesInvoiceLines:
// Service: Bulk
// Entity: SalesInvoiceSalesInvoiceLines
// URL: /api/v1/{division}/bulk/SalesInvoice/SalesInvoiceLines
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=BulkSalesInvoiceSalesInvoiceLines
type SalesInvoiceSalesInvoiceLines struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

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

	// CustomerItemCode:
	CustomerItemCode *string `json:"CustomerItemCode,omitempty"`

	// CustomField:
	CustomField *string `json:"CustomField,omitempty"`

	// DeliveryDate:
	DeliveryDate *types.Date `json:"DeliveryDate,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Employee:
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName:
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EndTime:
	EndTime *types.Date `json:"EndTime,omitempty"`

	// ExtraDutyAmountFC:
	ExtraDutyAmountFC *float64 `json:"ExtraDutyAmountFC,omitempty"`

	// ExtraDutyPercentage:
	ExtraDutyPercentage *float64 `json:"ExtraDutyPercentage,omitempty"`

	// GLAccount:
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// GLAccountDescription:
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// InvoiceID:
	InvoiceID *types.GUID `json:"InvoiceID,omitempty"`

	// Item:
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode:
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription:
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// LineNumber:
	LineNumber *int `json:"LineNumber,omitempty"`

	// NetPrice:
	NetPrice *float64 `json:"NetPrice,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// Pricelist:
	Pricelist *types.GUID `json:"Pricelist,omitempty"`

	// PricelistDescription:
	PricelistDescription *string `json:"PricelistDescription,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectWBS:
	ProjectWBS *types.GUID `json:"ProjectWBS,omitempty"`

	// ProjectWBSDescription:
	ProjectWBSDescription *string `json:"ProjectWBSDescription,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// SalesOrder:
	SalesOrder *types.GUID `json:"SalesOrder,omitempty"`

	// SalesOrderLine:
	SalesOrderLine *types.GUID `json:"SalesOrderLine,omitempty"`

	// SalesOrderLineNumber:
	SalesOrderLineNumber *int `json:"SalesOrderLineNumber,omitempty"`

	// SalesOrderNumber:
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// StartTime:
	StartTime *types.Date `json:"StartTime,omitempty"`

	// Subscription:
	Subscription *types.GUID `json:"Subscription,omitempty"`

	// SubscriptionDescription:
	SubscriptionDescription *string `json:"SubscriptionDescription,omitempty"`

	// TaxSchedule:
	TaxSchedule *types.GUID `json:"TaxSchedule,omitempty"`

	// TaxScheduleCode:
	TaxScheduleCode *string `json:"TaxScheduleCode,omitempty"`

	// TaxScheduleDescription:
	TaxScheduleDescription *string `json:"TaxScheduleDescription,omitempty"`

	// UnitCode:
	UnitCode *string `json:"UnitCode,omitempty"`

	// UnitDescription:
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// UnitPrice:
	UnitPrice *float64 `json:"UnitPrice,omitempty"`

	// VATAmountDC:
	VATAmountDC *float64 `json:"VATAmountDC,omitempty"`

	// VATAmountFC:
	VATAmountFC *float64 `json:"VATAmountFC,omitempty"`

	// VATCode:
	VATCode *string `json:"VATCode,omitempty"`

	// VATCodeDescription:
	VATCodeDescription *string `json:"VATCodeDescription,omitempty"`

	// VATPercentage:
	VATPercentage *float64 `json:"VATPercentage,omitempty"`
}

func (e *SalesInvoiceSalesInvoiceLines) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SalesInvoiceSalesInvoiceLinesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "SalesInvoice/SalesInvoiceLines", method)
}

// List the SalesInvoiceSalesInvoiceLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesInvoiceSalesInvoiceLinesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SalesInvoiceSalesInvoiceLines, error) {
	var entities []*SalesInvoiceSalesInvoiceLines
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/SalesInvoice/SalesInvoiceLines", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SalesInvoiceSalesInvoiceLines entitiy in the provided division.
func (s *SalesInvoiceSalesInvoiceLinesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SalesInvoiceSalesInvoiceLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/SalesInvoice/SalesInvoiceLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SalesInvoiceSalesInvoiceLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
