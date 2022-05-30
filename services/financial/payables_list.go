// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
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
	MetaData *api.MetaData `json:"__metadata,omitempty"`
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

	// AmountInTransit: Amount in transit - The amount that you have requested your bank to pay
	AmountInTransit *float64 `json:"AmountInTransit,omitempty"`

	// ApprovalStatus: Approval status: // 						null - Invoice was entered before approval functionality was activated (treated as Approved for payments) // 						1 - N/A (used for non-electronic payment methods) // 						2 - Awaiting review // 						3 - Awaiting approval // 						4 - Approved
	ApprovalStatus *int `json:"ApprovalStatus,omitempty"`

	// CurrencyCode: Code of Currency
	CurrencyCode *string `json:"CurrencyCode,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// DueDate: Date the invoice is due (This due date is not the discount due date)
	DueDate *types.Date `json:"DueDate,omitempty"`

	// EntryNumber: The entry number of this payment term corresponding purchase/cashflow entry
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// Id: Obsolete
	Id *types.GUID `json:"Id,omitempty"`

	// InvoiceDate: The date of the invoice or the date when money is paid to the supplier
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// InvoiceNumber: For purchase entry, it would be the entrynumber. For the cashflow entry, it will be entrynumber of an invoice which this payment is for
	InvoiceNumber *int `json:"InvoiceNumber,omitempty"`

	// JournalCode: Code of Journal
	JournalCode *string `json:"JournalCode,omitempty"`

	// JournalDescription: Description of Journal
	JournalDescription *string `json:"JournalDescription,omitempty"`

	// YourRef: Purchase invoice Your Reference number. Will be null if it is a cashflow entry
	YourRef *string `json:"YourRef,omitempty"`
}

func (e *PayablesList) GetPrimary() *int64 {
	return e.HID
}

func (s *PayablesListEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "financial/PayablesList", method)
}

// List the PayablesList entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PayablesListEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*PayablesList, error) {
	var entities []*PayablesList
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/PayablesList", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the PayablesList entitiy in the provided division.
func (s *PayablesListEndpoint) Get(ctx context.Context, division int, id *int64) (*PayablesList, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/PayablesList", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &PayablesList{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
