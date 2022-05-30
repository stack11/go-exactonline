// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
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

// CostcentersEndpoint is responsible for communicating with
// the Costcenters endpoint of the HRM service.
type CostcentersEndpoint service

// Costcenters:
// Service: HRM
// Entity: Costcenters
// URL: /api/v1/{division}/hrm/Costcenters
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMCostcenters
type Costcenters struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Active: Indicates if the cost center is active: 0 = inactive 1 = active
	Active *bool `json:"Active,omitempty"`

	// Code: Code (user-defined ID)
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

	// EndDate: The end date by which the cost center has to be inactive
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`
}

func (e *Costcenters) GetPrimary() *types.GUID {
	return e.ID
}

func (s *CostcentersEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "hrm/Costcenters", method)
}

// List the Costcenters entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CostcentersEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Costcenters, error) {
	var entities []*Costcenters
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costcenters", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Costcenters entitiy in the provided division.
func (s *CostcentersEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*Costcenters, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costcenters", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Costcenters{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty Costcenters entity
func (s *CostcentersEndpoint) New() *Costcenters {
	return &Costcenters{}
}

// Create the Costcenters entity in the provided division.
func (s *CostcentersEndpoint) Create(ctx context.Context, division int, entity *Costcenters) (*Costcenters, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costcenters", division) // #nosec
	e := &Costcenters{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the Costcenters entity in the provided division.
func (s *CostcentersEndpoint) Update(ctx context.Context, division int, entity *Costcenters) (*Costcenters, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costcenters", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &Costcenters{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the Costcenters entity in the provided division.
func (s *CostcentersEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Costcenters", division) // #nosec
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
