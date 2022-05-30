// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// RecentCostsEndpoint is responsible for communicating with
// the RecentCosts endpoint of the Project service.
type RecentCostsEndpoint service

// RecentCosts:
// Service: Project
// Entity: RecentCosts
// URL: /api/v1/{division}/read/project/RecentCosts
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectRecentCosts
type RecentCosts struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Id: Primary key
	Id *int `json:"Id,omitempty"`

	// AccountCode: Code of account linked to the project that hours are being entered to
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountId: ID of account linked to the project that hours are being entered to
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountName: Name of account linked to the project that hours are being entered to
	AccountName *string `json:"AccountName,omitempty"`

	// AmountApproved: The total amount of (Quantity * Sales price of cost item) that has been approved
	AmountApproved *float64 `json:"AmountApproved,omitempty"`

	// AmountDraft: The total amount of (Quantity * Sales price of cost item) that has been saved as draft
	AmountDraft *float64 `json:"AmountDraft,omitempty"`

	// AmountRejected: The total amount of (Quantity * Sales price of cost item) that has been rejected
	AmountRejected *float64 `json:"AmountRejected,omitempty"`

	// AmountSubmitted: The total amount of (Quantity * Sales price of cost item) that has been submitted
	AmountSubmitted *float64 `json:"AmountSubmitted,omitempty"`

	// CurrencyCode: Code of sales currency which is used in the cost item
	CurrencyCode *string `json:"CurrencyCode,omitempty"`

	// Date: Date of entry
	Date *types.Date `json:"Date,omitempty"`

	// EntryId: Entry ID of record
	EntryId *types.GUID `json:"EntryId,omitempty"`

	// Expense: The ID of the Expense that is linked to the project
	Expense *types.GUID `json:"Expense,omitempty"`

	// ExpenseDescription: The description of the Expense that is linked to the project
	ExpenseDescription *string `json:"ExpenseDescription,omitempty"`

	// ItemCode: Code of the item used for cost entry
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of the item used for cost entry
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemId: ID of the item used for cost entry
	ItemId *types.GUID `json:"ItemId,omitempty"`

	// Notes: Notes entered regarding the information of the cost entered
	Notes *string `json:"Notes,omitempty"`

	// ProjectCode: Code of project that the costs are entered on
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Description of project that the costs are entered on
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectId: ID of project that the costs are entered on
	ProjectId *types.GUID `json:"ProjectId,omitempty"`

	// QuantityApproved: Quantity of items that is used in cost entry that are approved
	QuantityApproved *float64 `json:"QuantityApproved,omitempty"`

	// QuantityDraft: Quantity of items that is used in cost entry that are saved as draft
	QuantityDraft *float64 `json:"QuantityDraft,omitempty"`

	// QuantityRejected: Quantity of items that is used in cost entry that are rejected
	QuantityRejected *float64 `json:"QuantityRejected,omitempty"`

	// QuantitySubmitted: Quantity of items that is used in cost entry that are submitted
	QuantitySubmitted *float64 `json:"QuantitySubmitted,omitempty"`

	// WeekNumber: The week number that the cost entries have been entered on
	WeekNumber *int `json:"WeekNumber,omitempty"`
}

func (e *RecentCosts) GetPrimary() *int {
	return e.Id
}

func (s *RecentCostsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/RecentCosts", method)
}

// List the RecentCosts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *RecentCostsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*RecentCosts, error) {
	var entities []*RecentCosts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/RecentCosts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the RecentCosts entitiy in the provided division.
func (s *RecentCostsEndpoint) Get(ctx context.Context, division int, id *int) (*RecentCosts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/RecentCosts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &RecentCosts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
