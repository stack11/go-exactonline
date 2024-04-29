// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package purchaseorder

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// PurchaseOrdersEndpoint is responsible for communicating with
// the PurchaseOrders endpoint of the PurchaseOrder service.
type PurchaseOrdersEndpoint service

// PurchaseOrders:
// Service: PurchaseOrder
// Entity: PurchaseOrders
// URL: /api/v1/{division}/purchaseorder/PurchaseOrders
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PurchaseOrderPurchaseOrders
type PurchaseOrders struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// PurchaseOrderID:
	PurchaseOrderID *types.GUID `json:"PurchaseOrderID,omitempty"`

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

	// ApprovalStatus:
	ApprovalStatus *int `json:"ApprovalStatus,omitempty"`

	// ApprovalStatusDescription:
	ApprovalStatusDescription *string `json:"ApprovalStatusDescription,omitempty"`

	// Approved:
	Approved *types.Date `json:"Approved,omitempty"`

	// Approver:
	Approver *types.GUID `json:"Approver,omitempty"`

	// ApproverFullName:
	ApproverFullName *string `json:"ApproverFullName,omitempty"`

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

	// DeliveryAccountName:
	DeliveryAccountName *string `json:"DeliveryAccountName,omitempty"`

	// DeliveryAddress:
	DeliveryAddress *types.GUID `json:"DeliveryAddress,omitempty"`

	// DeliveryContact:
	DeliveryContact *types.GUID `json:"DeliveryContact,omitempty"`

	// DeliveryContactPersonFullName:
	DeliveryContactPersonFullName *string `json:"DeliveryContactPersonFullName,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// DropShipment:
	DropShipment *bool `json:"DropShipment,omitempty"`

	// ExchangeRate:
	ExchangeRate *float64 `json:"ExchangeRate,omitempty"`

	// IncotermAddress:
	IncotermAddress *string `json:"IncotermAddress,omitempty"`

	// IncotermCode:
	IncotermCode *string `json:"IncotermCode,omitempty"`

	// IncotermVersion:
	IncotermVersion *int `json:"IncotermVersion,omitempty"`

	// InvoiceStatus:
	InvoiceStatus *int `json:"InvoiceStatus,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// OrderDate:
	OrderDate *types.Date `json:"OrderDate,omitempty"`

	// OrderNumber:
	OrderNumber *int `json:"OrderNumber,omitempty"`

	// OrderStatus:
	OrderStatus *int `json:"OrderStatus,omitempty"`

	// PaymentCondition:
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription:
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// PurchaseAgent:
	PurchaseAgent *types.GUID `json:"PurchaseAgent,omitempty"`

	// PurchaseAgentFullName:
	PurchaseAgentFullName *string `json:"PurchaseAgentFullName,omitempty"`

	// PurchaseOrderLineCount:
	PurchaseOrderLineCount *int `json:"PurchaseOrderLineCount,omitempty"`

	// PurchaseOrderLines:
	PurchaseOrderLines *json.RawMessage `json:"PurchaseOrderLines,omitempty"`

	// ReceiptDate:
	ReceiptDate *types.Date `json:"ReceiptDate,omitempty"`

	// ReceiptStatus:
	ReceiptStatus *int `json:"ReceiptStatus,omitempty"`

	// Remarks:
	Remarks *string `json:"Remarks,omitempty"`

	// SalesOrder:
	SalesOrder *types.GUID `json:"SalesOrder,omitempty"`

	// SalesOrderNumber:
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// SelectionCode:
	SelectionCode *types.GUID `json:"SelectionCode,omitempty"`

	// SelectionCodeCode:
	SelectionCodeCode *string `json:"SelectionCodeCode,omitempty"`

	// SelectionCodeDescription:
	SelectionCodeDescription *string `json:"SelectionCodeDescription,omitempty"`

	// ShippingMethod:
	ShippingMethod *types.GUID `json:"ShippingMethod,omitempty"`

	// ShippingMethodCode:
	ShippingMethodCode *string `json:"ShippingMethodCode,omitempty"`

	// ShippingMethodDescription:
	ShippingMethodDescription *string `json:"ShippingMethodDescription,omitempty"`

	// Source:
	Source *int `json:"Source,omitempty"`

	// Supplier:
	Supplier *types.GUID `json:"Supplier,omitempty"`

	// SupplierCode:
	SupplierCode *string `json:"SupplierCode,omitempty"`

	// SupplierContact:
	SupplierContact *types.GUID `json:"SupplierContact,omitempty"`

	// SupplierContactPersonFullName:
	SupplierContactPersonFullName *string `json:"SupplierContactPersonFullName,omitempty"`

	// SupplierName:
	SupplierName *string `json:"SupplierName,omitempty"`

	// VATAmount:
	VATAmount *float64 `json:"VATAmount,omitempty"`

	// Warehouse:
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode:
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription:
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`

	// YourRef:
	YourRef *string `json:"YourRef,omitempty"`
}

func (e *PurchaseOrders) GetPrimary() *types.GUID {
	return e.PurchaseOrderID
}

func (s *PurchaseOrdersEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "purchaseorder/PurchaseOrders", method)
}

// List the PurchaseOrders entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PurchaseOrdersEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*PurchaseOrders, error) {
	var entities []*PurchaseOrders
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/PurchaseOrders", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the PurchaseOrders entitiy in the provided division.
func (s *PurchaseOrdersEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*PurchaseOrders, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/PurchaseOrders", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &PurchaseOrders{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty PurchaseOrders entity
func (s *PurchaseOrdersEndpoint) New() *PurchaseOrders {
	return &PurchaseOrders{}
}

// Create the PurchaseOrders entity in the provided division.
func (s *PurchaseOrdersEndpoint) Create(ctx context.Context, division int, entity *PurchaseOrders) (*PurchaseOrders, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/PurchaseOrders", division) // #nosec
	e := &PurchaseOrders{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the PurchaseOrders entity in the provided division.
func (s *PurchaseOrdersEndpoint) Update(ctx context.Context, division int, entity *PurchaseOrders) (*PurchaseOrders, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/PurchaseOrders", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &PurchaseOrders{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the PurchaseOrders entity in the provided division.
func (s *PurchaseOrdersEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/PurchaseOrders", division) // #nosec
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
