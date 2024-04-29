// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import "github.com/stack11/go-exactonline/api"

type service struct {
	client *api.Client
}

// ManufacturingService is responsible for communication with the Manufacturing
// endpoints of the Exact Online API.
type ManufacturingService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	BillOfMaterialMaterials                  *BillOfMaterialMaterialsEndpoint
	BillOfMaterialRoutings                   *BillOfMaterialRoutingsEndpoint
	BillOfMaterialVersions                   *BillOfMaterialVersionsEndpoint
	ByProductReceipts                        *ByProductReceiptsEndpoint
	ByProductReversals                       *ByProductReversalsEndpoint
	ManufacturingSettings                    *ManufacturingSettingsEndpoint
	MaterialIssues                           *MaterialIssuesEndpoint
	MaterialReversals                        *MaterialReversalsEndpoint
	OperationResources                       *OperationResourcesEndpoint
	Operations                               *OperationsEndpoint
	ProductionAreas                          *ProductionAreasEndpoint
	RecentTimeTransactions                   *RecentTimeTransactionsEndpoint
	ShopOrderMaterialPlanDetails             *ShopOrderMaterialPlanDetailsEndpoint
	ShopOrderMaterialPlans                   *ShopOrderMaterialPlansEndpoint
	ShopOrderPriorities                      *ShopOrderPrioritiesEndpoint
	ShopOrderReceipts                        *ShopOrderReceiptsEndpoint
	ShopOrderReversals                       *ShopOrderReversalsEndpoint
	ShopOrderRoutingStepPlans                *ShopOrderRoutingStepPlansEndpoint
	ShopOrderRoutingStepPlansAvailableToWork *ShopOrderRoutingStepPlansAvailableToWorkEndpoint
	ShopOrders                               *ShopOrdersEndpoint
	StageForDeliveryReceipts                 *StageForDeliveryReceiptsEndpoint
	StageForDeliveryReversals                *StageForDeliveryReversalsEndpoint
	StartedTimedTimeTransactions             *StartedTimedTimeTransactionsEndpoint
	SubOrderReceipts                         *SubOrderReceiptsEndpoint
	SubOrderReversals                        *SubOrderReversalsEndpoint
	TimedTimeTransactions                    *TimedTimeTransactionsEndpoint
	TimeTransactions                         *TimeTransactionsEndpoint
	Workcenters                              *WorkcentersEndpoint
}

// NewManufacturingService creates a new initialized instance of the
// ManufacturingService.
func NewManufacturingService(apiClient *api.Client) *ManufacturingService {
	s := &ManufacturingService{client: apiClient}

	s.common.client = apiClient

	s.BillOfMaterialMaterials = (*BillOfMaterialMaterialsEndpoint)(&s.common)
	s.BillOfMaterialRoutings = (*BillOfMaterialRoutingsEndpoint)(&s.common)
	s.BillOfMaterialVersions = (*BillOfMaterialVersionsEndpoint)(&s.common)
	s.ByProductReceipts = (*ByProductReceiptsEndpoint)(&s.common)
	s.ByProductReversals = (*ByProductReversalsEndpoint)(&s.common)
	s.ManufacturingSettings = (*ManufacturingSettingsEndpoint)(&s.common)
	s.MaterialIssues = (*MaterialIssuesEndpoint)(&s.common)
	s.MaterialReversals = (*MaterialReversalsEndpoint)(&s.common)
	s.OperationResources = (*OperationResourcesEndpoint)(&s.common)
	s.Operations = (*OperationsEndpoint)(&s.common)
	s.ProductionAreas = (*ProductionAreasEndpoint)(&s.common)
	s.RecentTimeTransactions = (*RecentTimeTransactionsEndpoint)(&s.common)
	s.ShopOrderMaterialPlanDetails = (*ShopOrderMaterialPlanDetailsEndpoint)(&s.common)
	s.ShopOrderMaterialPlans = (*ShopOrderMaterialPlansEndpoint)(&s.common)
	s.ShopOrderPriorities = (*ShopOrderPrioritiesEndpoint)(&s.common)
	s.ShopOrderReceipts = (*ShopOrderReceiptsEndpoint)(&s.common)
	s.ShopOrderReversals = (*ShopOrderReversalsEndpoint)(&s.common)
	s.ShopOrderRoutingStepPlans = (*ShopOrderRoutingStepPlansEndpoint)(&s.common)
	s.ShopOrderRoutingStepPlansAvailableToWork = (*ShopOrderRoutingStepPlansAvailableToWorkEndpoint)(&s.common)
	s.ShopOrders = (*ShopOrdersEndpoint)(&s.common)
	s.StageForDeliveryReceipts = (*StageForDeliveryReceiptsEndpoint)(&s.common)
	s.StageForDeliveryReversals = (*StageForDeliveryReversalsEndpoint)(&s.common)
	s.StartedTimedTimeTransactions = (*StartedTimedTimeTransactionsEndpoint)(&s.common)
	s.SubOrderReceipts = (*SubOrderReceiptsEndpoint)(&s.common)
	s.SubOrderReversals = (*SubOrderReversalsEndpoint)(&s.common)
	s.TimedTimeTransactions = (*TimedTimeTransactionsEndpoint)(&s.common)
	s.TimeTransactions = (*TimeTransactionsEndpoint)(&s.common)
	s.Workcenters = (*WorkcentersEndpoint)(&s.common)

	return s
}
