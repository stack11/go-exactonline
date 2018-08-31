// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// PayrollTaxEmploymentEndFlexCodesService is responsible for communicating with
// the TaxEmploymentEndFlexCodes endpoint of the Payroll service.
type PayrollTaxEmploymentEndFlexCodesService service

// PayrollTaxEmploymentEndFlexCodes:
// Service: Payroll
// Entity: TaxEmploymentEndFlexCodes
// URL: /api/v1/{division}/payroll/TaxEmploymentEndFlexCodes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollTaxEmploymentEndFlexCodes
type PayrollTaxEmploymentEndFlexCodes struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Code: Code of flexible employment contract phase
	Code *string `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Description: Description of flexible employment contract phase
	Description *string `json:",omitempty"`

	// EndDate: End date of flexible employment contract
	EndDate *Date `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// StartDate: Start date of flexible employment contract phase
	StartDate *Date `json:",omitempty"`
}

func (s *PayrollTaxEmploymentEndFlexCodes) GetIdentifier() GUID {
	return *s.ID
}

// List the TaxEmploymentEndFlexCodes entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PayrollTaxEmploymentEndFlexCodesService) List(ctx context.Context, division int, all bool) ([]*PayrollTaxEmploymentEndFlexCodes, error) {
	var entities []*PayrollTaxEmploymentEndFlexCodes
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/TaxEmploymentEndFlexCodes?$select=*", division)
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
