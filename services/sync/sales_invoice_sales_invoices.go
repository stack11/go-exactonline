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

// SalesInvoiceSalesInvoicesEndpoint is responsible for communicating with
// the SalesInvoiceSalesInvoices endpoint of the Sync service.
type SalesInvoiceSalesInvoicesEndpoint service

// SalesInvoiceSalesInvoices:
// Service: Sync
// Entity: SalesInvoiceSalesInvoices
// URL: /api/v1/{division}/sync/SalesInvoice/SalesInvoices
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncSalesInvoiceSalesInvoices
type SalesInvoiceSalesInvoices struct {
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

	// AmountFCExclVat:
	AmountFCExclVat *float64 `json:"AmountFCExclVat,omitempty"`

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

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// CustomerItemCode:
	CustomerItemCode *string `json:"CustomerItemCode,omitempty"`

	// DeliverTo:
	DeliverTo *types.GUID `json:"DeliverTo,omitempty"`

	// DeliverToAddress:
	DeliverToAddress *types.GUID `json:"DeliverToAddress,omitempty"`

	// DeliverToContactPerson:
	DeliverToContactPerson *types.GUID `json:"DeliverToContactPerson,omitempty"`

	// DeliverToContactPersonFullName:
	DeliverToContactPersonFullName *string `json:"DeliverToContactPersonFullName,omitempty"`

	// DeliverToName:
	DeliverToName *string `json:"DeliverToName,omitempty"`

	// DeliveryDate:
	DeliveryDate *types.Date `json:"DeliveryDate,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// DiscountType:
	DiscountType *int `json:"DiscountType,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentNumber:
	DocumentNumber *int `json:"DocumentNumber,omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// DueDate:
	DueDate *types.Date `json:"DueDate,omitempty"`

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

	// GAccountAmountFC:
	GAccountAmountFC *float64 `json:"GAccountAmountFC,omitempty"`

	// GLAccount:
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// GLAccountDescription:
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// IncotermAddress:
	IncotermAddress *string `json:"IncotermAddress,omitempty"`

	// IncotermCode:
	IncotermCode *string `json:"IncotermCode,omitempty"`

	// IncotermVersion:
	IncotermVersion *int `json:"IncotermVersion,omitempty"`

	// InvoiceDate:
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// InvoiceID:
	InvoiceID *types.GUID `json:"InvoiceID,omitempty"`

	// InvoiceNumber:
	InvoiceNumber *int `json:"InvoiceNumber,omitempty"`

	// InvoiceTo:
	InvoiceTo *types.GUID `json:"InvoiceTo,omitempty"`

	// InvoiceToContactPerson:
	InvoiceToContactPerson *types.GUID `json:"InvoiceToContactPerson,omitempty"`

	// InvoiceToContactPersonFullName:
	InvoiceToContactPersonFullName *string `json:"InvoiceToContactPersonFullName,omitempty"`

	// InvoiceToName:
	InvoiceToName *string `json:"InvoiceToName,omitempty"`

	// IsExtraDuty:
	IsExtraDuty *bool `json:"IsExtraDuty,omitempty"`

	// Item:
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode:
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription:
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Journal:
	Journal *string `json:"Journal,omitempty"`

	// JournalDescription:
	JournalDescription *string `json:"JournalDescription,omitempty"`

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

	// OrderDate:
	OrderDate *types.Date `json:"OrderDate,omitempty"`

	// OrderedBy:
	OrderedBy *types.GUID `json:"OrderedBy,omitempty"`

	// OrderedByContactPerson:
	OrderedByContactPerson *types.GUID `json:"OrderedByContactPerson,omitempty"`

	// OrderedByContactPersonFullName:
	OrderedByContactPersonFullName *string `json:"OrderedByContactPersonFullName,omitempty"`

	// OrderedByName:
	OrderedByName *string `json:"OrderedByName,omitempty"`

	// OrderNumber:
	OrderNumber *int `json:"OrderNumber,omitempty"`

	// PaymentCondition:
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription:
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// PaymentReference:
	PaymentReference *string `json:"PaymentReference,omitempty"`

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

	// Remarks:
	Remarks *string `json:"Remarks,omitempty"`

	// SalesChannel:
	SalesChannel *types.GUID `json:"SalesChannel,omitempty"`

	// SalesChannelCode:
	SalesChannelCode *string `json:"SalesChannelCode,omitempty"`

	// SalesChannelDescription:
	SalesChannelDescription *string `json:"SalesChannelDescription,omitempty"`

	// SalesOrder:
	SalesOrder *types.GUID `json:"SalesOrder,omitempty"`

	// SalesOrderLine:
	SalesOrderLine *types.GUID `json:"SalesOrderLine,omitempty"`

	// SalesOrderLineNumber:
	SalesOrderLineNumber *int `json:"SalesOrderLineNumber,omitempty"`

	// SalesOrderNumber:
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// Salesperson:
	Salesperson *types.GUID `json:"Salesperson,omitempty"`

	// SalespersonFullName:
	SalespersonFullName *string `json:"SalespersonFullName,omitempty"`

	// StarterSalesInvoiceStatus:
	StarterSalesInvoiceStatus *int `json:"StarterSalesInvoiceStatus,omitempty"`

	// StarterSalesInvoiceStatusDescription:
	StarterSalesInvoiceStatusDescription *string `json:"StarterSalesInvoiceStatusDescription,omitempty"`

	// StartTime:
	StartTime *types.Date `json:"StartTime,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// StatusDescription:
	StatusDescription *string `json:"StatusDescription,omitempty"`

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

	// Type:
	Type *int `json:"Type,omitempty"`

	// TypeDescription:
	TypeDescription *string `json:"TypeDescription,omitempty"`

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

	// Warehouse:
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WithholdingTaxAmountFC:
	WithholdingTaxAmountFC *float64 `json:"WithholdingTaxAmountFC,omitempty"`

	// WithholdingTaxBaseAmount:
	WithholdingTaxBaseAmount *float64 `json:"WithholdingTaxBaseAmount,omitempty"`

	// WithholdingTaxPercentage:
	WithholdingTaxPercentage *float64 `json:"WithholdingTaxPercentage,omitempty"`

	// YourRef:
	YourRef *string `json:"YourRef,omitempty"`
}

func (e *SalesInvoiceSalesInvoices) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *SalesInvoiceSalesInvoicesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "SalesInvoice/SalesInvoices", method)
}

// List the SalesInvoiceSalesInvoices entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesInvoiceSalesInvoicesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SalesInvoiceSalesInvoices, error) {
	var entities []*SalesInvoiceSalesInvoices
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/SalesInvoice/SalesInvoices", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SalesInvoiceSalesInvoices entitiy in the provided division.
func (s *SalesInvoiceSalesInvoicesEndpoint) Get(ctx context.Context, division int, id *int64) (*SalesInvoiceSalesInvoices, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/SalesInvoice/SalesInvoices", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SalesInvoiceSalesInvoices{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
