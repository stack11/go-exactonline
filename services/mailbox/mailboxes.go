// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package mailbox

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// MailboxesEndpoint is responsible for communicating with
// the Mailboxes endpoint of the Mailbox service.
type MailboxesEndpoint service

// Mailboxes:
// Service: Mailbox
// Entity: Mailboxes
// URL: /api/v1/{division}/mailbox/Mailboxes
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=MailboxMailboxes
type Mailboxes struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: The account this mailbox belongs to. Can be empty if the owner of the mailbox isn&#39;t an Exact Online customer yet
	Account *types.GUID `json:"Account,omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:"AccountName,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Extra description of the mailbox
	Description *string `json:"Description,omitempty"`

	// ForDivision: Only used when this mailbox is used for one specific administration, for example invoices to this mailbox will only be booked in this administration
	ForDivision *int `json:"ForDivision,omitempty"`

	// ForDivisionDescription: Description of ForDivision
	ForDivisionDescription *string `json:"ForDivisionDescription,omitempty"`

	// Mailbox: E-mail address-like format, for example johndoe@exactonline.nl
	Mailbox *string `json:"Mailbox,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Publish: Customers can decide if they want this mailbox to be visible by all. i.e. some other customer can see this in address maintenance for digital postbox of type Exact
	Publish *byte `json:"Publish,omitempty"`

	// Type: Type of mailbox: 0-Unknown, 1-Exact, 2-Government, 3-Manual input
	Type *int `json:"Type,omitempty"`

	// ValidFrom: Date that this mailbox became valid
	ValidFrom *types.Date `json:"ValidFrom,omitempty"`

	// ValidTo: Date that this mailbox will not be valid anymore
	ValidTo *types.Date `json:"ValidTo,omitempty"`
}

func (e *Mailboxes) GetPrimary() *types.GUID {
	return e.ID
}

func (s *MailboxesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "mailbox/Mailboxes", method)
}

// List the Mailboxes entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *MailboxesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Mailboxes, error) {
	var entities []*Mailboxes
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/Mailboxes", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Mailboxes entitiy in the provided division.
func (s *MailboxesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*Mailboxes, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/Mailboxes", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Mailboxes{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty Mailboxes entity
func (s *MailboxesEndpoint) New() *Mailboxes {
	return &Mailboxes{}
}

// Create the Mailboxes entity in the provided division.
func (s *MailboxesEndpoint) Create(ctx context.Context, division int, entity *Mailboxes) (*Mailboxes, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/Mailboxes", division) // #nosec
	e := &Mailboxes{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the Mailboxes entity in the provided division.
func (s *MailboxesEndpoint) Update(ctx context.Context, division int, entity *Mailboxes) (*Mailboxes, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/Mailboxes", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &Mailboxes{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the Mailboxes entity in the provided division.
func (s *MailboxesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/Mailboxes", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return err
	}

	_, r, requestError := s.client.NewRequestAndDo(ctx, "DELETE", u.String(), nil, nil)
	if requestError != nil {
		return requestError
	}

	if r.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(r.Body) // #nosec
		return fmt.Errorf("Failed with status %v and body %v", r.StatusCode, body)
	}

	return nil
}
