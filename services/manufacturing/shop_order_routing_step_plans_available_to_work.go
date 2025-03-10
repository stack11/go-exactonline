// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// ShopOrderRoutingStepPlansAvailableToWorkEndpoint is responsible for communicating with
// the ShopOrderRoutingStepPlansAvailableToWork endpoint of the Manufacturing service.
type ShopOrderRoutingStepPlansAvailableToWorkEndpoint service

// ShopOrderRoutingStepPlansAvailableToWork:
// Service: Manufacturing
// Entity: ShopOrderRoutingStepPlansAvailableToWork
// URL: /api/v1/{division}/read/manufacturing/ShopOrderRoutingStepPlansAvailableToWork
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadManufacturingShopOrderRoutingStepPlansAvailableToWork
type ShopOrderRoutingStepPlansAvailableToWork struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// RoutingStep: Routing Step ID
	RoutingStep *types.GUID `json:"RoutingStep,omitempty"`

	// CustomerCode: Customer code
	CustomerCode *string `json:"CustomerCode,omitempty"`

	// CustomerCount: Count of customers
	CustomerCount *int `json:"CustomerCount,omitempty"`

	// CustomerName: Customer name
	CustomerName *string `json:"CustomerName,omitempty"`

	// DataType: Type of data returned by query - for internal use
	DataType *int `json:"DataType,omitempty"`

	// DateAscendingOrder: Planned dates ascending order
	DateAscendingOrder *int `json:"DateAscendingOrder,omitempty"`

	// DateDescendingOrder: Planned dates descending order
	DateDescendingOrder *int `json:"DateDescendingOrder,omitempty"`

	// ExtraDescription: Extra description
	ExtraDescription *string `json:"ExtraDescription,omitempty"`

	// IsFractionAllowedItem: Is fraction allowed item
	IsFractionAllowedItem *bool `json:"IsFractionAllowedItem,omitempty"`

	// IsReleased: Is released
	IsReleased *bool `json:"IsReleased,omitempty"`

	// IsRunOperationFinished: Is run operation finished
	IsRunOperationFinished *bool `json:"IsRunOperationFinished,omitempty"`

	// IsSetupOperationFinished: Is setup operation finished
	IsSetupOperationFinished *bool `json:"IsSetupOperationFinished,omitempty"`

	// Item: Item
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemCodeAscendingOrder: Shop order item code ascending order
	ItemCodeAscendingOrder *int `json:"ItemCodeAscendingOrder,omitempty"`

	// ItemCodeDescendingOrder: Shop order item code descending order
	ItemCodeDescendingOrder *int `json:"ItemCodeDescendingOrder,omitempty"`

	// ItemDescription: Item description
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemVersion: Item version ID
	ItemVersion *types.GUID `json:"ItemVersion,omitempty"`

	// ItemVersionNotes: Item version notes
	ItemVersionNotes *string `json:"ItemVersionNotes,omitempty"`

	// LineNumber: Sequence
	LineNumber *int `json:"LineNumber,omitempty"`

	// Mode: Mode of priority
	Mode *int `json:"Mode,omitempty"`

	// Notes: Shop order notes
	Notes *string `json:"Notes,omitempty"`

	// Operation: Operation
	Operation *types.GUID `json:"Operation,omitempty"`

	// OperationCode: Operation code
	OperationCode *string `json:"OperationCode,omitempty"`

	// PictureThumbnailPath: PictureThumbnailPath
	PictureThumbnailPath *string `json:"PictureThumbnailPath,omitempty"`

	// PlannedDate: Planned date
	PlannedDate *types.Date `json:"PlannedDate,omitempty"`

	// PlannedQuantity: Planned quantity
	PlannedQuantity *float64 `json:"PlannedQuantity,omitempty"`

	// PlannedSetupHours: Planned setup hours
	PlannedSetupHours *float64 `json:"PlannedSetupHours,omitempty"`

	// Priority: Priority of the shop order
	Priority *int `json:"Priority,omitempty"`

	// PriorityDescendingOrder: Priority of the shop order
	PriorityDescendingOrder *int `json:"PriorityDescendingOrder,omitempty"`

	// Project: Shop order project
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode: Shop order project code
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Project description
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// QuantityCompleted: QuantityCompleted
	QuantityCompleted *float64 `json:"QuantityCompleted,omitempty"`

	// RoutingStepDescription: Routing step description
	RoutingStepDescription *string `json:"RoutingStepDescription,omitempty"`

	// RoutingStepRealizationNotes: RoutingStepRealizationNotes
	RoutingStepRealizationNotes *string `json:"RoutingStepRealizationNotes,omitempty"`

	// RoutingStepStatus: Routing step status
	RoutingStepStatus *int `json:"RoutingStepStatus,omitempty"`

	// RoutingStepStatusDescription: Routing step status description
	RoutingStepStatusDescription *string `json:"RoutingStepStatusDescription,omitempty"`

	// RunStartTime: Run start time
	RunStartTime *types.Date `json:"RunStartTime,omitempty"`

	// RunStatus: Run timed status
	RunStatus *int `json:"RunStatus,omitempty"`

	// RunTimedTimeTransaction: Run timed time transaction
	RunTimedTimeTransaction *types.GUID `json:"RunTimedTimeTransaction,omitempty"`

	// SalesOrderCount: Count of Sales order
	SalesOrderCount *int `json:"SalesOrderCount,omitempty"`

	// SalesOrderLineNumber: Sales order line number
	SalesOrderLineNumber *int `json:"SalesOrderLineNumber,omitempty"`

	// SalesOrderNumber: Sales order number
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// SetupPercentComplete: SetupPercentComplete
	SetupPercentComplete *float64 `json:"SetupPercentComplete,omitempty"`

	// SetupStartTime: Setup start time
	SetupStartTime *types.Date `json:"SetupStartTime,omitempty"`

	// SetupStatus: Setup timed status
	SetupStatus *int `json:"SetupStatus,omitempty"`

	// SetupTimedTimeTransaction: Setup timed time transaction
	SetupTimedTimeTransaction *types.GUID `json:"SetupTimedTimeTransaction,omitempty"`

	// ShopOrder: Shop order ID
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// ShopOrderDescription: Shop order description
	ShopOrderDescription *string `json:"ShopOrderDescription,omitempty"`

	// ShopOrderNumber: Shop order number
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// ShopOrderNumberAscendingOrder: Shop order number ascending order
	ShopOrderNumberAscendingOrder *int `json:"ShopOrderNumberAscendingOrder,omitempty"`

	// ShopOrderNumberDescendingOrder: Shop order number descending order
	ShopOrderNumberDescendingOrder *int `json:"ShopOrderNumberDescendingOrder,omitempty"`

	// ShopOrderStatus: Shop order status
	ShopOrderStatus *int `json:"ShopOrderStatus,omitempty"`

	// Unit: Description of unit
	Unit *string `json:"Unit,omitempty"`

	// Warehouse: ID of warehouse where shop order is finished
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// Workcenter: Workcenter
	Workcenter *types.GUID `json:"Workcenter,omitempty"`

	// WorkcenterCode: Workcenter code
	WorkcenterCode *string `json:"WorkcenterCode,omitempty"`
}

func (e *ShopOrderRoutingStepPlansAvailableToWork) GetPrimary() *types.GUID {
	return e.RoutingStep
}

func (s *ShopOrderRoutingStepPlansAvailableToWorkEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "manufacturing/ShopOrderRoutingStepPlansAvailableToWork", method)
}

// List the ShopOrderRoutingStepPlansAvailableToWork entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ShopOrderRoutingStepPlansAvailableToWorkEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ShopOrderRoutingStepPlansAvailableToWork, error) {
	var entities []*ShopOrderRoutingStepPlansAvailableToWork
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/manufacturing/ShopOrderRoutingStepPlansAvailableToWork", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ShopOrderRoutingStepPlansAvailableToWork entitiy in the provided division.
func (s *ShopOrderRoutingStepPlansAvailableToWorkEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*ShopOrderRoutingStepPlansAvailableToWork, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/manufacturing/ShopOrderRoutingStepPlansAvailableToWork", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ShopOrderRoutingStepPlansAvailableToWork{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
