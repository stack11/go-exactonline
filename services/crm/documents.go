// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package crm

import (
	"context"
	"encoding/json"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// DocumentsEndpoint is responsible for communicating with
// the Documents endpoint of the CRM service.
type DocumentsEndpoint service

// Documents:
// Service: CRM
// Entity: Documents
// URL: /api/v1/{division}/read/crm/Documents
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadCRMDocuments
type Documents struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: ID of the related account of this document
	Account *types.GUID `json:"Account,omitempty"`

	// Attachments: Attachments linked to the document. Binaries are not sent in the response.
	Attachments *json.RawMessage `json:"Attachments,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// DocumentDate: Entry date of the incoming document
	DocumentDate *types.Date `json:"DocumentDate,omitempty"`

	// DocumentFolder: Id of document folder
	DocumentFolder *types.GUID `json:"DocumentFolder,omitempty"`

	// DocumentViewUrl: Url to view the document
	DocumentViewUrl *string `json:"DocumentViewUrl,omitempty"`

	// HasEmptyBody: Indicates that the document body is empty
	HasEmptyBody *bool `json:"HasEmptyBody,omitempty"`

	// HID: Human-readable ID, formatted as xx.xxx.xxx. Unique. May not be equal to zero
	HID *int `json:"HID,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// Opportunity: The opportunity linked to the document
	Opportunity *types.GUID `json:"Opportunity,omitempty"`

	// PurchaseInvoiceNumber: Purchase invoice number.
	PurchaseInvoiceNumber *int `json:"PurchaseInvoiceNumber,omitempty"`

	// PurchaseOrderNumber: Purchase order number.
	PurchaseOrderNumber *int `json:"PurchaseOrderNumber,omitempty"`

	// SalesInvoiceNumber: &#39;Our reference&#39; of the transaction that belongs to this document
	SalesInvoiceNumber *int `json:"SalesInvoiceNumber,omitempty"`

	// SalesOrderNumber: Number of the sales order
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// SendMethod: Send Method
	SendMethod *int `json:"SendMethod,omitempty"`

	// Subject: Subject of this document
	Subject *string `json:"Subject,omitempty"`

	// Type: The document type
	Type *int `json:"Type,omitempty"`

	// TypeDescription: Translated description of the document type. $filter and $orderby are not supported for this property.
	TypeDescription *string `json:"TypeDescription,omitempty"`
}

// List the Documents entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Documents, error) {
	var entities []*Documents
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/crm/Documents", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
