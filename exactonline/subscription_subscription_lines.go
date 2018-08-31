// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// SubscriptionSubscriptionLinesService is responsible for communicating with
// the SubscriptionLines endpoint of the Subscription service.
type SubscriptionSubscriptionLinesService service

// SubscriptionSubscriptionLines:
// Service: Subscription
// Entity: SubscriptionLines
// URL: /api/v1/{division}/subscription/SubscriptionLines
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SubscriptionSubscriptionLines
type SubscriptionSubscriptionLines struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// AmountDC: Amount in the default currency of the company
	AmountDC *float64 `json:",omitempty"`

	// AmountFC: Amount in the currency of the transaction
	AmountFC *float64 `json:",omitempty"`

	// Costcenter: Cost center
	Costcenter *string `json:",omitempty"`

	// Costunit: Cost unit
	Costunit *string `json:",omitempty"`

	// Description: Description
	Description *string `json:",omitempty"`

	// Discount: Discount percentage
	Discount *float64 `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// EntryID: Entry ID
	EntryID *GUID `json:",omitempty"`

	// FromDate: From date
	FromDate *Date `json:",omitempty"`

	// Item: Reference to Item
	Item *GUID `json:",omitempty"`

	// ItemDescription: Description of Item
	ItemDescription *string `json:",omitempty"`

	// LineNumber: Line number
	LineNumber *int `json:",omitempty"`

	// LineType: Reference to LineType
	LineType *int `json:",omitempty"`

	// LineTypeDescription: Description of LineType
	LineTypeDescription *string `json:",omitempty"`

	// NetPrice: Net price in the currency of the transaction
	NetPrice *float64 `json:",omitempty"`

	// Notes: Remarks
	Notes *string `json:",omitempty"`

	// Quantity: Quantity
	Quantity *float64 `json:",omitempty"`

	// ToDate: To date
	ToDate *Date `json:",omitempty"`

	// UnitCode: Unit code
	UnitCode *string `json:",omitempty"`

	// UnitDescription: Description of Unit
	UnitDescription *string `json:",omitempty"`

	// UnitPrice: Unit price in the currency of the transaction (price * unit factor)
	UnitPrice *float64 `json:",omitempty"`

	// VATAmountFC: Vat Amount in the currency of the transaction
	VATAmountFC *float64 `json:",omitempty"`

	// VATCode: VATCode
	VATCode *string `json:",omitempty"`

	// VATCodeDescription: Description of VATCode
	VATCodeDescription *string `json:",omitempty"`
}

func (s *SubscriptionSubscriptionLines) GetIdentifier() GUID {
	return *s.ID
}

// List the SubscriptionLines entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SubscriptionSubscriptionLinesService) List(ctx context.Context, division int, all bool) ([]*SubscriptionSubscriptionLines, error) {
	var entities []*SubscriptionSubscriptionLines
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/SubscriptionLines?$select=*", division)
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
