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

// SalesOrderGoodsDeliveriesEndpoint is responsible for communicating with
// the SalesOrderGoodsDeliveries endpoint of the Sync service.
type SalesOrderGoodsDeliveriesEndpoint service

// SalesOrderGoodsDeliveries:
// Service: Sync
// Entity: SalesOrderGoodsDeliveries
// URL: /api/v1/{division}/sync/SalesOrder/GoodsDeliveries
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncSalesOrderGoodsDeliveries
type SalesOrderGoodsDeliveries struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp:
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// DeliveryAccount:
	DeliveryAccount *types.GUID `json:"DeliveryAccount,omitempty"`

	// DeliveryAccountCode:
	DeliveryAccountCode *string `json:"DeliveryAccountCode,omitempty"`

	// DeliveryAccountName:
	DeliveryAccountName *string `json:"DeliveryAccountName,omitempty"`

	// DeliveryAddress:
	DeliveryAddress *types.GUID `json:"DeliveryAddress,omitempty"`

	// DeliveryContact:
	DeliveryContact *types.GUID `json:"DeliveryContact,omitempty"`

	// DeliveryContactPersonFullName:
	DeliveryContactPersonFullName *string `json:"DeliveryContactPersonFullName,omitempty"`

	// DeliveryDate:
	DeliveryDate *types.Date `json:"DeliveryDate,omitempty"`

	// DeliveryNumber:
	DeliveryNumber *int `json:"DeliveryNumber,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// EntryID:
	EntryID *types.GUID `json:"EntryID,omitempty"`

	// EntryNumber:
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Remarks:
	Remarks *string `json:"Remarks,omitempty"`

	// ShippingMethod:
	ShippingMethod *types.GUID `json:"ShippingMethod,omitempty"`

	// ShippingMethodCode:
	ShippingMethodCode *string `json:"ShippingMethodCode,omitempty"`

	// ShippingMethodDescription:
	ShippingMethodDescription *string `json:"ShippingMethodDescription,omitempty"`

	// TrackingNumber:
	TrackingNumber *string `json:"TrackingNumber,omitempty"`

	// Warehouse:
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode:
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription:
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (e *SalesOrderGoodsDeliveries) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *SalesOrderGoodsDeliveriesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "SalesOrder/GoodsDeliveries", method)
}

// List the SalesOrderGoodsDeliveries entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesOrderGoodsDeliveriesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SalesOrderGoodsDeliveries, error) {
	var entities []*SalesOrderGoodsDeliveries
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/SalesOrder/GoodsDeliveries", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SalesOrderGoodsDeliveries entitiy in the provided division.
func (s *SalesOrderGoodsDeliveriesEndpoint) Get(ctx context.Context, division int, id *int64) (*SalesOrderGoodsDeliveries, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/SalesOrder/GoodsDeliveries", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SalesOrderGoodsDeliveries{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
