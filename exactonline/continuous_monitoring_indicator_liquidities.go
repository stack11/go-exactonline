// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// ContinuousMonitoringIndicatorLiquiditiesService is responsible for communicating with
// the IndicatorLiquidities endpoint of the ContinuousMonitoring service.
type ContinuousMonitoringIndicatorLiquiditiesService service

// ContinuousMonitoringIndicatorLiquidities:
// Service: ContinuousMonitoring
// Entity: IndicatorLiquidities
// URL: /api/v1/beta/{division}/continuousmonitoring/IndicatorLiquidities
// HasWebhook: false
// IsInBeta: true
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ContinuousMonitoringIndicatorLiquidities
type ContinuousMonitoringIndicatorLiquidities struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Active: Indicates if this indicator is active or inactive
	Active *byte `json:",omitempty"`

	// Classification: Indicator classification (1 = Quality, 2 = Advice). Default = 1
	Classification *int `json:",omitempty"`

	// Created: Creation date
	Created *Date `json:",omitempty"`

	// CreateSignal: Indicates whether a signal is created
	CreateSignal *byte `json:",omitempty"`

	// Creator: User ID of creator
	Creator *GUID `json:",omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:",omitempty"`

	// Description: Description of indicator
	Description *string `json:",omitempty"`

	// Division: Division code
	Division *int `json:",omitempty"`

	// ExternalCode: External code
	ExternalCode *string `json:",omitempty"`

	// Modified: Last modified date
	Modified *Date `json:",omitempty"`

	// Modifier: User ID of modifier
	Modifier *GUID `json:",omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:",omitempty"`

	// Severity: Severity of the indicators (1 = Low, 2 = Medium, 3 = High, 4 = Critical)
	Severity *int `json:",omitempty"`

	// Type: Indicator type (1 = Balance G/L account per financial year, 2 = Usage of journals, 3 = Deviating amount entered, 4 = Liquidity, 5 = VAT Return deadline, 6 = Difference result in percentage, 7 = Different VAT code used)
	Type *int `json:",omitempty"`
}

func (s *ContinuousMonitoringIndicatorLiquidities) GetIdentifier() GUID {
	return *s.ID
}

// List the IndicatorLiquidities entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ContinuousMonitoringIndicatorLiquiditiesService) List(ctx context.Context, division int, all bool) ([]*ContinuousMonitoringIndicatorLiquidities, error) {
	var entities []*ContinuousMonitoringIndicatorLiquidities
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorLiquidities?$select=*", division)
	if err != nil {
		return nil, err
	}
	if all {
		err = s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err = s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}
