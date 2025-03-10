// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package documents

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

func DocumentTypeCategoriesPrimaryPropertySample() *int {
	v := 100
	return &v
}

func DocumentTypeCategoriesEntityWithPopulatedPrimaryProperty() *DocumentTypeCategories {
	return &DocumentTypeCategories{ID: DocumentTypeCategoriesPrimaryPropertySample()}
}

func DocumentTypeCategoriesStringOfPrimaryProperty(v *int) string {
	return strconv.Itoa(*v)
}

func DocumentTypeCategoriesStringJSONOfPrimaryProperty(v *int) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestDocumentTypeCategoriesEndpoint_GetPrimary(t *testing.T) {
	var want int
	n := &DocumentTypeCategories{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("DocumentTypeCategoriesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestDocumentTypeCategoriesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentTypeCategoriesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'documents/DocumentTypeCategories'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.DocumentTypeCategories.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.DocumentTypeCategories.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.DocumentTypeCategories.UserHasRights should return true, got: %v", got)
	}
}

func TestDocumentTypeCategoriesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypeCategories", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentTypeCategoriesEndpoint.List returned error: %v, with url /api/v1/{division}/documents/DocumentTypeCategories?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypeCategories", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentTypeCategoriesEndpoint.List returned error: %v, with url /api/v1/{division}/documents/DocumentTypeCategories?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := DocumentTypeCategoriesPrimaryPropertySample()
	gs := DocumentTypeCategoriesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.DocumentTypeCategories.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("DocumentTypeCategoriesEndpoint.List returned error: %v", err)
	}

	want := []*DocumentTypeCategories{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("DocumentTypeCategoriesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestDocumentTypeCategoriesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypeCategories", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentTypeCategoriesEndpoint.List returned error: %v, with url /api/v1/{division}/documents/DocumentTypeCategories", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypeCategories", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentTypeCategoriesEndpoint.List returned error: %v, with url /api/v1/{division}/documents/DocumentTypeCategories", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := DocumentTypeCategoriesPrimaryPropertySample()
	gs := DocumentTypeCategoriesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.DocumentTypeCategories.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("DocumentTypeCategoriesEndpoint.List returned error: %v", err)
	}

	want := []*DocumentTypeCategories{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("DocumentTypeCategoriesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestDocumentTypeCategoriesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := DocumentTypeCategoriesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *int
	}
	tests := []struct {
		name    string
		args    args
		want    *DocumentTypeCategories
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&DocumentTypeCategories{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/documents/DocumentTypeCategories", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in DocumentTypeCategoriesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/documents/DocumentTypeCategories", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in DocumentTypeCategoriesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.DocumentTypeCategories.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentTypeCategoriesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentTypeCategoriesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
