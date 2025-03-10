// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package sync

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// DocumentsDocumentAttachmentsEndpoint is responsible for communicating with
// the DocumentsDocumentAttachments endpoint of the Sync service.
type DocumentsDocumentAttachmentsEndpoint service

// DocumentsDocumentAttachments:
// Service: Sync
// Entity: DocumentsDocumentAttachments
// URL: /api/v1/{division}/sync/Documents/DocumentAttachments
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncDocumentsDocumentAttachments
type DocumentsDocumentAttachments struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp:
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// Attachment:
	Attachment *[]byte `json:"Attachment,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// FileName:
	FileName *string `json:"FileName,omitempty"`

	// FileSize:
	FileSize *float64 `json:"FileSize,omitempty"`

	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Url:
	Url *string `json:"Url,omitempty"`
}

func (e *DocumentsDocumentAttachments) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *DocumentsDocumentAttachmentsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "Documents/DocumentAttachments", method)
}

// List the DocumentsDocumentAttachments entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentsDocumentAttachmentsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*DocumentsDocumentAttachments, error) {
	var entities []*DocumentsDocumentAttachments
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Documents/DocumentAttachments", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the DocumentsDocumentAttachments entitiy in the provided division.
func (s *DocumentsDocumentAttachmentsEndpoint) Get(ctx context.Context, division int, id *int64) (*DocumentsDocumentAttachments, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Documents/DocumentAttachments", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &DocumentsDocumentAttachments{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
