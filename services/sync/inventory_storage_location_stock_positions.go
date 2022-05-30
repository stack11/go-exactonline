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

// InventoryStorageLocationStockPositionsEndpoint is responsible for communicating with
// the InventoryStorageLocationStockPositions endpoint of the Sync service.
type InventoryStorageLocationStockPositionsEndpoint service

// InventoryStorageLocationStockPositions:
// Service: Sync
// Entity: InventoryStorageLocationStockPositions
// URL: /api/v1/{division}/sync/Inventory/StorageLocationStockPositions
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncInventoryStorageLocationStockPositions
type InventoryStorageLocationStockPositions struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp: Timestamp
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Item: Item
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of the item
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of the item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Stock: Stock
	Stock *float64 `json:"Stock,omitempty"`

	// StorageLocation: Storage location
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Code of the storage location
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Description of the storage location
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// UnitCode: Code of the unit for the item
	UnitCode *string `json:"UnitCode,omitempty"`

	// UnitDescription: Description of the unit for the item
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// Warehouse: Warehouse
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of the warehouse
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of the warehouse
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (e *InventoryStorageLocationStockPositions) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *InventoryStorageLocationStockPositionsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "Inventory/StorageLocationStockPositions", method)
}

// List the InventoryStorageLocationStockPositions entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InventoryStorageLocationStockPositionsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*InventoryStorageLocationStockPositions, error) {
	var entities []*InventoryStorageLocationStockPositions
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Inventory/StorageLocationStockPositions", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the InventoryStorageLocationStockPositions entitiy in the provided division.
func (s *InventoryStorageLocationStockPositionsEndpoint) Get(ctx context.Context, division int, id *int64) (*InventoryStorageLocationStockPositions, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Inventory/StorageLocationStockPositions", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &InventoryStorageLocationStockPositions{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
