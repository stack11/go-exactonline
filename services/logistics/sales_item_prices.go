// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package logistics

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// SalesItemPricesEndpoint is responsible for communicating with
// the SalesItemPrices endpoint of the Logistics service.
type SalesItemPricesEndpoint service

// SalesItemPrices:
// Service: Logistics
// Entity: SalesItemPrices
// URL: /api/v1/{division}/logistics/SalesItemPrices
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsSalesItemPrices
type SalesItemPrices struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

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

func (e *SalesItemPrices) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SalesItemPricesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "logistics/SalesItemPrices", method)
}

// List the SalesItemPrices entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesItemPricesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SalesItemPrices, error) {
	var entities []*SalesItemPrices
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SalesItemPrices", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SalesItemPrices entitiy in the provided division.
func (s *SalesItemPricesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SalesItemPrices, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SalesItemPrices", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SalesItemPrices{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty SalesItemPrices entity
func (s *SalesItemPricesEndpoint) New() *SalesItemPrices {
	return &SalesItemPrices{}
}

// Create the SalesItemPrices entity in the provided division.
func (s *SalesItemPricesEndpoint) Create(ctx context.Context, division int, entity *SalesItemPrices) (*SalesItemPrices, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SalesItemPrices", division) // #nosec
	e := &SalesItemPrices{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the SalesItemPrices entity in the provided division.
func (s *SalesItemPricesEndpoint) Update(ctx context.Context, division int, entity *SalesItemPrices) (*SalesItemPrices, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SalesItemPrices", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &SalesItemPrices{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the SalesItemPrices entity in the provided division.
func (s *SalesItemPricesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SalesItemPrices", division) // #nosec
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
