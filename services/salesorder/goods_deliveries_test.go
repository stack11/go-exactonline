// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesorder

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

func GoodsDeliveriesPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func GoodsDeliveriesEntityWithPopulatedPrimaryProperty() *GoodsDeliveries {
	return &GoodsDeliveries{EntryID: GoodsDeliveriesPrimaryPropertySample()}
}

func GoodsDeliveriesStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func GoodsDeliveriesStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestGoodsDeliveriesEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &GoodsDeliveries{EntryID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("GoodsDeliveriesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestGoodsDeliveriesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in GoodsDeliveriesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'salesorder/GoodsDeliveries'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.GoodsDeliveries.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.GoodsDeliveries.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.GoodsDeliveries.UserHasRights should return true, got: %v", got)
	}
}

func TestGoodsDeliveriesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveries", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in GoodsDeliveriesEndpoint.List returned error: %v, with url /api/v1/{division}/salesorder/GoodsDeliveries?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveries", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in GoodsDeliveriesEndpoint.List returned error: %v, with url /api/v1/{division}/salesorder/GoodsDeliveries?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := GoodsDeliveriesPrimaryPropertySample()
	gs := GoodsDeliveriesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "EntryID": `+gs+`}]}}`)
		}
	})

	entities, err := s.GoodsDeliveries.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("GoodsDeliveriesEndpoint.List returned error: %v", err)
	}

	want := []*GoodsDeliveries{{EntryID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("GoodsDeliveriesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestGoodsDeliveriesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveries", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in GoodsDeliveriesEndpoint.List returned error: %v, with url /api/v1/{division}/salesorder/GoodsDeliveries", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveries", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in GoodsDeliveriesEndpoint.List returned error: %v, with url /api/v1/{division}/salesorder/GoodsDeliveries", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := GoodsDeliveriesPrimaryPropertySample()
	gs := GoodsDeliveriesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "EntryID": `+gs+`}]}}`)
	})

	entities, err := s.GoodsDeliveries.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("GoodsDeliveriesEndpoint.List returned error: %v", err)
	}

	want := []*GoodsDeliveries{{EntryID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("GoodsDeliveriesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestGoodsDeliveriesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := GoodsDeliveriesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *GoodsDeliveries
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&GoodsDeliveries{EntryID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveries", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in GoodsDeliveriesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/salesorder/GoodsDeliveries", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in GoodsDeliveriesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.GoodsDeliveries.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoodsDeliveriesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoodsDeliveriesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoodsDeliveriesEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.GoodsDeliveries.New()
	want := &GoodsDeliveries{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GoodsDeliveriesEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestGoodsDeliveriesEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *GoodsDeliveries
	}
	tests := []struct {
		name    string
		args    args
		want    *GoodsDeliveries
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &GoodsDeliveries{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&GoodsDeliveries{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveries", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in GoodsDeliveriesEndpoint.Create returned error: %v, with url /api/v1/{division}/salesorder/GoodsDeliveries", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.GoodsDeliveries.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoodsDeliveriesEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoodsDeliveriesEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoodsDeliveriesEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *GoodsDeliveries
	}
	s1 := GoodsDeliveriesPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *GoodsDeliveries
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &GoodsDeliveries{EntryID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&GoodsDeliveries{EntryID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/GoodsDeliveries", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in GoodsDeliveriesEndpoint.Update returned error: %v, with url /api/v1/{division}/salesorder/GoodsDeliveries", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in GoodsDeliveriesEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.GoodsDeliveries.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoodsDeliveriesEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoodsDeliveriesEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
