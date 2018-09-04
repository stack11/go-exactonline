// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package inventory

import (
	"context"
	"encoding/json"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// BatchNumbersEndpoint is responsible for communicating with
// the BatchNumbers endpoint of the Inventory service.
type BatchNumbersEndpoint service

// BatchNumbers:
// Service: Inventory
// Entity: BatchNumbers
// URL: /api/v1/{division}/inventory/BatchNumbers
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryBatchNumbers
type BatchNumbers struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AvailableQuantity: Available quantity of this batch number
	AvailableQuantity *float64 `json:"AvailableQuantity,omitempty"`

	// BatchNumber: Human readable batch number
	BatchNumber *string `json:"BatchNumber,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// ExpiryDate: Expiry date of effective period for batch number
	ExpiryDate *types.Date `json:"ExpiryDate,omitempty"`

	// IsBlocked: Boolean value indicating whether or not the batch number is blocked
	IsBlocked *byte `json:"IsBlocked,omitempty"`

	// Item: Item
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Remarks: Remarks
	Remarks *string `json:"Remarks,omitempty"`

	// StorageLocations: Total quantity available per location
	StorageLocations *json.RawMessage `json:"StorageLocations,omitempty"`

	// Warehouses: Total quantity available per warehouse
	Warehouses *json.RawMessage `json:"Warehouses,omitempty"`
}

// List the BatchNumbers entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *BatchNumbersEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*BatchNumbers, error) {
	var entities []*BatchNumbers
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/BatchNumbers", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
