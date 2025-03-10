// Copyright 2022 The go-exactonline AUTHORS. All rights reserved.
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

func ShopOrderMaterialPlansPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func ShopOrderMaterialPlansEntityWithPopulatedPrimaryProperty() *ShopOrderMaterialPlans {
	return &ShopOrderMaterialPlans{ID: ShopOrderMaterialPlansPrimaryPropertySample()}
}

func ShopOrderMaterialPlansStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func ShopOrderMaterialPlansStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestShopOrderMaterialPlansEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &ShopOrderMaterialPlans{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("ShopOrderMaterialPlansEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestShopOrderMaterialPlansEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'manufacturing/ShopOrderMaterialPlans'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.ShopOrderMaterialPlans.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.ShopOrderMaterialPlans.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.ShopOrderMaterialPlans.UserHasRights should return true, got: %v", got)
	}
}

func TestShopOrderMaterialPlansEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.List returned error: %v, with url /api/v1/{division}/manufacturing/ShopOrderMaterialPlans?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.List returned error: %v, with url /api/v1/{division}/manufacturing/ShopOrderMaterialPlans?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ShopOrderMaterialPlansPrimaryPropertySample()
	gs := ShopOrderMaterialPlansStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.ShopOrderMaterialPlans.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("ShopOrderMaterialPlansEndpoint.List returned error: %v", err)
	}

	want := []*ShopOrderMaterialPlans{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ShopOrderMaterialPlansEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestShopOrderMaterialPlansEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.List returned error: %v, with url /api/v1/{division}/manufacturing/ShopOrderMaterialPlans", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.List returned error: %v, with url /api/v1/{division}/manufacturing/ShopOrderMaterialPlans", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := ShopOrderMaterialPlansPrimaryPropertySample()
	gs := ShopOrderMaterialPlansStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.ShopOrderMaterialPlans.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("ShopOrderMaterialPlansEndpoint.List returned error: %v", err)
	}

	want := []*ShopOrderMaterialPlans{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("ShopOrderMaterialPlansEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestShopOrderMaterialPlansEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := ShopOrderMaterialPlansPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *ShopOrderMaterialPlans
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&ShopOrderMaterialPlans{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.Delete() returned error: %v, with url /api/v1/{division}/manufacturing/ShopOrderMaterialPlans", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ShopOrderMaterialPlansEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.ShopOrderMaterialPlans.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShopOrderMaterialPlansEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShopOrderMaterialPlansEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShopOrderMaterialPlansEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.ShopOrderMaterialPlans.New()
	want := &ShopOrderMaterialPlans{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ShopOrderMaterialPlansEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestShopOrderMaterialPlansEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *ShopOrderMaterialPlans
	}
	tests := []struct {
		name    string
		args    args
		want    *ShopOrderMaterialPlans
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &ShopOrderMaterialPlans{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&ShopOrderMaterialPlans{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.Create returned error: %v, with url /api/v1/{division}/manufacturing/ShopOrderMaterialPlans", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.ShopOrderMaterialPlans.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShopOrderMaterialPlansEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShopOrderMaterialPlansEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShopOrderMaterialPlansEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *ShopOrderMaterialPlans
	}
	s1 := ShopOrderMaterialPlansPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *ShopOrderMaterialPlans
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &ShopOrderMaterialPlans{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&ShopOrderMaterialPlans{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.Update returned error: %v, with url /api/v1/{division}/manufacturing/ShopOrderMaterialPlans", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ShopOrderMaterialPlansEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.ShopOrderMaterialPlans.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShopOrderMaterialPlansEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShopOrderMaterialPlansEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestShopOrderMaterialPlansEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, ShopOrderMaterialPlansPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderMaterialPlans", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in ShopOrderMaterialPlansEndpoint.Delete() returned error: %v, with url /api/v1/{division}/manufacturing/ShopOrderMaterialPlans", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in ShopOrderMaterialPlansEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.ShopOrderMaterialPlans.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShopOrderMaterialPlansEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
