// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesorder

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// PlannedSalesReturnsEndpoint is responsible for communicating with
// the PlannedSalesReturns endpoint of the SalesOrder service.
type PlannedSalesReturnsEndpoint service

// PlannedSalesReturns:
// Service: SalesOrder
// Entity: PlannedSalesReturns
// URL: /api/v1/{division}/salesorder/PlannedSalesReturns
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderPlannedSalesReturns
type PlannedSalesReturns struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// PlannedSalesReturnID: Primary key
	PlannedSalesReturnID *types.GUID `json:"PlannedSalesReturnID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of the creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// DeliveredTo: Reference to the customer where item is delivered to/customer returning items
	DeliveredTo *types.GUID `json:"DeliveredTo,omitempty"`

	// DeliveredToContactPerson: Reference to the contact person of customer where item is delivered to/customer returning items
	DeliveredToContactPerson *types.GUID `json:"DeliveredToContactPerson,omitempty"`

	// DeliveredToContactPersonFullName: Name of contact person of delivered to customer
	DeliveredToContactPersonFullName *string `json:"DeliveredToContactPersonFullName,omitempty"`

	// DeliveredToName: Name of delivered to customer
	DeliveredToName *string `json:"DeliveredToName,omitempty"`

	// DeliveryAddress: Delivered to address
	DeliveryAddress *types.GUID `json:"DeliveryAddress,omitempty"`

	// Description: Description of the planned sales return
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Document: Document that is manually linked to the planned sales return
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentSubject: Subject of Document
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// PlannedSalesReturnLines: Collection of planned sales return lines
	PlannedSalesReturnLines *json.RawMessage `json:"PlannedSalesReturnLines,omitempty"`

	// Remarks: Remarks
	Remarks *string `json:"Remarks,omitempty"`

	// ReturnDate: Date of planned sales return
	ReturnDate *types.Date `json:"ReturnDate,omitempty"`

	// ReturnNumber: Human readable id of the planned sales return
	ReturnNumber *int `json:"ReturnNumber,omitempty"`

	// Source: Source of planned sales return entry: 1-Manual entry, 2-Web service
	Source *int `json:"Source,omitempty"`

	// Status: Planned sales return status: 20-Open, 30-Confirmed, 50-Processed
	Status *int `json:"Status,omitempty"`

	// Warehouse: ID of warehouse to receive the returning items
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of warehouse to receive the returning items
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of warehouse to receive the returning items
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

func (e *PlannedSalesReturns) GetPrimary() *types.GUID {
	return e.PlannedSalesReturnID
}

func (s *PlannedSalesReturnsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "salesorder/PlannedSalesReturns", method)
}

// List the PlannedSalesReturns entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PlannedSalesReturnsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*PlannedSalesReturns, error) {
	var entities []*PlannedSalesReturns
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturns", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the PlannedSalesReturns entitiy in the provided division.
func (s *PlannedSalesReturnsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*PlannedSalesReturns, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturns", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &PlannedSalesReturns{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty PlannedSalesReturns entity
func (s *PlannedSalesReturnsEndpoint) New() *PlannedSalesReturns {
	return &PlannedSalesReturns{}
}

// Create the PlannedSalesReturns entity in the provided division.
func (s *PlannedSalesReturnsEndpoint) Create(ctx context.Context, division int, entity *PlannedSalesReturns) (*PlannedSalesReturns, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturns", division) // #nosec
	e := &PlannedSalesReturns{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the PlannedSalesReturns entity in the provided division.
func (s *PlannedSalesReturnsEndpoint) Update(ctx context.Context, division int, entity *PlannedSalesReturns) (*PlannedSalesReturns, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturns", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &PlannedSalesReturns{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the PlannedSalesReturns entity in the provided division.
func (s *PlannedSalesReturnsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturns", division) // #nosec
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
