// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ManufacturingProductionAreasService is responsible for communicating with
// the ProductionAreas endpoint of the Manufacturing service.
type ManufacturingProductionAreasService service

// ManufacturingProductionAreas:
// Service: Manufacturing
// Entity: ProductionAreas
// URL: /api/v1/{division}/manufacturing/ProductionAreas
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingProductionAreas
type ManufacturingProductionAreas struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Code: Code of the production area group
	Code *string `json:",omitempty"`

	// Costcenter: Reference to Cost center
	Costcenter *string `json:",omitempty"`

	// CostcenterDescription: Description of Costcenter
	CostcenterDescription *string `json:",omitempty"`

	// Costunit: Reference to Cost unit
	Costunit *string `json:",omitempty"`

	// CostunitDescription: Description of Costunit
	CostunitDescription *string `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Description: Description of the production area
	Description *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// IsDefault: Indicates if this is the default production area. This will be used when creating a new production area
	IsDefault *byte `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// Notes: Production area notes
	Notes *string `json:",omitempty"`
}

func (s *ManufacturingProductionAreas) GetIdentifier() GUID {
	return *s.ID
}

// List the ProductionAreas entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ManufacturingProductionAreasService) List(ctx context.Context, division int, all bool) ([]*ManufacturingProductionAreas, error) {
	var entities []*ManufacturingProductionAreas
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ProductionAreas?$select=*", division)
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
