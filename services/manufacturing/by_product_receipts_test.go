// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

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

func ByProductReceiptsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func ByProductReceiptsEntityWithPopulatedPrimaryProperty() *ByProductReceipts {
	return &ByProductReceipts{StockTransactionId: ByProductReceiptsPrimaryPropertySample()}
}

func ByProductReceiptsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func ByProductReceiptsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestByProductReceiptsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &ByProductReceipts{StockTransactionId: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("ByProductReceiptsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestByProductReceiptsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ByProductReceiptsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'manufacturing/ByProductReceipts'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.ByProductReceipts.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.ByProductReceipts.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.ByProductReceipts.UserHasRights should return true, got: %v", got)
	}
}

func TestByProductReceiptsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReceipts", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ByProductReceiptsEndpoint.List returned error: %v, with url /api/v1/{division}/manufacturing/ByProductReceipts?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReceipts", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ByProductReceiptsEndpoint.List returned error: %v, with url /api/v1/{division}/manufacturing/ByProductReceipts?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ByProductReceiptsPrimaryPropertySample()
	gs := ByProductReceiptsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "StockTransactionId": `+gs+`}]}}`)
		}
	})

	entities, err := s.ByProductReceipts.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("ByProductReceiptsEndpoint.List returned error: %v", err)
	}

	want := []*ByProductReceipts{{StockTransactionId: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ByProductReceiptsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestByProductReceiptsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReceipts", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ByProductReceiptsEndpoint.List returned error: %v, with url /api/v1/{division}/manufacturing/ByProductReceipts", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReceipts", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ByProductReceiptsEndpoint.List returned error: %v, with url /api/v1/{division}/manufacturing/ByProductReceipts", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ByProductReceiptsPrimaryPropertySample()
	gs := ByProductReceiptsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "StockTransactionId": `+gs+`}]}}`)
	})

	entities, err := s.ByProductReceipts.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("ByProductReceiptsEndpoint.List returned error: %v", err)
	}

	want := []*ByProductReceipts{{StockTransactionId: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ByProductReceiptsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestByProductReceiptsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := ByProductReceiptsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *ByProductReceipts
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&ByProductReceipts{StockTransactionId: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReceipts", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ByProductReceiptsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/manufacturing/ByProductReceipts", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ByProductReceiptsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.ByProductReceipts.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByProductReceiptsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByProductReceiptsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByProductReceiptsEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.ByProductReceipts.New()
	want := &ByProductReceipts{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ByProductReceiptsEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestByProductReceiptsEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *ByProductReceipts
	}
	tests := []struct {
		name    string
		args    args
		want    *ByProductReceipts
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &ByProductReceipts{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&ByProductReceipts{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ByProductReceipts", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ByProductReceiptsEndpoint.Create returned error: %v, with url /api/v1/{division}/manufacturing/ByProductReceipts", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.ByProductReceipts.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByProductReceiptsEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByProductReceiptsEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
