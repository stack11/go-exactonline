// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// CashflowImportNotificationsService is responsible for communicating with
// the ImportNotifications endpoint of the Cashflow service.
type CashflowImportNotificationsService service

// CashflowImportNotifications:
// Service: Cashflow
// Entity: ImportNotifications
// URL: /api/v1/{division}/cashflow/ImportNotifications
// HasWebhook: false
// IsInBeta: false
// Methods: GET PUT
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CashflowImportNotifications
type CashflowImportNotifications struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// BankAccount: Number of the bank account related to the notification.
	BankAccount *string `json:",omitempty"`

	// BankAccountID: ID of the bank account related to the notification.
	BankAccountID *GUID `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// Document: ID of the document related to the notification.
	Document *GUID `json:",omitempty"`

	// Result: Code of the import result
	Result *int `json:",omitempty"`

	// ResultDescription: Description of the import result
	ResultDescription *string `json:",omitempty"`

	// RetriedBy: ID of the user that requested a retry of the import.
	RetriedBy *GUID `json:",omitempty"`

	// RetriedOn: Date when the retry was requested.
	RetriedOn *Date `json:",omitempty"`
}

func (s *CashflowImportNotifications) GetIdentifier() GUID {
	return *s.ID
}

// List the ImportNotifications entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CashflowImportNotificationsService) List(ctx context.Context, division int, all bool) ([]*CashflowImportNotifications, error) {
	var entities []*CashflowImportNotifications
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/ImportNotifications?$select=*", division)
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
