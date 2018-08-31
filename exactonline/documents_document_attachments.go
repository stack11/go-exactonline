// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// DocumentsDocumentAttachmentsService is responsible for communicating with
// the DocumentAttachments endpoint of the Documents service.
type DocumentsDocumentAttachmentsService service

// DocumentsDocumentAttachments:
// Service: Documents
// Entity: DocumentAttachments
// URL: /api/v1/{division}/documents/DocumentAttachments
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=DocumentsDocumentAttachments
type DocumentsDocumentAttachments struct {
	// ID:
	ID *GUID `json:",omitempty"`

	// Attachment:
	Attachment *[]byte `json:",omitempty"`

	// Document:
	Document *GUID `json:",omitempty"`

	// FileName:
	FileName *string `json:",omitempty"`

	// FileSize:
	FileSize *float64 `json:",omitempty"`

	// Url:
	Url *string `json:",omitempty"`
}

func (s *DocumentsDocumentAttachments) GetIdentifier() GUID {
	return *s.ID
}

// List the DocumentAttachments entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentsDocumentAttachmentsService) List(ctx context.Context, division int, all bool) ([]*DocumentsDocumentAttachments, error) {
	var entities []*DocumentsDocumentAttachments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentAttachments?$select=*", division)
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
