// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package activities

import (
	"context"
	"encoding/json"

	"github.com/mcnijman/go-exactonline/api"
	"github.com/mcnijman/go-exactonline/types"
)

// ComplaintsEndpoint is responsible for communicating with
// the Complaints endpoint of the Activities service.
type ComplaintsEndpoint service

// Complaints:
// Service: Activities
// Entity: Complaints
// URL: /api/v1/{division}/activities/Complaints
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesComplaints
type Complaints struct {
	// ID: The Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: The account that is related to the complaint
	Account *types.GUID `json:"Account,omitempty"`

	// AccountName: The name of the account
	AccountName *string `json:"AccountName,omitempty"`

	// AssignedTo: The user that the complaint is assigned to
	AssignedTo *types.GUID `json:"AssignedTo,omitempty"`

	// AssignedToFullName: The user name
	AssignedToFullName *string `json:"AssignedToFullName,omitempty"`

	// Attachments: Attachments linked to the complaint
	Attachments *json.RawMessage `json:"Attachments,omitempty"`

	// Complaint: A short description of the complaint
	Complaint *string `json:"Complaint,omitempty"`

	// Contact: The contact person that is related to the complaint
	Contact *types.GUID `json:"Contact,omitempty"`

	// ContactFullName: The name of the contact person
	ContactFullName *string `json:"ContactFullName,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of the creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: The division
	Division *int `json:"Division,omitempty"`

	// Document: The document that is linked to the complaint
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentSubject: The subject of the document
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// HID: The human readable key
	HID *int `json:"HID,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of the last modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// NextAction: The date indicating by when the next action has to be taken
	NextAction *types.Date `json:"NextAction,omitempty"`

	// Notes: The notes of the complaint
	Notes *string `json:"Notes,omitempty"`

	// ReceiptDate: The date the complaint was received
	ReceiptDate *types.Date `json:"ReceiptDate,omitempty"`

	// Status: Status: 0 = Void, 5 = Rejected, 10 = Draft, 20 = Open, 30 = Approved, 40 = Realized, 50 = Processed
	Status *int `json:"Status,omitempty"`

	// StatusDescription: The description of the status
	StatusDescription *string `json:"StatusDescription,omitempty"`
}

// List the Complaints entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ComplaintsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Complaints, error) {
	var entities []*Complaints
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/activities/Complaints", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
