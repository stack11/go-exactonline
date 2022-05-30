// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package mailbox

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// PreferredMailboxEndpoint is responsible for communicating with
// the PreferredMailbox endpoint of the Mailbox service.
type PreferredMailboxEndpoint service

// PreferredMailbox:
// Service: Mailbox
// Entity: PreferredMailbox
// URL: /api/v1/{division}/read/mailbox/PreferredMailbox
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadMailboxPreferredMailbox
type PreferredMailbox struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// Description: Extra description of the mailbox
	Description *string `json:"Description,omitempty"`

	// ForDivision: Only used when this mailbox is used for one specific administration, for example invoices to this mailbox will only be booked in this administration
	ForDivision *int `json:"ForDivision,omitempty"`

	// IsScanServiceMailbox: Indicates whether this service is used for messages returned by the scanning service
	IsScanServiceMailbox *bool `json:"IsScanServiceMailbox,omitempty"`

	// Mailbox: E-mail address-like format, for example johndoe@exactonline.nl
	Mailbox *string `json:"Mailbox,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ValidFrom: Date that this mailbox became valid
	ValidFrom *types.Date `json:"ValidFrom,omitempty"`

	// ValidTo: Date that this mailbox will not be valid anymore
	ValidTo *types.Date `json:"ValidTo,omitempty"`
}

func (e *PreferredMailbox) GetPrimary() *types.GUID {
	return e.ID
}

func (s *PreferredMailboxEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "mailbox/PreferredMailbox", method)
}

// List the PreferredMailbox entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *PreferredMailboxEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*PreferredMailbox, error) {
	var entities []*PreferredMailbox
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/mailbox/PreferredMailbox", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the PreferredMailbox entitiy in the provided division.
func (s *PreferredMailboxEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*PreferredMailbox, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/mailbox/PreferredMailbox", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &PreferredMailbox{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
