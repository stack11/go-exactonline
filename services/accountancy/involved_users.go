// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package accountancy

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// InvolvedUsersEndpoint is responsible for communicating with
// the InvolvedUsers endpoint of the Accountancy service.
type InvolvedUsersEndpoint service

// InvolvedUsers:
// Service: Accountancy
// Entity: InvolvedUsers
// URL: /api/v1/{division}/accountancy/InvolvedUsers
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUsers
type InvolvedUsers struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: ID of the account the user is involved with
	Account *types.GUID `json:"Account,omitempty"`

	// AccountCity: City of the account
	AccountCity *string `json:"AccountCity,omitempty"`

	// AccountCode: Code of the account
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountIsSupplier: Supplier flag of the account
	AccountIsSupplier *bool `json:"AccountIsSupplier,omitempty"`

	// AccountLogoThumbnailUrl: Logo thumbnail url of the account
	AccountLogoThumbnailUrl *string `json:"AccountLogoThumbnailUrl,omitempty"`

	// AccountName: Name of the account
	AccountName *string `json:"AccountName,omitempty"`

	// AccountStatus: Status of the account
	AccountStatus *string `json:"AccountStatus,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// InvolvedUserRole: ID of the user role
	InvolvedUserRole *types.GUID `json:"InvolvedUserRole,omitempty"`

	// InvolvedUserRoleDescription: Description of the user role
	InvolvedUserRoleDescription *string `json:"InvolvedUserRoleDescription,omitempty"`

	// IsMainContact: Main contact flag of the involved user
	IsMainContact *bool `json:"IsMainContact,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// PersonEmail: Email of a person
	PersonEmail *string `json:"PersonEmail,omitempty"`

	// PersonPhone: Phone of a person
	PersonPhone *string `json:"PersonPhone,omitempty"`

	// PersonPhoneExtension: Phone extension of a person
	PersonPhoneExtension *string `json:"PersonPhoneExtension,omitempty"`

	// PersonPictureThumbnailUrl: Picture thumbnail url of a person
	PersonPictureThumbnailUrl *string `json:"PersonPictureThumbnailUrl,omitempty"`

	// User: ID of the involved user
	User *types.GUID `json:"User,omitempty"`

	// UserFullName: User name of creator
	UserFullName *string `json:"UserFullName,omitempty"`
}

// List the InvolvedUsers entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InvolvedUsersEndpoint) List(ctx context.Context, division int, all bool) ([]*InvolvedUsers, error) {
	var entities []*InvolvedUsers
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/accountancy/InvolvedUsers?$select=*", division) // #nosec
	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err := s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
