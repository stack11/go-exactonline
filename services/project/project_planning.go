// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
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

// ProjectPlanningEndpoint is responsible for communicating with
// the ProjectPlanning endpoint of the Project service.
type ProjectPlanningEndpoint service

// ProjectPlanning:
// Service: Project
// Entity: ProjectPlanning
// URL: /api/v1/{division}/project/ProjectPlanning
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjectPlanning
type ProjectPlanning struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Account:
	Account *types.GUID `json:"Account,omitempty"`

	// AccountCode:
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// BGTStatus:
	BGTStatus *int `json:"BGTStatus,omitempty"`

	// CommunicationErrorStatus:
	CommunicationErrorStatus *int `json:"CommunicationErrorStatus,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Employee:
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeCode:
	EmployeeCode *string `json:"EmployeeCode,omitempty"`

	// EmployeeHID:
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// EndDate:
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Hours:
	Hours *float64 `json:"Hours,omitempty"`

	// HourType:
	HourType *types.GUID `json:"HourType,omitempty"`

	// HourTypeCode:
	HourTypeCode *string `json:"HourTypeCode,omitempty"`

	// HourTypeDescription:
	HourTypeDescription *string `json:"HourTypeDescription,omitempty"`

	// IsBrokenRecurrence:
	IsBrokenRecurrence *bool `json:"IsBrokenRecurrence,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// OverAllocate:
	OverAllocate *bool `json:"OverAllocate,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode:
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectPlanningRecurring:
	ProjectPlanningRecurring *types.GUID `json:"ProjectPlanningRecurring,omitempty"`

	// ProjectWBS:
	ProjectWBS *types.GUID `json:"ProjectWBS,omitempty"`

	// ProjectWBSDescription:
	ProjectWBSDescription *string `json:"ProjectWBSDescription,omitempty"`

	// StartDate:
	StartDate *types.Date `json:"StartDate,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`
}

func (e *ProjectPlanning) GetPrimary() *types.GUID {
	return e.ID
}

func (s *ProjectPlanningEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/ProjectPlanning", method)
}

// List the ProjectPlanning entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectPlanningEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*ProjectPlanning, error) {
	var entities []*ProjectPlanning
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectPlanning", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the ProjectPlanning entitiy in the provided division.
func (s *ProjectPlanningEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*ProjectPlanning, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectPlanning", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &ProjectPlanning{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty ProjectPlanning entity
func (s *ProjectPlanningEndpoint) New() *ProjectPlanning {
	return &ProjectPlanning{}
}

// Create the ProjectPlanning entity in the provided division.
func (s *ProjectPlanningEndpoint) Create(ctx context.Context, division int, entity *ProjectPlanning) (*ProjectPlanning, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectPlanning", division) // #nosec
	e := &ProjectPlanning{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the ProjectPlanning entity in the provided division.
func (s *ProjectPlanningEndpoint) Update(ctx context.Context, division int, entity *ProjectPlanning) (*ProjectPlanning, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectPlanning", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &ProjectPlanning{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the ProjectPlanning entity in the provided division.
func (s *ProjectPlanningEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectPlanning", division) // #nosec
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
