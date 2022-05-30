// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// ShopOrdersEndpoint is responsible for communicating with
// the ShopOrders endpoint of the Manufacturing service.
type ShopOrdersEndpoint service

// ShopOrders:
// Service: Manufacturing
// Entity: ShopOrders
// URL: /api/v1/{division}/manufacturing/ShopOrders
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingShopOrders
type ShopOrders struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// CADDrawingURL: URL to CAD Drawing Specified on General tab shoporder
	CADDrawingURL *string `json:"CADDrawingURL,omitempty"`

	// Costcenter: The cost center linked to the shop order
	Costcenter *string `json:"Costcenter,omitempty"`

	// CostcenterDescription: Description of Costcenter
	CostcenterDescription *string `json:"CostcenterDescription,omitempty"`

	// Costunit: The cost unit linked to the shop order
	Costunit *string `json:"Costunit,omitempty"`

	// CostunitDescription: Description of Costunit
	CostunitDescription *string `json:"CostunitDescription,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the shop order
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EntryDate: Date on which the shop order was entered, does not need to be the same as syscreated date but is usually the same.
	EntryDate *types.Date `json:"EntryDate,omitempty"`

	// IsBatch: Indicates if the item created by this shoporder is a batch item or not
	IsBatch *byte `json:"IsBatch,omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the item created by this shoporder
	IsFractionAllowedItem *byte `json:"IsFractionAllowedItem,omitempty"`

	// IsInPlanning: Indicator that Shop order is in planning
	IsInPlanning *byte `json:"IsInPlanning,omitempty"`

	// IsOnHold: Indicator if the Shop order is on hold
	IsOnHold *byte `json:"IsOnHold,omitempty"`

	// IsReleased: Indicator that the Shop order has been released to production
	IsReleased *byte `json:"IsReleased,omitempty"`

	// IsSerial: Does the material plan&#39;s item use serial numbers
	IsSerial *byte `json:"IsSerial,omitempty"`

	// Item: Reference to the item created by this shoporder
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of the item created by this shop order
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of the item created by this shop order
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemPictureUrl: URL of the picture linked to the item created by this shop order
	ItemPictureUrl *string `json:"ItemPictureUrl,omitempty"`

	// ItemVersion: Reference to ItemVersion
	ItemVersion *types.GUID `json:"ItemVersion,omitempty"`

	// ItemVersionDescription: Description of Item Version
	ItemVersionDescription *string `json:"ItemVersionDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes - only viewed internally
	Notes *string `json:"Notes,omitempty"`

	// PlannedDate: Planned end date of this shop order
	PlannedDate *types.Date `json:"PlannedDate,omitempty"`

	// PlannedQuantity: Planned quantity
	PlannedQuantity *float64 `json:"PlannedQuantity,omitempty"`

	// PlannedStartDate: Planned start date of this shop order
	PlannedStartDate *types.Date `json:"PlannedStartDate,omitempty"`

	// ProducedQuantity: Quantity finished
	ProducedQuantity *float64 `json:"ProducedQuantity,omitempty"`

	// ProductionLeadDays: Production lead days
	ProductionLeadDays *int `json:"ProductionLeadDays,omitempty"`

	// Project: Reference to Project
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectDescription: Description of Project
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ReadyToShipQuantity: Quantity ready to ship
	ReadyToShipQuantity *float64 `json:"ReadyToShipQuantity,omitempty"`

	// SalesOrderLineCount: Number of sales orders linked to this shop order
	SalesOrderLineCount *int `json:"SalesOrderLineCount,omitempty"`

	// SalesOrderLines: Collection of Sales order lines
	SalesOrderLines *json.RawMessage `json:"SalesOrderLines,omitempty"`

	// SelectionCode: ID of selection code. Only supported by the Advanced and Premium editions for Wholesale &amp; Distribution and Manufacturing
	SelectionCode *types.GUID `json:"SelectionCode,omitempty"`

	// SelectionCodeCode: Code of Selection code
	SelectionCodeCode *string `json:"SelectionCodeCode,omitempty"`

	// SelectionCodeDescription: Description of Selection code
	SelectionCodeDescription *string `json:"SelectionCodeDescription,omitempty"`

	// ShopOrderByProductPlanBackflushCount: Number of shop order by-product plans, which are backflushed, for this shop order
	ShopOrderByProductPlanBackflushCount *int `json:"ShopOrderByProductPlanBackflushCount,omitempty"`

	// ShopOrderByProductPlanCount: Number of shop order by-product plans for this shop order
	ShopOrderByProductPlanCount *int `json:"ShopOrderByProductPlanCount,omitempty"`

	// ShopOrderMain: Shop order main
	ShopOrderMain *types.GUID `json:"ShopOrderMain,omitempty"`

	// ShopOrderMainNumber: Shop order main number
	ShopOrderMainNumber *int `json:"ShopOrderMainNumber,omitempty"`

	// ShopOrderMaterialPlanBackflushCount: Number of shop order material plans, which are backflushed, for this shop order
	ShopOrderMaterialPlanBackflushCount *int `json:"ShopOrderMaterialPlanBackflushCount,omitempty"`

	// ShopOrderMaterialPlanCount: Number of shop order material plans for this shop order
	ShopOrderMaterialPlanCount *int `json:"ShopOrderMaterialPlanCount,omitempty"`

	// ShopOrderMaterialPlans: Collection of Shop order Material plans
	ShopOrderMaterialPlans *json.RawMessage `json:"ShopOrderMaterialPlans,omitempty"`

	// ShopOrderNumber: Unique number to indentify the shop order
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// ShopOrderNumberString: Unique number to indentify the shop order (as a string to allow OData filtering, e.g. $filter=substringof(&#39;123&#39;,ShopOrderNumberString) eq true
	ShopOrderNumberString *string `json:"ShopOrderNumberString,omitempty"`

	// ShopOrderParent: Shop order parent
	ShopOrderParent *types.GUID `json:"ShopOrderParent,omitempty"`

	// ShopOrderParentNumber: Shop order parent number
	ShopOrderParentNumber *int `json:"ShopOrderParentNumber,omitempty"`

	// ShopOrderRoutingStepPlanCount: Number of shop order routing step plans for this shop order
	ShopOrderRoutingStepPlanCount *int `json:"ShopOrderRoutingStepPlanCount,omitempty"`

	// ShopOrderRoutingStepPlans: Collection of Shop order Routing step plans
	ShopOrderRoutingStepPlans *json.RawMessage `json:"ShopOrderRoutingStepPlans,omitempty"`

	// Status: Indicates the status of  the Shop Order: 10 Open, 20 In process, 30 Finished, 40 Completed
	Status *int `json:"Status,omitempty"`

	// SubShopOrderCount: The count of material lines of this shop order, which have been linked to a sub order
	SubShopOrderCount *int `json:"SubShopOrderCount,omitempty"`

	// Type: Type of shoporder: always 9040 Regular
	Type *int `json:"Type,omitempty"`

	// Unit: Unit of the item created by this shop order
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription: Unit description of the unit of the item created by this shop order
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// Warehouse: Reference to the Warehouse associated with the Shop order
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// YourRef: Your reference (of the customer)
	YourRef *string `json:"YourRef,omitempty"`
}

func (e *ShopOrders) GetPrimary() *types.GUID {
	return e.ID
}

func (s *ShopOrdersEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "manufacturing/ShopOrders", method)
}

// List the ShopOrders entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ShopOrdersEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ShopOrders, error) {
	var entities []*ShopOrders
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrders", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ShopOrders entitiy in the provided division.
func (s *ShopOrdersEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*ShopOrders, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrders", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ShopOrders{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty ShopOrders entity
func (s *ShopOrdersEndpoint) New() *ShopOrders {
	return &ShopOrders{}
}

// Create the ShopOrders entity in the provided division.
func (s *ShopOrdersEndpoint) Create(ctx context.Context, division int, entity *ShopOrders) (*ShopOrders, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrders", division) // #nosec
	e := &ShopOrders{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the ShopOrders entity in the provided division.
func (s *ShopOrdersEndpoint) Update(ctx context.Context, division int, entity *ShopOrders) (*ShopOrders, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrders", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &ShopOrders{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the ShopOrders entity in the provided division.
func (s *ShopOrdersEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrders", division) // #nosec
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
