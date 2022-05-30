// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package hrm

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
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
	MetaData *api.MetaData `json:"__metadata,omitempty"`
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

func (e *Costunits) GetPrimary() *types.GUID {
	return e.ID
}

func (s *CostunitsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "hrm/Costunits", method)
}

// List the Costunits entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CostunitsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Costunits, error) {
	var entities []*Costunits
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costunits", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Costunits entitiy in the provided division.
func (s *CostunitsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*Costunits, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costunits", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Costunits{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty Costunits entity
func (s *CostunitsEndpoint) New() *Costunits {
	return &Costunits{}
}

// Create the Costunits entity in the provided division.
func (s *CostunitsEndpoint) Create(ctx context.Context, division int, entity *Costunits) (*Costunits, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costunits", division) // #nosec
	e := &Costunits{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the Costunits entity in the provided division.
func (s *CostunitsEndpoint) Update(ctx context.Context, division int, entity *Costunits) (*Costunits, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costunits", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &Costunits{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the Costunits entity in the provided division.
func (s *CostunitsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costunits", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return err
	}

	_, r, requestError := s.client.NewRequestAndDo(ctx, "DELETE", u.String(), nil, nil)
	if requestError != nil {
		return requestError
	}

	if r.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(r.Body) // #nosec
		return fmt.Errorf("Failed with status %v and body %v", r.StatusCode, body)
	}

	return nil
}
