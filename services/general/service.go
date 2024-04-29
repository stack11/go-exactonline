// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package general

import "github.com/stack11/go-exactonline/api"

type service struct {
	client *api.Client
}

// GeneralService is responsible for communication with the General
// endpoints of the Exact Online API.
type GeneralService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	Currencies *CurrenciesEndpoint
	Layouts    *LayoutsEndpoint
}

// NewGeneralService creates a new initialized instance of the
// GeneralService.
func NewGeneralService(apiClient *api.Client) *GeneralService {
	s := &GeneralService{client: apiClient}

	s.common.client = apiClient

	s.Currencies = (*CurrenciesEndpoint)(&s.common)
	s.Layouts = (*LayoutsEndpoint)(&s.common)

	return s
}
