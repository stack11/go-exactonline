// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package subscription

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// SubscriptionTypesEndpoint is responsible for communicating with
// the SubscriptionTypes endpoint of the Subscription service.
type SubscriptionTypesEndpoint service

// SubscriptionTypes:
// Service: Subscription
// Entity: SubscriptionTypes
// URL: /api/v1/{division}/subscription/SubscriptionTypes
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SubscriptionSubscriptionTypes
type SubscriptionTypes struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AutomaticGenerateInvoiceDays: Number of days before or after generating the subscription invoice. Company settings and automatic  // 						generate invoice type need to be enabled before subscription invoice generated automatically
	AutomaticGenerateInvoiceDays *int `json:"AutomaticGenerateInvoiceDays,omitempty"`

	// AutomaticGenerateInvoiceDescription: Description of the automatic generated invoice
	AutomaticGenerateInvoiceDescription *string `json:"AutomaticGenerateInvoiceDescription,omitempty"`

	// AutomaticGenerateInvoiceType: Type of automatic generate invoice: 1=Never, 2=Before the subscription period, 3=After the subscription period.  // 						Company settings need to be enabled before subscription invoice generated automatically
	AutomaticGenerateInvoiceType *int `json:"AutomaticGenerateInvoiceType,omitempty"`

	// AutomaticSendInvoiceDays: Number of days after sending the subscription invoice. Company settings and automatic  // 						sending invoice type need to be enabled before subscription invoice sent automatically
	AutomaticSendInvoiceDays *int `json:"AutomaticSendInvoiceDays,omitempty"`

	// AutomaticSendInvoiceMethod: Method of automatic send invoice: 1=Send based on account, 2=Send via email, 3=Create documents, 4=Send via digital postbox,  // 						5=Send and track, 6=Send via peppol. Company settings need to be enabled before subscription invoice sent automatically
	AutomaticSendInvoiceMethod *int `json:"AutomaticSendInvoiceMethod,omitempty"`

	// AutomaticSendInvoiceSender: Sender&#39;s email of automatic send invoice: 1=Company email address, 2=Main user email address.  // 						Company settings need to be enabled before subscription invoice sent automatically
	AutomaticSendInvoiceSender *int `json:"AutomaticSendInvoiceSender,omitempty"`

	// AutomaticSendInvoiceSenderMailbox: ID of automatic send invoice sender&#39;s mailbox. Company settings need to be enabled before subscription invoice sent automatically
	AutomaticSendInvoiceSenderMailbox *types.GUID `json:"AutomaticSendInvoiceSenderMailbox,omitempty"`

	// AutomaticSendInvoiceType: Type of automatic send invoice: 1=Never, 2=When available.  // 						Company settings need to be enabled before subscription invoice sent automatically
	AutomaticSendInvoiceType *int `json:"AutomaticSendInvoiceType,omitempty"`

	// CancellationPeriod: Cancellation period of subscription type
	CancellationPeriod *int `json:"CancellationPeriod,omitempty"`

	// CancellationPeriodUnit: Unit of cancellation period: wk=Week, mm=Month, yy=Year, hy=Half-year, qt=Quarter
	CancellationPeriodUnit *string `json:"CancellationPeriodUnit,omitempty"`

	// Code: Code
	Code *string `json:"Code,omitempty"`

	// Created: Date and time when the subscription type was created
	Created *types.Date `json:"Created,omitempty"`

	// Creator: ID of user that created the subscription type
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Full name of user that created the subscription type
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// CustomField: Custom field endpoint. Provided only for the Exact Online Premium users.
	CustomField *string `json:"CustomField,omitempty"`

	// Description: Description of subscription type
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EnablePaymentLink: Enable payment link: 0=Never, 1=Always, 2=Based on account
	EnablePaymentLink *int `json:"EnablePaymentLink,omitempty"`

	// InvoiceCorrectionMethod: Invoice correction method: 1=Ratio based, 2=Zero Invoice, 3=Never invoiced
	InvoiceCorrectionMethod *int `json:"InvoiceCorrectionMethod,omitempty"`

	// InvoicePeriod: Invoice period of subscription type
	InvoicePeriod *int `json:"InvoicePeriod,omitempty"`

	// InvoicePeriodUnit: Unit of invoice period: wk=Week, mm=Month, yy=Year, hy=Half-year, qt=Quarter
	InvoicePeriodUnit *string `json:"InvoicePeriodUnit,omitempty"`

	// ManualRenewalMethod: Manual renewal method: 1=Use item prices, 2=Use current subscription prices
	ManualRenewalMethod *int `json:"ManualRenewalMethod,omitempty"`

	// Modified: Last modified date of subscription type
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: ID of user that modified the subscription type
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Full name of user that modified the subscription type
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Additional information about subscription type
	Notes *string `json:"Notes,omitempty"`

	// ProlongationType: Prolongation type: 0=No, 1=Manual, 2=Automatic
	ProlongationType *byte `json:"ProlongationType,omitempty"`

	// RenewalCancellationPeriod: Renewal cancellation period of subscription type
	RenewalCancellationPeriod *int `json:"RenewalCancellationPeriod,omitempty"`

	// RenewalCancellationPeriodUnit: Unit of renewal cancellation period: wk=Week, mm=Month, yy=Year, hy=Half-year, qt=Quarter
	RenewalCancellationPeriodUnit *string `json:"RenewalCancellationPeriodUnit,omitempty"`

	// RenewalPeriod: Renewal period of subscription type
	RenewalPeriod *int `json:"RenewalPeriod,omitempty"`

	// RenewalPeriodUnit: Unit of renewal period: wk=Week, mm=Month, yy=Year, hy=Half-year, qt=Quarter
	RenewalPeriodUnit *string `json:"RenewalPeriodUnit,omitempty"`

	// SubscriptionPeriod: Subscription period of subscription type
	SubscriptionPeriod *int `json:"SubscriptionPeriod,omitempty"`

	// SubscriptionPeriodUnit: Unit of subscription period: wk=Week, mm=Month, yy=Year, hy=Half-year, qt=Quarter
	SubscriptionPeriodUnit *string `json:"SubscriptionPeriodUnit,omitempty"`
}

func (e *SubscriptionTypes) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SubscriptionTypesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "subscription/SubscriptionTypes", method)
}

// List the SubscriptionTypes entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SubscriptionTypesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*SubscriptionTypes, error) {
	var entities []*SubscriptionTypes
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionTypes", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the SubscriptionTypes entitiy in the provided division.
func (s *SubscriptionTypesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*SubscriptionTypes, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionTypes", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &SubscriptionTypes{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
