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

// LogisticsItemsEndpoint is responsible for communicating with
// the LogisticsItems endpoint of the Sync service.
type LogisticsItemsEndpoint service

// LogisticsItems:
// Service: Sync
// Entity: LogisticsItems
// URL: /api/v1/{division}/sync/Logistics/Items
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncLogisticsItems
type LogisticsItems struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp:
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// AverageCost:
	AverageCost *float64 `json:"AverageCost,omitempty"`

	// Barcode:
	Barcode *string `json:"Barcode,omitempty"`

	// Class_01:
	Class_01 *string `json:"Class_01,omitempty"`

	// Class_02:
	Class_02 *string `json:"Class_02,omitempty"`

	// Class_03:
	Class_03 *string `json:"Class_03,omitempty"`

	// Class_04:
	Class_04 *string `json:"Class_04,omitempty"`

	// Class_05:
	Class_05 *string `json:"Class_05,omitempty"`

	// Class_06:
	Class_06 *string `json:"Class_06,omitempty"`

	// Class_07:
	Class_07 *string `json:"Class_07,omitempty"`

	// Class_08:
	Class_08 *string `json:"Class_08,omitempty"`

	// Class_09:
	Class_09 *string `json:"Class_09,omitempty"`

	// Class_10:
	Class_10 *string `json:"Class_10,omitempty"`

	// Code:
	Code *string `json:"Code,omitempty"`

	// CopyRemarks:
	CopyRemarks *byte `json:"CopyRemarks,omitempty"`

	// CostPriceCurrency:
	CostPriceCurrency *string `json:"CostPriceCurrency,omitempty"`

	// CostPriceNew:
	CostPriceNew *float64 `json:"CostPriceNew,omitempty"`

	// CostPriceStandard:
	CostPriceStandard *float64 `json:"CostPriceStandard,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// CustomField:
	CustomField *string `json:"CustomField,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// EndDate:
	EndDate *types.Date `json:"EndDate,omitempty"`

	// ExtraDescription:
	ExtraDescription *string `json:"ExtraDescription,omitempty"`

	// FreeBoolField_01:
	FreeBoolField_01 *bool `json:"FreeBoolField_01,omitempty"`

	// FreeBoolField_02:
	FreeBoolField_02 *bool `json:"FreeBoolField_02,omitempty"`

	// FreeBoolField_03:
	FreeBoolField_03 *bool `json:"FreeBoolField_03,omitempty"`

	// FreeBoolField_04:
	FreeBoolField_04 *bool `json:"FreeBoolField_04,omitempty"`

	// FreeBoolField_05:
	FreeBoolField_05 *bool `json:"FreeBoolField_05,omitempty"`

	// FreeDateField_01:
	FreeDateField_01 *types.Date `json:"FreeDateField_01,omitempty"`

	// FreeDateField_02:
	FreeDateField_02 *types.Date `json:"FreeDateField_02,omitempty"`

	// FreeDateField_03:
	FreeDateField_03 *types.Date `json:"FreeDateField_03,omitempty"`

	// FreeDateField_04:
	FreeDateField_04 *types.Date `json:"FreeDateField_04,omitempty"`

	// FreeDateField_05:
	FreeDateField_05 *types.Date `json:"FreeDateField_05,omitempty"`

	// FreeNumberField_01:
	FreeNumberField_01 *float64 `json:"FreeNumberField_01,omitempty"`

	// FreeNumberField_02:
	FreeNumberField_02 *float64 `json:"FreeNumberField_02,omitempty"`

	// FreeNumberField_03:
	FreeNumberField_03 *float64 `json:"FreeNumberField_03,omitempty"`

	// FreeNumberField_04:
	FreeNumberField_04 *float64 `json:"FreeNumberField_04,omitempty"`

	// FreeNumberField_05:
	FreeNumberField_05 *float64 `json:"FreeNumberField_05,omitempty"`

	// FreeNumberField_06:
	FreeNumberField_06 *float64 `json:"FreeNumberField_06,omitempty"`

	// FreeNumberField_07:
	FreeNumberField_07 *float64 `json:"FreeNumberField_07,omitempty"`

	// FreeNumberField_08:
	FreeNumberField_08 *float64 `json:"FreeNumberField_08,omitempty"`

	// FreeTextField_01:
	FreeTextField_01 *string `json:"FreeTextField_01,omitempty"`

	// FreeTextField_02:
	FreeTextField_02 *string `json:"FreeTextField_02,omitempty"`

	// FreeTextField_03:
	FreeTextField_03 *string `json:"FreeTextField_03,omitempty"`

	// FreeTextField_04:
	FreeTextField_04 *string `json:"FreeTextField_04,omitempty"`

	// FreeTextField_05:
	FreeTextField_05 *string `json:"FreeTextField_05,omitempty"`

	// FreeTextField_06:
	FreeTextField_06 *string `json:"FreeTextField_06,omitempty"`

	// FreeTextField_07:
	FreeTextField_07 *string `json:"FreeTextField_07,omitempty"`

	// FreeTextField_08:
	FreeTextField_08 *string `json:"FreeTextField_08,omitempty"`

	// FreeTextField_09:
	FreeTextField_09 *string `json:"FreeTextField_09,omitempty"`

	// FreeTextField_10:
	FreeTextField_10 *string `json:"FreeTextField_10,omitempty"`

	// GLCosts:
	GLCosts *types.GUID `json:"GLCosts,omitempty"`

	// GLCostsCode:
	GLCostsCode *string `json:"GLCostsCode,omitempty"`

	// GLCostsDescription:
	GLCostsDescription *string `json:"GLCostsDescription,omitempty"`

	// GLRevenue:
	GLRevenue *types.GUID `json:"GLRevenue,omitempty"`

	// GLRevenueCode:
	GLRevenueCode *string `json:"GLRevenueCode,omitempty"`

	// GLRevenueDescription:
	GLRevenueDescription *string `json:"GLRevenueDescription,omitempty"`

	// GLStock:
	GLStock *types.GUID `json:"GLStock,omitempty"`

	// GLStockCode:
	GLStockCode *string `json:"GLStockCode,omitempty"`

	// GLStockDescription:
	GLStockDescription *string `json:"GLStockDescription,omitempty"`

	// GrossWeight:
	GrossWeight *float64 `json:"GrossWeight,omitempty"`

	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// IsBatchItem:
	IsBatchItem *byte `json:"IsBatchItem,omitempty"`

	// IsFractionAllowedItem:
	IsFractionAllowedItem *bool `json:"IsFractionAllowedItem,omitempty"`

	// IsMakeItem:
	IsMakeItem *byte `json:"IsMakeItem,omitempty"`

	// IsNewContract:
	IsNewContract *byte `json:"IsNewContract,omitempty"`

	// IsOnDemandItem:
	IsOnDemandItem *byte `json:"IsOnDemandItem,omitempty"`

	// IsPackageItem:
	IsPackageItem *bool `json:"IsPackageItem,omitempty"`

	// IsPurchaseItem:
	IsPurchaseItem *bool `json:"IsPurchaseItem,omitempty"`

	// IsRegistrationCodeItem:
	IsRegistrationCodeItem *byte `json:"IsRegistrationCodeItem,omitempty"`

	// IsSalesItem:
	IsSalesItem *bool `json:"IsSalesItem,omitempty"`

	// IsSerialItem:
	IsSerialItem *bool `json:"IsSerialItem,omitempty"`

	// IsStockItem:
	IsStockItem *bool `json:"IsStockItem,omitempty"`

	// IsSubcontractedItem:
	IsSubcontractedItem *bool `json:"IsSubcontractedItem,omitempty"`

	// IsTaxableItem:
	IsTaxableItem *byte `json:"IsTaxableItem,omitempty"`

	// IsTime:
	IsTime *byte `json:"IsTime,omitempty"`

	// IsWebshopItem:
	IsWebshopItem *byte `json:"IsWebshopItem,omitempty"`

	// ItemGroup:
	ItemGroup *types.GUID `json:"ItemGroup,omitempty"`

	// ItemGroupCode:
	ItemGroupCode *string `json:"ItemGroupCode,omitempty"`

	// ItemGroupDescription:
	ItemGroupDescription *string `json:"ItemGroupDescription,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// NetWeight:
	NetWeight *float64 `json:"NetWeight,omitempty"`

	// NetWeightUnit:
	NetWeightUnit *string `json:"NetWeightUnit,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// PictureName:
	PictureName *string `json:"PictureName,omitempty"`

	// PictureThumbnailUrl:
	PictureThumbnailUrl *string `json:"PictureThumbnailUrl,omitempty"`

	// PictureUrl:
	PictureUrl *string `json:"PictureUrl,omitempty"`

	// SalesVatCode:
	SalesVatCode *string `json:"SalesVatCode,omitempty"`

	// SalesVatCodeDescription:
	SalesVatCodeDescription *string `json:"SalesVatCodeDescription,omitempty"`

	// SearchCode:
	SearchCode *string `json:"SearchCode,omitempty"`

	// SecurityLevel:
	SecurityLevel *int `json:"SecurityLevel,omitempty"`

	// StartDate:
	StartDate *types.Date `json:"StartDate,omitempty"`

	// StatisticalCode:
	StatisticalCode *string `json:"StatisticalCode,omitempty"`

	// StatisticalNetWeight:
	StatisticalNetWeight *float64 `json:"StatisticalNetWeight,omitempty"`

	// StatisticalUnits:
	StatisticalUnits *float64 `json:"StatisticalUnits,omitempty"`

	// StatisticalValue:
	StatisticalValue *float64 `json:"StatisticalValue,omitempty"`

	// Stock:
	Stock *float64 `json:"Stock,omitempty"`

	// Unit:
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription:
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// UnitType:
	UnitType *string `json:"UnitType,omitempty"`
}

func (e *LogisticsItems) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *LogisticsItemsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "Logistics/Items", method)
}

// List the LogisticsItems entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *LogisticsItemsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*LogisticsItems, error) {
	var entities []*LogisticsItems
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Logistics/Items", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the LogisticsItems entitiy in the provided division.
func (s *LogisticsItemsEndpoint) Get(ctx context.Context, division int, id *int64) (*LogisticsItems, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Logistics/Items", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &LogisticsItems{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
