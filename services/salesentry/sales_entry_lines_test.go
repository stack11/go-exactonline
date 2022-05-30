// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesentry

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

func SalesEntryLinesPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func SalesEntryLinesEntityWithPopulatedPrimaryProperty() *SalesEntryLines {
	return &SalesEntryLines{ID: SalesEntryLinesPrimaryPropertySample()}
}

func SalesEntryLinesStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func SalesEntryLinesStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestSalesEntryLinesEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &SalesEntryLines{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("SalesEntryLinesEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestSalesEntryLinesEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'salesentry/SalesEntryLines'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.SalesEntryLines.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.SalesEntryLines.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.SalesEntryLines.UserHasRights should return true, got: %v", got)
	}
}

func TestSalesEntryLinesEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SalesEntryLinesPrimaryPropertySample()
	gs := SalesEntryLinesStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.SalesEntryLines.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("SalesEntryLinesEndpoint.List returned error: %v", err)
	}

	want := []*SalesEntryLines{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SalesEntryLinesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSalesEntryLinesEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.List returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SalesEntryLinesPrimaryPropertySample()
	gs := SalesEntryLinesStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.SalesEntryLines.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("SalesEntryLinesEndpoint.List returned error: %v", err)
	}

	want := []*SalesEntryLines{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SalesEntryLinesEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSalesEntryLinesEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := SalesEntryLinesPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *SalesEntryLines
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&SalesEntryLines{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SalesEntryLinesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.SalesEntryLines.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesEntryLinesEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesEntryLinesEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesEntryLinesEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.SalesEntryLines.New()
	want := &SalesEntryLines{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SalesEntryLinesEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestSalesEntryLinesEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *SalesEntryLines
	}
	tests := []struct {
		name    string
		args    args
		want    *SalesEntryLines
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &SalesEntryLines{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&SalesEntryLines{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.Create returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.SalesEntryLines.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesEntryLinesEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesEntryLinesEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesEntryLinesEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *SalesEntryLines
	}
	s1 := SalesEntryLinesPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *SalesEntryLines
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &SalesEntryLines{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&SalesEntryLines{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.Update returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SalesEntryLinesEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.SalesEntryLines.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesEntryLinesEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SalesEntryLinesEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestSalesEntryLinesEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, SalesEntryLinesPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/salesentry/SalesEntryLines", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SalesEntryLinesEndpoint.Delete() returned error: %v, with url /api/v1/{division}/salesentry/SalesEntryLines", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SalesEntryLinesEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.SalesEntryLines.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SalesEntryLinesEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
