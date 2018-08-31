// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// OpeningBalanceCurrentYearProcessedService is responsible for communicating with
// the CurrentYearProcessed endpoint of the OpeningBalance service.
type OpeningBalanceCurrentYearProcessedService service

// OpeningBalanceCurrentYearProcessed:
// Service: OpeningBalance
// Entity: CurrentYearProcessed
// URL: /api/v1/{division}/openingbalance/CurrentYear/Processed
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=OpeningBalanceCurrentYearProcessed
type OpeningBalanceCurrentYearProcessed struct {
	// Division: Division code.
	Division *int `json:",omitempty"`

	// GLAccount: The balance sheet account.
	GLAccount *GUID `json:",omitempty"`

	// ReportingYear: The reporting year of the opening balance.
	ReportingYear *int `json:",omitempty"`

	// Amount: The opening balance amount of the G/L account.
	Amount *float64 `json:",omitempty"`

	// BalanceSide: Indicates whether the G/L account is a debit or credit account. D = Debit, C = Credit.
	BalanceSide *string `json:",omitempty"`

	// GLAccountCode: The code of the G/L account.
	GLAccountCode *string `json:",omitempty"`

	// GLAccountDescription: The description of the G/L account.
	GLAccountDescription *string `json:",omitempty"`
}

func (s *OpeningBalanceCurrentYearProcessed) GetIdentifier() int {
	return *s.Division
}

// List the CurrentYearProcessed entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *OpeningBalanceCurrentYearProcessedService) List(ctx context.Context, division int, all bool) ([]*OpeningBalanceCurrentYearProcessed, error) {
	var entities []*OpeningBalanceCurrentYearProcessed
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/openingbalance/CurrentYear/Processed?$select=*", division)
	if err != nil {
		return nil, err
	}
	if all {
		err = s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err = s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
