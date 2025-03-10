// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package payroll

import "github.com/stack11/go-exactonline/api"

type service struct {
	client *api.Client
}

// PayrollService is responsible for communication with the Payroll
// endpoints of the Exact Online API.
type PayrollService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	ActiveEmployments               *ActiveEmploymentsEndpoint
	Employees                       *EmployeesEndpoint
	EmploymentConditionGroups       *EmploymentConditionGroupsEndpoint
	EmploymentContractFlexPhases    *EmploymentContractFlexPhasesEndpoint
	EmploymentContracts             *EmploymentContractsEndpoint
	EmploymentEndReasons            *EmploymentEndReasonsEndpoint
	EmploymentOrganizations         *EmploymentOrganizationsEndpoint
	Employments                     *EmploymentsEndpoint
	EmploymentSalaries              *EmploymentSalariesEndpoint
	EmploymentTaxAuthoritiesGeneral *EmploymentTaxAuthoritiesGeneralEndpoint
	TaxEmploymentEndFlexCodes       *TaxEmploymentEndFlexCodesEndpoint
}

// NewPayrollService creates a new initialized instance of the
// PayrollService.
func NewPayrollService(apiClient *api.Client) *PayrollService {
	s := &PayrollService{client: apiClient}

	s.common.client = apiClient

	s.ActiveEmployments = (*ActiveEmploymentsEndpoint)(&s.common)
	s.Employees = (*EmployeesEndpoint)(&s.common)
	s.EmploymentConditionGroups = (*EmploymentConditionGroupsEndpoint)(&s.common)
	s.EmploymentContractFlexPhases = (*EmploymentContractFlexPhasesEndpoint)(&s.common)
	s.EmploymentContracts = (*EmploymentContractsEndpoint)(&s.common)
	s.EmploymentEndReasons = (*EmploymentEndReasonsEndpoint)(&s.common)
	s.EmploymentOrganizations = (*EmploymentOrganizationsEndpoint)(&s.common)
	s.Employments = (*EmploymentsEndpoint)(&s.common)
	s.EmploymentSalaries = (*EmploymentSalariesEndpoint)(&s.common)
	s.EmploymentTaxAuthoritiesGeneral = (*EmploymentTaxAuthoritiesGeneralEndpoint)(&s.common)
	s.TaxEmploymentEndFlexCodes = (*TaxEmploymentEndFlexCodesEndpoint)(&s.common)

	return s
}
