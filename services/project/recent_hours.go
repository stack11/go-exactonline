// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// RecentHoursEndpoint is responsible for communicating with
// the RecentHours endpoint of the Project service.
type RecentHoursEndpoint service

// RecentHours:
// Service: Project
// Entity: RecentHours
// URL: /api/v1/{division}/read/project/RecentHours
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectRecentHours
type RecentHours struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Id: Primary key
	Id *int `json:"Id,omitempty"`

	// AccountCode: Code of account linked to the project that hours are being entered to
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountId: ID of account linked to the project that hours are being entered to
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountName: Name of account linked to the project that hours are being entered to
	AccountName *string `json:"AccountName,omitempty"`

	// Activity: The activity of the project that the hours are entered to
	Activity *types.GUID `json:"Activity,omitempty"`

	// ActivityDescription: Name of activity of the project that the hours are entered to
	ActivityDescription *string `json:"ActivityDescription,omitempty"`

	// Date: Date of hour entry record by week
	Date *types.Date `json:"Date,omitempty"`

	// EntryId: Entry ID of record
	EntryId *types.GUID `json:"EntryId,omitempty"`

	// HoursApproved: Hours approved
	HoursApproved *float64 `json:"HoursApproved,omitempty"`

	// HoursApprovedBillable: Billable hours that is approved
	HoursApprovedBillable *float64 `json:"HoursApprovedBillable,omitempty"`

	// HoursDraft: Hours saved as draft
	HoursDraft *float64 `json:"HoursDraft,omitempty"`

	// HoursDraftBillable: Billable hours saved as draft
	HoursDraftBillable *float64 `json:"HoursDraftBillable,omitempty"`

	// HoursRejected: Hours that are rejected
	HoursRejected *float64 `json:"HoursRejected,omitempty"`

	// HoursRejectedBillable: Billable hours that are rejected
	HoursRejectedBillable *float64 `json:"HoursRejectedBillable,omitempty"`

	// HoursSubmitted: Hours that are submitted
	HoursSubmitted *float64 `json:"HoursSubmitted,omitempty"`

	// HoursSubmittedBillable: Billable hours that are submitted
	HoursSubmittedBillable *float64 `json:"HoursSubmittedBillable,omitempty"`

	// ItemCode: Code of the item that is used for hours
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of the item that is used for hours
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemId: ID of the item that is used for hours
	ItemId *types.GUID `json:"ItemId,omitempty"`

	// Notes: Notes entered regarding the information of the hours entered
	Notes *string `json:"Notes,omitempty"`

	// ProjectCode: Code of project that the hours are entered on
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Description of project that the hours are entered on
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectId: ID of project that the hours are entered on
	ProjectId *types.GUID `json:"ProjectId,omitempty"`

	// WeekNumber: The week number that the hours are entered on
	WeekNumber *int `json:"WeekNumber,omitempty"`
}

func (e *RecentHours) GetPrimary() *int {
	return e.Id
}

func (s *RecentHoursEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "project/RecentHours", method)
}

// List the RecentHours entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *RecentHoursEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*RecentHours, error) {
	var entities []*RecentHours
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/RecentHours", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the RecentHours entitiy in the provided division.
func (s *RecentHoursEndpoint) Get(ctx context.Context, division int, id *int) (*RecentHours, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/RecentHours", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &RecentHours{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
