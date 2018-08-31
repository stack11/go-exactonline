// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ProjectProjectBudgetTypesService is responsible for communicating with
// the ProjectBudgetTypes endpoint of the Project service.
type ProjectProjectBudgetTypesService service

// ProjectProjectBudgetTypes:
// Service: Project
// Entity: ProjectBudgetTypes
// URL: /api/v1/{division}/project/ProjectBudgetTypes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjectBudgetTypes
type ProjectProjectBudgetTypes struct {
	// ID: Primary key
	ID *int `json:",omitempty"`

	// Description: Description
	Description *string `json:",omitempty"`
}

func (s *ProjectProjectBudgetTypes) GetIdentifier() int {
	return *s.ID
}

// List the ProjectBudgetTypes entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectProjectBudgetTypesService) List(ctx context.Context, division int, all bool) ([]*ProjectProjectBudgetTypes, error) {
	var entities []*ProjectProjectBudgetTypes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectBudgetTypes?$select=*", division)
	if err != nil {
		return nil, err
	}
	if all {
		err = s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err = s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
