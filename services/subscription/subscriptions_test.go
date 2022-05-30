// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package subscription

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

func SubscriptionsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func SubscriptionsEntityWithPopulatedPrimaryProperty() *Subscriptions {
	return &Subscriptions{EntryID: SubscriptionsPrimaryPropertySample()}
}

func SubscriptionsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func SubscriptionsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestSubscriptionsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &Subscriptions{EntryID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("SubscriptionsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestSubscriptionsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'subscription/Subscriptions'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.Subscriptions.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.Subscriptions.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.Subscriptions.UserHasRights should return true, got: %v", got)
	}
}

func TestSubscriptionsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.List returned error: %v, with url /api/v1/{division}/subscription/Subscriptions?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.List returned error: %v, with url /api/v1/{division}/subscription/Subscriptions?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SubscriptionsPrimaryPropertySample()
	gs := SubscriptionsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "EntryID": `+gs+`}]}}`)
		}
	})

	entities, err := s.Subscriptions.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("SubscriptionsEndpoint.List returned error: %v", err)
	}

	want := []*Subscriptions{{EntryID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SubscriptionsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSubscriptionsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.List returned error: %v, with url /api/v1/{division}/subscription/Subscriptions", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.List returned error: %v, with url /api/v1/{division}/subscription/Subscriptions", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := SubscriptionsPrimaryPropertySample()
	gs := SubscriptionsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "EntryID": `+gs+`}]}}`)
	})

	entities, err := s.Subscriptions.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("SubscriptionsEndpoint.List returned error: %v", err)
	}

	want := []*Subscriptions{{EntryID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("SubscriptionsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestSubscriptionsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := SubscriptionsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *Subscriptions
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&Subscriptions{EntryID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/subscription/Subscriptions", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SubscriptionsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.Subscriptions.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscriptionsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionsEndpoint_New(t *testing.T) {
	s, _, _, teardown := setup()
	defer teardown()
	got := s.Subscriptions.New()
	want := &Subscriptions{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SubscriptionsEndpoint.New() expected to return %v, got %v", want, got)
	}
}

func TestSubscriptionsEndpoint_Create(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *Subscriptions
	}
	tests := []struct {
		name    string
		args    args
		want    *Subscriptions
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &Subscriptions{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&Subscriptions{MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.Create returned error: %v, with url /api/v1/{division}/subscription/Subscriptions", e)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				testBody(t, r, `{"__metadata":{"uri":"https://start.exactonline.nl"}}`+"\n")
				fmt.Fprint(w, `{ "d": { "__metadata": { "uri": "https://start.exactonline.nl"}}}`)
			})

			got, err := s.Subscriptions.Create(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionsEndpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscriptionsEndpoint.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionsEndpoint_Update(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	type args struct {
		ctx      context.Context
		division int
		entity   *Subscriptions
	}
	s1 := SubscriptionsPrimaryPropertySample()
	tests := []struct {
		name    string
		args    args
		want    *Subscriptions
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, &Subscriptions{EntryID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}}},
			&Subscriptions{EntryID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.Update returned error: %v, with url /api/v1/{division}/subscription/Subscriptions", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.entity.GetPrimary())
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SubscriptionsEndpoint.Update returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PUT")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testHeader(t, r, "Content-Type", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.args.entity)
				testBody(t, r, string(b)+"\n")
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.Subscriptions.Update(tt.args.ctx, tt.args.division, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionsEndpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscriptionsEndpoint.Update() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestSubscriptionsEndpoint_Delete(t *testing.T) {
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
			args{context.Background(), 0, SubscriptionsPrimaryPropertySample()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/subscription/Subscriptions", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in SubscriptionsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/subscription/Subscriptions", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in SubscriptionsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				testBody(t, r, "")
				w.WriteHeader(http.StatusNoContent)
			})

			err := s.Subscriptions.Delete(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionsEndpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
