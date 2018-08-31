// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// InventoryItemWarehouseStorageLocationsService is responsible for communicating with
// the ItemWarehouseStorageLocations endpoint of the Inventory service.
type InventoryItemWarehouseStorageLocationsService service

// InventoryItemWarehouseStorageLocations:
// Service: Inventory
// Entity: ItemWarehouseStorageLocations
// URL: /api/v1/{division}/inventory/ItemWarehouseStorageLocations
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryItemWarehouseStorageLocations
type InventoryItemWarehouseStorageLocations struct {
	// ID: Uniquely identifies the item, warehouse, storage location combination
	ID *GUID `json:",omitempty"`

	// IsFractionAllowedItem: Does the item allow partial quantities (1.75 meters)
	IsFractionAllowedItem *byte `json:",omitempty"`

	// Item: Item
	Item *GUID `json:",omitempty"`

	// ItemCode: Code of the item of this stock quantity
	ItemCode *string `json:",omitempty"`

	// ItemDescription: Description of the item of this stock quantity
	ItemDescription *string `json:",omitempty"`

	// ItemUnit: Unit of the item
	ItemUnit *string `json:",omitempty"`

	// ItemUnitDescription: Unit description of the item
	ItemUnitDescription *string `json:",omitempty"`

	// Stock: Number of items in stock
	Stock *float64 `json:",omitempty"`

	// StorageLocation: Storage location of this stock
	StorageLocation *GUID `json:",omitempty"`

	// StorageLocationCode: Code of the storage location of this stock quantity
	StorageLocationCode *string `json:",omitempty"`

	// StorageLocationDescription: Description of the storage location of this stock quantity
	StorageLocationDescription *string `json:",omitempty"`

	// Warehouse: ID of Warehouse
	Warehouse *GUID `json:",omitempty"`

	// WarehouseCode: Code of the warehouse of this stock quantity
	WarehouseCode *string `json:",omitempty"`

	// WarehouseDescription: Description of the warehouse of this stock quantity
	WarehouseDescription *string `json:",omitempty"`
}

func (s *InventoryItemWarehouseStorageLocations) GetIdentifier() GUID {
	return *s.ID
}

// List the ItemWarehouseStorageLocations entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InventoryItemWarehouseStorageLocationsService) List(ctx context.Context, division int, all bool) ([]*InventoryItemWarehouseStorageLocations, error) {
	var entities []*InventoryItemWarehouseStorageLocations
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/ItemWarehouseStorageLocations?$select=*", division)
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
