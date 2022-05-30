// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package system

import (
	"context"

	"github.com/stack11/go-exactonline/api"
)

// AvailableFeaturesEndpoint is responsible for communicating with
// the AvailableFeatures endpoint of the System service.
type AvailableFeaturesEndpoint service

// AvailableFeatures:
// Service: System
// Entity: AvailableFeatures
// URL: /api/v1/{division}/system/AvailableFeatures
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SystemSystemAvailableFeatures
type AvailableFeatures struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: The ID of the feature.
	ID *int `json:"ID,omitempty"`

	// Description: The description of the feature.
	Description *string `json:"Description,omitempty"`
}

func (e *AvailableFeatures) GetPrimary() *int {
	return e.ID
}

func (s *AvailableFeaturesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "system/AvailableFeatures", method)
}

// List the AvailableFeatures entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AvailableFeaturesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*AvailableFeatures, error) {
	var entities []*AvailableFeatures
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/system/AvailableFeatures", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the AvailableFeatures entitiy in the provided division.
func (s *AvailableFeaturesEndpoint) Get(ctx context.Context, division int, id *int) (*AvailableFeatures, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/system/AvailableFeatures", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &AvailableFeatures{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
