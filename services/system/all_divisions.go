// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package system

import (
	"context"
	"encoding/json"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// AllDivisionsEndpoint is responsible for communicating with
// the AllDivisions endpoint of the System service.
type AllDivisionsEndpoint service

// AllDivisions:
// Service: System
// Entity: AllDivisions
// URL: /api/v1/{division}/system/AllDivisions
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SystemSystemAllDivisions
type AllDivisions struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Code: Primary key
	Code *int `json:"Code,omitempty"`

	// AddressLine1: Address line 1
	AddressLine1 *string `json:"AddressLine1,omitempty"`

	// AddressLine2: Address line 2
	AddressLine2 *string `json:"AddressLine2,omitempty"`

	// AddressLine3: Address line 3
	AddressLine3 *string `json:"AddressLine3,omitempty"`

	// ArchiveDate: Date on which the division is archived
	ArchiveDate *types.Date `json:"ArchiveDate,omitempty"`

	// BlockingStatus: Values: 0 = Not blocked, 1 = Backup/restore, 2 = Conversion busy, 3 = Conversion shadow, 4 = Conversion waiting, 5 = Copy data waiting, 6 = Copy data busy
	BlockingStatus *int `json:"BlockingStatus,omitempty"`

	// BusinessTypeCode: Business type code
	BusinessTypeCode *string `json:"BusinessTypeCode,omitempty"`

	// BusinessTypeDescription: Business type description
	BusinessTypeDescription *string `json:"BusinessTypeDescription,omitempty"`

	// ChamberOfCommerceEstablishment: Chamber of commerce establishment
	ChamberOfCommerceEstablishment *string `json:"ChamberOfCommerceEstablishment,omitempty"`

	// ChamberOfCommerceNumber: Chamber of commerce number
	ChamberOfCommerceNumber *string `json:"ChamberOfCommerceNumber,omitempty"`

	// City: City
	City *string `json:"City,omitempty"`

	// Class_01: First division classification. User should have access rights to view division classifications.
	Class_01 *json.RawMessage `json:"Class_01,omitempty"`

	// Class_02: Second division classification. User should have access rights to view division classifications.
	Class_02 *json.RawMessage `json:"Class_02,omitempty"`

	// Class_03: Third division classification. User should have access rights to view division classifications.
	Class_03 *json.RawMessage `json:"Class_03,omitempty"`

	// Class_04: Fourth division classification. User should have access rights to view division classifications.
	Class_04 *json.RawMessage `json:"Class_04,omitempty"`

	// Class_05: Fifth division classification. User should have access rights to view division classifications.
	Class_05 *json.RawMessage `json:"Class_05,omitempty"`

	// CompanySizeCode: Company size code
	CompanySizeCode *string `json:"CompanySizeCode,omitempty"`

	// CompanySizeDescription: Company size description
	CompanySizeDescription *string `json:"CompanySizeDescription,omitempty"`

	// Country: Country of the division. Is used for determination of legislation
	Country *string `json:"Country,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency: Default currency
	Currency *string `json:"Currency,omitempty"`

	// Current: True when this division is most recently used by the API
	Current *bool `json:"Current,omitempty"`

	// Customer: Owner account of the division
	Customer *types.GUID `json:"Customer,omitempty"`

	// CustomerCode: Owner account code of the division
	CustomerCode *string `json:"CustomerCode,omitempty"`

	// CustomerName: Owner account name of the division
	CustomerName *string `json:"CustomerName,omitempty"`

	// DatevAccountantNumber: Accountant number DATEV (Germany)
	DatevAccountantNumber *string `json:"DatevAccountantNumber,omitempty"`

	// DatevClientNumber: Client number DATEV (Germany)
	DatevClientNumber *string `json:"DatevClientNumber,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// DivisionMoveDate: Date when the division was moved. Please resync all data when this value changes because value of Timestamp is regenerated.
	DivisionMoveDate *types.Date `json:"DivisionMoveDate,omitempty"`

	// Email: Email address
	Email *string `json:"Email,omitempty"`

	// Fax: Fax number
	Fax *string `json:"Fax,omitempty"`

	// Hid: Company number that is assigned by the customer
	Hid *int64 `json:"Hid,omitempty"`

	// IsDossierDivision: True if the division is a dossier division
	IsDossierDivision *bool `json:"IsDossierDivision,omitempty"`

	// IsMainDivision: True if the division is the main division
	IsMainDivision *bool `json:"IsMainDivision,omitempty"`

	// IsPracticeDivision: True if the division is a practice division
	IsPracticeDivision *bool `json:"IsPracticeDivision,omitempty"`

	// Legislation: Legislation
	Legislation *string `json:"Legislation,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// OBNumber: The soletrader VAT number used for offical returns to tax authority
	OBNumber *string `json:"OBNumber,omitempty"`

	// Phone: Phone number
	Phone *string `json:"Phone,omitempty"`

	// Postcode: Postcode
	Postcode *string `json:"Postcode,omitempty"`

	// SbiCode: SBI code
	SbiCode *string `json:"SbiCode,omitempty"`

	// SbiDescription: SBI description
	SbiDescription *string `json:"SbiDescription,omitempty"`

	// SectorCode: Sector code
	SectorCode *string `json:"SectorCode,omitempty"`

	// SectorDescription: Sector description
	SectorDescription *string `json:"SectorDescription,omitempty"`

	// ShareCapital: the part of the capital of a company that comes from the issue of shares (France)
	ShareCapital *float64 `json:"ShareCapital,omitempty"`

	// SiretNumber: An INSEE code which allows the geographic identification of the company. (France)
	SiretNumber *string `json:"SiretNumber,omitempty"`

	// StartDate: Date on which the division becomes active
	StartDate *types.Date `json:"StartDate,omitempty"`

	// State: State/Province code
	State *string `json:"State,omitempty"`

	// Status: Follow the Division Status 0 for Inactive, 1 for Active and 2 for Archived Divisions
	Status *int `json:"Status,omitempty"`

	// SubsectorCode: Subsector code
	SubsectorCode *string `json:"SubsectorCode,omitempty"`

	// SubsectorDescription: Subsector description
	SubsectorDescription *string `json:"SubsectorDescription,omitempty"`

	// TaxOfficeNumber: Number of your local tax authority (Germany)
	TaxOfficeNumber *string `json:"TaxOfficeNumber,omitempty"`

	// TaxReferenceNumber: Local tax reference number (Germany)
	TaxReferenceNumber *string `json:"TaxReferenceNumber,omitempty"`

	// TemplateCode: Division template code
	TemplateCode *string `json:"TemplateCode,omitempty"`

	// VATNumber: The number under which the account is known at the Value Added Tax collection agency
	VATNumber *string `json:"VATNumber,omitempty"`

	// Website: Customer value, hyperlink to external website
	Website *string `json:"Website,omitempty"`
}

func (e *AllDivisions) GetPrimary() *int {
	return e.Code
}

func (s *AllDivisionsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "system/AllDivisions", method)
}

// List the AllDivisions entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AllDivisionsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*AllDivisions, error) {
	var entities []*AllDivisions
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/system/AllDivisions", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the AllDivisions entitiy in the provided division.
func (s *AllDivisionsEndpoint) Get(ctx context.Context, division int, id *int) (*AllDivisions, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/system/AllDivisions", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &AllDivisions{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
