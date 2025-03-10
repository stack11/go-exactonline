// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// TimeAndBillingEntryRecentProjectsEndpoint is responsible for communicating with
// the TimeAndBillingEntryRecentProjects endpoint of the Project service.
type TimeAndBillingEntryRecentProjectsEndpoint service

// TimeAndBillingEntryRecentProjects:
// Service: Project
// Entity: TimeAndBillingEntryRecentProjects
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryRecentProjects
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryRecentProjects
type TimeAndBillingEntryRecentProjects struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ProjectId: The Id of the project that hours entries are entered
	ProjectId *types.GUID `json:"ProjectId,omitempty"`

	// DateLastUsed: The datetime the hour entries have been entered on the project
	DateLastUsed *types.Date `json:"DateLastUsed,omitempty"`

	// ProjectCode: The code of the project that the hour entries have been entered on
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: The description of the project that the hour entries have been entered on
	ProjectDescription *string `json:"ProjectDescription,omitempty"`
}

func (e *TimeAndBillingEntryRecentProjects) GetPrimary() *types.GUID {
	return e.ProjectId
}

func (s *TimeAndBillingEntryRecentProjectsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/TimeAndBillingEntryRecentProjects", method)
}

// List the TimeAndBillingEntryRecentProjects entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingEntryRecentProjectsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*TimeAndBillingEntryRecentProjects, error) {
	var entities []*TimeAndBillingEntryRecentProjects
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentProjects", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the TimeAndBillingEntryRecentProjects entitiy in the provided division.
func (s *TimeAndBillingEntryRecentProjectsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*TimeAndBillingEntryRecentProjects, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentProjects", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &TimeAndBillingEntryRecentProjects{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
