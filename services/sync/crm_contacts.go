// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package sync

import (
	"context"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// CRMContactsEndpoint is responsible for communicating with
// the CRMContacts endpoint of the Sync service.
type CRMContactsEndpoint service

// CRMContacts:
// Service: Sync
// Entity: CRMContacts
// URL: /api/v1/{division}/sync/CRM/Contacts
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncCRMContacts
type CRMContacts struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp:
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// Account:
	Account *types.GUID `json:"Account,omitempty"`

	// AccountIsCustomer:
	AccountIsCustomer *bool `json:"AccountIsCustomer,omitempty"`

	// AccountIsSupplier:
	AccountIsSupplier *bool `json:"AccountIsSupplier,omitempty"`

	// AccountMainContact:
	AccountMainContact *types.GUID `json:"AccountMainContact,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// AddressLine2:
	AddressLine2 *string `json:"AddressLine2,omitempty"`

	// AddressStreet:
	AddressStreet *string `json:"AddressStreet,omitempty"`

	// AddressStreetNumber:
	AddressStreetNumber *string `json:"AddressStreetNumber,omitempty"`

	// AddressStreetNumberSuffix:
	AddressStreetNumberSuffix *string `json:"AddressStreetNumberSuffix,omitempty"`

	// AllowMailing:
	AllowMailing *int `json:"AllowMailing,omitempty"`

	// BirthDate:
	BirthDate *types.Date `json:"BirthDate,omitempty"`

	// BirthName:
	BirthName *string `json:"BirthName,omitempty"`

	// BirthNamePrefix:
	BirthNamePrefix *string `json:"BirthNamePrefix,omitempty"`

	// BirthPlace:
	BirthPlace *string `json:"BirthPlace,omitempty"`

	// BusinessEmail:
	BusinessEmail *string `json:"BusinessEmail,omitempty"`

	// BusinessFax:
	BusinessFax *string `json:"BusinessFax,omitempty"`

	// BusinessMobile:
	BusinessMobile *string `json:"BusinessMobile,omitempty"`

	// BusinessPhone:
	BusinessPhone *string `json:"BusinessPhone,omitempty"`

	// BusinessPhoneExtension:
	BusinessPhoneExtension *string `json:"BusinessPhoneExtension,omitempty"`

	// City:
	City *string `json:"City,omitempty"`

	// Code:
	Code *string `json:"Code,omitempty"`

	// Country:
	Country *string `json:"Country,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// CustomField:
	CustomField *string `json:"CustomField,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Email:
	Email *string `json:"Email,omitempty"`

	// EndDate:
	EndDate *types.Date `json:"EndDate,omitempty"`

	// FirstName:
	FirstName *string `json:"FirstName,omitempty"`

	// FullName:
	FullName *string `json:"FullName,omitempty"`

	// Gender:
	Gender *string `json:"Gender,omitempty"`

	// HID:
	HID *int `json:"HID,omitempty"`

	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// IdentificationDate:
	IdentificationDate *types.Date `json:"IdentificationDate,omitempty"`

	// IdentificationDocument:
	IdentificationDocument *types.GUID `json:"IdentificationDocument,omitempty"`

	// IdentificationUser:
	IdentificationUser *types.GUID `json:"IdentificationUser,omitempty"`

	// Initials:
	Initials *string `json:"Initials,omitempty"`

	// IsAnonymised:
	IsAnonymised *byte `json:"IsAnonymised,omitempty"`

	// IsMailingExcluded:
	IsMailingExcluded *bool `json:"IsMailingExcluded,omitempty"`

	// IsMainContact:
	IsMainContact *bool `json:"IsMainContact,omitempty"`

	// JobTitleDescription:
	JobTitleDescription *string `json:"JobTitleDescription,omitempty"`

	// Language:
	Language *string `json:"Language,omitempty"`

	// LastName:
	LastName *string `json:"LastName,omitempty"`

	// LeadPurpose:
	LeadPurpose *types.GUID `json:"LeadPurpose,omitempty"`

	// LeadSource:
	LeadSource *types.GUID `json:"LeadSource,omitempty"`

	// MarketingNotes:
	MarketingNotes *string `json:"MarketingNotes,omitempty"`

	// MiddleName:
	MiddleName *string `json:"MiddleName,omitempty"`

	// Mobile:
	Mobile *string `json:"Mobile,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Nationality:
	Nationality *string `json:"Nationality,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// PartnerName:
	PartnerName *string `json:"PartnerName,omitempty"`

	// PartnerNamePrefix:
	PartnerNamePrefix *string `json:"PartnerNamePrefix,omitempty"`

	// Person:
	Person *types.GUID `json:"Person,omitempty"`

	// Phone:
	Phone *string `json:"Phone,omitempty"`

	// PhoneExtension:
	PhoneExtension *string `json:"PhoneExtension,omitempty"`

	// Picture:
	Picture *[]byte `json:"Picture,omitempty"`

	// PictureName:
	PictureName *string `json:"PictureName,omitempty"`

	// PictureThumbnailUrl:
	PictureThumbnailUrl *string `json:"PictureThumbnailUrl,omitempty"`

	// PictureUrl:
	PictureUrl *string `json:"PictureUrl,omitempty"`

	// Postcode:
	Postcode *string `json:"Postcode,omitempty"`

	// SocialSecurityNumber:
	SocialSecurityNumber *string `json:"SocialSecurityNumber,omitempty"`

	// StartDate:
	StartDate *types.Date `json:"StartDate,omitempty"`

	// State:
	State *string `json:"State,omitempty"`

	// Title:
	Title *string `json:"Title,omitempty"`
}

func (e *CRMContacts) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *CRMContactsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "CRM/Contacts", method)
}

// List the CRMContacts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CRMContactsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*CRMContacts, error) {
	var entities []*CRMContacts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/CRM/Contacts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the CRMContacts entitiy in the provided division.
func (s *CRMContactsEndpoint) Get(ctx context.Context, division int, id *int64) (*CRMContacts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/CRM/Contacts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &CRMContacts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
