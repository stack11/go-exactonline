// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package accountancy

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// SolutionLinksEndpoint is responsible for communicating with
// the SolutionLinks endpoint of the Accountancy service.
type SolutionLinksEndpoint service

// SolutionLinks:
// Service: Accountancy
// Entity: SolutionLinks
// URL: /api/v1/{division}/accountancy/SolutionLinks
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancySolutionLinks
type SolutionLinks struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: ID of account to which solution is linked
	Account *types.GUID `json:"Account,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// Division: Accountant main division
	Division *int `json:"Division,omitempty"`

	// ExternalSolutionCode: If type is external predefined, represents ID of PracticeManagementExternalSolutions (mandatory for External solution)
	ExternalSolutionCode *int `json:"ExternalSolutionCode,omitempty"`

	// ExternalSolutionName: Name of the external solution
	ExternalSolutionName *string `json:"ExternalSolutionName,omitempty"`

	// ExternalSolutionUrl: Customer URl in external solution, like solution.com/id123 (mandatory for External and ExternalOther solution)
	ExternalSolutionUrl *string `json:"ExternalSolutionUrl,omitempty"`

	// InternalSolutionDivision: Division code of linked internal solution (mandatory for Internal solution)
	InternalSolutionDivision *int `json:"InternalSolutionDivision,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// Name: Name of the solution link
	Name *string `json:"Name,omitempty"`

	// OtherExternalSolutionName: Name of the custom external solution (mandatory for ExternalOther solution)
	OtherExternalSolutionName *string `json:"OtherExternalSolutionName,omitempty"`

	// SolutionType: Type of solution: 0 - Internal(EOL), 1 - External(Wellknown solution), 2 - ExternalOther
	SolutionType *int `json:"SolutionType,omitempty"`

	// Status: Link status: 0 - Active, 1 - Inactive, 2 -Archived
	Status *int `json:"Status,omitempty"`
}

func (e *SolutionLinks) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SolutionLinksEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "accountancy/SolutionLinks", method)
}

// List the SolutionLinks entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SolutionLinksEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SolutionLinks, error) {
	var entities []*SolutionLinks
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/SolutionLinks", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SolutionLinks entitiy in the provided division.
func (s *SolutionLinksEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SolutionLinks, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/SolutionLinks", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SolutionLinks{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty SolutionLinks entity
func (s *SolutionLinksEndpoint) New() *SolutionLinks {
	return &SolutionLinks{}
}

// Create the SolutionLinks entity in the provided division.
func (s *SolutionLinksEndpoint) Create(ctx context.Context, division int, entity *SolutionLinks) (*SolutionLinks, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/SolutionLinks", division) // #nosec
	e := &SolutionLinks{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the SolutionLinks entity in the provided division.
func (s *SolutionLinksEndpoint) Update(ctx context.Context, division int, entity *SolutionLinks) (*SolutionLinks, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/SolutionLinks", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &SolutionLinks{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the SolutionLinks entity in the provided division.
func (s *SolutionLinksEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/SolutionLinks", division) // #nosec
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
