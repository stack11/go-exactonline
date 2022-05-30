// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import "github.com/stack11/go-exactonline/api"

type service struct {
	client *api.Client
}

// ProjectService is responsible for communication with the Project
// endpoints of the Exact Online API.
type ProjectService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	CostTransactions                               *CostTransactionsEndpoint
	HourCostTypes                                  *HourCostTypesEndpoint
	InvoiceTerms                                   *InvoiceTermsEndpoint
	ProjectBudgetTypes                             *ProjectBudgetTypesEndpoint
	ProjectHourBudgets                             *ProjectHourBudgetsEndpoint
	ProjectPlanning                                *ProjectPlanningEndpoint
	ProjectPlanningRecurring                       *ProjectPlanningRecurringEndpoint
	ProjectRestrictionEmployees                    *ProjectRestrictionEmployeesEndpoint
	ProjectRestrictionItems                        *ProjectRestrictionItemsEndpoint
	ProjectRestrictionRebillings                   *ProjectRestrictionRebillingsEndpoint
	Projects                                       *ProjectsEndpoint
	RecentCosts                                    *RecentCostsEndpoint
	RecentHours                                    *RecentHoursEndpoint
	TimeAndBillingAccountDetails                   *TimeAndBillingAccountDetailsEndpoint
	TimeAndBillingActivitiesAndExpenses            *TimeAndBillingActivitiesAndExpensesEndpoint
	TimeAndBillingEntryAccounts                    *TimeAndBillingEntryAccountsEndpoint
	TimeAndBillingEntryProjects                    *TimeAndBillingEntryProjectsEndpoint
	TimeAndBillingEntryRecentAccounts              *TimeAndBillingEntryRecentAccountsEndpoint
	TimeAndBillingEntryRecentActivitiesAndExpenses *TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint
	TimeAndBillingEntryRecentHourCostTypes         *TimeAndBillingEntryRecentHourCostTypesEndpoint
	TimeAndBillingEntryRecentProjects              *TimeAndBillingEntryRecentProjectsEndpoint
	TimeAndBillingItemDetails                      *TimeAndBillingItemDetailsEndpoint
	TimeAndBillingProjectDetails                   *TimeAndBillingProjectDetailsEndpoint
	TimeCorrections                                *TimeCorrectionsEndpoint
	TimeTransactions                               *TimeTransactionsEndpoint
}

// NewProjectService creates a new initialized instance of the
// ProjectService.
func NewProjectService(apiClient *api.Client) *ProjectService {
	s := &ProjectService{client: apiClient}

	s.common.client = apiClient

	s.CostTransactions = (*CostTransactionsEndpoint)(&s.common)
	s.HourCostTypes = (*HourCostTypesEndpoint)(&s.common)
	s.InvoiceTerms = (*InvoiceTermsEndpoint)(&s.common)
	s.ProjectBudgetTypes = (*ProjectBudgetTypesEndpoint)(&s.common)
	s.ProjectHourBudgets = (*ProjectHourBudgetsEndpoint)(&s.common)
	s.ProjectPlanning = (*ProjectPlanningEndpoint)(&s.common)
	s.ProjectPlanningRecurring = (*ProjectPlanningRecurringEndpoint)(&s.common)
	s.ProjectRestrictionEmployees = (*ProjectRestrictionEmployeesEndpoint)(&s.common)
	s.ProjectRestrictionItems = (*ProjectRestrictionItemsEndpoint)(&s.common)
	s.ProjectRestrictionRebillings = (*ProjectRestrictionRebillingsEndpoint)(&s.common)
	s.Projects = (*ProjectsEndpoint)(&s.common)
	s.RecentCosts = (*RecentCostsEndpoint)(&s.common)
	s.RecentHours = (*RecentHoursEndpoint)(&s.common)
	s.TimeAndBillingAccountDetails = (*TimeAndBillingAccountDetailsEndpoint)(&s.common)
	s.TimeAndBillingActivitiesAndExpenses = (*TimeAndBillingActivitiesAndExpensesEndpoint)(&s.common)
	s.TimeAndBillingEntryAccounts = (*TimeAndBillingEntryAccountsEndpoint)(&s.common)
	s.TimeAndBillingEntryProjects = (*TimeAndBillingEntryProjectsEndpoint)(&s.common)
	s.TimeAndBillingEntryRecentAccounts = (*TimeAndBillingEntryRecentAccountsEndpoint)(&s.common)
	s.TimeAndBillingEntryRecentActivitiesAndExpenses = (*TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint)(&s.common)
	s.TimeAndBillingEntryRecentHourCostTypes = (*TimeAndBillingEntryRecentHourCostTypesEndpoint)(&s.common)
	s.TimeAndBillingEntryRecentProjects = (*TimeAndBillingEntryRecentProjectsEndpoint)(&s.common)
	s.TimeAndBillingItemDetails = (*TimeAndBillingItemDetailsEndpoint)(&s.common)
	s.TimeAndBillingProjectDetails = (*TimeAndBillingProjectDetailsEndpoint)(&s.common)
	s.TimeCorrections = (*TimeCorrectionsEndpoint)(&s.common)
	s.TimeTransactions = (*TimeTransactionsEndpoint)(&s.common)

	return s
}
