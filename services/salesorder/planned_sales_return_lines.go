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

// PlannedSalesReturnLinesEndpoint is responsible for communicating with
// the PlannedSalesReturnLines endpoint of the SalesOrder service.
type PlannedSalesReturnLinesEndpoint service

// PlannedSalesReturnLines:
// Service: SalesOrder
// Entity: PlannedSalesReturnLines
// URL: /api/v1/{division}/salesorder/PlannedSalesReturnLines
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderPlannedSalesReturnLines
type PlannedSalesReturnLines struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// BatchNumbers: The collection of batch numbers that belong to the items included in this planned sales return
	BatchNumbers *json.RawMessage `json:"BatchNumbers,omitempty"`

	// CreateCredit: Option to redeliver to replace the goods or to create a credit note for the returned item: 0-Redelivery, 1-Credit Note
	CreateCredit *byte `json:"CreateCredit,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// GoodDeliveryLineID: Goods delivery line of the particular item
	GoodDeliveryLineID *types.GUID `json:"GoodDeliveryLineID,omitempty"`

	// Item: Item ID. This item must link to the provided warehouse for POST.
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of item
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// LineNumber: Line number
	LineNumber *int `json:"LineNumber,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`

	// PlannedReturnQuantity: Expected quantity to be returned
	PlannedReturnQuantity *float64 `json:"PlannedReturnQuantity,omitempty"`

	// PlannedSalesReturnID: Entry number of the planned sales return
	PlannedSalesReturnID *types.GUID `json:"PlannedSalesReturnID,omitempty"`

	// ReceivedQuantity: Actual quantity returned
	ReceivedQuantity *float64 `json:"ReceivedQuantity,omitempty"`

	// ReturnReasonCodeCode: Code of ReasonCode
	ReturnReasonCodeCode *string `json:"ReturnReasonCodeCode,omitempty"`

	// ReturnReasonCodeDescription: Description of ReasonCode
	ReturnReasonCodeDescription *string `json:"ReturnReasonCodeDescription,omitempty"`

	// ReturnReasonCodeID: Indicates the reason why the planned sales was returned
	ReturnReasonCodeID *types.GUID `json:"ReturnReasonCodeID,omitempty"`

	// SalesOrderLineID: Sales order line of the particular item
	SalesOrderLineID *types.GUID `json:"SalesOrderLineID,omitempty"`

	// SalesOrderNumber: Saler order of the particular item
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// SerialNumbers: The collection of serial numbers that belong to the items included in this planned sales return
	SerialNumbers *json.RawMessage `json:"SerialNumbers,omitempty"`

	// StockTransactionEntryID: Entry number of the stock transaction
	StockTransactionEntryID *types.GUID `json:"StockTransactionEntryID,omitempty"`

	// StorageLocation: Storage location
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Storage location code
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Storage location description
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// StorageLocationSequenceNumber: Sequence number of planned sales return (Premium Only)
	StorageLocationSequenceNumber *int `json:"StorageLocationSequenceNumber,omitempty"`

	// UnitCode: Code of item&#39;s sales unit
	UnitCode *string `json:"UnitCode,omitempty"`

	// UnitDescription: Description of item&#39;s sales unit
	UnitDescription *string `json:"UnitDescription,omitempty"`
}

func (e *PlannedSalesReturnLines) GetPrimary() *types.GUID {
	return e.ID
}

func (s *PlannedSalesReturnLinesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "salesorder/PlannedSalesReturnLines", method)
}

// List the PlannedSalesReturnLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PlannedSalesReturnLinesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*PlannedSalesReturnLines, error) {
	var entities []*PlannedSalesReturnLines
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturnLines", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the PlannedSalesReturnLines entitiy in the provided division.
func (s *PlannedSalesReturnLinesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*PlannedSalesReturnLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturnLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &PlannedSalesReturnLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty PlannedSalesReturnLines entity
func (s *PlannedSalesReturnLinesEndpoint) New() *PlannedSalesReturnLines {
	return &PlannedSalesReturnLines{}
}

// Create the PlannedSalesReturnLines entity in the provided division.
func (s *PlannedSalesReturnLinesEndpoint) Create(ctx context.Context, division int, entity *PlannedSalesReturnLines) (*PlannedSalesReturnLines, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturnLines", division) // #nosec
	e := &PlannedSalesReturnLines{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the PlannedSalesReturnLines entity in the provided division.
func (s *PlannedSalesReturnLinesEndpoint) Update(ctx context.Context, division int, entity *PlannedSalesReturnLines) (*PlannedSalesReturnLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturnLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &PlannedSalesReturnLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the PlannedSalesReturnLines entity in the provided division.
func (s *PlannedSalesReturnLinesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/PlannedSalesReturnLines", division) // #nosec
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
