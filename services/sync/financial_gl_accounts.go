// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package sync

import (
	"context"
	"encoding/json"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// FinancialGLAccountsEndpoint is responsible for communicating with
// the FinancialGLAccounts endpoint of the Sync service.
type FinancialGLAccountsEndpoint service

// FinancialGLAccounts:
// Service: Sync
// Entity: FinancialGLAccounts
// URL: /api/v1/{division}/sync/Financial/GLAccounts
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncFinancialGLAccounts
type FinancialGLAccounts struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp:
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// AllowCostsInSales:
	AllowCostsInSales *byte `json:"AllowCostsInSales,omitempty"`

	// AssimilatedVATBox:
	AssimilatedVATBox *int `json:"AssimilatedVATBox,omitempty"`

	// BalanceSide:
	BalanceSide *string `json:"BalanceSide,omitempty"`

	// BalanceType:
	BalanceType *string `json:"BalanceType,omitempty"`

	// BelcotaxType:
	BelcotaxType *int `json:"BelcotaxType,omitempty"`

	// Code:
	Code *string `json:"Code,omitempty"`

	// Compress:
	Compress *bool `json:"Compress,omitempty"`

	// Costcenter:
	Costcenter *string `json:"Costcenter,omitempty"`

	// CostcenterDescription:
	CostcenterDescription *string `json:"CostcenterDescription,omitempty"`

	// Costunit:
	Costunit *string `json:"Costunit,omitempty"`

	// CostunitDescription:
	CostunitDescription *string `json:"CostunitDescription,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// CustomField:
	CustomField *string `json:"CustomField,omitempty"`

	// DeductibilityPercentages:
	DeductibilityPercentages *json.RawMessage `json:"DeductibilityPercentages,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// DescriptionTermID:
	DescriptionTermID *int `json:"DescriptionTermID,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// ExcludeVATListing:
	ExcludeVATListing *byte `json:"ExcludeVATListing,omitempty"`

	// ExpenseNonDeductiblePercentage:
	ExpenseNonDeductiblePercentage *float64 `json:"ExpenseNonDeductiblePercentage,omitempty"`

	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// IsBlocked:
	IsBlocked *bool `json:"IsBlocked,omitempty"`

	// Matching:
	Matching *bool `json:"Matching,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// PrivateGLAccount:
	PrivateGLAccount *types.GUID `json:"PrivateGLAccount,omitempty"`

	// PrivatePercentage:
	PrivatePercentage *float64 `json:"PrivatePercentage,omitempty"`

	// ReportingCode:
	ReportingCode *string `json:"ReportingCode,omitempty"`

	// RevalueCurrency:
	RevalueCurrency *bool `json:"RevalueCurrency,omitempty"`

	// SearchCode:
	SearchCode *string `json:"SearchCode,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`

	// TypeDescription:
	TypeDescription *string `json:"TypeDescription,omitempty"`

	// UseCostcenter:
	UseCostcenter *byte `json:"UseCostcenter,omitempty"`

	// UseCostunit:
	UseCostunit *byte `json:"UseCostunit,omitempty"`

	// VATCode:
	VATCode *string `json:"VATCode,omitempty"`

	// VATDescription:
	VATDescription *string `json:"VATDescription,omitempty"`

	// VATGLAccountType:
	VATGLAccountType *string `json:"VATGLAccountType,omitempty"`

	// VATNonDeductibleGLAccount:
	VATNonDeductibleGLAccount *types.GUID `json:"VATNonDeductibleGLAccount,omitempty"`

	// VATNonDeductiblePercentage:
	VATNonDeductiblePercentage *float64 `json:"VATNonDeductiblePercentage,omitempty"`

	// VATSystem:
	VATSystem *string `json:"VATSystem,omitempty"`

	// YearEndCostGLAccount:
	YearEndCostGLAccount *types.GUID `json:"YearEndCostGLAccount,omitempty"`

	// YearEndReflectionGLAccount:
	YearEndReflectionGLAccount *types.GUID `json:"YearEndReflectionGLAccount,omitempty"`
}

func (e *FinancialGLAccounts) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *FinancialGLAccountsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "Financial/GLAccounts", method)
}

// List the FinancialGLAccounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *FinancialGLAccountsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*FinancialGLAccounts, error) {
	var entities []*FinancialGLAccounts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Financial/GLAccounts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the FinancialGLAccounts entitiy in the provided division.
func (s *FinancialGLAccountsEndpoint) Get(ctx context.Context, division int, id *int64) (*FinancialGLAccounts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/Financial/GLAccounts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &FinancialGLAccounts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
