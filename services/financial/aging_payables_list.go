// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// AgingPayablesListEndpoint is responsible for communicating with
// the AgingPayablesList endpoint of the Financial service.
type AgingPayablesListEndpoint service

// AgingPayablesList:
// Service: Financial
// Entity: AgingPayablesList
// URL: /api/v1/{division}/read/financial/AgingPayablesList
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialAgingPayablesList
type AgingPayablesList struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// AccountId: Primary key
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountCode: Code of Account
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:"AccountName,omitempty"`

	// AgeGroup1: Age group 1
	AgeGroup1 *int `json:"AgeGroup1,omitempty"`

	// AgeGroup1Amount: Amount of age group 1
	AgeGroup1Amount *float64 `json:"AgeGroup1Amount,omitempty"`

	// AgeGroup1Description: Description of AgeGroup1
	AgeGroup1Description *string `json:"AgeGroup1Description,omitempty"`

	// AgeGroup2: Age group 2
	AgeGroup2 *int `json:"AgeGroup2,omitempty"`

	// AgeGroup2Amount: Amount of age group 2
	AgeGroup2Amount *float64 `json:"AgeGroup2Amount,omitempty"`

	// AgeGroup2Description: Description of AgeGroup2
	AgeGroup2Description *string `json:"AgeGroup2Description,omitempty"`

	// AgeGroup3: Age group 3
	AgeGroup3 *int `json:"AgeGroup3,omitempty"`

	// AgeGroup3Amount: Amount of age group 3
	AgeGroup3Amount *float64 `json:"AgeGroup3Amount,omitempty"`

	// AgeGroup3Description: Description of AgeGroup3
	AgeGroup3Description *string `json:"AgeGroup3Description,omitempty"`

	// AgeGroup4: Age group 4
	AgeGroup4 *int `json:"AgeGroup4,omitempty"`

	// AgeGroup4Amount: Amount of age group 4
	AgeGroup4Amount *float64 `json:"AgeGroup4Amount,omitempty"`

	// AgeGroup4Description: Description of AgeGroup4
	AgeGroup4Description *string `json:"AgeGroup4Description,omitempty"`

	// CurrencyCode: Code of Currency
	CurrencyCode *string `json:"CurrencyCode,omitempty"`

	// TotalAmount: Total amount of all age groups
	TotalAmount *float64 `json:"TotalAmount,omitempty"`
}

func (e *AgingPayablesList) GetPrimary() *types.GUID {
	return e.AccountId
}

func (s *AgingPayablesListEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "financial/AgingPayablesList", method)
}

// List the AgingPayablesList entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AgingPayablesListEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*AgingPayablesList, error) {
	var entities []*AgingPayablesList
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/AgingPayablesList", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the AgingPayablesList entitiy in the provided division.
func (s *AgingPayablesListEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*AgingPayablesList, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/AgingPayablesList", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &AgingPayablesList{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
