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

// LogisticsSalesItemPricesEndpoint is responsible for communicating with
// the LogisticsSalesItemPrices endpoint of the Sync service.
type LogisticsSalesItemPricesEndpoint service

// LogisticsSalesItemPrices:
// Service: Sync
// Entity: LogisticsSalesItemPrices
// URL: /api/v1/{division}/sync/Logistics/SalesItemPrices
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncLogisticsSalesItemPrices
type LogisticsSalesItemPrices struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp: Timestamp
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// Account: ID of the customer
	Account *types.GUID `json:"Account,omitempty"`

	// AccountName: Name of the customer account
	AccountName *string `json:"AccountName,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency: The currency of the price
	Currency *string `json:"Currency,omitempty"`

	// DefaultItemUnit: The default unit of the item
	DefaultItemUnit *string `json:"DefaultItemUnit,omitempty"`

	// DefaultItemUnitDescription: The description of the default item unit
	DefaultItemUnitDescription *string `json:"DefaultItemUnitDescription,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EndDate: Together with StartDate this determines whether the price is active
	EndDate *types.Date `json:"EndDate,omitempty"`

	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Item: Item ID
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of Item
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of Item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// NumberOfItemsPerUnit: This is the multiplication factor when going from default item unit to the unit of this price.For example if the default item unit is &#39;gram&#39; and the price unit is &#39;kilogram&#39; then the value of this property is 1000.
	NumberOfItemsPerUnit *float64 `json:"NumberOfItemsPerUnit,omitempty"`

	// Price: The actual price of this sales item
	Price *float64 `json:"Price,omitempty"`

	// Quantity: Minimum quantity to which the price is applicable
	Quantity *float64 `json:"Quantity,omitempty"`

	// StartDate: Together with EndDate this determines whether the price is active
	StartDate *types.Date `json:"StartDate,omitempty"`

	// Unit: The unit code of the price
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription: Description of the price unit
	UnitDescription *string `json:"UnitDescription,omitempty"`
}

func (e *LogisticsSalesItemPrices) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *LogisticsSalesItemPricesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "Logistics/SalesItemPrices", method)
}

// List the LogisticsSalesItemPrices entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *LogisticsSalesItemPricesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*LogisticsSalesItemPrices, error) {
	var entities []*LogisticsSalesItemPrices
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Logistics/SalesItemPrices", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the LogisticsSalesItemPrices entitiy in the provided division.
func (s *LogisticsSalesItemPricesEndpoint) Get(ctx context.Context, division int, id *int64) (*LogisticsSalesItemPrices, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Logistics/SalesItemPrices", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &LogisticsSalesItemPrices{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
