// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
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

// TimeAndBillingEntryRecentAccountsEndpoint is responsible for communicating with
// the TimeAndBillingEntryRecentAccounts endpoint of the Project service.
type TimeAndBillingEntryRecentAccountsEndpoint service

// TimeAndBillingEntryRecentAccounts:
// Service: Project
// Entity: TimeAndBillingEntryRecentAccounts
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryRecentAccounts
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryRecentAccounts
type TimeAndBillingEntryRecentAccounts struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// AccountId: Primary key
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountName: Name of account
	AccountName *string `json:"AccountName,omitempty"`

	// DateLastUsed: Date last used
	DateLastUsed *types.Date `json:"DateLastUsed,omitempty"`
}

func (e *TimeAndBillingEntryRecentAccounts) GetPrimary() *types.GUID {
	return e.AccountId
}

func (s *TimeAndBillingEntryRecentAccountsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/TimeAndBillingEntryRecentAccounts", method)
}

// List the TimeAndBillingEntryRecentAccounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingEntryRecentAccountsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*TimeAndBillingEntryRecentAccounts, error) {
	var entities []*TimeAndBillingEntryRecentAccounts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentAccounts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the TimeAndBillingEntryRecentAccounts entitiy in the provided division.
func (s *TimeAndBillingEntryRecentAccountsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*TimeAndBillingEntryRecentAccounts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentAccounts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &TimeAndBillingEntryRecentAccounts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
