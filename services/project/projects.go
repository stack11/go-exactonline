// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// ProjectsEndpoint is responsible for communicating with
// the Projects endpoint of the Project service.
type ProjectsEndpoint service

// Projects:
// Service: Project
// Entity: Projects
// URL: /api/v1/{division}/project/Projects
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectProjects
type Projects struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: The account for this project
	Account *types.GUID `json:"Account,omitempty"`

	// AccountCode: Code of Account
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountContact: Contact person of Account
	AccountContact *types.GUID `json:"AccountContact,omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:"AccountName,omitempty"`

	// AllowAdditionalInvoicing: Indicates if additional invoice is allowed for project
	AllowAdditionalInvoicing *bool `json:"AllowAdditionalInvoicing,omitempty"`

	// AllowNonMemberEntry: Allow non member to create time or cost entry
	AllowNonMemberEntry *bool `json:"AllowNonMemberEntry,omitempty"`

	// BlockEntry: Block time and cost entries
	BlockEntry *bool `json:"BlockEntry,omitempty"`

	// BlockPurchasing: Block purchasing
	BlockPurchasing *bool `json:"BlockPurchasing,omitempty"`

	// BlockRebilling: Block rebilling
	BlockRebilling *bool `json:"BlockRebilling,omitempty"`

	// BudgetedAmount: Budgeted amount of sales in the default currency of the company
	BudgetedAmount *float64 `json:"BudgetedAmount,omitempty"`

	// BudgetedCosts: Budgeted amount of costs in the default currency of the company
	BudgetedCosts *float64 `json:"BudgetedCosts,omitempty"`

	// BudgetedHoursPerHourType: Collection of budgeted hours
	BudgetedHoursPerHourType *json.RawMessage `json:"BudgetedHoursPerHourType,omitempty"`

	// BudgetedRevenue: Budgeted amount of revenue in the default currency of the company
	BudgetedRevenue *float64 `json:"BudgetedRevenue,omitempty"`

	// BudgetOverrunHours: BudgetOverrunHours: 10-Allowed, 20-Not Allowed
	BudgetOverrunHours *byte `json:"BudgetOverrunHours,omitempty"`

	// BudgetType: Budget type
	BudgetType *int `json:"BudgetType,omitempty"`

	// BudgetTypeDescription: Budget type description
	BudgetTypeDescription *string `json:"BudgetTypeDescription,omitempty"`

	// Classification: Used only for PSA to link a project classification to the project
	Classification *types.GUID `json:"Classification,omitempty"`

	// ClassificationDescription: Description of Classification
	ClassificationDescription *string `json:"ClassificationDescription,omitempty"`

	// Code: Code
	Code *string `json:"Code,omitempty"`

	// CostsAmountFC: Used only for PSA to store the budgetted costs of a project (except for project type Campaign and Non-billable). Positive quantities only
	CostsAmountFC *float64 `json:"CostsAmountFC,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// CustomerPOnumber: Used only for PSA to store the customer&#39;s PO number
	CustomerPOnumber *string `json:"CustomerPOnumber,omitempty"`

	// CustomField: Custom field endpoint. Provided only for the Exact Online Premium users.
	CustomField *string `json:"CustomField,omitempty"`

	// Description: Description of the project
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// DivisionName: Name of Division
	DivisionName *string `json:"DivisionName,omitempty"`

	// EndDate: End date of the project. In combination with the start date the status is determined
	EndDate *types.Date `json:"EndDate,omitempty"`

	// FixedPriceItem: Item used for fixed price invoicing. To be defined per project. If empty the functionality relies on the setting
	FixedPriceItem *types.GUID `json:"FixedPriceItem,omitempty"`

	// FixedPriceItemDescription: Description of FixedPriceItem
	FixedPriceItemDescription *string `json:"FixedPriceItemDescription,omitempty"`

	// HasWBSLines: Indicates if whether the Project has WBS
	HasWBSLines *bool `json:"HasWBSLines,omitempty"`

	// IncludeInvoiceSpecification: Include invoice specification. E.g: 1 = Based on account, 2 = Always, 3 = Never
	IncludeInvoiceSpecification *int `json:"IncludeInvoiceSpecification,omitempty"`

	// IncludeSpecificationInInvoicePdf: Indicates whether to include invoice specification in invoice PDF
	IncludeSpecificationInInvoicePdf *bool `json:"IncludeSpecificationInInvoicePdf,omitempty"`

	// InternalNotes: Internal notes not to be printed in invoice
	InternalNotes *string `json:"InternalNotes,omitempty"`

	// InvoiceAddress: Invoice address
	InvoiceAddress *types.GUID `json:"InvoiceAddress,omitempty"`

	// InvoiceAsQuoted: Indicates whether the project is invoice as quoted
	InvoiceAsQuoted *bool `json:"InvoiceAsQuoted,omitempty"`

	// InvoiceTerms: Collection of invoice terms
	InvoiceTerms *json.RawMessage `json:"InvoiceTerms,omitempty"`

	// Manager: Responsible person for this project
	Manager *types.GUID `json:"Manager,omitempty"`

	// ManagerFullname: Name of Manager
	ManagerFullname *string `json:"ManagerFullname,omitempty"`

	// MarkupPercentage: Purchase markup percentage
	MarkupPercentage *float64 `json:"MarkupPercentage,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: For additional information about projects
	Notes *string `json:"Notes,omitempty"`

	// PrepaidItem: Used only for PSA. This item is used for prepaid invoicing. If left empty, the functionality relies on a setting
	PrepaidItem *types.GUID `json:"PrepaidItem,omitempty"`

	// PrepaidItemDescription: Description of PrepaidItem
	PrepaidItemDescription *string `json:"PrepaidItemDescription,omitempty"`

	// PrepaidType: PrepaidType: 1-Retainer, 2-Hour type bundle
	PrepaidType *int `json:"PrepaidType,omitempty"`

	// PrepaidTypeDescription: Description of PrepaidType
	PrepaidTypeDescription *string `json:"PrepaidTypeDescription,omitempty"`

	// ProjectRestrictionEmployees: Collection of employee restrictions
	ProjectRestrictionEmployees *json.RawMessage `json:"ProjectRestrictionEmployees,omitempty"`

	// ProjectRestrictionItems: Collection of item restrictions
	ProjectRestrictionItems *json.RawMessage `json:"ProjectRestrictionItems,omitempty"`

	// ProjectRestrictionRebillings: Collection of rebilling restrictions
	ProjectRestrictionRebillings *json.RawMessage `json:"ProjectRestrictionRebillings,omitempty"`

	// SalesTimeQuantity: Budgeted time. Total number of hours estimated for the fixed price project
	SalesTimeQuantity *float64 `json:"SalesTimeQuantity,omitempty"`

	// SourceQuotation: Source quotation
	SourceQuotation *types.GUID `json:"SourceQuotation,omitempty"`

	// StartDate: Start date of a project. In combination with the end date the status is determined
	StartDate *types.Date `json:"StartDate,omitempty"`

	// TimeQuantityToAlert: Alert when exceeding (Hours)
	TimeQuantityToAlert *float64 `json:"TimeQuantityToAlert,omitempty"`

	// Type: Reference to ProjectTypes. E.g: 1 = Campaign , 2 = Fixed Price, 3 = Time and Material, 4 = Non billable, 5 = Prepaid
	Type *int `json:"Type,omitempty"`

	// TypeDescription: Description of Type
	TypeDescription *string `json:"TypeDescription,omitempty"`

	// UseBillingMilestones: Indicates whether the Project is using billing milestones
	UseBillingMilestones *bool `json:"UseBillingMilestones,omitempty"`
}

func (e *Projects) GetPrimary() *types.GUID {
	return e.ID
}

func (s *ProjectsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/Projects", method)
}

// List the Projects entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ProjectsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Projects, error) {
	var entities []*Projects
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/Projects", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Projects entitiy in the provided division.
func (s *ProjectsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*Projects, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/Projects", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Projects{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty Projects entity
func (s *ProjectsEndpoint) New() *Projects {
	return &Projects{}
}

// Create the Projects entity in the provided division.
func (s *ProjectsEndpoint) Create(ctx context.Context, division int, entity *Projects) (*Projects, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/Projects", division) // #nosec
	e := &Projects{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the Projects entity in the provided division.
func (s *ProjectsEndpoint) Update(ctx context.Context, division int, entity *Projects) (*Projects, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/Projects", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &Projects{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the Projects entity in the provided division.
func (s *ProjectsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/Projects", division) // #nosec
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
