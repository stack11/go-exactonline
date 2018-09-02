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

// DivisionClassValuesEndpoint is responsible for communicating with
// the DivisionClassValues endpoint of the HRM service.
type DivisionClassValuesEndpoint service

// DivisionClassValues:
// Service: HRM
// Entity: DivisionClassValues
// URL: /api/v1/{division}/hrm/DivisionClassValues
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMDivisionClassValues
type DivisionClassValues struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Class_01: First classification
	Class_01 *[]byte `json:"Class_01,omitempty"`

	// Class_01_ID: First classification ID
	Class_01_ID *types.GUID `json:"Class_01_ID,omitempty"`

	// Class_02: Second classification
	Class_02 *[]byte `json:"Class_02,omitempty"`

	// Class_02_ID: Second classification ID
	Class_02_ID *types.GUID `json:"Class_02_ID,omitempty"`

	// Class_03: Third classification
	Class_03 *[]byte `json:"Class_03,omitempty"`

	// Class_03_ID: Third classification ID
	Class_03_ID *types.GUID `json:"Class_03_ID,omitempty"`

	// Class_04: Fourth classification
	Class_04 *[]byte `json:"Class_04,omitempty"`

	// Class_04_ID: Fourth classification ID
	Class_04_ID *types.GUID `json:"Class_04_ID,omitempty"`

	// Class_05: Fifth classification
	Class_05 *[]byte `json:"Class_05,omitempty"`

	// Class_05_ID: Fifth classification ID
	Class_05_ID *types.GUID `json:"Class_05_ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Customer: ID of customer
	Customer *types.GUID `json:"Customer,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`
}

// List the DivisionClassValues entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DivisionClassValuesEndpoint) List(ctx context.Context, division int, all bool) ([]*DivisionClassValues, error) {
	var entities []*DivisionClassValues
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/DivisionClassValues?$select=*", division) // #nosec
	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
