// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// CashflowImportNotificationDetailsService is responsible for communicating with
// the ImportNotificationDetails endpoint of the Cashflow service.
type CashflowImportNotificationDetailsService service

// CashflowImportNotificationDetails:
// Service: Cashflow
// Entity: ImportNotificationDetails
// URL: /api/v1/{division}/cashflow/ImportNotificationDetails
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CashflowImportNotificationDetails
type CashflowImportNotificationDetails struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// CashflowImportNotification: ID of the notification these details belong to.
	CashflowImportNotification *GUID `json:",omitempty"`

	// CashflowTransactionFeed: ID of the cashflow transaction feed related to this notification.
	CashflowTransactionFeed *GUID `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// Message: Termed message
	Message *string `json:",omitempty"`

	// ResponseCode: Response code
	ResponseCode *int `json:",omitempty"`

	// ResponseCodeArguments: Additional information about the response
	ResponseCodeArguments *string `json:",omitempty"`
}

func (s *CashflowImportNotificationDetails) GetIdentifier() GUID {
	return *s.ID
}

// List the ImportNotificationDetails entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CashflowImportNotificationDetailsService) List(ctx context.Context, division int, all bool) ([]*CashflowImportNotificationDetails, error) {
	var entities []*CashflowImportNotificationDetails
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/ImportNotificationDetails?$select=*", division)
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
