// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// FinancialTransactionCashEntriesService is responsible for communicating with
// the CashEntries endpoint of the FinancialTransaction service.
type FinancialTransactionCashEntriesService service

// FinancialTransactionCashEntries:
// Service: FinancialTransaction
// Entity: CashEntries
// URL: /api/v1/{division}/financialtransaction/CashEntries
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialTransactionCashEntries
type FinancialTransactionCashEntries struct {
	// EntryID:
	EntryID *GUID `json:",omitempty"`

	// CashEntryLines:
	CashEntryLines *[]byte `json:",omitempty"`

	// ClosingBalanceFC:
	ClosingBalanceFC *float64 `json:",omitempty"`

	// Created:
	Created *Date `json:",omitempty"`

	// Currency:
	Currency *string `json:",omitempty"`

	// Division:
	Division *int `json:",omitempty"`

	// EntryNumber:
	EntryNumber *int `json:",omitempty"`

	// FinancialPeriod:
	FinancialPeriod *int `json:",omitempty"`

	// FinancialYear:
	FinancialYear *int `json:",omitempty"`

	// JournalCode:
	JournalCode *string `json:",omitempty"`

	// JournalDescription:
	JournalDescription *string `json:",omitempty"`

	// Modified:
	Modified *Date `json:",omitempty"`

	// OpeningBalanceFC:
	OpeningBalanceFC *float64 `json:",omitempty"`

	// Status:
	Status *int `json:",omitempty"`

	// StatusDescription:
	StatusDescription *string `json:",omitempty"`
}

func (s *FinancialTransactionCashEntries) GetIdentifier() GUID {
	return *s.EntryID
}

// List the CashEntries entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *FinancialTransactionCashEntriesService) List(ctx context.Context, division int, all bool) ([]*FinancialTransactionCashEntries, error) {
	var entities []*FinancialTransactionCashEntries
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/CashEntries?$select=*", division)
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
