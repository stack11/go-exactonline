// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package payroll

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// EmploymentTaxAuthoritiesGeneralEndpoint is responsible for communicating with
// the EmploymentTaxAuthoritiesGeneral endpoint of the Payroll service.
type EmploymentTaxAuthoritiesGeneralEndpoint service

// EmploymentTaxAuthoritiesGeneral:
// Service: Payroll
// Entity: EmploymentTaxAuthoritiesGeneral
// URL: /api/v1/{division}/payroll/EmploymentTaxAuthoritiesGeneral
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollEmploymentTaxAuthoritiesGeneral
type EmploymentTaxAuthoritiesGeneral struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: ID of the account
	Account *types.GUID `json:"Account,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee ID
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Name of employee
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Employee number
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// Employment: Employment
	Employment *types.GUID `json:"Employment,omitempty"`

	// EmploymentHID: Employment number
	EmploymentHID *int `json:"EmploymentHID,omitempty"`

	// EndDate: End date of employment agencies
	EndDate *types.Date `json:"EndDate,omitempty"`

	// InfluenceInsuranceObligation: Influence insurance obligation
	InfluenceInsuranceObligation *string `json:"InfluenceInsuranceObligation,omitempty"`

	// InfluenceInsuranceObligationDescription: Influence insurance obligation description
	InfluenceInsuranceObligationDescription *string `json:"InfluenceInsuranceObligationDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// NatureOfWorkRelationship: Nature of work relationship
	NatureOfWorkRelationship *string `json:"NatureOfWorkRelationship,omitempty"`

	// NatureOfWorkRelationshipDescription: Nature of work relationship description
	NatureOfWorkRelationshipDescription *string `json:"NatureOfWorkRelationshipDescription,omitempty"`

	// PayrollTaxesNumber: Payroll taxes number
	PayrollTaxesNumber *string `json:"PayrollTaxesNumber,omitempty"`

	// StartDate: Start date of employment agencies
	StartDate *types.Date `json:"StartDate,omitempty"`

	// TypeOfIncome: Type of income
	TypeOfIncome *string `json:"TypeOfIncome,omitempty"`

	// TypeOfIncomeDescription: Type of income description
	TypeOfIncomeDescription *string `json:"TypeOfIncomeDescription,omitempty"`
}

func (e *EmploymentTaxAuthoritiesGeneral) GetPrimary() *types.GUID {
	return e.ID
}

func (s *EmploymentTaxAuthoritiesGeneralEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "payroll/EmploymentTaxAuthoritiesGeneral", method)
}

// List the EmploymentTaxAuthoritiesGeneral entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *EmploymentTaxAuthoritiesGeneralEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*EmploymentTaxAuthoritiesGeneral, error) {
	var entities []*EmploymentTaxAuthoritiesGeneral
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/EmploymentTaxAuthoritiesGeneral", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the EmploymentTaxAuthoritiesGeneral entitiy in the provided division.
func (s *EmploymentTaxAuthoritiesGeneralEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*EmploymentTaxAuthoritiesGeneral, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/EmploymentTaxAuthoritiesGeneral", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &EmploymentTaxAuthoritiesGeneral{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
