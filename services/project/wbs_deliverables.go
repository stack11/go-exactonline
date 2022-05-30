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

// WBSDeliverablesEndpoint is responsible for communicating with
// the WBSDeliverables endpoint of the Project service.
type WBSDeliverablesEndpoint service

// WBSDeliverables:
// Service: Project
// Entity: WBSDeliverables
// URL: /api/v1/{division}/project/WBSDeliverables
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectWBSDeliverables
type WBSDeliverables struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Completed: To indicated if the WBS deliverable is completed. WBS deliverable can only be set to completed when an invoice term is linked
	Completed *bool `json:"Completed,omitempty"`

	// Created: The date and time when the WBS deliverable was created
	Created *types.Date `json:"Created,omitempty"`

	// Creator: The ID of the user that created the WBS deliverable
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: The full name of the user that created the WBS deliverable
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the WBS deliverable
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Milestone: To indicated if this is a milestone
	Milestone *bool `json:"Milestone,omitempty"`

	// Modified: The date when the WBS deliverable was modified
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: The ID of the user that modified the WBS deliverable
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: The full name of the user that modified the WBS deliverable
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: For additional information
	Notes *string `json:"Notes,omitempty"`

	// PartOf: ID of the WBS deliverable part of
	PartOf *types.GUID `json:"PartOf,omitempty"`

	// PartOfDescription: Description of part of
	PartOfDescription *string `json:"PartOfDescription,omitempty"`

	// Project: ID of the project that linked to WBS deliverable
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectDescription: Project description that is linked to WBS deliverable
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectTerm: ID of invoice term that linked to the WBS deliverable. Invoice term can only be linked when the project is using billing milestones and the deliverable is set as a milestone
	ProjectTerm *types.GUID `json:"ProjectTerm,omitempty"`

	// ProjectTermDescription: Description of invoice term that linked to the WBS deliverable
	ProjectTermDescription *string `json:"ProjectTermDescription,omitempty"`

	// ReleaseInvoiceTerm: Action to release the invoice term. You can only release a WBS deliverable&#39;s invoice term once and it cannot be undo
	ReleaseInvoiceTerm *bool `json:"ReleaseInvoiceTerm,omitempty"`

	// ReleaseInvoiceTermDate: Release invoice term date. The linked invoice term date can be updated by using this property. The update will only happen when releasing a WBS deliverable&#39;s invoice term
	ReleaseInvoiceTermDate *types.Date `json:"ReleaseInvoiceTermDate,omitempty"`

	// ReleaseInvoiceTermHasSpecifyDate: Release invoice term has specify date
	ReleaseInvoiceTermHasSpecifyDate *bool `json:"ReleaseInvoiceTermHasSpecifyDate,omitempty"`

	// SequenceNumber: Sequence number of the WBS deliverable. Last sequence will be selected if not specified
	SequenceNumber *int `json:"SequenceNumber,omitempty"`

	// Type: The type of project WBS. E.g: 1 = Deliverable, 2 = Activity, 3 = Expense
	Type *int `json:"Type,omitempty"`

	// UpdateAction: Update action
	UpdateAction *int `json:"UpdateAction,omitempty"`
}

func (e *WBSDeliverables) GetPrimary() *types.GUID {
	return e.ID
}

func (s *WBSDeliverablesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/WBSDeliverables", method)
}

// List the WBSDeliverables entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *WBSDeliverablesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*WBSDeliverables, error) {
	var entities []*WBSDeliverables
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSDeliverables", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the WBSDeliverables entitiy in the provided division.
func (s *WBSDeliverablesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*WBSDeliverables, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSDeliverables", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &WBSDeliverables{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty WBSDeliverables entity
func (s *WBSDeliverablesEndpoint) New() *WBSDeliverables {
	return &WBSDeliverables{}
}

// Create the WBSDeliverables entity in the provided division.
func (s *WBSDeliverablesEndpoint) Create(ctx context.Context, division int, entity *WBSDeliverables) (*WBSDeliverables, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSDeliverables", division) // #nosec
	e := &WBSDeliverables{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the WBSDeliverables entity in the provided division.
func (s *WBSDeliverablesEndpoint) Update(ctx context.Context, division int, entity *WBSDeliverables) (*WBSDeliverables, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSDeliverables", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &WBSDeliverables{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the WBSDeliverables entity in the provided division.
func (s *WBSDeliverablesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/WBSDeliverables", division) // #nosec
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
