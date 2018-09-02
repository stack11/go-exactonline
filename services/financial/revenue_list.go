// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import "context"

// RevenueListEndpoint is responsible for communicating with
// the RevenueList endpoint of the Financial service.
type RevenueListEndpoint service

// RevenueList:
// Service: Financial
// Entity: RevenueList
// URL: /api/v1/{division}/read/financial/RevenueList
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialRevenueList
type RevenueList struct {
	// Period: Reporting period
	Period *int `json:"Period,omitempty"`

	// Year: Current Reporting year
	Year *int `json:"Year,omitempty"`

	// Amount: Total amount in the default currency of the company
	Amount *float64 `json:"Amount,omitempty"`
}

// List the RevenueList entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *RevenueListEndpoint) List(ctx context.Context, division int, all bool) ([]*RevenueList, error) {
	var entities []*RevenueList
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/RevenueList?$select=*", division) // #nosec
	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
