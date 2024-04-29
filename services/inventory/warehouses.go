// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package inventory

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

// WarehousesEndpoint is responsible for communicating with
// the Warehouses endpoint of the Inventory service.
type WarehousesEndpoint service

// Warehouses:
// Service: Inventory
// Entity: Warehouses
// URL: /api/v1/{division}/inventory/Warehouses
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryWarehouses
type Warehouses struct {
	MetaData *api.MetaData `json:"__metadata,omitempty"`
	// ID: A guid that is the unique identifier of the warehouse
	ID *types.GUID `json:"ID,omitempty"`

	// Code: Code of the warehouse
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// DefaultStorageLocation: The default storage location of this warehouse. Warehouses can have a default storage location in packages Manufacturing Professional &amp; Premium or Wholesale Professional &amp; Premium
	DefaultStorageLocation *types.GUID `json:"DefaultStorageLocation,omitempty"`

	// DefaultStorageLocationCode: Default storage location&#39;s code
	DefaultStorageLocationCode *string `json:"DefaultStorageLocationCode,omitempty"`

	// DefaultStorageLocationDescription: Default storage location&#39;s description
	DefaultStorageLocationDescription *string `json:"DefaultStorageLocationDescription,omitempty"`

	// Description: The description of the warehouse
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EMail: Email address
	EMail *string `json:"EMail,omitempty"`

	// Main: Indicates if this is the main warehouse. There&#39;s always exactly one main warehouse per administration
	Main *byte `json:"Main,omitempty"`

	// ManagerUser: User reponsible for the warehouse
	ManagerUser *types.GUID `json:"ManagerUser,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// UseStorageLocations: Indicates if this warehouse is using storage locations. The storage locations will not be removed when when this is deactivated
	UseStorageLocations *byte `json:"UseStorageLocations,omitempty"`
}

func (e *Warehouses) GetPrimary() *types.GUID {
	return e.ID
}

func (s *WarehousesEndpoint) UserHasRights(ctx context.Context, division int, method string) (bool, error) {
	return s.client.UserHasRights(ctx, division, "inventory/Warehouses", method)
}

// List the Warehouses entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *WarehousesEndpoint) List(ctx context.Context, division int, all bool, o *api.ListOptions) ([]*Warehouses, error) {
	var entities []*Warehouses
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/Warehouses", division) // #nosec
	api.AddListOptionsToURL(u, o)

	if all {
		err := s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, err := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, &entities)
	return entities, err
}

// Get the Warehouses entitiy in the provided division.
func (s *WarehousesEndpoint) Get(ctx context.Context, division int, id *types.GUID) (*Warehouses, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/Warehouses", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return nil, err
	}

	e := &Warehouses{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "GET", u.String(), nil, e)
	return e, requestError
}

// New returns an empty Warehouses entity
func (s *WarehousesEndpoint) New() *Warehouses {
	return &Warehouses{}
}

// Create the Warehouses entity in the provided division.
func (s *WarehousesEndpoint) Create(ctx context.Context, division int, entity *Warehouses) (*Warehouses, error) {
	u, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/Warehouses", division) // #nosec
	e := &Warehouses{}
	_, _, err := s.client.NewRequestAndDo(ctx, "POST", u.String(), entity, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Update the Warehouses entity in the provided division.
func (s *WarehousesEndpoint) Update(ctx context.Context, division int, entity *Warehouses) (*Warehouses, error) {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/Warehouses", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, entity.GetPrimary())
	if err != nil {
		return nil, err
	}

	e := &Warehouses{}
	_, _, requestError := s.client.NewRequestAndDo(ctx, "PUT", u.String(), entity, e)
	return e, requestError
}

// Delete the Warehouses entity in the provided division.
func (s *WarehousesEndpoint) Delete(ctx context.Context, division int, id *types.GUID) error {
	b, _ := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/Warehouses", division) // #nosec
	u, err := api.AddOdataKeyToURL(b, id)
	if err != nil {
		return err
	}

	_, r, requestError := s.client.NewRequestAndDo(ctx, "DELETE", u.String(), nil, nil)
	if requestError != nil {
		return requestError
	}

	if r.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(r.Body) // #nosec
		return fmt.Errorf("Failed with status %v and body %v", r.StatusCode, body)
	}

	return nil
}
