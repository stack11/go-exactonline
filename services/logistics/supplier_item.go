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

// SupplierItemEndpoint is responsible for communicating with
// the SupplierItem endpoint of the Logistics service.
type SupplierItemEndpoint service

// SupplierItem:
// Service: Logistics
// Entity: SupplierItem
// URL: /api/v1/{division}/logistics/SupplierItem
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsSupplierItem
type SupplierItem struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// CopyRemarks: Copy purchase remarks to purchase lines
	CopyRemarks *byte `json:"CopyRemarks,omitempty"`

	// CountryOfOrigin: Country of origin code
	CountryOfOrigin *string `json:"CountryOfOrigin,omitempty"`

	// CountryOfOriginDescription: Description of country of origin
	CountryOfOriginDescription *string `json:"CountryOfOriginDescription,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency: Currency of item price
	Currency *string `json:"Currency,omitempty"`

	// CurrencyDescription: Description of currency of item price
	CurrencyDescription *string `json:"CurrencyDescription,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// DropShipment: Indicates that the supplier will deliver the item directly to customer. Values: 0 = No, 1 = Yes, 2 = Optional
	DropShipment *byte `json:"DropShipment,omitempty"`

	// EndDate: Together with StartDate this determines whether the price is active
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Item: Item ID
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of Item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemUnit: Item Unit
	ItemUnit *types.GUID `json:"ItemUnit,omitempty"`

	// ItemUnitCode: Item Unit Code
	ItemUnitCode *string `json:"ItemUnitCode,omitempty"`

	// ItemUnitDescription: Item Unit Description
	ItemUnitDescription *string `json:"ItemUnitDescription,omitempty"`

	// MainSupplier: Indicates this is a main supplier
	MainSupplier *bool `json:"MainSupplier,omitempty"`

	// MinimumQuantity: Minimum quantity of the item for purchase, only available for Wholesale &amp; Distribution (Professional and Premium only)
	MinimumQuantity *float64 `json:"MinimumQuantity,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`

	// PurchaseLeadTime: The number of days between placing an order with a supplier and receiving items from the supplier
	PurchaseLeadTime *int `json:"PurchaseLeadTime,omitempty"`

	// PurchaseLotSize: Lot size of the item for purchase, only available for Wholesale &amp; Distribution (Premium only)
	PurchaseLotSize *int `json:"PurchaseLotSize,omitempty"`

	// PurchasePrice: Purchase price. If neither active nor future price exists, it shows 0 when GET
	PurchasePrice *float64 `json:"PurchasePrice,omitempty"`

	// PurchaseUnit: Unit code
	PurchaseUnit *string `json:"PurchaseUnit,omitempty"`

	// PurchaseUnitDescription: Description of unit
	PurchaseUnitDescription *string `json:"PurchaseUnitDescription,omitempty"`

	// PurchaseUnitFactor: This is the multiplication factor when going from default item unit to the unit of this price
	PurchaseUnitFactor *float64 `json:"PurchaseUnitFactor,omitempty"`

	// PurchaseVATCode: VAT code
	PurchaseVATCode *string `json:"PurchaseVATCode,omitempty"`

	// PurchaseVATCodeDescription: Description of VAT
	PurchaseVATCodeDescription *string `json:"PurchaseVATCodeDescription,omitempty"`

	// StartDate: Together with EndDate this determines whether the price is active
	StartDate *types.Date `json:"StartDate,omitempty"`

	// Supplier: Supplier ID
	Supplier *types.GUID `json:"Supplier,omitempty"`

	// SupplierCode: Supplier code
	SupplierCode *string `json:"SupplierCode,omitempty"`

	// SupplierDescription: Description of supplier
	SupplierDescription *string `json:"SupplierDescription,omitempty"`

	// SupplierItemCode: Supplier&#39;s item code
	SupplierItemCode *string `json:"SupplierItemCode,omitempty"`
}

func (e *SupplierItem) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SupplierItemEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "logistics/SupplierItem", method)
}

// List the SupplierItem entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SupplierItemEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SupplierItem, error) {
	var entities []*SupplierItem
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SupplierItem", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SupplierItem entitiy in the provided division.
func (s *SupplierItemEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SupplierItem, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SupplierItem", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SupplierItem{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty SupplierItem entity
func (s *SupplierItemEndpoint) New() *SupplierItem {
	return &SupplierItem{}
}

// Create the SupplierItem entity in the provided division.
func (s *SupplierItemEndpoint) Create(ctx context.Context, division int, entity *SupplierItem) (*SupplierItem, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SupplierItem", division) // #nosec
	e := &SupplierItem{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the SupplierItem entity in the provided division.
func (s *SupplierItemEndpoint) Update(ctx context.Context, division int, entity *SupplierItem) (*SupplierItem, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SupplierItem", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &SupplierItem{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the SupplierItem entity in the provided division.
func (s *SupplierItemEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/SupplierItem", division) // #nosec
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
