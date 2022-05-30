// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package continuousmonitoring

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// IndicatorStatesEndpoint is responsible for communicating with
// the IndicatorStates endpoint of the ContinuousMonitoring service.
type IndicatorStatesEndpoint service

// IndicatorStates:
// Service: ContinuousMonitoring
// Entity: IndicatorStates
// URL: /api/v1/beta/{division}/continuousmonitoring/IndicatorStates
// HasWebhook: false
// IsInBeta: true
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ContinuousMonitoringIndicatorStates
type IndicatorStates struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Count: To store the number (e.g. 2) of occurrences of an indicator (e.g. Number of deviating entries: 2)
	Count *int `json:"Count,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Indicator: ID of Indicators
	Indicator *types.GUID `json:"Indicator,omitempty"`

	// IndicatorDescription: Indicator type description
	IndicatorDescription *string `json:"IndicatorDescription,omitempty"`

	// IndicatorType: Indicator type (1 = Balance G/L account per financial year, 2 = Usage of journals, 3 = Deviating amount entered, 4 = Liquidity, 5 = VAT Return deadline, 6 = Difference result in percentage, 7 = Different VAT code used)
	IndicatorType *int `json:"IndicatorType,omitempty"`

	// LastUpdated: Last update date
	LastUpdated *types.Date `json:"LastUpdated,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// ReportingYear: Financial year
	ReportingYear *int `json:"ReportingYear,omitempty"`

	// Status: Indicator status (1 = OK, 2 = Warning, 3 = Exception)
	Status *int `json:"Status,omitempty"`

	// Value: To store a value (e.g. -1234.56) that will be used by the indicators current situation (e.g. Lowest expected balance of liquid assets will be: -1,234.56)
	Value *float64 `json:"Value,omitempty"`
}

func (e *IndicatorStates) GetPrimary() *types.GUID {
	return e.ID
}

func (s *IndicatorStatesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "continuousmonitoring/IndicatorStates", method)
}

// List the IndicatorStates entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *IndicatorStatesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*IndicatorStates, error) {
	var entities []*IndicatorStates
	u, _ := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorStates", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the IndicatorStates entitiy in the provided division.
func (s *IndicatorStatesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*IndicatorStates, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorStates", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &IndicatorStates{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
