// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package continuousmonitoring

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// IndicatorGLAccountsEndpoint is responsible for communicating with
// the IndicatorGLAccounts endpoint of the ContinuousMonitoring service.
type IndicatorGLAccountsEndpoint service

// IndicatorGLAccounts:
// Service: ContinuousMonitoring
// Entity: IndicatorGLAccounts
// URL: /api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts
// HasWebhook: false
// IsInBeta: true
// Methods: GET POST DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ContinuousMonitoringIndicatorGLAccounts
type IndicatorGLAccounts struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// GLAccount: ID of GLAccount
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// GLAccountCode: GL account code
	GLAccountCode *string `json:"GLAccountCode,omitempty"`

	// GLAccountDescription: Description of GLAccount
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// Indicator: ID of Indicators
	Indicator *types.GUID `json:"Indicator,omitempty"`
}

func (e *IndicatorGLAccounts) GetPrimary() *types.GUID {
	return e.ID
}

func (s *IndicatorGLAccountsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "continuousmonitoring/IndicatorGLAccounts", method)
}

// List the IndicatorGLAccounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *IndicatorGLAccountsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*IndicatorGLAccounts, error) {
	var entities []*IndicatorGLAccounts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the IndicatorGLAccounts entitiy in the provided division.
func (s *IndicatorGLAccountsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*IndicatorGLAccounts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &IndicatorGLAccounts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty IndicatorGLAccounts entity
func (s *IndicatorGLAccountsEndpoint) New() *IndicatorGLAccounts {
	return &IndicatorGLAccounts{}
}

// Create the IndicatorGLAccounts entity in the provided division.
func (s *IndicatorGLAccountsEndpoint) Create(ctx context.Context, division int, entity *IndicatorGLAccounts) (*IndicatorGLAccounts, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts", division) // #nosec
	e := &IndicatorGLAccounts{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Delete the IndicatorGLAccounts entity in the provided division.
func (s *IndicatorGLAccountsEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorGLAccounts", division) // #nosec
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
