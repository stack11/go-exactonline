// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package hrm

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// SchedulesEndpoint is responsible for communicating with
// the Schedules endpoint of the HRM service.
type SchedulesEndpoint service

// Schedules:
// Service: HRM
// Entity: Schedules
// URL: /api/v1/{division}/hrm/Schedules
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMSchedules
type Schedules struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Active: Obsolete
	Active *byte `json:"Active,omitempty"`

	// AverageHours: Average hours per week in a schedule
	AverageHours *float64 `json:"AverageHours,omitempty"`

	// BillabilityTarget: Billability target
	BillabilityTarget *float64 `json:"BillabilityTarget,omitempty"`

	// Code: Schedule code
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Days: Average days per week in the schedule
	Days *float64 `json:"Days,omitempty"`

	// Description: Description of the schedule
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employment: Employment ID for schedule
	Employment *types.GUID `json:"Employment,omitempty"`

	// EmploymentHID: Employment number
	EmploymentHID *int `json:"EmploymentHID,omitempty"`

	// EndDate: End date of the schedule
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Hours: Number of hours per week in a CLA for which the schedule is built
	Hours *float64 `json:"Hours,omitempty"`

	// LeaveHoursCompensation: Number of hours which are built up each week for later leave
	LeaveHoursCompensation *float64 `json:"LeaveHoursCompensation,omitempty"`

	// Main: Indication if the schedule is a main schedule for a CLA. 1 = Yes, 0 = No
	Main *byte `json:"Main,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// PaymentParttimeFactor: Part-time factor for payroll calculation. Value between 0 and 1
	PaymentParttimeFactor *float64 `json:"PaymentParttimeFactor,omitempty"`

	// ScheduleType: Type of schedule. 1 = Hours and average days, 2 = Hours and specific days, 3 = Hours per day, 4 = Time frames per day
	ScheduleType *int `json:"ScheduleType,omitempty"`

	// ScheduleTypeDescription: Description of the schedule type
	ScheduleTypeDescription *string `json:"ScheduleTypeDescription,omitempty"`

	// StartDate: Week in the schedule which is used to start with. By default the number will be 1.
	StartDate *types.Date `json:"StartDate,omitempty"`

	// StartWeek: Week to start the schedule from for an employee
	StartWeek *int `json:"StartWeek,omitempty"`
}

func (e *Schedules) GetPrimary() *types.GUID {
	return e.ID
}

func (s *SchedulesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "hrm/Schedules", method)
}

// List the Schedules entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SchedulesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Schedules, error) {
	var entities []*Schedules
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Schedules", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Schedules entitiy in the provided division.
func (s *SchedulesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*Schedules, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/Schedules", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Schedules{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
