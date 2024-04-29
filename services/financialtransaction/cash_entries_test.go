// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financialtransaction

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

func CashEntriesPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func CashEntriesEntityWithPopulatedPrimaryProperty() *CashEntries {
	return &CashEntries{EntryID: CashEntriesPrimaryPropertySample()}
}

func CashEntriesStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func CashEntriesStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestCashEntriesEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &CashEntries{EntryID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("CashEntriesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestCashEntriesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CashEntriesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'financialtransaction/CashEntries'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.CashEntries.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.CashEntries.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.CashEntries.UserHasRights should return true, got: %v", got)
	}
}

func TestCashEntriesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/CashEntries", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CashEntriesEndpoint.List returned error: %v, with url /api/v1/{division}/financialtransaction/CashEntries?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/CashEntries", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CashEntriesEndpoint.List returned error: %v, with url /api/v1/{division}/financialtransaction/CashEntries?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := CashEntriesPrimaryPropertySample()
	gs := CashEntriesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "EntryID": `+gs+`}]}}`)
		}
	})

	entities, err := s.CashEntries.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("CashEntriesEndpoint.List returned error: %v", err)
	}

	want := []*CashEntries{{EntryID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("CashEntriesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestCashEntriesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/CashEntries", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CashEntriesEndpoint.List returned error: %v, with url /api/v1/{division}/financialtransaction/CashEntries", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/CashEntries", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CashEntriesEndpoint.List returned error: %v, with url /api/v1/{division}/financialtransaction/CashEntries", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := CashEntriesPrimaryPropertySample()
	gs := CashEntriesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "EntryID": `+gs+`}]}}`)
	})

	entities, err := s.CashEntries.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("CashEntriesEndpoint.List returned error: %v", err)
	}

	want := []*CashEntries{{EntryID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("CashEntriesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestCashEntriesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := CashEntriesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *CashEntries
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&CashEntries{EntryID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/CashEntries", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in CashEntriesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/financialtransaction/CashEntries", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in CashEntriesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.CashEntries.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CashEntriesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CashEntriesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCashEntriesEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.CashEntries.New()
	want := &CashEntries{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CashEntriesEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestCashEntriesEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *CashEntries
	}
	tests := []struct {
		name    string
		args    args
		want    *CashEntries
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &CashEntries{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&CashEntries{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/CashEntries", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in CashEntriesEndpoint.Create returned error: %v, with url /api/v1/{division}/financialtransaction/CashEntries", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.CashEntries.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("CashEntriesEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CashEntriesEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCashEntriesEndpoint_Delete(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, CashEntriesPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/financialtransaction/CashEntries", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in CashEntriesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/financialtransaction/CashEntries", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in CashEntriesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.CashEntries.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CashEntriesEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
