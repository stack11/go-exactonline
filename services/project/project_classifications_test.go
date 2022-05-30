// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
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
	"strings"
	"testing"

	"github.com/stack11/go-exactonline/api"
	"github.com/stack11/go-exactonline/types"
)

func ProjectClassificationsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func ProjectClassificationsEntityWithPopulatedPrimaryProperty() *ProjectClassifications {
	return &ProjectClassifications{ID: ProjectClassificationsPrimaryPropertySample()}
}

func ProjectClassificationsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func ProjectClassificationsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestProjectClassificationsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &ProjectClassifications{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("ProjectClassificationsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestProjectClassificationsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'project/ProjectClassifications'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.ProjectClassifications.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.ProjectClassifications.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.ProjectClassifications.UserHasRights should return true, got: %v", got)
	}
}

func TestProjectClassificationsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectClassifications", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectClassifications?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectClassifications", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectClassifications?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ProjectClassificationsPrimaryPropertySample()
	gs := ProjectClassificationsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.ProjectClassifications.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("ProjectClassificationsEndpoint.List returned error: %v", err)
	}

	want := []*ProjectClassifications{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ProjectClassificationsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestProjectClassificationsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectClassifications", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectClassifications", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectClassifications", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.List returned error: %v, with url /api/v1/{division}/project/ProjectClassifications", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ProjectClassificationsPrimaryPropertySample()
	gs := ProjectClassificationsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.ProjectClassifications.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("ProjectClassificationsEndpoint.List returned error: %v", err)
	}

	want := []*ProjectClassifications{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ProjectClassificationsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestProjectClassificationsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := ProjectClassificationsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *ProjectClassifications
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&ProjectClassifications{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectClassifications", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/project/ProjectClassifications", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ProjectClassificationsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.ProjectClassifications.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectClassificationsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectClassificationsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectClassificationsEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.ProjectClassifications.New()
	want := &ProjectClassifications{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProjectClassificationsEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestProjectClassificationsEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *ProjectClassifications
	}
	tests := []struct {
		name    string
		args    args
		want    *ProjectClassifications
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &ProjectClassifications{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&ProjectClassifications{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectClassifications", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.Create returned error: %v, with url /api/v1/{division}/project/ProjectClassifications", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.ProjectClassifications.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectClassificationsEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectClassificationsEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectClassificationsEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *ProjectClassifications
	}
	s1 := ProjectClassificationsPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *ProjectClassifications
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &ProjectClassifications{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&ProjectClassifications{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectClassifications", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.Update returned error: %v, with url /api/v1/{division}/project/ProjectClassifications", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ProjectClassificationsEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.ProjectClassifications.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectClassificationsEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProjectClassificationsEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestProjectClassificationsEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, ProjectClassificationsPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/project/ProjectClassifications", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ProjectClassificationsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/project/ProjectClassifications", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ProjectClassificationsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.ProjectClassifications.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectClassificationsEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
