// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package crm

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

func AddressesPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func AddressesEntityWithPopulatedPrimaryProperty() *Addresses {
	return &Addresses{ID: AddressesPrimaryPropertySample()}
}

func AddressesStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func AddressesStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestAddressesEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &Addresses{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("AddressesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestAddressesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'crm/Addresses'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.Addresses.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.Addresses.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.Addresses.UserHasRights should return true, got: %v", got)
	}
}

func TestAddressesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.List returned error: %v, with url /api/v1/{division}/crm/Addresses?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.List returned error: %v, with url /api/v1/{division}/crm/Addresses?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := AddressesPrimaryPropertySample()
	gs := AddressesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.Addresses.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("AddressesEndpoint.List returned error: %v", err)
	}

	want := []*Addresses{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("AddressesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestAddressesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.List returned error: %v, with url /api/v1/{division}/crm/Addresses", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.List returned error: %v, with url /api/v1/{division}/crm/Addresses", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := AddressesPrimaryPropertySample()
	gs := AddressesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.Addresses.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("AddressesEndpoint.List returned error: %v", err)
	}

	want := []*Addresses{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("AddressesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestAddressesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := AddressesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *Addresses
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&Addresses{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/crm/Addresses", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in AddressesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.Addresses.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressesEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.Addresses.New()
	want := &Addresses{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddressesEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestAddressesEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *Addresses
	}
	tests := []struct {
		name    string
		args    args
		want    *Addresses
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &Addresses{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&Addresses{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.Create returned error: %v, with url /api/v1/{division}/crm/Addresses", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.Addresses.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressesEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressesEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressesEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *Addresses
	}
	s1 := AddressesPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *Addresses
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &Addresses{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&Addresses{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.Update returned error: %v, with url /api/v1/{division}/crm/Addresses", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in AddressesEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.Addresses.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressesEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressesEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestAddressesEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, AddressesPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Addresses", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in AddressesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/crm/Addresses", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in AddressesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.Addresses.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressesEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
