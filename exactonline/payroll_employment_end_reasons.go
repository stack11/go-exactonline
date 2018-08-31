// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// PayrollEmploymentEndReasonsService is responsible for communicating with
// the EmploymentEndReasons endpoint of the Payroll service.
type PayrollEmploymentEndReasonsService service

// PayrollEmploymentEndReasons:
// Service: Payroll
// Entity: EmploymentEndReasons
// URL: /api/v1/{division}/payroll/EmploymentEndReasons
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollEmploymentEndReasons
type PayrollEmploymentEndReasons struct {
	// ID: Primary key
	ID *int `json:",omitempty"`

	// Description: Employment end reason description
	Description *string `json:",omitempty"`
}

func (s *PayrollEmploymentEndReasons) GetIdentifier() int {
	return *s.ID
}

// List the EmploymentEndReasons entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PayrollEmploymentEndReasonsService) List(ctx context.Context, division int, all bool) ([]*PayrollEmploymentEndReasons, error) {
	var entities []*PayrollEmploymentEndReasons
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/EmploymentEndReasons?$select=*", division)
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
