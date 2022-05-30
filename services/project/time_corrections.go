// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// TimeCorrectionsEndpoint is responsible for communicating with
// the TimeCorrections endpoint of the Project service.
type TimeCorrectionsEndpoint service

// TimeCorrections:
// Service: Project
// Entity: TimeCorrections
// URL: /api/v1/{division}/project/TimeCorrections
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectTimeCorrections
type TimeCorrections struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Id
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`

	// OriginalEntryId: Reference to the time entry that this corrects for
	OriginalEntryId *types.GUID `json:"OriginalEntryId,omitempty"`

	// Quantity: Quantity has to be negative value. E.g.: If original quantity is 10 and the correct quantity is 4, this quantity is -6
	Quantity *float64 `json:"Quantity,omitempty"`
}

func (e *TimeCorrections) GetPrimary() *types.GUID {
	return e.ID
}

func (s *TimeCorrectionsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/TimeCorrections", method)
}

// List the TimeCorrections entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeCorrectionsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*TimeCorrections, error) {
	var entities []*TimeCorrections
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeCorrections", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the TimeCorrections entitiy in the provided division.
func (s *TimeCorrectionsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*TimeCorrections, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeCorrections", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &TimeCorrections{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty TimeCorrections entity
func (s *TimeCorrectionsEndpoint) New() *TimeCorrections {
	return &TimeCorrections{}
}

// Create the TimeCorrections entity in the provided division.
func (s *TimeCorrectionsEndpoint) Create(ctx context.Context, division int, entity *TimeCorrections) (*TimeCorrections, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeCorrections", division) // #nosec
	e := &TimeCorrections{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the TimeCorrections entity in the provided division.
func (s *TimeCorrectionsEndpoint) Update(ctx context.Context, division int, entity *TimeCorrections) (*TimeCorrections, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeCorrections", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &TimeCorrections{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the TimeCorrections entity in the provided division.
func (s *TimeCorrectionsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/TimeCorrections", division) // #nosec
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
