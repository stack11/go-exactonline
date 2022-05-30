// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package bulk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

func SalesOrderGoodsDeliveriesPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func SalesOrderGoodsDeliveriesEntityWithPopulatedPrimaryProperty() *SalesOrderGoodsDeliveries {
	return &SalesOrderGoodsDeliveries{EntryID: SalesOrderGoodsDeliveriesPrimaryPropertySample()}
}

func SalesOrderGoodsDeliveriesStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func SalesOrderGoodsDeliveriesStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestSalesOrderGoodsDeliveriesEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &SalesOrderGoodsDeliveries{EntryID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("SalesOrderGoodsDeliveriesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestSalesOrderGoodsDeliveriesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesOrderGoodsDeliveriesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'SalesOrder/GoodsDeliveries'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.SalesOrderGoodsDeliveries.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.SalesOrderGoodsDeliveries.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.SalesOrderGoodsDeliveries.UserHasRights should return true, got: %v", got)
	}
}

func TestSalesOrderGoodsDeliveriesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/SalesOrder/GoodsDeliveries", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesOrderGoodsDeliveriesEndpoint.List returned error: %v, with url /api/v1/{division}/bulk/SalesOrder/GoodsDeliveries?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/SalesOrder/GoodsDeliveries", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesOrderGoodsDeliveriesEndpoint.List returned error: %v, with url /api/v1/{division}/bulk/SalesOrder/GoodsDeliveries?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SalesOrderGoodsDeliveriesPrimaryPropertySample()
	gs := SalesOrderGoodsDeliveriesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "EntryID": `+gs+`}]}}`)
		}
	})

	entities, err := s.SalesOrderGoodsDeliveries.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("SalesOrderGoodsDeliveriesEndpoint.List returned error: %v", err)
	}

	want := []*SalesOrderGoodsDeliveries{{EntryID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SalesOrderGoodsDeliveriesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSalesOrderGoodsDeliveriesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/SalesOrder/GoodsDeliveries", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesOrderGoodsDeliveriesEndpoint.List returned error: %v, with url /api/v1/{division}/bulk/SalesOrder/GoodsDeliveries", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/SalesOrder/GoodsDeliveries", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesOrderGoodsDeliveriesEndpoint.List returned error: %v, with url /api/v1/{division}/bulk/SalesOrder/GoodsDeliveries", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SalesOrderGoodsDeliveriesPrimaryPropertySample()
	gs := SalesOrderGoodsDeliveriesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "EntryID": `+gs+`}]}}`)
	})

	entities, err := s.SalesOrderGoodsDeliveries.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("SalesOrderGoodsDeliveriesEndpoint.List returned error: %v", err)
	}

	want := []*SalesOrderGoodsDeliveries{{EntryID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SalesOrderGoodsDeliveriesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSalesOrderGoodsDeliveriesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := SalesOrderGoodsDeliveriesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *SalesOrderGoodsDeliveries
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&SalesOrderGoodsDeliveries{EntryID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/bulk/SalesOrder/GoodsDeliveries", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesOrderGoodsDeliveriesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/bulk/SalesOrder/GoodsDeliveries", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SalesOrderGoodsDeliveriesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.SalesOrderGoodsDeliveries.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesOrderGoodsDeliveriesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesOrderGoodsDeliveriesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
