// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// LogisticsItemsService is responsible for communicating with
// the Items endpoint of the Logistics service.
type LogisticsItemsService service

// LogisticsItems:
// Service: Logistics
// Entity: Items
// URL: /api/v1/{division}/logistics/Items
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsItems
type LogisticsItems struct {
	// ID:
	ID *GUID `json:",omitempty"`

	// Barcode:
	Barcode *string `json:",omitempty"`

	// Class_01:
	Class_01 *string `json:",omitempty"`

	// Class_02:
	Class_02 *string `json:",omitempty"`

	// Class_03:
	Class_03 *string `json:",omitempty"`

	// Class_04:
	Class_04 *string `json:",omitempty"`

	// Class_05:
	Class_05 *string `json:",omitempty"`

	// Class_06:
	Class_06 *string `json:",omitempty"`

	// Class_07:
	Class_07 *string `json:",omitempty"`

	// Class_08:
	Class_08 *string `json:",omitempty"`

	// Class_09:
	Class_09 *string `json:",omitempty"`

	// Class_10:
	Class_10 *string `json:",omitempty"`

	// Code:
	Code *string `json:",omitempty"`

	// CopyRemarks:
	CopyRemarks *byte `json:",omitempty"`

	// CostPriceCurrency:
	CostPriceCurrency *string `json:",omitempty"`

	// CostPriceNew:
	CostPriceNew *float64 `json:",omitempty"`

	// CostPriceStandard:
	CostPriceStandard *float64 `json:",omitempty"`

	// Created:
	Created *Date `json:",omitempty"`

	// Creator:
	Creator *GUID `json:",omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:",omitempty"`

	// Description:
	Description *string `json:",omitempty"`

	// Division:
	Division *int `json:",omitempty"`

	// EndDate:
	EndDate *Date `json:",omitempty"`

	// ExtraDescription:
	ExtraDescription *string `json:",omitempty"`

	// FreeBoolField_01:
	FreeBoolField_01 *bool `json:",omitempty"`

	// FreeBoolField_02:
	FreeBoolField_02 *bool `json:",omitempty"`

	// FreeBoolField_03:
	FreeBoolField_03 *bool `json:",omitempty"`

	// FreeBoolField_04:
	FreeBoolField_04 *bool `json:",omitempty"`

	// FreeBoolField_05:
	FreeBoolField_05 *bool `json:",omitempty"`

	// FreeDateField_01:
	FreeDateField_01 *Date `json:",omitempty"`

	// FreeDateField_02:
	FreeDateField_02 *Date `json:",omitempty"`

	// FreeDateField_03:
	FreeDateField_03 *Date `json:",omitempty"`

	// FreeDateField_04:
	FreeDateField_04 *Date `json:",omitempty"`

	// FreeDateField_05:
	FreeDateField_05 *Date `json:",omitempty"`

	// FreeNumberField_01:
	FreeNumberField_01 *float64 `json:",omitempty"`

	// FreeNumberField_02:
	FreeNumberField_02 *float64 `json:",omitempty"`

	// FreeNumberField_03:
	FreeNumberField_03 *float64 `json:",omitempty"`

	// FreeNumberField_04:
	FreeNumberField_04 *float64 `json:",omitempty"`

	// FreeNumberField_05:
	FreeNumberField_05 *float64 `json:",omitempty"`

	// FreeNumberField_06:
	FreeNumberField_06 *float64 `json:",omitempty"`

	// FreeNumberField_07:
	FreeNumberField_07 *float64 `json:",omitempty"`

	// FreeNumberField_08:
	FreeNumberField_08 *float64 `json:",omitempty"`

	// FreeTextField_01:
	FreeTextField_01 *string `json:",omitempty"`

	// FreeTextField_02:
	FreeTextField_02 *string `json:",omitempty"`

	// FreeTextField_03:
	FreeTextField_03 *string `json:",omitempty"`

	// FreeTextField_04:
	FreeTextField_04 *string `json:",omitempty"`

	// FreeTextField_05:
	FreeTextField_05 *string `json:",omitempty"`

	// FreeTextField_06:
	FreeTextField_06 *string `json:",omitempty"`

	// FreeTextField_07:
	FreeTextField_07 *string `json:",omitempty"`

	// FreeTextField_08:
	FreeTextField_08 *string `json:",omitempty"`

	// FreeTextField_09:
	FreeTextField_09 *string `json:",omitempty"`

	// FreeTextField_10:
	FreeTextField_10 *string `json:",omitempty"`

	// GLCosts:
	GLCosts *GUID `json:",omitempty"`

	// GLCostsCode:
	GLCostsCode *string `json:",omitempty"`

	// GLCostsDescription:
	GLCostsDescription *string `json:",omitempty"`

	// GLRevenue:
	GLRevenue *GUID `json:",omitempty"`

	// GLRevenueCode:
	GLRevenueCode *string `json:",omitempty"`

	// GLRevenueDescription:
	GLRevenueDescription *string `json:",omitempty"`

	// GLStock:
	GLStock *GUID `json:",omitempty"`

	// GLStockCode:
	GLStockCode *string `json:",omitempty"`

	// GLStockDescription:
	GLStockDescription *string `json:",omitempty"`

	// GrossWeight:
	GrossWeight *float64 `json:",omitempty"`

	// IsBatchItem:
	IsBatchItem *byte `json:",omitempty"`

	// IsBatchNumberItem:
	IsBatchNumberItem *byte `json:",omitempty"`

	// IsFractionAllowedItem:
	IsFractionAllowedItem *bool `json:",omitempty"`

	// IsMakeItem:
	IsMakeItem *byte `json:",omitempty"`

	// IsNewContract:
	IsNewContract *byte `json:",omitempty"`

	// IsOnDemandItem:
	IsOnDemandItem *byte `json:",omitempty"`

	// IsPackageItem:
	IsPackageItem *bool `json:",omitempty"`

	// IsPurchaseItem:
	IsPurchaseItem *bool `json:",omitempty"`

	// IsRegistrationCodeItem:
	IsRegistrationCodeItem *byte `json:",omitempty"`

	// IsSalesItem:
	IsSalesItem *bool `json:",omitempty"`

	// IsSerialItem:
	IsSerialItem *bool `json:",omitempty"`

	// IsSerialNumberItem:
	IsSerialNumberItem *bool `json:",omitempty"`

	// IsStockItem:
	IsStockItem *bool `json:",omitempty"`

	// IsSubcontractedItem:
	IsSubcontractedItem *bool `json:",omitempty"`

	// IsTaxableItem:
	IsTaxableItem *byte `json:",omitempty"`

	// IsTime:
	IsTime *byte `json:",omitempty"`

	// IsWebshopItem:
	IsWebshopItem *byte `json:",omitempty"`

	// ItemGroup:
	ItemGroup *GUID `json:",omitempty"`

	// ItemGroupCode:
	ItemGroupCode *string `json:",omitempty"`

	// ItemGroupDescription:
	ItemGroupDescription *string `json:",omitempty"`

	// Modified:
	Modified *Date `json:",omitempty"`

	// Modifier:
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:",omitempty"`

	// NetWeight:
	NetWeight *float64 `json:",omitempty"`

	// NetWeightUnit:
	NetWeightUnit *string `json:",omitempty"`

	// Notes:
	Notes *string `json:",omitempty"`

	// PictureName:
	PictureName *string `json:",omitempty"`

	// PictureThumbnailUrl:
	PictureThumbnailUrl *string `json:",omitempty"`

	// PictureUrl:
	PictureUrl *string `json:",omitempty"`

	// SalesVatCode:
	SalesVatCode *string `json:",omitempty"`

	// SalesVatCodeDescription:
	SalesVatCodeDescription *string `json:",omitempty"`

	// SearchCode:
	SearchCode *string `json:",omitempty"`

	// SecurityLevel:
	SecurityLevel *int `json:",omitempty"`

	// StartDate:
	StartDate *Date `json:",omitempty"`

	// Stock:
	Stock *float64 `json:",omitempty"`

	// Unit:
	Unit *string `json:",omitempty"`

	// UnitDescription:
	UnitDescription *string `json:",omitempty"`

	// UnitType:
	UnitType *string `json:",omitempty"`
}

func (s *LogisticsItems) GetIdentifier() GUID {
	return *s.ID
}

// List the Items entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *LogisticsItemsService) List(ctx context.Context, division int, all bool) ([]*LogisticsItems, error) {
	var entities []*LogisticsItems
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/Items?$select=*", division)
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
