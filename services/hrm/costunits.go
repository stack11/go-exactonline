// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package hrm

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// CostunitsEndpoint is responsible for communicating with
// the Costunits endpoint of the HRM service.
type CostunitsEndpoint service

// Costunits:
// Service: HRM
// Entity: Costunits
// URL: /api/v1/{division}/hrm/Costunits
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMCostunits
type Costunits struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Code: Code (user-defined)
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description (text)
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EndDate: The end date by which the cost unit has to be inactive
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`
}

// List the Costunits entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CostunitsEndpoint) List(ctx context.Context, division int, all bool) ([]*Costunits, error) {
	var entities []*Costunits
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costunits?$select=*", division) // #nosec
	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
