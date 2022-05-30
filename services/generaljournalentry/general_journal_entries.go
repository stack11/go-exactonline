// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package generaljournalentry

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// GeneralJournalEntriesEndpoint is responsible for communicating with
// the GeneralJournalEntries endpoint of the GeneralJournalEntry service.
type GeneralJournalEntriesEndpoint service

// GeneralJournalEntries:
// Service: GeneralJournalEntry
// Entity: GeneralJournalEntries
// URL: /api/v1/{division}/generaljournalentry/GeneralJournalEntries
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=GeneralJournalEntryGeneralJournalEntries
type GeneralJournalEntries struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// EntryID:
	EntryID *types.GUID `json:"EntryID,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// EntryNumber:
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// ExchangeRate:
	ExchangeRate *float64 `json:"ExchangeRate,omitempty"`

	// FinancialPeriod:
	FinancialPeriod *int `json:"FinancialPeriod,omitempty"`

	// FinancialYear:
	FinancialYear *int `json:"FinancialYear,omitempty"`

	// GeneralJournalEntryLines:
	GeneralJournalEntryLines *json.RawMessage `json:"GeneralJournalEntryLines,omitempty"`

	// JournalCode:
	JournalCode *string `json:"JournalCode,omitempty"`

	// JournalDescription:
	JournalDescription *string `json:"JournalDescription,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Reversal:
	Reversal *bool `json:"Reversal,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// StatusDescription:
	StatusDescription *string `json:"StatusDescription,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`

	// TypeDescription:
	TypeDescription *string `json:"TypeDescription,omitempty"`
}

func (e *GeneralJournalEntries) GetPrimary() *types.GUID {
	return e.EntryID
}

func (s *GeneralJournalEntriesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "generaljournalentry/GeneralJournalEntries", method)
}

// List the GeneralJournalEntries entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *GeneralJournalEntriesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*GeneralJournalEntries, error) {
	var entities []*GeneralJournalEntries
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/generaljournalentry/GeneralJournalEntries", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the GeneralJournalEntries entitiy in the provided division.
func (s *GeneralJournalEntriesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*GeneralJournalEntries, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/generaljournalentry/GeneralJournalEntries", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &GeneralJournalEntries{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty GeneralJournalEntries entity
func (s *GeneralJournalEntriesEndpoint) New() *GeneralJournalEntries {
	return &GeneralJournalEntries{}
}

// Create the GeneralJournalEntries entity in the provided division.
func (s *GeneralJournalEntriesEndpoint) Create(ctx context.Context, division int, entity *GeneralJournalEntries) (*GeneralJournalEntries, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/generaljournalentry/GeneralJournalEntries", division) // #nosec
	e := &GeneralJournalEntries{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Delete the GeneralJournalEntries entity in the provided division.
func (s *GeneralJournalEntriesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/generaljournalentry/GeneralJournalEntries", division) // #nosec
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
