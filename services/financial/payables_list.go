// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// PayablesListEndpoint is responsible for communicating with
// the PayablesList endpoint of the Financial service.
type PayablesListEndpoint service

// PayablesList:
// Service: Financial
// Entity: PayablesList
// URL: /api/v1/{division}/read/financial/PayablesList
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialPayablesList
type PayablesList struct {
	// HID: Primary key, human readable ID
	HID *int64 `json:"HID,omitempty"`

	// AccountCode: Code of Account
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountId: Reference to the account
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:"AccountName,omitempty"`

	// Amount: Amount
	Amount *float64 `json:"Amount,omitempty"`

	// AmountInTransit: Amount in transit
	AmountInTransit *float64 `json:"AmountInTransit,omitempty"`

	// ApprovalStatus: Approval status: // 						null - Invoice was entered before approval functionality was activated (treated as Approved for payments) // 						1 - N/A (used for non-electronic payment methods) // 						2 - Awaiting review // 						3 - Awaiting approval // 						4 - Approved
	ApprovalStatus *int `json:"ApprovalStatus,omitempty"`

	// CurrencyCode: Code of Currency
	CurrencyCode *string `json:"CurrencyCode,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// DueDate: Date the invoice is due (This due date is not the discount due date)
	DueDate *types.Date `json:"DueDate,omitempty"`

	// EntryNumber: Entry number
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// Id: Obsolete
	Id *types.GUID `json:"Id,omitempty"`

	// InvoiceDate: Invoice date
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// InvoiceNumber: Invoice number. The value is 0 when the invoice number of the linked transaction is empty.
	InvoiceNumber *int `json:"InvoiceNumber,omitempty"`

	// JournalCode: Code of Journal
	JournalCode *string `json:"JournalCode,omitempty"`

	// JournalDescription: Description of Journal
	JournalDescription *string `json:"JournalDescription,omitempty"`

	// YourRef: Your reference
	YourRef *string `json:"YourRef,omitempty"`
}

// List the PayablesList entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PayablesListEndpoint) List(ctx context.Context, division int, all bool) ([]*PayablesList, error) {
	var entities []*PayablesList
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/PayablesList?$select=*", division) // #nosec
	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
