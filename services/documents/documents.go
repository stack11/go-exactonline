// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package documents

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// DocumentsEndpoint is responsible for communicating with
// the Documents endpoint of the Documents service.
type DocumentsEndpoint service

// Documents:
// Service: Documents
// Entity: Documents
// URL: /api/v1/{division}/documents/Documents
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=DocumentsDocuments
type Documents struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Account:
	Account *types.GUID `json:"Account,omitempty"`

	// AccountCode:
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// Body:
	Body *string `json:"Body,omitempty"`

	// Category:
	Category *types.GUID `json:"Category,omitempty"`

	// CategoryDescription:
	CategoryDescription *string `json:"CategoryDescription,omitempty"`

	// Contact:
	Contact *types.GUID `json:"Contact,omitempty"`

	// ContactFullName:
	ContactFullName *string `json:"ContactFullName,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// DocumentDate:
	DocumentDate *types.Date `json:"DocumentDate,omitempty"`

	// DocumentFolder:
	DocumentFolder *types.GUID `json:"DocumentFolder,omitempty"`

	// DocumentFolderCode:
	DocumentFolderCode *string `json:"DocumentFolderCode,omitempty"`

	// DocumentFolderDescription:
	DocumentFolderDescription *string `json:"DocumentFolderDescription,omitempty"`

	// DocumentViewUrl:
	DocumentViewUrl *string `json:"DocumentViewUrl,omitempty"`

	// FinancialTransactionEntryID:
	FinancialTransactionEntryID *types.GUID `json:"FinancialTransactionEntryID,omitempty"`

	// HasEmptyBody:
	HasEmptyBody *bool `json:"HasEmptyBody,omitempty"`

	// HID:
	HID *int `json:"HID,omitempty"`

	// Language:
	Language *string `json:"Language,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Opportunity:
	Opportunity *types.GUID `json:"Opportunity,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode:
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// SalesInvoiceNumber:
	SalesInvoiceNumber *int `json:"SalesInvoiceNumber,omitempty"`

	// SalesOrderNumber:
	SalesOrderNumber *int `json:"SalesOrderNumber,omitempty"`

	// SendMethod:
	SendMethod *int `json:"SendMethod,omitempty"`

	// ShopOrderNumber:
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// Subject:
	Subject *string `json:"Subject,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`

	// TypeDescription:
	TypeDescription *string `json:"TypeDescription,omitempty"`
}

// List the Documents entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *DocumentsEndpoint) List(ctx context.Context, division int, all bool) ([]*Documents, error) {
	var entities []*Documents
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/Documents?$select=*", division) // #nosec
	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
