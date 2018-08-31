// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// SalesOrderSalesOrdersService is responsible for communicating with
// the SalesOrders endpoint of the SalesOrder service.
type SalesOrderSalesOrdersService service

// SalesOrderSalesOrders:
// Service: SalesOrder
// Entity: SalesOrders
// URL: /api/v1/{division}/salesorder/SalesOrders
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderSalesOrders
type SalesOrderSalesOrders struct {
	// OrderID:
	OrderID *GUID `json:",omitempty"`

	// AmountDC:
	AmountDC *float64 `json:",omitempty"`

	// AmountDiscount:
	AmountDiscount *float64 `json:",omitempty"`

	// AmountDiscountExclVat:
	AmountDiscountExclVat *float64 `json:",omitempty"`

	// AmountFC:
	AmountFC *float64 `json:",omitempty"`

	// AmountFCExclVat:
	AmountFCExclVat *float64 `json:",omitempty"`

	// ApprovalStatus:
	ApprovalStatus *int `json:",omitempty"`

	// ApprovalStatusDescription:
	ApprovalStatusDescription *string `json:",omitempty"`

	// Approved:
	Approved *Date `json:",omitempty"`

	// Approver:
	Approver *GUID `json:",omitempty"`

	// ApproverFullName:
	ApproverFullName *string `json:",omitempty"`

	// Created:
	Created *Date `json:",omitempty"`

	// Creator:
	Creator *GUID `json:",omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:",omitempty"`

	// Currency:
	Currency *string `json:",omitempty"`

	// DeliverTo:
	DeliverTo *GUID `json:",omitempty"`

	// DeliverToContactPerson:
	DeliverToContactPerson *GUID `json:",omitempty"`

	// DeliverToContactPersonFullName:
	DeliverToContactPersonFullName *string `json:",omitempty"`

	// DeliverToName:
	DeliverToName *string `json:",omitempty"`

	// DeliveryAddress:
	DeliveryAddress *GUID `json:",omitempty"`

	// DeliveryDate:
	DeliveryDate *Date `json:",omitempty"`

	// DeliveryStatus:
	DeliveryStatus *int `json:",omitempty"`

	// DeliveryStatusDescription:
	DeliveryStatusDescription *string `json:",omitempty"`

	// Description:
	Description *string `json:",omitempty"`

	// Discount:
	Discount *float64 `json:",omitempty"`

	// Division:
	Division *int `json:",omitempty"`

	// Document:
	Document *GUID `json:",omitempty"`

	// DocumentNumber:
	DocumentNumber *int `json:",omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:",omitempty"`

	// InvoiceStatus:
	InvoiceStatus *int `json:",omitempty"`

	// InvoiceStatusDescription:
	InvoiceStatusDescription *string `json:",omitempty"`

	// InvoiceTo:
	InvoiceTo *GUID `json:",omitempty"`

	// InvoiceToContactPerson:
	InvoiceToContactPerson *GUID `json:",omitempty"`

	// InvoiceToContactPersonFullName:
	InvoiceToContactPersonFullName *string `json:",omitempty"`

	// InvoiceToName:
	InvoiceToName *string `json:",omitempty"`

	// Modified:
	Modified *Date `json:",omitempty"`

	// Modifier:
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:",omitempty"`

	// OrderDate:
	OrderDate *Date `json:",omitempty"`

	// OrderedBy:
	OrderedBy *GUID `json:",omitempty"`

	// OrderedByContactPerson:
	OrderedByContactPerson *GUID `json:",omitempty"`

	// OrderedByContactPersonFullName:
	OrderedByContactPersonFullName *string `json:",omitempty"`

	// OrderedByName:
	OrderedByName *string `json:",omitempty"`

	// OrderNumber:
	OrderNumber *int `json:",omitempty"`

	// PaymentCondition:
	PaymentCondition *string `json:",omitempty"`

	// PaymentConditionDescription:
	PaymentConditionDescription *string `json:",omitempty"`

	// PaymentReference:
	PaymentReference *string `json:",omitempty"`

	// Remarks:
	Remarks *string `json:",omitempty"`

	// SalesOrderLines:
	SalesOrderLines *[]byte `json:",omitempty"`

	// Salesperson:
	Salesperson *GUID `json:",omitempty"`

	// SalespersonFullName:
	SalespersonFullName *string `json:",omitempty"`

	// ShippingMethod:
	ShippingMethod *GUID `json:",omitempty"`

	// ShippingMethodDescription:
	ShippingMethodDescription *string `json:",omitempty"`

	// Status:
	Status *int `json:",omitempty"`

	// StatusDescription:
	StatusDescription *string `json:",omitempty"`

	// TaxSchedule:
	TaxSchedule *GUID `json:",omitempty"`

	// TaxScheduleCode:
	TaxScheduleCode *string `json:",omitempty"`

	// TaxScheduleDescription:
	TaxScheduleDescription *string `json:",omitempty"`

	// WarehouseCode:
	WarehouseCode *string `json:",omitempty"`

	// WarehouseDescription:
	WarehouseDescription *string `json:",omitempty"`

	// WarehouseID:
	WarehouseID *GUID `json:",omitempty"`

	// YourRef:
	YourRef *string `json:",omitempty"`
}

func (s *SalesOrderSalesOrders) GetIdentifier() GUID {
	return *s.OrderID
}

// List the SalesOrders entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesOrderSalesOrdersService) List(ctx context.Context, division int, all bool) ([]*SalesOrderSalesOrders, error) {
	var entities []*SalesOrderSalesOrders
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/SalesOrders?$select=*", division)
	if err != nil {
		return nil, err
	}
	if all {
		err = s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err = s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
