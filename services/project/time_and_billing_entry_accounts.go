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

// TimeAndBillingEntryAccountsEndpoint is responsible for communicating with
// the TimeAndBillingEntryAccounts endpoint of the Project service.
type TimeAndBillingEntryAccountsEndpoint service

// TimeAndBillingEntryAccounts:
// Service: Project
// Entity: TimeAndBillingEntryAccounts
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryAccounts
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryAccounts
type TimeAndBillingEntryAccounts struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// AccountId: ID of account used for entries
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountName: Name of account
	AccountName *string `json:"AccountName,omitempty"`
}

func (e *TimeAndBillingEntryAccounts) GetPrimary() *types.GUID {
	return e.AccountId
}

func (s *TimeAndBillingEntryAccountsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/TimeAndBillingEntryAccounts", method)
}

// List the TimeAndBillingEntryAccounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingEntryAccountsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*TimeAndBillingEntryAccounts, error) {
	var entities []*TimeAndBillingEntryAccounts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryAccounts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the TimeAndBillingEntryAccounts entitiy in the provided division.
func (s *TimeAndBillingEntryAccountsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*TimeAndBillingEntryAccounts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryAccounts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &TimeAndBillingEntryAccounts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
