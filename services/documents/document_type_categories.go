// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package documents

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// DocumentTypeCategoriesEndpoint is responsible for communicating with
// the DocumentTypeCategories endpoint of the Documents service.
type DocumentTypeCategoriesEndpoint service

// DocumentTypeCategories:
// Service: Documents
// Entity: DocumentTypeCategories
// URL: /api/v1/{division}/documents/DocumentTypeCategories
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=DocumentsDocumentTypeCategories
type DocumentTypeCategories struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *int `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Description: Document category type description
	Description *string `json:"Description,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`
}

func (e *DocumentTypeCategories) GetPrimary() *int {
	return e.ID
}

func (s *DocumentTypeCategoriesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "documents/DocumentTypeCategories", method)
}

// List the DocumentTypeCategories entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentTypeCategoriesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*DocumentTypeCategories, error) {
	var entities []*DocumentTypeCategories
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypeCategories", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the DocumentTypeCategories entitiy in the provided division.
func (s *DocumentTypeCategoriesEndpoint) Get(ctx context.Context, division int, id *int) (*DocumentTypeCategories, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypeCategories", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &DocumentTypeCategories{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
