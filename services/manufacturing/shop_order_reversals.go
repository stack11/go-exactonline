// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// ShopOrderReversalsEndpoint is responsible for communicating with
// the ShopOrderReversals endpoint of the Manufacturing service.
type ShopOrderReversalsEndpoint service

// ShopOrderReversals:
// Service: Manufacturing
// Entity: ShopOrderReversals
// URL: /api/v1/{division}/manufacturing/ShopOrderReversals
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingShopOrderReversals
type ShopOrderReversals struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ReversalStockTransactionId: ID of stock transaction related to this ShopOrderReversal
	ReversalStockTransactionId *types.GUID `json:"ReversalStockTransactionId,omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *types.GUID `json:"CreatedBy,omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:"CreatedByFullName,omitempty"`

	// CreatedDate: Date of this reversal
	CreatedDate *types.Date `json:"CreatedDate,omitempty"`

	// IsBatch: Does the ShopOrderReversal&#39;s item use batch numbers
	IsBatch *byte `json:"IsBatch,omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the ShopOrderReversal&#39;s item
	IsFractionAllowedItem *byte `json:"IsFractionAllowedItem,omitempty"`

	// IsSerial: Does the ShopOrderReversal&#39;s item use serial numbers
	IsSerial *byte `json:"IsSerial,omitempty"`

	// Item: Item reversed
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of item reversed
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item reversed
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemPictureUrl: Picture url of shop order item
	ItemPictureUrl *string `json:"ItemPictureUrl,omitempty"`

	// Note: Notes associated with this reversal
	Note *string `json:"Note,omitempty"`

	// OriginalStockTransactionId: ID of the original stock transaction, which was reversed
	OriginalStockTransactionId *types.GUID `json:"OriginalStockTransactionId,omitempty"`

	// Quantity: Quantity reversed
	Quantity *float64 `json:"Quantity,omitempty"`

	// ShopOrder: Shop order being reversed
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// ShopOrderNumber: Number of shop order being reversed
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// StorageLocation: ID of storage location reversed from
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Code of storage location reversed from
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Description of storage location reversed from
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// TransactionDate: Effective date of this ShopOrderReversal
	TransactionDate *types.Date `json:"TransactionDate,omitempty"`

	// Unit: Unit of measurement abbreviation of item reversed
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription: Unit of measurement of item reversed
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// Warehouse: ID of warehouse reversed from
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of warehouse reversed from
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of warehouse reversed from
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (e *ShopOrderReversals) GetPrimary() *types.GUID {
	return e.ReversalStockTransactionId
}

func (s *ShopOrderReversalsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "manufacturing/ShopOrderReversals", method)
}

// List the ShopOrderReversals entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ShopOrderReversalsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ShopOrderReversals, error) {
	var entities []*ShopOrderReversals
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderReversals", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ShopOrderReversals entitiy in the provided division.
func (s *ShopOrderReversalsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*ShopOrderReversals, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderReversals", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ShopOrderReversals{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty ShopOrderReversals entity
func (s *ShopOrderReversalsEndpoint) New() *ShopOrderReversals {
	return &ShopOrderReversals{}
}

// Create the ShopOrderReversals entity in the provided division.
func (s *ShopOrderReversalsEndpoint) Create(ctx context.Context, division int, entity *ShopOrderReversals) (*ShopOrderReversals, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderReversals", division) // #nosec
	e := &ShopOrderReversals{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
