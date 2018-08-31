// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// LogisticsItemGroupsService is responsible for communicating with
// the ItemGroups endpoint of the Logistics service.
type LogisticsItemGroupsService service

// LogisticsItemGroups:
// Service: Logistics
// Entity: ItemGroups
// URL: /api/v1/{division}/logistics/ItemGroups
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsItemGroups
type LogisticsItemGroups struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Code: Code of the item group
	Code *string `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Description: Description of the item group
	Description *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// GLCosts: GL account on which the costs of items of this group will be booked
	GLCosts *GUID `json:",omitempty"`

	// GLCostsCode: Code of GLCosts
	GLCostsCode *string `json:",omitempty"`

	// GLCostsDescription: Description of GLCosts
	GLCostsDescription *string `json:",omitempty"`

	// GLPurchaseAccount: GL Purchase account for purchase invoicing according to (non-) perpetual inventory method
	GLPurchaseAccount *GUID `json:",omitempty"`

	// GLPurchaseAccountCode: Code of GLPurchase
	GLPurchaseAccountCode *string `json:",omitempty"`

	// GLPurchaseAccountDescription: Description of GLPurchaseAccount
	GLPurchaseAccountDescription *string `json:",omitempty"`

	// GLPurchasePriceDifference: GL account that will be used for the &#39;Standard cost price&#39; valuation method to balance the difference between purchase price and cost price
	GLPurchasePriceDifference *GUID `json:",omitempty"`

	// GLPurchasePriceDifferenceCode: Code of GLPurchasePriceDifference
	GLPurchasePriceDifferenceCode *string `json:",omitempty"`

	// GLPurchasePriceDifferenceDescr: Description of GLPurchasePriceDifference
	GLPurchasePriceDifferenceDescr *string `json:",omitempty"`

	// GLRevenue: GL account on which the revenue for items of this group will be booked
	GLRevenue *GUID `json:",omitempty"`

	// GLRevenueCode: Code of GLRevenue
	GLRevenueCode *string `json:",omitempty"`

	// GLRevenueDescription: Description of GLRevenue
	GLRevenueDescription *string `json:",omitempty"`

	// GLStock: GL account on which stock entries will be booked for items of this group
	GLStock *GUID `json:",omitempty"`

	// GLStockCode: Code of GLStock
	GLStockCode *string `json:",omitempty"`

	// GLStockDescription: Description of GLStock
	GLStockDescription *string `json:",omitempty"`

	// GLStockVariance: GL stock variance account for perpetual inventory
	GLStockVariance *GUID `json:",omitempty"`

	// GLStockVarianceCode: Code of GLStockVariance
	GLStockVarianceCode *string `json:",omitempty"`

	// GLStockVarianceDescription: Description of GLStockVariance
	GLStockVarianceDescription *string `json:",omitempty"`

	// IsDefault: Indicates if this is the default item group that will be assigned when a new item is created
	IsDefault *byte `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// Notes: Notes
	Notes *string `json:",omitempty"`
}

func (s *LogisticsItemGroups) GetIdentifier() GUID {
	return *s.ID
}

// List the ItemGroups entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *LogisticsItemGroupsService) List(ctx context.Context, division int, all bool) ([]*LogisticsItemGroups, error) {
	var entities []*LogisticsItemGroups
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/ItemGroups?$select=*", division)
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
