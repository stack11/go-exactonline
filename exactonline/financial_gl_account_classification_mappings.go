// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
)

// FinancialGLAccountClassificationMappingsService is responsible for communicating with
// the GLAccountClassificationMappings endpoint of the Financial service.
type FinancialGLAccountClassificationMappingsService service

// FinancialGLAccountClassificationMappings:
// Service: Financial
// Entity: GLAccountClassificationMappings
// URL: /api/v1/beta/{division}/financial/GLAccountClassificationMappings
// HasWebhook: false
// IsInBeta: true
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialGLAccountClassificationMappings
type FinancialGLAccountClassificationMappings struct {
	// ID: Primary key
	ID *GUID `json:",omitempty"`

	// Classification: ID of the classification
	Classification *GUID `json:",omitempty"`

	// ClassificationCode: Code of the classification
	ClassificationCode *string `json:",omitempty"`

	// ClassificationDescription: Description of the classification
	ClassificationDescription *string `json:",omitempty"`

	// Division: Division of the classification mapping
	Division *int `json:",omitempty"`

	// GLAccount: ID of the general ledger account
	GLAccount *GUID `json:",omitempty"`

	// GLAccountCode: Code of the general ledger account
	GLAccountCode *string `json:",omitempty"`

	// GLAccountDescription: Description of the general ledger account
	GLAccountDescription *string `json:",omitempty"`

	// GLSchemeCode: Code of the general ledger scheme
	GLSchemeCode *string `json:",omitempty"`

	// GLSchemeDescription: Description of the general ledger scheme
	GLSchemeDescription *string `json:",omitempty"`

	// GLSchemeID: General ledger scheme ID of the element
	GLSchemeID *GUID `json:",omitempty"`
}

func (s *FinancialGLAccountClassificationMappings) GetIdentifier() GUID {
	return *s.ID
}

// List the GLAccountClassificationMappings entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *FinancialGLAccountClassificationMappingsService) List(ctx context.Context, division int, all bool) ([]*FinancialGLAccountClassificationMappings, error) {
	var entities []*FinancialGLAccountClassificationMappings
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/financial/GLAccountClassificationMappings?$select=*", division)
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
