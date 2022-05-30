// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesorder

import "github.com/stack11/go-exactonline/api"

type service struct {
	client *api.Client
}

// SalesOrderService is responsible for communication with the SalesOrder
// endpoints of the Exact Online API.
type SalesOrderService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	GoodsDeliveries    *GoodsDeliveriesEndpoint
	GoodsDeliveryLines *GoodsDeliveryLinesEndpoint
	SalesOrderLines    *SalesOrderLinesEndpoint
	SalesOrders        *SalesOrdersEndpoint
}

// NewSalesOrderService creates a new initialized instance of the
// SalesOrderService.
func NewSalesOrderService(apiClient *api.Client) *SalesOrderService {
	s := &SalesOrderService{client: apiClient}

	s.common.client = apiClient

	s.GoodsDeliveries = (*GoodsDeliveriesEndpoint)(&s.common)
	s.GoodsDeliveryLines = (*GoodsDeliveryLinesEndpoint)(&s.common)
	s.SalesOrderLines = (*SalesOrderLinesEndpoint)(&s.common)
	s.SalesOrders = (*SalesOrdersEndpoint)(&s.common)

	return s
}
