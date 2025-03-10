// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financialtransaction

import "github.com/stack11/go-exactonline/api"

type service struct {
	client *api.Client
}

// FinancialTransactionService is responsible for communication with the FinancialTransaction
// endpoints of the Exact Online API.
type FinancialTransactionService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	BankEntries      *BankEntriesEndpoint
	BankEntryLines   *BankEntryLinesEndpoint
	CashEntries      *CashEntriesEndpoint
	CashEntryLines   *CashEntryLinesEndpoint
	TransactionLines *TransactionLinesEndpoint
}

// NewFinancialTransactionService creates a new initialized instance of the
// FinancialTransactionService.
func NewFinancialTransactionService(apiClient *api.Client) *FinancialTransactionService {
	s := &FinancialTransactionService{client: apiClient}

	s.common.client = apiClient

	s.BankEntries = (*BankEntriesEndpoint)(&s.common)
	s.BankEntryLines = (*BankEntryLinesEndpoint)(&s.common)
	s.CashEntries = (*CashEntriesEndpoint)(&s.common)
	s.CashEntryLines = (*CashEntryLinesEndpoint)(&s.common)
	s.TransactionLines = (*TransactionLinesEndpoint)(&s.common)

	return s
}
