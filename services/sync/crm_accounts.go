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

// CRMAccountsEndpoint is responsible for communicating with
// the CRMAccounts endpoint of the Sync service.
type CRMAccountsEndpoint service

// CRMAccounts:
// Service: Sync
// Entity: CRMAccounts
// URL: /api/v1/{division}/sync/CRM/Accounts
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SyncCRMAccounts
type CRMAccounts struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// Timestamp:
	Timestamp *int64 `json:"Timestamp,omitempty"`

	// Accountant:
	Accountant *types.GUID `json:"Accountant,omitempty"`

	// AccountManager:
	AccountManager *types.GUID `json:"AccountManager,omitempty"`

	// AccountManagerFullName:
	AccountManagerFullName *string `json:"AccountManagerFullName,omitempty"`

	// AccountManagerHID:
	AccountManagerHID *int `json:"AccountManagerHID,omitempty"`

	// ActivitySector:
	ActivitySector *types.GUID `json:"ActivitySector,omitempty"`

	// ActivitySubSector:
	ActivitySubSector *types.GUID `json:"ActivitySubSector,omitempty"`

	// AddressLine1:
	AddressLine1 *string `json:"AddressLine1,omitempty"`

	// AddressLine2:
	AddressLine2 *string `json:"AddressLine2,omitempty"`

	// AddressLine3:
	AddressLine3 *string `json:"AddressLine3,omitempty"`

	// BankAccounts:
	BankAccounts *json.RawMessage `json:"BankAccounts,omitempty"`

	// Blocked:
	Blocked *bool `json:"Blocked,omitempty"`

	// BRIN:
	BRIN *types.GUID `json:"BRIN,omitempty"`

	// BSN:
	BSN *string `json:"BSN,omitempty"`

	// BusinessType:
	BusinessType *types.GUID `json:"BusinessType,omitempty"`

	// CanDropShip:
	CanDropShip *bool `json:"CanDropShip,omitempty"`

	// ChamberOfCommerce:
	ChamberOfCommerce *string `json:"ChamberOfCommerce,omitempty"`

	// City:
	City *string `json:"City,omitempty"`

	// Classification:
	Classification *string `json:"Classification,omitempty"`

	// Classification1:
	Classification1 *types.GUID `json:"Classification1,omitempty"`

	// Classification2:
	Classification2 *types.GUID `json:"Classification2,omitempty"`

	// Classification3:
	Classification3 *types.GUID `json:"Classification3,omitempty"`

	// Classification4:
	Classification4 *types.GUID `json:"Classification4,omitempty"`

	// Classification5:
	Classification5 *types.GUID `json:"Classification5,omitempty"`

	// Classification6:
	Classification6 *types.GUID `json:"Classification6,omitempty"`

	// Classification7:
	Classification7 *types.GUID `json:"Classification7,omitempty"`

	// Classification8:
	Classification8 *types.GUID `json:"Classification8,omitempty"`

	// ClassificationDescription:
	ClassificationDescription *string `json:"ClassificationDescription,omitempty"`

	// Code:
	Code *string `json:"Code,omitempty"`

	// CodeAtSupplier:
	CodeAtSupplier *string `json:"CodeAtSupplier,omitempty"`

	// CompanySize:
	CompanySize *types.GUID `json:"CompanySize,omitempty"`

	// ConsolidationScenario:
	ConsolidationScenario *byte `json:"ConsolidationScenario,omitempty"`

	// ControlledDate:
	ControlledDate *types.Date `json:"ControlledDate,omitempty"`

	// Costcenter:
	Costcenter *string `json:"Costcenter,omitempty"`

	// CostcenterDescription:
	CostcenterDescription *string `json:"CostcenterDescription,omitempty"`

	// CostPaid:
	CostPaid *byte `json:"CostPaid,omitempty"`

	// Country:
	Country *string `json:"Country,omitempty"`

	// CountryName:
	CountryName *string `json:"CountryName,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// CreditLinePurchase:
	CreditLinePurchase *float64 `json:"CreditLinePurchase,omitempty"`

	// CreditLineSales:
	CreditLineSales *float64 `json:"CreditLineSales,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// CustomerSince:
	CustomerSince *types.Date `json:"CustomerSince,omitempty"`

	// CustomField:
	CustomField *string `json:"CustomField,omitempty"`

	// DatevCreditorCode:
	DatevCreditorCode *string `json:"DatevCreditorCode,omitempty"`

	// DatevDebtorCode:
	DatevDebtorCode *string `json:"DatevDebtorCode,omitempty"`

	// DeliveryAdvice:
	DeliveryAdvice *byte `json:"DeliveryAdvice,omitempty"`

	// DiscountPurchase:
	DiscountPurchase *float64 `json:"DiscountPurchase,omitempty"`

	// DiscountSales:
	DiscountSales *float64 `json:"DiscountSales,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DunsNumber:
	DunsNumber *string `json:"DunsNumber,omitempty"`

	// Email:
	Email *string `json:"Email,omitempty"`

	// EndDate:
	EndDate *types.Date `json:"EndDate,omitempty"`

	// EORINumber:
	EORINumber *string `json:"EORINumber,omitempty"`

	// EstablishedDate:
	EstablishedDate *types.Date `json:"EstablishedDate,omitempty"`

	// Fax:
	Fax *string `json:"Fax,omitempty"`

	// GLAccountPurchase:
	GLAccountPurchase *types.GUID `json:"GLAccountPurchase,omitempty"`

	// GLAccountSales:
	GLAccountSales *types.GUID `json:"GLAccountSales,omitempty"`

	// GLAP:
	GLAP *types.GUID `json:"GLAP,omitempty"`

	// GLAR:
	GLAR *types.GUID `json:"GLAR,omitempty"`

	// GlnNumber:
	GlnNumber *string `json:"GlnNumber,omitempty"`

	// HasWithholdingTaxSales:
	HasWithholdingTaxSales *bool `json:"HasWithholdingTaxSales,omitempty"`

	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// IgnoreDatevWarningMessage:
	IgnoreDatevWarningMessage *bool `json:"IgnoreDatevWarningMessage,omitempty"`

	// IncotermAddressPurchase:
	IncotermAddressPurchase *string `json:"IncotermAddressPurchase,omitempty"`

	// IncotermAddressSales:
	IncotermAddressSales *string `json:"IncotermAddressSales,omitempty"`

	// IncotermCodePurchase:
	IncotermCodePurchase *string `json:"IncotermCodePurchase,omitempty"`

	// IncotermCodeSales:
	IncotermCodeSales *string `json:"IncotermCodeSales,omitempty"`

	// IncotermVersionPurchase:
	IncotermVersionPurchase *int `json:"IncotermVersionPurchase,omitempty"`

	// IncotermVersionSales:
	IncotermVersionSales *int `json:"IncotermVersionSales,omitempty"`

	// IntraStatArea:
	IntraStatArea *string `json:"IntraStatArea,omitempty"`

	// IntraStatDeliveryTerm:
	IntraStatDeliveryTerm *string `json:"IntraStatDeliveryTerm,omitempty"`

	// IntraStatSystem:
	IntraStatSystem *string `json:"IntraStatSystem,omitempty"`

	// IntraStatTransactionA:
	IntraStatTransactionA *string `json:"IntraStatTransactionA,omitempty"`

	// IntraStatTransactionB:
	IntraStatTransactionB *string `json:"IntraStatTransactionB,omitempty"`

	// IntraStatTransportMethod:
	IntraStatTransportMethod *string `json:"IntraStatTransportMethod,omitempty"`

	// InvoiceAccount:
	InvoiceAccount *types.GUID `json:"InvoiceAccount,omitempty"`

	// InvoiceAccountCode:
	InvoiceAccountCode *string `json:"InvoiceAccountCode,omitempty"`

	// InvoiceAccountName:
	InvoiceAccountName *string `json:"InvoiceAccountName,omitempty"`

	// InvoiceAttachmentType:
	InvoiceAttachmentType *int `json:"InvoiceAttachmentType,omitempty"`

	// InvoicingMethod:
	InvoicingMethod *int `json:"InvoicingMethod,omitempty"`

	// IsAccountant:
	IsAccountant *byte `json:"IsAccountant,omitempty"`

	// IsAgency:
	IsAgency *byte `json:"IsAgency,omitempty"`

	// IsAnonymised:
	IsAnonymised *byte `json:"IsAnonymised,omitempty"`

	// IsBank:
	IsBank *bool `json:"IsBank,omitempty"`

	// IsCompetitor:
	IsCompetitor *byte `json:"IsCompetitor,omitempty"`

	// IsExtraDuty:
	IsExtraDuty *bool `json:"IsExtraDuty,omitempty"`

	// IsMailing:
	IsMailing *byte `json:"IsMailing,omitempty"`

	// IsMember:
	IsMember *bool `json:"IsMember,omitempty"`

	// IsPilot:
	IsPilot *bool `json:"IsPilot,omitempty"`

	// IsPurchase:
	IsPurchase *bool `json:"IsPurchase,omitempty"`

	// IsReseller:
	IsReseller *bool `json:"IsReseller,omitempty"`

	// IsSales:
	IsSales *bool `json:"IsSales,omitempty"`

	// IsSupplier:
	IsSupplier *bool `json:"IsSupplier,omitempty"`

	// Language:
	Language *string `json:"Language,omitempty"`

	// LanguageDescription:
	LanguageDescription *string `json:"LanguageDescription,omitempty"`

	// Latitude:
	Latitude *float64 `json:"Latitude,omitempty"`

	// LeadPurpose:
	LeadPurpose *types.GUID `json:"LeadPurpose,omitempty"`

	// LeadSource:
	LeadSource *types.GUID `json:"LeadSource,omitempty"`

	// Logo:
	Logo *[]byte `json:"Logo,omitempty"`

	// LogoFileName:
	LogoFileName *string `json:"LogoFileName,omitempty"`

	// LogoThumbnailUrl:
	LogoThumbnailUrl *string `json:"LogoThumbnailUrl,omitempty"`

	// LogoUrl:
	LogoUrl *string `json:"LogoUrl,omitempty"`

	// Longitude:
	Longitude *float64 `json:"Longitude,omitempty"`

	// MainContact:
	MainContact *types.GUID `json:"MainContact,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Name:
	Name *string `json:"Name,omitempty"`

	// OINNumber:
	OINNumber *string `json:"OINNumber,omitempty"`

	// Parent:
	Parent *types.GUID `json:"Parent,omitempty"`

	// PayAsYouEarn:
	PayAsYouEarn *string `json:"PayAsYouEarn,omitempty"`

	// PaymentConditionPurchase:
	PaymentConditionPurchase *string `json:"PaymentConditionPurchase,omitempty"`

	// PaymentConditionPurchaseDescription:
	PaymentConditionPurchaseDescription *string `json:"PaymentConditionPurchaseDescription,omitempty"`

	// PaymentConditionSales:
	PaymentConditionSales *string `json:"PaymentConditionSales,omitempty"`

	// PaymentConditionSalesDescription:
	PaymentConditionSalesDescription *string `json:"PaymentConditionSalesDescription,omitempty"`

	// Phone:
	Phone *string `json:"Phone,omitempty"`

	// PhoneExtension:
	PhoneExtension *string `json:"PhoneExtension,omitempty"`

	// Postcode:
	Postcode *string `json:"Postcode,omitempty"`

	// PriceList:
	PriceList *types.GUID `json:"PriceList,omitempty"`

	// PurchaseCurrency:
	PurchaseCurrency *string `json:"PurchaseCurrency,omitempty"`

	// PurchaseCurrencyDescription:
	PurchaseCurrencyDescription *string `json:"PurchaseCurrencyDescription,omitempty"`

	// PurchaseLeadDays:
	PurchaseLeadDays *int `json:"PurchaseLeadDays,omitempty"`

	// PurchaseVATCode:
	PurchaseVATCode *string `json:"PurchaseVATCode,omitempty"`

	// PurchaseVATCodeDescription:
	PurchaseVATCodeDescription *string `json:"PurchaseVATCodeDescription,omitempty"`

	// RecepientOfCommissions:
	RecepientOfCommissions *bool `json:"RecepientOfCommissions,omitempty"`

	// Remarks:
	Remarks *string `json:"Remarks,omitempty"`

	// Reseller:
	Reseller *types.GUID `json:"Reseller,omitempty"`

	// ResellerCode:
	ResellerCode *string `json:"ResellerCode,omitempty"`

	// ResellerName:
	ResellerName *string `json:"ResellerName,omitempty"`

	// RSIN:
	RSIN *string `json:"RSIN,omitempty"`

	// SalesCurrency:
	SalesCurrency *string `json:"SalesCurrency,omitempty"`

	// SalesCurrencyDescription:
	SalesCurrencyDescription *string `json:"SalesCurrencyDescription,omitempty"`

	// SalesTaxSchedule:
	SalesTaxSchedule *types.GUID `json:"SalesTaxSchedule,omitempty"`

	// SalesTaxScheduleCode:
	SalesTaxScheduleCode *string `json:"SalesTaxScheduleCode,omitempty"`

	// SalesTaxScheduleDescription:
	SalesTaxScheduleDescription *string `json:"SalesTaxScheduleDescription,omitempty"`

	// SalesVATCode:
	SalesVATCode *string `json:"SalesVATCode,omitempty"`

	// SalesVATCodeDescription:
	SalesVATCodeDescription *string `json:"SalesVATCodeDescription,omitempty"`

	// SearchCode:
	SearchCode *string `json:"SearchCode,omitempty"`

	// SecurityLevel:
	SecurityLevel *int `json:"SecurityLevel,omitempty"`

	// SeparateInvPerProject:
	SeparateInvPerProject *byte `json:"SeparateInvPerProject,omitempty"`

	// SeparateInvPerSubscription:
	SeparateInvPerSubscription *byte `json:"SeparateInvPerSubscription,omitempty"`

	// ShippingLeadDays:
	ShippingLeadDays *int `json:"ShippingLeadDays,omitempty"`

	// ShippingMethod:
	ShippingMethod *types.GUID `json:"ShippingMethod,omitempty"`

	// ShowRemarkForSales:
	ShowRemarkForSales *bool `json:"ShowRemarkForSales,omitempty"`

	// StartDate:
	StartDate *types.Date `json:"StartDate,omitempty"`

	// State:
	State *string `json:"State,omitempty"`

	// StateName:
	StateName *string `json:"StateName,omitempty"`

	// Status:
	Status *string `json:"Status,omitempty"`

	// StatusSince:
	StatusSince *types.Date `json:"StatusSince,omitempty"`

	// TradeName:
	TradeName *string `json:"TradeName,omitempty"`

	// Type:
	Type *string `json:"Type,omitempty"`

	// UniqueTaxpayerReference:
	UniqueTaxpayerReference *string `json:"UniqueTaxpayerReference,omitempty"`

	// VATLiability:
	VATLiability *string `json:"VATLiability,omitempty"`

	// VATNumber:
	VATNumber *string `json:"VATNumber,omitempty"`

	// Website:
	Website *string `json:"Website,omitempty"`
}

func (e *CRMAccounts) GetPrimary() *int64 {
	return e.Timestamp
}

func (s *CRMAccountsEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "CRM/Accounts", method)
}

// List the CRMAccounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *CRMAccountsEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*CRMAccounts, error) {
	var entities []*CRMAccounts
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/CRM/Accounts", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the CRMAccounts entitiy in the provided division.
func (s *CRMAccountsEndpoint) Get(ctx context.Context, division int, id *int64) (*CRMAccounts, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/sync/CRM/Accounts", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &CRMAccounts{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}
