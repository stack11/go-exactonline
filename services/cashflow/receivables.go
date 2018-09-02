// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package cashflow

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// ReceivablesEndpoint is responsible for communicating with
// the Receivables endpoint of the Cashflow service.
type ReceivablesEndpoint service

// Receivables:
// Service: Cashflow
// Entity: Receivables
// URL: /api/v1/{division}/cashflow/Receivables
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CashflowReceivables
type Receivables struct {
	// ID: Identifier of the receivable.
	ID *types.GUID `json:"ID,omitempty"`

	// Account: The customer from which the receivable will come.
	Account *types.GUID `json:"Account,omitempty"`

	// AccountBankAccountID: The bank account of the customer, from which the receivable will come.
	AccountBankAccountID *types.GUID `json:"AccountBankAccountID,omitempty"`

	// AccountBankAccountNumber: The bank account number of the customer, from which the receivable will come.
	AccountBankAccountNumber *string `json:"AccountBankAccountNumber,omitempty"`

	// AccountCode: The code of the customer from which the receivable will come.
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountContact: Contact person copied from the purchase invoice linked to the related purchase entry.
	AccountContact *types.GUID `json:"AccountContact,omitempty"`

	// AccountContactName: Name of the contact person of the customer.
	AccountContactName *string `json:"AccountContactName,omitempty"`

	// AccountCountry: Country code of the customer.
	AccountCountry *string `json:"AccountCountry,omitempty"`

	// AccountName: Name of the customer.
	AccountName *string `json:"AccountName,omitempty"`

	// AmountDC: The amount in default currency (division currency). Receivables are matched on this amount.
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountDiscountDC: The amount of the discount in the default currency.
	AmountDiscountDC *float64 `json:"AmountDiscountDC,omitempty"`

	// AmountDiscountFC: The amount of the discount. This is in the amount of the selected currency.
	AmountDiscountFC *float64 `json:"AmountDiscountFC,omitempty"`

	// AmountFC: The amount of the receivable. This is in the amount of the selected currency.
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// BankAccountID: Own bank account to which the receivable will be done.
	BankAccountID *types.GUID `json:"BankAccountID,omitempty"`

	// BankAccountNumber: Own bank account number to which the receivable will be done.
	BankAccountNumber *string `json:"BankAccountNumber,omitempty"`

	// CashflowTransactionBatchCode: When processing receivables, all receivable with the same processing data are put in a batch. This field contains the code of that batch.
	CashflowTransactionBatchCode *string `json:"CashflowTransactionBatchCode,omitempty"`

	// Created: Creation date.
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of the creator.
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator.
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency: The currency of the receivable. This currency can only deviate from the division currency if the module Currency is in the license.
	Currency *string `json:"Currency,omitempty"`

	// Description: Description.
	Description *string `json:"Description,omitempty"`

	// DirectDebitMandate: Direct Debit Mandate used to collect the receivable.
	DirectDebitMandate *types.GUID `json:"DirectDebitMandate,omitempty"`

	// DirectDebitMandateDescription: Description of the mandate.
	DirectDebitMandateDescription *string `json:"DirectDebitMandateDescription,omitempty"`

	// DirectDebitMandatePaymentType: Payment type of the mandate. 0 = One off payment 1 = Recurrent payment.
	DirectDebitMandatePaymentType *int `json:"DirectDebitMandatePaymentType,omitempty"`

	// DirectDebitMandateReference: Unique mandate reference.
	DirectDebitMandateReference *string `json:"DirectDebitMandateReference,omitempty"`

	// DirectDebitMandateType: Type of the mandate. 0 = Core 1 = Business-to-business.
	DirectDebitMandateType *int `json:"DirectDebitMandateType,omitempty"`

	// DiscountDueDate: Date before which the payment by the customer must be done to be eligible for discount.
	DiscountDueDate *types.Date `json:"DiscountDueDate,omitempty"`

	// Division: Division code.
	Division *int `json:"Division,omitempty"`

	// Document: Document that is created when processing collections.  The bank export file is attached to the document.
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentNumber: Number of the document.
	DocumentNumber *int `json:"DocumentNumber,omitempty"`

	// DocumentSubject: Subject of the document.
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// DueDate: Date before which the payment by the customer must be done.
	DueDate *types.Date `json:"DueDate,omitempty"`

	// EndDate: Date since when the receivable is no longer an outstanding item. This is the highest invoice date of all matched receivables.
	EndDate *types.Date `json:"EndDate,omitempty"`

	// EndPeriod: Period since when the receivable is no longer an outstanding item. This is the highest period of all matched receivables.
	EndPeriod *int `json:"EndPeriod,omitempty"`

	// EndToEndID: The value of the tag &#39;EndToEndID&#39; when generating a SEPA file.
	EndToEndID *string `json:"EndToEndID,omitempty"`

	// EndYear: Year (of period) since when the receivable is no longer an outstanding item. This is the highest year of all matched receivables. Used in combination with EndPeriod.
	EndYear *int `json:"EndYear,omitempty"`

	// EntryDate: Processing date of the receivable.
	EntryDate *types.Date `json:"EntryDate,omitempty"`

	// EntryID: The unique identifier for a set of receivables. A receivable can be split so that one part is received on a different date. In that case the two records get a different EntryID.
	EntryID *types.GUID `json:"EntryID,omitempty"`

	// EntryNumber: Entry number of the linked transaction.
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// GLAccount: G/L account of the payment. Must be of type 20 (Accounts receivable).
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// GLAccountCode: Code of the G/L account.
	GLAccountCode *string `json:"GLAccountCode,omitempty"`

	// GLAccountDescription: Description of the G/L account.
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// InvoiceDate: Invoice date of the linked transaction.
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// InvoiceNumber: Invoice number of the linked transaction.
	InvoiceNumber *int `json:"InvoiceNumber,omitempty"`

	// IsBatchBooking: Boolean indicating whether the receivable is part of a batch booking.
	IsBatchBooking *byte `json:"IsBatchBooking,omitempty"`

	// IsFullyPaid: Boolean indicating whether the receivable was fully paid by the customer.
	IsFullyPaid *bool `json:"IsFullyPaid,omitempty"`

	// Journal: Journal of the linked transaction.
	Journal *string `json:"Journal,omitempty"`

	// JournalDescription: Description of the journal.
	JournalDescription *string `json:"JournalDescription,omitempty"`

	// LastPaymentDate: Last payment date.
	LastPaymentDate *types.Date `json:"LastPaymentDate,omitempty"`

	// Modified: Last modified date.
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier.
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier.
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// PaymentCondition: Payment condition of the linked transaction.
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription: Description of the payment condition.
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// PaymentDays: Number of days between invoice date and due date.
	PaymentDays *int `json:"PaymentDays,omitempty"`

	// PaymentDaysDiscount: Number of days between invoice date and due date of the discount.
	PaymentDaysDiscount *int `json:"PaymentDaysDiscount,omitempty"`

	// PaymentDiscountPercentage: Payment discount percentage.
	PaymentDiscountPercentage *float64 `json:"PaymentDiscountPercentage,omitempty"`

	// PaymentInformationID: PaymentInformationID tag from the SEPA xml file.
	PaymentInformationID *string `json:"PaymentInformationID,omitempty"`

	// PaymentMethod: Method of payment. B = On credit (default) I = Collection K = Cash V = Credit card
	PaymentMethod *string `json:"PaymentMethod,omitempty"`

	// PaymentReference: Payment reference for the receivable that may be included In the bank export file
	PaymentReference *string `json:"PaymentReference,omitempty"`

	// RateFC: Exchange rate from receivable currency to division currency. AmountFC * RateFC = AmountDC.
	RateFC *float64 `json:"RateFC,omitempty"`

	// ReceivableBatchNumber: Number assigned during the processing of receivables.
	ReceivableBatchNumber *int `json:"ReceivableBatchNumber,omitempty"`

	// ReceivableSelected: Date and time since when the receivable is selected to be collected.
	ReceivableSelected *types.Date `json:"ReceivableSelected,omitempty"`

	// ReceivableSelector: User who selected the receivable to be collected.
	ReceivableSelector *types.GUID `json:"ReceivableSelector,omitempty"`

	// ReceivableSelectorFullName: Name of the receivable selector.
	ReceivableSelectorFullName *string `json:"ReceivableSelectorFullName,omitempty"`

	// Source: The source of the receivable. 1 = manual 2 = reconcile 3 = match 4 = import 5 = process
	Source *int `json:"Source,omitempty"`

	// Status: The status of the receivable. 20 = open 30 = selected - receivable is selected to be collected 40 = processed - collection has been done 50 = matched - receivable is matched with one or more other outstanding items or financial statement lines
	Status *int `json:"Status,omitempty"`

	// TransactionAmountDC: Total amount of the linked transaction in default currency (division currency).
	TransactionAmountDC *float64 `json:"TransactionAmountDC,omitempty"`

	// TransactionAmountFC: Total amount of the linked transaction in the selected currency.
	TransactionAmountFC *float64 `json:"TransactionAmountFC,omitempty"`

	// TransactionDueDate: Due date of the linked transaction.
	TransactionDueDate *types.Date `json:"TransactionDueDate,omitempty"`

	// TransactionEntryID: Linked transaction. Use this as reference to SalesEntries.
	TransactionEntryID *types.GUID `json:"TransactionEntryID,omitempty"`

	// TransactionID: Linked transaction line. Use this as reference to BankEntryLines and CashEntryLines.
	TransactionID *types.GUID `json:"TransactionID,omitempty"`

	// TransactionIsReversal: Indicates if the linked transaction is a reversal entry.
	TransactionIsReversal *bool `json:"TransactionIsReversal,omitempty"`

	// TransactionReportingPeriod: Period of the linked transaction.
	TransactionReportingPeriod *int `json:"TransactionReportingPeriod,omitempty"`

	// TransactionReportingYear: Year of the linked transaction.
	TransactionReportingYear *int `json:"TransactionReportingYear,omitempty"`

	// TransactionStatus: Status of the linked transaction.
	TransactionStatus *int `json:"TransactionStatus,omitempty"`

	// TransactionType: Type of the linked transaction.
	TransactionType *int `json:"TransactionType,omitempty"`

	// YourRef: Invoice number. In case the receivable belongs to a bank entry line and is matched with one invoice, YourRef is filled with the YourRef of this invoice.
	YourRef *string `json:"YourRef,omitempty"`
}

// List the Receivables entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ReceivablesEndpoint) List(ctx context.Context, division int, all bool) ([]*Receivables, error) {
	var entities []*Receivables
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/cashflow/Receivables?$select=*", division) // #nosec
	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
