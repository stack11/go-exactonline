// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package payroll

import (
	"context"

	"github.com/stack11/go-exactonline/api"
)

// EmploymentContractFlexPhasesEndpoint is responsible for communicating with
// the EmploymentContractFlexPhases endpoint of the Payroll service.
type EmploymentContractFlexPhasesEndpoint service

// EmploymentContractFlexPhases:
// Service: Payroll
// Entity: EmploymentContractFlexPhases
// URL: /api/v1/{division}/payroll/EmploymentContractFlexPhases
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollEmploymentContractFlexPhases
type EmploymentContractFlexPhases struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *int `json:"ID,omitempty"`

	// Description: Flexible employment contract phase description
	Description *string `json:"Description,omitempty"`
}

func (e *EmploymentContractFlexPhases) GetPrimary() *int {
	return e.ID
}

func (s *EmploymentContractFlexPhasesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "payroll/EmploymentContractFlexPhases", method)
}

// List the EmploymentContractFlexPhases entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *EmploymentContractFlexPhasesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*EmploymentContractFlexPhases, error) {
	var entities []*EmploymentContractFlexPhases
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/EmploymentContractFlexPhases", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the EmploymentContractFlexPhases entitiy in the provided division.
func (s *EmploymentContractFlexPhasesEndpoint) Get(ctx context.Context, division int, id *int) (*EmploymentContractFlexPhases, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/EmploymentContractFlexPhases", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &EmploymentContractFlexPhases{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
