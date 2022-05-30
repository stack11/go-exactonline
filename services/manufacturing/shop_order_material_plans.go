// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// ShopOrderMaterialPlansEndpoint is responsible for communicating with
// the ShopOrderMaterialPlans endpoint of the Manufacturing service.
type ShopOrderMaterialPlansEndpoint service

// ShopOrderMaterialPlans:
// Service: Manufacturing
// Entity: ShopOrderMaterialPlans
// URL: /api/v1/{division}/manufacturing/ShopOrderMaterialPlans
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingShopOrderMaterialPlans
type ShopOrderMaterialPlans struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Backflush: Indicates if this is a backflush step
	Backflush *byte `json:"Backflush,omitempty"`

	// CalculatorType: Calculator type
	CalculatorType *int `json:"CalculatorType,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the material
	Description *string `json:"Description,omitempty"`

	// DetailDrawing: Detail drawing reference
	DetailDrawing *string `json:"DetailDrawing,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Item: Reference to Items table
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item Code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of Item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemPictureUrl: URL of the material item&#39;s picture
	ItemPictureUrl *string `json:"ItemPictureUrl,omitempty"`

	// LineNumber: Line number
	LineNumber *int `json:"LineNumber,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Line notes
	Notes *string `json:"Notes,omitempty"`

	// PlannedAmountFC: Planned amount in the currency of the transaction
	PlannedAmountFC *float64 `json:"PlannedAmountFC,omitempty"`

	// PlannedDate: Date that the material is required.
	PlannedDate *types.Date `json:"PlannedDate,omitempty"`

	// PlannedPriceFC: Planned price of the material
	PlannedPriceFC *float64 `json:"PlannedPriceFC,omitempty"`

	// PlannedQuantity: Intended quantity
	PlannedQuantity *float64 `json:"PlannedQuantity,omitempty"`

	// PlannedQuantityFactor: Intended quantity unit factor
	PlannedQuantityFactor *float64 `json:"PlannedQuantityFactor,omitempty"`

	// ShopOrder: Reference to ShopOrders table
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// Status: Line status
	Status *int `json:"Status,omitempty"`

	// StatusDescription: Description of Status
	StatusDescription *string `json:"StatusDescription,omitempty"`

	// Type: Type
	Type *int `json:"Type,omitempty"`

	// Unit: Unit
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription: Unit description
	UnitDescription *string `json:"UnitDescription,omitempty"`
}

func (e *ShopOrderMaterialPlans) GetPrimary() *types.GUID {
	return e.ID
}

func (s *ShopOrderMaterialPlansEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "manufacturing/ShopOrderMaterialPlans", method)
}

// List the ShopOrderMaterialPlans entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ShopOrderMaterialPlansEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ShopOrderMaterialPlans, error) {
	var entities []*ShopOrderMaterialPlans
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ShopOrderMaterialPlans entitiy in the provided division.
func (s *ShopOrderMaterialPlansEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*ShopOrderMaterialPlans, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ShopOrderMaterialPlans{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty ShopOrderMaterialPlans entity
func (s *ShopOrderMaterialPlansEndpoint) New() *ShopOrderMaterialPlans {
	return &ShopOrderMaterialPlans{}
}

// Create the ShopOrderMaterialPlans entity in the provided division.
func (s *ShopOrderMaterialPlansEndpoint) Create(ctx context.Context, division int, entity *ShopOrderMaterialPlans) (*ShopOrderMaterialPlans, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", division) // #nosec
	e := &ShopOrderMaterialPlans{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the ShopOrderMaterialPlans entity in the provided division.
func (s *ShopOrderMaterialPlansEndpoint) Update(ctx context.Context, division int, entity *ShopOrderMaterialPlans) (*ShopOrderMaterialPlans, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &ShopOrderMaterialPlans{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the ShopOrderMaterialPlans entity in the provided division.
func (s *ShopOrderMaterialPlansEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", division) // #nosec
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
