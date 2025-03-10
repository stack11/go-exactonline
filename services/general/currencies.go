// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package general

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// CurrenciesEndpoint is responsible for communicating with
// the Currencies endpoint of the General service.
type CurrenciesEndpoint service

// Currencies:
// Service: General
// Entity: Currencies
// URL: /api/v1/{division}/general/Currencies
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=GeneralCurrencies
type Currencies struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Code: Primary key
	Code *string `json:"Code,omitempty"`

	// AmountPrecision: Defines the number of decimals of the exchange rate used to calculate the amount in domestic currency from the amount in foreign currency
	AmountPrecision *float64 `json:"AmountPrecision,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Description: Description of the currency
	Description *string `json:"Description,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// PricePrecision: Defines the number of decimals used to calculate the item price in created invoices
	PricePrecision *float64 `json:"PricePrecision,omitempty"`
}

func (e *Currencies) GetPrimary() *string {
	return e.Code
}

func (s *CurrenciesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "general/Currencies", method)
}

// List the Currencies entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CurrenciesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Currencies, error) {
	var entities []*Currencies
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/general/Currencies", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Currencies entitiy in the provided division.
func (s *CurrenciesEndpoint) Get(ctx context.Context, division int, id *string) (*Currencies, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/general/Currencies", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Currencies{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
