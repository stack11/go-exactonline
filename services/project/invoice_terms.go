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

// InvoiceTermsEndpoint is responsible for communicating with
// the InvoiceTerms endpoint of the Project service.
type InvoiceTermsEndpoint service

// InvoiceTerms:
// Service: Project
// Entity: InvoiceTerms
// URL: /api/v1/{division}/project/InvoiceTerms
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectInvoiceTerms
type InvoiceTerms struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Amount: Amount in the currency of the transaction
	Amount *float64 `json:"Amount,omitempty"`

	// Created: Date and time when the invoice term was created
	Created *types.Date `json:"Created,omitempty"`

	// Creator: ID of user that created the invoice term
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Full name of user that created the record
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Deliverable: WBS&#39;s deliverable linked to the invoice term
	Deliverable *string `json:"Deliverable,omitempty"`

	// Description: Description of invoice term
	Description *string `json:"Description,omitempty"`

	// Division: Division number
	Division *int `json:"Division,omitempty"`

	// ExecutionFromDate: Execution date: From of invoice term
	ExecutionFromDate *types.Date `json:"ExecutionFromDate,omitempty"`

	// ExecutionToDate: Execution date: To of invoice term
	ExecutionToDate *types.Date `json:"ExecutionToDate,omitempty"`

	// InvoiceDate: Invoice date of invoice term
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// Item: Item that linked to the invoice term
	Item *types.GUID `json:"Item,omitempty"`

	// ItemDescription: Description of item that linked to the invoice term
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Modified: Last modified date of invoice term
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: ID of user that modified the record
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Full name of user that modified the record
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes linked to the invoice term for providing additional information
	Notes *string `json:"Notes,omitempty"`

	// Percentage: Percentage of amount per project&#39;s budgeted amount
	Percentage *float64 `json:"Percentage,omitempty"`

	// Project: ID of project that linked to the invoice term
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectDescription: Project description that linked to the invoice term
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// VATCode: VAT code that used in the invoice term
	VATCode *string `json:"VATCode,omitempty"`

	// VATCodeDescription: Description of VAT code that used in the invoice term
	VATCodeDescription *string `json:"VATCodeDescription,omitempty"`

	// VATPercentage: Percentage of VAT code that used in the invoice term
	VATPercentage *float64 `json:"VATPercentage,omitempty"`

	// WBS: ID of WBS that linked to the invoice term
	WBS *types.GUID `json:"WBS,omitempty"`
}

func (e *InvoiceTerms) GetPrimary() *types.GUID {
	return e.ID
}

func (s *InvoiceTermsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/InvoiceTerms", method)
}

// List the InvoiceTerms entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InvoiceTermsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*InvoiceTerms, error) {
	var entities []*InvoiceTerms
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/InvoiceTerms", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the InvoiceTerms entitiy in the provided division.
func (s *InvoiceTermsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*InvoiceTerms, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/InvoiceTerms", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &InvoiceTerms{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty InvoiceTerms entity
func (s *InvoiceTermsEndpoint) New() *InvoiceTerms {
	return &InvoiceTerms{}
}

// Create the InvoiceTerms entity in the provided division.
func (s *InvoiceTermsEndpoint) Create(ctx context.Context, division int, entity *InvoiceTerms) (*InvoiceTerms, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/InvoiceTerms", division) // #nosec
	e := &InvoiceTerms{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the InvoiceTerms entity in the provided division.
func (s *InvoiceTermsEndpoint) Update(ctx context.Context, division int, entity *InvoiceTerms) (*InvoiceTerms, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/InvoiceTerms", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &InvoiceTerms{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the InvoiceTerms entity in the provided division.
func (s *InvoiceTermsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/project/InvoiceTerms", division) // #nosec
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
