// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
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

// TimeAndBillingEntryProjectsEndpoint is responsible for communicating with
// the TimeAndBillingEntryProjects endpoint of the Project service.
type TimeAndBillingEntryProjectsEndpoint service

// TimeAndBillingEntryProjects:
// Service: Project
// Entity: TimeAndBillingEntryProjects
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryProjects
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryProjects
type TimeAndBillingEntryProjects struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ProjectId: Primary key
	ProjectId *types.GUID `json:"ProjectId,omitempty"`

	// ProjectCode: Code
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Description
	ProjectDescription *string `json:"ProjectDescription,omitempty"`
}

func (e *TimeAndBillingEntryProjects) GetPrimary() *types.GUID {
	return e.ProjectId
}

func (s *TimeAndBillingEntryProjectsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/TimeAndBillingEntryProjects", method)
}

// List the TimeAndBillingEntryProjects entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingEntryProjectsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*TimeAndBillingEntryProjects, error) {
	var entities []*TimeAndBillingEntryProjects
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryProjects", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the TimeAndBillingEntryProjects entitiy in the provided division.
func (s *TimeAndBillingEntryProjectsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*TimeAndBillingEntryProjects, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryProjects", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &TimeAndBillingEntryProjects{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
