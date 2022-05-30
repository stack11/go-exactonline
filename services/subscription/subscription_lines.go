// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package subscription

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// SubscriptionLinesEndpoint is responsible for communicating with
// the SubscriptionLines endpoint of the Subscription service.
type SubscriptionLinesEndpoint service

// SubscriptionLines:
// Service: Subscription
// Entity: SubscriptionLines
// URL: /api/v1/{division}/subscription/SubscriptionLines
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SubscriptionSubscriptionLines
type SubscriptionLines struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AmountDC: Amount in the default currency of the company
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC: Amount in the currency of the transaction
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// Costcenter: Cost center linked to the subscription line
	Costcenter *string `json:"Costcenter,omitempty"`

	// Costunit: Cost unit linked to the subscription line
	Costunit *string `json:"Costunit,omitempty"`

	// Description: Description of the subscription line
	Description *string `json:"Description,omitempty"`

	// Discount: Discount percentage of the subscription line
	Discount *float64 `json:"Discount,omitempty"`

	// Division: Code of division the subscription line is made
	Division *int `json:"Division,omitempty"`

	// EntryID: Entry ID referencing to the subscription
	EntryID *types.GUID `json:"EntryID,omitempty"`

	// FromDate: The date which the subscription line starts
	FromDate *types.Date `json:"FromDate,omitempty"`

	// Item: The item that is used by the subscription line for sales details. Reference to Item
	Item *types.GUID `json:"Item,omitempty"`

	// ItemDescription: Description of Item used by the subscription line
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// LineNumber: Line number of the subscription line per subscription
	LineNumber *int `json:"LineNumber,omitempty"`

	// LineType: Reference to LineType endpoint
	LineType *int `json:"LineType,omitempty"`

	// LineTypeDescription: Description of LineType
	LineTypeDescription *string `json:"LineTypeDescription,omitempty"`

	// Modified: Date and time when the subscription line has been modified
	Modified *types.Date `json:"Modified,omitempty"`

	// NetPrice: Net price in the currency of the transaction
	NetPrice *float64 `json:"NetPrice,omitempty"`

	// Notes: To add or retrieve additional information in the subscription line
	Notes *string `json:"Notes,omitempty"`

	// Quantity: Quantity of item used in the subscription line
	Quantity *float64 `json:"Quantity,omitempty"`

	// SubscriptionNumber: Subscription number of the subscription line per subscription
	SubscriptionNumber *int `json:"SubscriptionNumber,omitempty"`

	// ToDate: The date the subscription line ends
	ToDate *types.Date `json:"ToDate,omitempty"`

	// UnitCode: The code of the unit used in the subscription line. E.g: kg, meter
	UnitCode *string `json:"UnitCode,omitempty"`

	// UnitDescription: Description of Unit used in the subscription line
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// UnitPrice: Unit price in the currency of the transaction (price * unit factor)
	UnitPrice *float64 `json:"UnitPrice,omitempty"`

	// VATAmountFC: Vat Amount in the currency of the transaction
	VATAmountFC *float64 `json:"VATAmountFC,omitempty"`

	// VATCode: VAT code that is used in the subscription line
	VATCode *string `json:"VATCode,omitempty"`

	// VATCodeDescription: Description of VAT code that is used in the subscription line
	VATCodeDescription *string `json:"VATCodeDescription,omitempty"`
}

func (e *SubscriptionLines) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SubscriptionLinesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "subscription/SubscriptionLines", method)
}

// List the SubscriptionLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SubscriptionLinesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SubscriptionLines, error) {
	var entities []*SubscriptionLines
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionLines", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SubscriptionLines entitiy in the provided division.
func (s *SubscriptionLinesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SubscriptionLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SubscriptionLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty SubscriptionLines entity
func (s *SubscriptionLinesEndpoint) New() *SubscriptionLines {
	return &SubscriptionLines{}
}

// Create the SubscriptionLines entity in the provided division.
func (s *SubscriptionLinesEndpoint) Create(ctx context.Context, division int, entity *SubscriptionLines) (*SubscriptionLines, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionLines", division) // #nosec
	e := &SubscriptionLines{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the SubscriptionLines entity in the provided division.
func (s *SubscriptionLinesEndpoint) Update(ctx context.Context, division int, entity *SubscriptionLines) (*SubscriptionLines, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionLines", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &SubscriptionLines{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the SubscriptionLines entity in the provided division.
func (s *SubscriptionLinesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionLines", division) // #nosec
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
