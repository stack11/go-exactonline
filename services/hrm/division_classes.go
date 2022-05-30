// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package hrm

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// DivisionClassesEndpoint is responsible for communicating with
// the DivisionClasses endpoint of the HRM service.
type DivisionClassesEndpoint service

// DivisionClasses:
// Service: HRM
// Entity: DivisionClasses
// URL: /api/v1/{division}/hrm/DivisionClasses
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMDivisionClasses
type DivisionClasses struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// ClassNameCustomer: Classification customer ID
	ClassNameCustomer *types.GUID `json:"ClassNameCustomer,omitempty"`

	// ClassNameDescription: Related classification description
	ClassNameDescription *string `json:"ClassNameDescription,omitempty"`

	// ClassNameID: Related classification ID
	ClassNameID *types.GUID `json:"ClassNameID,omitempty"`

	// Code: Property code
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Property description
	Description *string `json:"Description,omitempty"`

	// DescriptionTermID: Property description term ID
	DescriptionTermID *int `json:"DescriptionTermID,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// SequenceNr: Related classification sequence number
	SequenceNr *int `json:"SequenceNr,omitempty"`
}

func (e *DivisionClasses) GetPrimary() *types.GUID {
	return e.ID
}

func (s *DivisionClassesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "hrm/DivisionClasses", method)
}

// List the DivisionClasses entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DivisionClassesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*DivisionClasses, error) {
	var entities []*DivisionClasses
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/DivisionClasses", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the DivisionClasses entitiy in the provided division.
func (s *DivisionClassesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*DivisionClasses, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/DivisionClasses", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &DivisionClasses{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
