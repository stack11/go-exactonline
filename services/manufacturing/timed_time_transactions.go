// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
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

// TimedTimeTransactionsEndpoint is responsible for communicating with
// the TimedTimeTransactions endpoint of the Manufacturing service.
type TimedTimeTransactionsEndpoint service

// TimedTimeTransactions:
// Service: Manufacturing
// Entity: TimedTimeTransactions
// URL: /api/v1/{division}/manufacturing/TimedTimeTransactions
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingTimedTimeTransactions
type TimedTimeTransactions struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: ID of employee
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Name of employee
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EndTime: Time that operation was stopped
	EndTime *types.Date `json:"EndTime,omitempty"`

	// IsOperationFinished: Is the operation finished?
	IsOperationFinished *byte `json:"IsOperationFinished,omitempty"`

	// LaborHours: Adjustable labor hours
	LaborHours *float64 `json:"LaborHours,omitempty"`

	// MachineHours: Adjustable machine hours
	MachineHours *float64 `json:"MachineHours,omitempty"`

	// Modified: Modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes - viewable in data collection
	Notes *string `json:"Notes,omitempty"`

	// Operation: ID of operation
	Operation *types.GUID `json:"Operation,omitempty"`

	// OperationCode: Code of operation
	OperationCode *string `json:"OperationCode,omitempty"`

	// OperationDescription: Description of operation
	OperationDescription *string `json:"OperationDescription,omitempty"`

	// PercentComplete: Percentage of operation completed within time period
	PercentComplete *float64 `json:"PercentComplete,omitempty"`

	// ProducedQuantity: Quantity of make item produced within time period
	ProducedQuantity *float64 `json:"ProducedQuantity,omitempty"`

	// ProductionArea: Production area of the work center
	ProductionArea *types.GUID `json:"ProductionArea,omitempty"`

	// ProductionAreaCode: Code of production area of the work center
	ProductionAreaCode *string `json:"ProductionAreaCode,omitempty"`

	// ProductionAreaDescription: Description of production area of the work center
	ProductionAreaDescription *string `json:"ProductionAreaDescription,omitempty"`

	// ShopOrder: ID of shop order
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// ShopOrderDescription: Description of shop order
	ShopOrderDescription *string `json:"ShopOrderDescription,omitempty"`

	// ShopOrderNumber: Number of shop order
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// ShopOrderRoutingStepPlan: Shop order routing step where work occurred
	ShopOrderRoutingStepPlan *types.GUID `json:"ShopOrderRoutingStepPlan,omitempty"`

	// ShopOrderRoutingStepPlanDescription: Description of the routing step plan
	ShopOrderRoutingStepPlanDescription *string `json:"ShopOrderRoutingStepPlanDescription,omitempty"`

	// ShopOrderRoutingStepPlanRemainingRunHours: Remaining planned run hours of the routing step plan
	ShopOrderRoutingStepPlanRemainingRunHours *float64 `json:"ShopOrderRoutingStepPlanRemainingRunHours,omitempty"`

	// ShopOrderRoutingStepPlanRemainingSetupHours: Remaining planned setup hours of the routing step plan
	ShopOrderRoutingStepPlanRemainingSetupHours *float64 `json:"ShopOrderRoutingStepPlanRemainingSetupHours,omitempty"`

	// Source: Source of the timed time transaction
	Source *int `json:"Source,omitempty"`

	// StartTime: Time that operation was started
	StartTime *types.Date `json:"StartTime,omitempty"`

	// Status: Status of the timed time transaction
	Status *int `json:"Status,omitempty"`

	// Type: Type of the timed time transaction: Setup = 10, Run = 20
	Type *int `json:"Type,omitempty"`

	// Workcenter: Work center where work occurred
	Workcenter *types.GUID `json:"Workcenter,omitempty"`

	// WorkcenterCode: Code of the work center
	WorkcenterCode *string `json:"WorkcenterCode,omitempty"`

	// WorkcenterDescription: Description of the work center
	WorkcenterDescription *string `json:"WorkcenterDescription,omitempty"`
}

func (e *TimedTimeTransactions) GetPrimary() *types.GUID {
	return e.ID
}

func (s *TimedTimeTransactionsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "manufacturing/TimedTimeTransactions", method)
}

// List the TimedTimeTransactions entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimedTimeTransactionsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*TimedTimeTransactions, error) {
	var entities []*TimedTimeTransactions
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/TimedTimeTransactions", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the TimedTimeTransactions entitiy in the provided division.
func (s *TimedTimeTransactionsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*TimedTimeTransactions, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/TimedTimeTransactions", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &TimedTimeTransactions{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty TimedTimeTransactions entity
func (s *TimedTimeTransactionsEndpoint) New() *TimedTimeTransactions {
	return &TimedTimeTransactions{}
}

// Create the TimedTimeTransactions entity in the provided division.
func (s *TimedTimeTransactionsEndpoint) Create(ctx context.Context, division int, entity *TimedTimeTransactions) (*TimedTimeTransactions, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/TimedTimeTransactions", division) // #nosec
	e := &TimedTimeTransactions{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the TimedTimeTransactions entity in the provided division.
func (s *TimedTimeTransactionsEndpoint) Update(ctx context.Context, division int, entity *TimedTimeTransactions) (*TimedTimeTransactions, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/TimedTimeTransactions", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &TimedTimeTransactions{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the TimedTimeTransactions entity in the provided division.
func (s *TimedTimeTransactionsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/TimedTimeTransactions", division) // #nosec
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
