// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package logistics

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

func CustomerItemsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func CustomerItemsEntityWithPopulatedPrimaryProperty() *CustomerItems {
	return &CustomerItems{ID: CustomerItemsPrimaryPropertySample()}
}

func CustomerItemsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func CustomerItemsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestCustomerItemsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &CustomerItems{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("CustomerItemsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestCustomerItemsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'logistics/CustomerItems'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.CustomerItems.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.CustomerItems.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.CustomerItems.UserHasRights should return true, got: %v", got)
	}
}

func TestCustomerItemsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/CustomerItems", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.List returned error: %v, with url /api/v1/{division}/logistics/CustomerItems?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/CustomerItems", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.List returned error: %v, with url /api/v1/{division}/logistics/CustomerItems?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := CustomerItemsPrimaryPropertySample()
	gs := CustomerItemsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.CustomerItems.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("CustomerItemsEndpoint.List returned error: %v", err)
	}

	want := []*CustomerItems{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("CustomerItemsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestCustomerItemsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/CustomerItems", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.List returned error: %v, with url /api/v1/{division}/logistics/CustomerItems", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/CustomerItems", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.List returned error: %v, with url /api/v1/{division}/logistics/CustomerItems", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := CustomerItemsPrimaryPropertySample()
	gs := CustomerItemsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.CustomerItems.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("CustomerItemsEndpoint.List returned error: %v", err)
	}

	want := []*CustomerItems{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("CustomerItemsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestCustomerItemsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := CustomerItemsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *CustomerItems
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&CustomerItems{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/CustomerItems", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/logistics/CustomerItems", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in CustomerItemsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.CustomerItems.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerItemsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerItemsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerItemsEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.CustomerItems.New()
	want := &CustomerItems{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CustomerItemsEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestCustomerItemsEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *CustomerItems
	}
	tests := []struct {
		name    string
		args    args
		want    *CustomerItems
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &CustomerItems{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&CustomerItems{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/CustomerItems", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.Create returned error: %v, with url /api/v1/{division}/logistics/CustomerItems", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.CustomerItems.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerItemsEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerItemsEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerItemsEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *CustomerItems
	}
	s1 := CustomerItemsPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *CustomerItems
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &CustomerItems{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&CustomerItems{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/CustomerItems", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.Update returned error: %v, with url /api/v1/{division}/logistics/CustomerItems", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in CustomerItemsEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.CustomerItems.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerItemsEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerItemsEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestCustomerItemsEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, CustomerItemsPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/CustomerItems", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in CustomerItemsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/logistics/CustomerItems", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in CustomerItemsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.CustomerItems.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerItemsEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
