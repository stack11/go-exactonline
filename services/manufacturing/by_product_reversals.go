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

// ByProductReversalsEndpoint is responsible for communicating with
// the ByProductReversals endpoint of the Manufacturing service.
type ByProductReversalsEndpoint service

// ByProductReversals:
// Service: Manufacturing
// Entity: ByProductReversals
// URL: /api/v1/{division}/manufacturing/ByProductReversals
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingByProductReversals
type ByProductReversals struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ReversalStockTransactionId: ID of stock transaction related to this ByProductReversal
	ReversalStockTransactionId *types.GUID `json:"ReversalStockTransactionId,omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *types.GUID `json:"CreatedBy,omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:"CreatedByFullName,omitempty"`

	// CreatedDate: Date of this reversal
	CreatedDate *types.Date `json:"CreatedDate,omitempty"`

	// IsBackflush: Boolean indicating if this reversal was the result of shop order backflushing, processed during a ShopOrderReversal
	IsBackflush *bool `json:"IsBackflush,omitempty"`

	// IsBatch: Does the ByProductReversal&#39;s item use batch numbers
	IsBatch *byte `json:"IsBatch,omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the ByProductReversal&#39;s item
	IsFractionAllowedItem *byte `json:"IsFractionAllowedItem,omitempty"`

	// IsSerial: Does the ByProductReversal&#39;s item use serial numbers
	IsSerial *byte `json:"IsSerial,omitempty"`

	// Item: Item reversed
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of item reversed
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item reversed
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemPictureUrl: Picture url of by-product item
	ItemPictureUrl *string `json:"ItemPictureUrl,omitempty"`

	// Note: Notes associated with this reversal
	Note *string `json:"Note,omitempty"`

	// OriginalStockTransactionId: ID of the original stock transaction, which was reversed
	OriginalStockTransactionId *types.GUID `json:"OriginalStockTransactionId,omitempty"`

	// Quantity: Quantity reversed
	Quantity *float64 `json:"Quantity,omitempty"`

	// ShopOrder: Shop order being reversed to
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// ShopOrderMaterialPlan: ID of shop order material plan
	ShopOrderMaterialPlan *types.GUID `json:"ShopOrderMaterialPlan,omitempty"`

	// ShopOrderNumber: Number of shop order being reversed to
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// StorageLocation: ID of storage location reversed from
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Code of storage location reversed from
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Description of storage location reversed from
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// TransactionDate: Effective date of this ByProductReversal
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

func (e *ByProductReversals) GetPrimary() *types.GUID {
	return e.ReversalStockTransactionId
}

func (s *ByProductReversalsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "manufacturing/ByProductReversals", method)
}

// List the ByProductReversals entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ByProductReversalsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ByProductReversals, error) {
	var entities []*ByProductReversals
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReversals", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ByProductReversals entitiy in the provided division.
func (s *ByProductReversalsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*ByProductReversals, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReversals", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ByProductReversals{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty ByProductReversals entity
func (s *ByProductReversalsEndpoint) New() *ByProductReversals {
	return &ByProductReversals{}
}

// Create the ByProductReversals entity in the provided division.
func (s *ByProductReversalsEndpoint) Create(ctx context.Context, division int, entity *ByProductReversals) (*ByProductReversals, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReversals", division) // #nosec
	e := &ByProductReversals{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
