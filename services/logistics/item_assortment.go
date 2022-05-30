// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package logistics

import (
	"context"
	"encoding/json"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// ItemAssortmentEndpoint is responsible for communicating with
// the ItemAssortment endpoint of the Logistics service.
type ItemAssortmentEndpoint service

// ItemAssortment:
// Service: Logistics
// Entity: ItemAssortment
// URL: /api/v1/{division}/logistics/ItemAssortment
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsItemAssortment
type ItemAssortment struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: ID of ItemAssortment
	ID *types.GUID `json:"ID,omitempty"`

	// Code: Code of ItemAssortment
	Code *int `json:"Code,omitempty"`

	// Description: Description of ItemAssortment
	Description *string `json:"Description,omitempty"`

	// Division: Division
	Division *int `json:"Division,omitempty"`

	// Properties: Properties of this ItemAssortment
	Properties *json.RawMessage `json:"Properties,omitempty"`
}

func (e *ItemAssortment) GetPrimary() *types.GUID {
	return e.ID
}

func (s *ItemAssortmentEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "logistics/ItemAssortment", method)
}

// List the ItemAssortment entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ItemAssortmentEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ItemAssortment, error) {
	var entities []*ItemAssortment
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/ItemAssortment", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ItemAssortment entitiy in the provided division.
func (s *ItemAssortmentEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*ItemAssortment, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/ItemAssortment", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ItemAssortment{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
