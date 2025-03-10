// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package sales

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// SalesPriceListDetailsEndpoint is responsible for communicating with
// the SalesPriceListDetails endpoint of the Sales service.
type SalesPriceListDetailsEndpoint service

// SalesPriceListDetails:
// Service: Sales
// Entity: SalesPriceListDetails
// URL: /api/v1/{division}/sales/SalesPriceListDetails
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesSalesPriceListDetails
type SalesPriceListDetails struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: Customer account Id
	Account *types.GUID `json:"Account,omitempty"`

	// AccountName: Customer account name
	AccountName *string `json:"AccountName,omitempty"`

	// BasePrice: ID of the base price.  If base price = use the standard sales price, it shows null.  If base price = set sales price, it shows ID of the sales price within this volume discount.
	BasePrice *types.GUID `json:"BasePrice,omitempty"`

	// BasePriceAmount: Amount of the base price.  If base price = use the standard sales price, it shows the latest item sales price. If base price = set sales price, it shows the base price which defined in price list.
	BasePriceAmount *float64 `json:"BasePriceAmount,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency: Currency
	Currency *string `json:"Currency,omitempty"`

	// Discount: Discount
	Discount *float64 `json:"Discount,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EndDate: End date
	EndDate *types.Date `json:"EndDate,omitempty"`

	// EntryMethod: Indicates whether discount or the new price is leading : 1-Discount, 2-New price.  Scenario  1. When entry method is Discount and use base price, Discounted price = (1 - SalesPriceListDetails.Discount) * SalesPriceListDetails.BasePriceAmount  2. When entry method is Discount and use Item&#39;s standard sales price, Discounted price = (1 - SalesPriceListDetails.Discount) * SalesItemPrices.Price  3. When entry method is New price, Discounted price = SalesPriceListDetails.NewPrice
	EntryMethod *int `json:"EntryMethod,omitempty"`

	// Item: Item
	Item *types.GUID `json:"Item,omitempty"`

	// ItemDescription: Description of the item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemGroup: ItemGroup
	ItemGroup *types.GUID `json:"ItemGroup,omitempty"`

	// ItemUnit: Default sales unit of the item
	ItemUnit *string `json:"ItemUnit,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// NewPrice: New price after discount
	NewPrice *float64 `json:"NewPrice,omitempty"`

	// NumberOfItemsPerUnit: Number of the item per unit
	NumberOfItemsPerUnit *float64 `json:"NumberOfItemsPerUnit,omitempty"`

	// PriceListCode: Code of the PriceList
	PriceListCode *string `json:"PriceListCode,omitempty"`

	// PriceListId: Id of the PriceList
	PriceListId *types.GUID `json:"PriceListId,omitempty"`

	// Quantity: Quantity
	Quantity *float64 `json:"Quantity,omitempty"`

	// StartDate: Start date
	StartDate *types.Date `json:"StartDate,omitempty"`

	// Unit: Unit
	Unit *string `json:"Unit,omitempty"`
}

func (e *SalesPriceListDetails) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SalesPriceListDetailsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "sales/SalesPriceListDetails", method)
}

// List the SalesPriceListDetails entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesPriceListDetailsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SalesPriceListDetails, error) {
	var entities []*SalesPriceListDetails
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sales/SalesPriceListDetails", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SalesPriceListDetails entitiy in the provided division.
func (s *SalesPriceListDetailsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SalesPriceListDetails, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sales/SalesPriceListDetails", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SalesPriceListDetails{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
