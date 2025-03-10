// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package purchaseorder

import (
	"context"
	"encoding/json"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// GoodsReceiptsEndpoint is responsible for communicating with
// the GoodsReceipts endpoint of the PurchaseOrder service.
type GoodsReceiptsEndpoint service

// GoodsReceipts:
// Service: PurchaseOrder
// Entity: GoodsReceipts
// URL: /api/v1/{division}/purchaseorder/GoodsReceipts
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PurchaseOrderGoodsReceipts
type GoodsReceipts struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of the creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the goods receipt
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Document: Document that is linked to the goods receipt
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentSubject: Document subject
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// EntryNumber: Entry number of the resulting stock entry
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// GoodsReceiptLineCount: Total row count of lines
	GoodsReceiptLineCount *int `json:"GoodsReceiptLineCount,omitempty"`

	// GoodsReceiptLines: Collection of receipt lines
	GoodsReceiptLines *json.RawMessage `json:"GoodsReceiptLines,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of the last modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// ReceiptDate: Date of the goods receipt
	ReceiptDate *types.Date `json:"ReceiptDate,omitempty"`

	// ReceiptNumber: Receipt number
	ReceiptNumber *int `json:"ReceiptNumber,omitempty"`

	// Remarks: Receipt note
	Remarks *string `json:"Remarks,omitempty"`

	// Supplier: Account ID of the supplier
	Supplier *types.GUID `json:"Supplier,omitempty"`

	// SupplierCode: Supplier code
	SupplierCode *string `json:"SupplierCode,omitempty"`

	// SupplierContact: ID of the contact person at the supplier
	SupplierContact *types.GUID `json:"SupplierContact,omitempty"`

	// SupplierContactFullName: Name of the contact person at the supplier
	SupplierContactFullName *string `json:"SupplierContactFullName,omitempty"`

	// SupplierName: Supplier name
	SupplierName *string `json:"SupplierName,omitempty"`

	// Warehouse: Warehouse ID
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Warehouse code
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of the warehouse
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`

	// YourRef: The purchase invoice number provided by the supplier
	YourRef *string `json:"YourRef,omitempty"`
}

func (e *GoodsReceipts) GetPrimary() *types.GUID {
	return e.ID
}

func (s *GoodsReceiptsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "purchaseorder/GoodsReceipts", method)
}

// List the GoodsReceipts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *GoodsReceiptsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*GoodsReceipts, error) {
	var entities []*GoodsReceipts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/GoodsReceipts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the GoodsReceipts entitiy in the provided division.
func (s *GoodsReceiptsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*GoodsReceipts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/GoodsReceipts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &GoodsReceipts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty GoodsReceipts entity
func (s *GoodsReceiptsEndpoint) New() *GoodsReceipts {
	return &GoodsReceipts{}
}

// Create the GoodsReceipts entity in the provided division.
func (s *GoodsReceiptsEndpoint) Create(ctx context.Context, division int, entity *GoodsReceipts) (*GoodsReceipts, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/GoodsReceipts", division) // #nosec
	e := &GoodsReceipts{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the GoodsReceipts entity in the provided division.
func (s *GoodsReceiptsEndpoint) Update(ctx context.Context, division int, entity *GoodsReceipts) (*GoodsReceipts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/purchaseorder/GoodsReceipts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &GoodsReceipts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}
