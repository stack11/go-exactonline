// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package openingbalance

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// CurrentYearProcessedEndpoint is responsible for communicating with
// the CurrentYearProcessed endpoint of the OpeningBalance service.
type CurrentYearProcessedEndpoint service

// CurrentYearProcessed:
// Service: OpeningBalance
// Entity: CurrentYearProcessed
// URL: /api/v1/{division}/openingbalance/CurrentYear/Processed
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=OpeningBalanceCurrentYearProcessed
type CurrentYearProcessed struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Division: Division code.
	Division *int `json:"Division,omitempty"`

	// GLAccount: The balance sheet account.
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// ReportingYear: The reporting year of the opening balance.
	ReportingYear *int `json:"ReportingYear,omitempty"`

	// Amount: The opening balance amount of the G/L account.
	Amount *float64 `json:"Amount,omitempty"`

	// BalanceSide: Indicates whether the G/L account is a debit or credit account. D = Debit, C = Credit.
	BalanceSide *string `json:"BalanceSide,omitempty"`

	// GLAccountCode: The code of the G/L account.
	GLAccountCode *string `json:"GLAccountCode,omitempty"`

	// GLAccountDescription: The description of the G/L account.
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`
}

func (e *CurrentYearProcessed) GetPrimary() *int {
	return e.Division
}

func (s *CurrentYearProcessedEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "CurrentYear/Processed", method)
}

// List the CurrentYearProcessed entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CurrentYearProcessedEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*CurrentYearProcessed, error) {
	var entities []*CurrentYearProcessed
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/openingbalance/CurrentYear/Processed", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the CurrentYearProcessed entitiy in the provided division.
func (s *CurrentYearProcessedEndpoint) Get(ctx context.Context, division int, id *int) (*CurrentYearProcessed, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/openingbalance/CurrentYear/Processed", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &CurrentYearProcessed{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
