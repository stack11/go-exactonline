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

// LeaveRegistrationsEndpoint is responsible for communicating with
// the LeaveRegistrations endpoint of the HRM service.
type LeaveRegistrationsEndpoint service

// LeaveRegistrations:
// Service: HRM
// Entity: LeaveRegistrations
// URL: /api/v1/{division}/hrm/LeaveRegistrations
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMLeaveRegistrations
type LeaveRegistrations struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of leave
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee linked to the leave registration
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Employee full name
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Numeric ID of the employee
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// EndDate: End date of leave
	EndDate *types.Date `json:"EndDate,omitempty"`

	// EndTime: End time on the last day of leave stored as DateTime, and the date should be ignored
	EndTime *types.Date `json:"EndTime,omitempty"`

	// Hours: Total number of leave hours
	Hours *float64 `json:"Hours,omitempty"`

	// HoursFirstDay: Hours of leave on the first day
	HoursFirstDay *float64 `json:"HoursFirstDay,omitempty"`

	// HoursLastDay: Hours of leave on the last day
	HoursLastDay *float64 `json:"HoursLastDay,omitempty"`

	// LeaveType: Type of leave
	LeaveType *types.GUID `json:"LeaveType,omitempty"`

	// LeaveTypeCode: Code for type of leave
	LeaveTypeCode *string `json:"LeaveTypeCode,omitempty"`

	// LeaveTypeDescription: Description for type of leave
	LeaveTypeDescription *string `json:"LeaveTypeDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Extra information for leave
	Notes *string `json:"Notes,omitempty"`

	// StartDate: Start date of leave
	StartDate *types.Date `json:"StartDate,omitempty"`

	// StartTime: Start time on the first day of leave stored as DateTime, and the date should be ignored
	StartTime *types.Date `json:"StartTime,omitempty"`

	// Status: Status of leave, 1 = Submitted, 2 = Approved, 3 = Rejected
	Status *int `json:"Status,omitempty"`
}

func (e *LeaveRegistrations) GetPrimary() *types.GUID {
	return e.ID
}

func (s *LeaveRegistrationsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "hrm/LeaveRegistrations", method)
}

// List the LeaveRegistrations entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *LeaveRegistrationsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*LeaveRegistrations, error) {
	var entities []*LeaveRegistrations
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/LeaveRegistrations", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the LeaveRegistrations entitiy in the provided division.
func (s *LeaveRegistrationsEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*LeaveRegistrations, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/LeaveRegistrations", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &LeaveRegistrations{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
