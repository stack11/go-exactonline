// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// FinancialAgingReceivablesListService is responsible for communicating with
// the AgingReceivablesList endpoint of the Financial service.
type FinancialAgingReceivablesListService service

// FinancialAgingReceivablesList:
// Service: Financial
// Entity: AgingReceivablesList
// URL: /api/v1/{division}/read/financial/AgingReceivablesList
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialAgingReceivablesList
type FinancialAgingReceivablesList struct {
	// AccountId: Primary key
	AccountId *GUID `json:",omitempty"`

	// AccountCode: Code of Account
	AccountCode *string `json:",omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:",omitempty"`

	// AgeGroup1: Age group 1
	AgeGroup1 *int `json:",omitempty"`

	// AgeGroup1Amount: Amount of age group 1
	AgeGroup1Amount *float64 `json:",omitempty"`

	// AgeGroup1Description: Description of AgeGroup1
	AgeGroup1Description *string `json:",omitempty"`

	// AgeGroup2: Age group 2
	AgeGroup2 *int `json:",omitempty"`

	// AgeGroup2Amount: Amount of age group 2
	AgeGroup2Amount *float64 `json:",omitempty"`

	// AgeGroup2Description: Description of AgeGroup2
	AgeGroup2Description *string `json:",omitempty"`

	// AgeGroup3: Age group 3
	AgeGroup3 *int `json:",omitempty"`

	// AgeGroup3Amount: Amount of age group 3
	AgeGroup3Amount *float64 `json:",omitempty"`

	// AgeGroup3Description: Description of AgeGroup3
	AgeGroup3Description *string `json:",omitempty"`

	// AgeGroup4: Age group 4
	AgeGroup4 *int `json:",omitempty"`

	// AgeGroup4Amount: Amount of age group 4
	AgeGroup4Amount *float64 `json:",omitempty"`

	// AgeGroup4Description: Description of AgeGroup4
	AgeGroup4Description *string `json:",omitempty"`

	// CurrencyCode: Code of Currency
	CurrencyCode *string `json:",omitempty"`

	// TotalAmount: Total amount of all age groups
	TotalAmount *float64 `json:",omitempty"`
}

func (s *FinancialAgingReceivablesList) GetIdentifier() GUID {
	return *s.AccountId
}

// List the AgingReceivablesList entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *FinancialAgingReceivablesListService) List(ctx context.Context, division int, all bool) ([]*FinancialAgingReceivablesList, error) {
	var entities []*FinancialAgingReceivablesList
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/AgingReceivablesList?$select=*", division)
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
