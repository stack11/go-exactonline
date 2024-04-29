// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// WBSActivitiesEndpoint is responsible for communicating with
// the WBSActivities endpoint of the Project service.
type WBSActivitiesEndpoint service

// WBSActivities:
// Service: Project
// Entity: WBSActivities
// URL: /api/v1/{division}/project/WBSActivities
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectWBSActivities
type WBSActivities struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AutoCreateInvoiceTerm: To indicated auto create invoice term when invoice method is Fixed price
	AutoCreateInvoiceTerm *bool `json:"AutoCreateInvoiceTerm,omitempty"`

	// BlockEntry: To indicated if time and cost entries is blocked
	BlockEntry *bool `json:"BlockEntry,omitempty"`

	// BudgetedCost: Budget cost of the WBS activity
	BudgetedCost *float64 `json:"BudgetedCost,omitempty"`

	// BudgetedHours: Budget hours of the WBS
	BudgetedHours *float64 `json:"BudgetedHours,omitempty"`

	// BudgetedRevenue: Revenue of the WBS activity
	BudgetedRevenue *float64 `json:"BudgetedRevenue,omitempty"`

	// BudgetOverrunHours: BudgetOverrunHours: 10-Allowed, 20-Not Allowed
	BudgetOverrunHours *byte `json:"BudgetOverrunHours,omitempty"`

	// Completed: To indicated if the WBS activity is completed
	Completed *bool `json:"Completed,omitempty"`

	// Created: The date and time when the WBS activity was created
	Created *types.Date `json:"Created,omitempty"`

	// Creator: The ID of the user that created the WBS activity
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: The full name of the user that created the WBS activity
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// CustomField: Custom field endpoint. Provided only for the Exact Online Premium users.
	CustomField *string `json:"CustomField,omitempty"`

	// DefaultItem: Default Item to used for time entry
	DefaultItem *types.GUID `json:"DefaultItem,omitempty"`

	// DefaultItemIsMandatory: To indicated if only default item is allowed
	DefaultItemIsMandatory *bool `json:"DefaultItemIsMandatory,omitempty"`

	// Description: Description of the WBS activity
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EndDate: End date of the WBS activity
	EndDate *types.Date `json:"EndDate,omitempty"`

	// InvoiceDate: The invoice date of the WBS when auto create invoice term
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// InvoiceMethod: The invoice method of the WBS. E.g: 2 = Fixed price, 3 = Time and Material, 4 = Non billable, 5 = Prepaid
	InvoiceMethod *int `json:"InvoiceMethod,omitempty"`

	// InvoiceSeparately: To indicated if additional invoice is allowed on this WBS expense. Additional invoice can only be set when the project type is fixed price or prepaid and the project allow additional invoices.
	InvoiceSeparately *bool `json:"InvoiceSeparately,omitempty"`

	// InvoiceTerm: ID of the invoice term that linked to the WBS
	InvoiceTerm *types.GUID `json:"InvoiceTerm,omitempty"`

	// Modified: The date when the WBS activity was modified
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: The ID of the user that modified the WBS activity
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: The full name of the user that modified the WBS activity
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: For additional information
	Notes *string `json:"Notes,omitempty"`

	// PartOf: ID of the WBS activity part of
	PartOf *types.GUID `json:"PartOf,omitempty"`

	// PartOfDescription: Description of part of
	PartOfDescription *string `json:"PartOfDescription,omitempty"`

	// Project: ID of the project that linked to WBS activity
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectDescription: Project description that is linked to WBS activity
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectTermAction: Project term action
	ProjectTermAction *int `json:"ProjectTermAction,omitempty"`

	// ReleaseInvoiceTerm: Action to release the invoice term. You can only release a WBS activity&#39;s invoice term once and it cannot be undo
	ReleaseInvoiceTerm *bool `json:"ReleaseInvoiceTerm,omitempty"`

	// ReleaseInvoiceTermDate: Release invoice term date. The linked invoice term date can be updated by using this property. The update will only happen when releasing a WBS activity&#39;s invoice term
	ReleaseInvoiceTermDate *types.Date `json:"ReleaseInvoiceTermDate,omitempty"`

	// ReleaseInvoiceTermHasSpecifyDate: Release invoice term has specify date
	ReleaseInvoiceTermHasSpecifyDate *bool `json:"ReleaseInvoiceTermHasSpecifyDate,omitempty"`

	// SequenceNumber: Sequence number of the WBS activity. Last sequence will be selected if not specified
	SequenceNumber *int `json:"SequenceNumber,omitempty"`

	// StartDate: Start date of the WBS activity
	StartDate *types.Date `json:"StartDate,omitempty"`

	// TimeQuantityToAlert: Alert when exceeding this time quantity
	TimeQuantityToAlert *float64 `json:"TimeQuantityToAlert,omitempty"`

	// Type: The type of project WBS. E.g: 1 = Deliverable, 2 = Activity, 3 = Expense
	Type *int `json:"Type,omitempty"`

	// UpdateAction: Update action
	UpdateAction *int `json:"UpdateAction,omitempty"`
}

func (e *WBSActivities) GetPrimary() *types.GUID {
	return e.ID
}

func (s *WBSActivitiesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/WBSActivities", method)
}

// List the WBSActivities entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *WBSActivitiesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*WBSActivities, error) {
	var entities []*WBSActivities
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSActivities", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the WBSActivities entitiy in the provided division.
func (s *WBSActivitiesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*WBSActivities, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSActivities", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &WBSActivities{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty WBSActivities entity
func (s *WBSActivitiesEndpoint) New() *WBSActivities {
	return &WBSActivities{}
}

// Create the WBSActivities entity in the provided division.
func (s *WBSActivitiesEndpoint) Create(ctx context.Context, division int, entity *WBSActivities) (*WBSActivities, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSActivities", division) // #nosec
	e := &WBSActivities{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the WBSActivities entity in the provided division.
func (s *WBSActivitiesEndpoint) Update(ctx context.Context, division int, entity *WBSActivities) (*WBSActivities, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSActivities", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &WBSActivities{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the WBSActivities entity in the provided division.
func (s *WBSActivitiesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSActivities", division) // #nosec
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
