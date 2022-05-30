// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

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

func ProjectBudgetTypesPrimaryPropertySample() *int {
	v := 100
	return &v
}

func ProjectBudgetTypesEntityWithPopulatedPrimaryProperty() *ProjectBudgetTypes {
	return &ProjectBudgetTypes{ID: ProjectBudgetTypesPrimaryPropertySample()}
}

func ProjectBudgetTypesStringOfPrimaryProperty(v *int) string {
	return strconv.Itoa(*v)
}

func ProjectBudgetTypesStringJSONOfPrimaryProperty(v *int) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestProjectBudgetTypesEndpoint_GetPrimary(t *testing.T) {
	var want int
	n := &ProjectBudgetTypes{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("ProjectBudgetTypesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestProjectBudgetTypesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectBudgetTypesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'project/ProjectBudgetTypes'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.ProjectBudgetTypes.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.ProjectBudgetTypes.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.ProjectBudgetTypes.UserHasRights should return true, got: %v", got)
	}
}

func TestProjectBudgetTypesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectBudgetTypes", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectBudgetTypesEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectBudgetTypes?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectBudgetTypes", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectBudgetTypesEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectBudgetTypes?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ProjectBudgetTypesPrimaryPropertySample()
	gs := ProjectBudgetTypesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.ProjectBudgetTypes.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("ProjectBudgetTypesEndpoint.List returned error: %v", err)
	}

	want := []*ProjectBudgetTypes{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ProjectBudgetTypesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestProjectBudgetTypesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectBudgetTypes", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectBudgetTypesEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectBudgetTypes", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectBudgetTypes", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectBudgetTypesEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectBudgetTypes", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ProjectBudgetTypesPrimaryPropertySample()
	gs := ProjectBudgetTypesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.ProjectBudgetTypes.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("ProjectBudgetTypesEndpoint.List returned error: %v", err)
	}

	want := []*ProjectBudgetTypes{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ProjectBudgetTypesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestProjectBudgetTypesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := ProjectBudgetTypesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *int
	}
	tests := []struct {
		name    string
		args    args
		want    *ProjectBudgetTypes
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&ProjectBudgetTypes{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectBudgetTypes", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectBudgetTypesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/project/ProjectBudgetTypes", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ProjectBudgetTypesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.ProjectBudgetTypes.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectBudgetTypesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectBudgetTypesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
