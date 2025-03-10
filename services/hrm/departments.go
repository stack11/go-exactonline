// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
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

// DepartmentsEndpoint is responsible for communicating with
// the Departments endpoint of the HRM service.
type DepartmentsEndpoint service

// Departments:
// Service: HRM
// Entity: Departments
// URL: /api/v1/{division}/hrm/Departments
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMDepartments
type Departments struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Code: Department Code
	Code *string `json:"Code,omitempty"`

	// Costcenter: Cost center Code
	Costcenter *string `json:"Costcenter,omitempty"`

	// CostcenterDescription: Cost center description
	CostcenterDescription *string `json:"CostcenterDescription,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Department description
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Explanation or extra information can be stored in the notes
	Notes *string `json:"Notes,omitempty"`
}

func (e *Departments) GetPrimary() *types.GUID {
	return e.ID
}

func (s *DepartmentsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "hrm/Departments", method)
}

// List the Departments entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DepartmentsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Departments, error) {
	var entities []*Departments
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Departments", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Departments entitiy in the provided division.
func (s *DepartmentsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*Departments, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Departments", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Departments{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
