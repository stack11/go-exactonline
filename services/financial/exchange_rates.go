// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// ExchangeRatesEndpoint is responsible for communicating with
// the ExchangeRates endpoint of the Financial service.
type ExchangeRatesEndpoint service

// ExchangeRates:
// Service: Financial
// Entity: ExchangeRates
// URL: /api/v1/{division}/financial/ExchangeRates
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialExchangeRates
type ExchangeRates struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Rate: The exchange rate is stored as 1 TARGET CURRENCY = x SOURCE CURRENCY
	Rate *float64 `json:"Rate,omitempty"`

	// SourceCurrency: The foreign currency
	SourceCurrency *string `json:"SourceCurrency,omitempty"`

	// SourceCurrencyDescription: Description of SourceCurrency
	SourceCurrencyDescription *string `json:"SourceCurrencyDescription,omitempty"`

	// StartDate: The date as of which the rate is valid. The rate is valid until a next rate is defined
	StartDate *types.Date `json:"StartDate,omitempty"`

	// TargetCurrency: The default currency of the division
	TargetCurrency *string `json:"TargetCurrency,omitempty"`

	// TargetCurrencyDescription: Description of TargetCurrency
	TargetCurrencyDescription *string `json:"TargetCurrencyDescription,omitempty"`
}

// List the ExchangeRates entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ExchangeRatesEndpoint) List(ctx context.Context, division int, all bool) ([]*ExchangeRates, error) {
	var entities []*ExchangeRates
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/financial/ExchangeRates?$select=*", division) // #nosec
	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
