// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package assets

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// CommercialBuildingValuesEndpoint is responsible for communicating with
// the CommercialBuildingValues endpoint of the Assets service.
type CommercialBuildingValuesEndpoint service

// CommercialBuildingValues:
// Service: Assets
// Entity: CommercialBuildingValues
// URL: /api/v1/{division}/assets/CommercialBuildingValues
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AssetsCommercialBuildingValues
type CommercialBuildingValues struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Asset: Asset
	Asset *types.GUID `json:"Asset,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EndDate: End date of the date range during which this percentage is valid
	EndDate *types.Date `json:"EndDate,omitempty"`

	// LineNumber: Line number
	LineNumber *int `json:"LineNumber,omitempty"`

	// MinimumValue: Minimum Value
	MinimumValue *float64 `json:"MinimumValue,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// PropertyValue: Property Value
	PropertyValue *float64 `json:"PropertyValue,omitempty"`

	// PropertyValueOption: Property Value Option
	PropertyValueOption *int `json:"PropertyValueOption,omitempty"`

	// StartDate: Start date of the date range during which this percentage is valid
	StartDate *types.Date `json:"StartDate,omitempty"`
}

func (e *CommercialBuildingValues) GetPrimary() *types.GUID {
	return e.ID
}

func (s *CommercialBuildingValuesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "assets/CommercialBuildingValues", method)
}

// List the CommercialBuildingValues entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CommercialBuildingValuesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*CommercialBuildingValues, error) {
	var entities []*CommercialBuildingValues
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/assets/CommercialBuildingValues", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the CommercialBuildingValues entitiy in the provided division.
func (s *CommercialBuildingValuesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*CommercialBuildingValues, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/assets/CommercialBuildingValues", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &CommercialBuildingValues{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
