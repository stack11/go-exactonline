// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
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

func DocumentsAttachmentsPrimaryPropertySample() *types.GUID {
	v := types.NewGUID()
	return &v
}

func DocumentsAttachmentsEntityWithPopulatedPrimaryProperty() *DocumentsAttachments {
	return &DocumentsAttachments{ID: DocumentsAttachmentsPrimaryPropertySample()}
}

func DocumentsAttachmentsStringOfPrimaryProperty(v *types.GUID) string {
	return v.String()
}

func DocumentsAttachmentsStringJSONOfPrimaryProperty(v *types.GUID) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestDocumentsAttachmentsEndpoint_GetPrimary(t *testing.T) {
	var want types.GUID
	n := &DocumentsAttachments{ID: &want}

	if got := n.GetPrimary(); !reflect.DeepEqual(*got, want) {
		t.Errorf("DocumentsAttachmentsEndpoint.GetPrimary() failed, got: %v, want: %v", *got, want)
	}
}

func TestDocumentsAttachmentsEndpoint_UserHasRights(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserHasRights", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentsAttachmentsEndpoint.List returned error: %v", e)
	}

	acceptHeaders := []string{"application/json"}

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))

		q := r.URL.Query()

		if got, want := q.Get("endpoint"), "'crm/DocumentsAttachments'"; got != want {
			t.Errorf("endpoint query param doesn't match, got: %v, want: %v", got, want)
		}

		if got, want := q.Get("method"), "GET"; got != want {
			t.Errorf("method query param doesn't match, got: %v, want: %v", got, want)
		}

		fmt.Fprint(w, `{ "d": { "UserHasRights": true } }`)
	})

	got, err := s.DocumentsAttachments.UserHasRights(context.Background(), 0, "GET")
	if err != nil {
		t.Errorf("s.DocumentsAttachments.UserHasRights should not return an error = %v", err)
	}

	if got != true {
		t.Errorf("s.DocumentsAttachments.UserHasRights should return true, got: %v", got)
	}
}

func TestDocumentsAttachmentsEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/crm/DocumentsAttachments", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentsAttachmentsEndpoint.List returned error: %v, with url /api/v1/{division}/read/crm/DocumentsAttachments?$select=*", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/read/crm/DocumentsAttachments", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentsAttachmentsEndpoint.List returned error: %v, with url /api/v1/{division}/read/crm/DocumentsAttachments?$skiptoken=foo", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := DocumentsAttachmentsPrimaryPropertySample()
	gs := DocumentsAttachmentsStringJSONOfPrimaryProperty(g)

	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
		}
	})

	entities, err := s.DocumentsAttachments.List(context.Background(), 0, true, opts1)
	if err != nil {
		t.Errorf("DocumentsAttachmentsEndpoint.List returned error: %v", err)
	}

	want := []*DocumentsAttachments{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("DocumentsAttachmentsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestDocumentsAttachmentsEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	opts1 := api.NewListOptions()
	opts1.Select.Add("*")
	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/crm/DocumentsAttachments", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentsAttachmentsEndpoint.List returned error: %v, with url /api/v1/{division}/read/crm/DocumentsAttachments", e)
	}
	api.AddListOptionsToURL(u, opts1)

	opts2 := api.NewListOptions()
	opts2.Select.Add("*")
	opts2.SkipToken.Set(types.NewGUID())
	u2, e2 := s.client.ResolvePathWithDivision("/api/v1/{division}/read/crm/DocumentsAttachments", 0)
	if e2 != nil {
		t.Errorf("s.client.ResolvePathWithDivision in DocumentsAttachmentsEndpoint.List returned error: %v, with url /api/v1/{division}/read/crm/DocumentsAttachments", e2)
	}
	api.AddListOptionsToURL(u2, opts2)

	g := DocumentsAttachmentsPrimaryPropertySample()
	gs := DocumentsAttachmentsStringJSONOfPrimaryProperty(g)
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": `+gs+`}]}}`)
	})

	entities, err := s.DocumentsAttachments.List(context.Background(), 0, false, opts1)
	if err != nil {
		t.Errorf("DocumentsAttachmentsEndpoint.List returned error: %v", err)
	}

	want := []*DocumentsAttachments{{ID: g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("DocumentsAttachmentsEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestDocumentsAttachmentsEndpoint_Get(t *testing.T) {
	acceptHeaders := []string{"application/json"}
	s1 := DocumentsAttachmentsPrimaryPropertySample()
	type args struct {
		ctx      context.Context
		division int
		id       *types.GUID
	}
	tests := []struct {
		name    string
		args    args
		want    *DocumentsAttachments
		wantErr bool
	}{
		{
			"1",
			args{context.Background(), 0, s1},
			&DocumentsAttachments{ID: s1, MetaData: &api.MetaData{URI: &types.URL{URL: &url.URL{Scheme: "https", Host: "start.exactonline.nl"}}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, mux, _, teardown := setup()
			defer teardown()

			b, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/crm/DocumentsAttachments", 0)
			if e != nil {
				t.Errorf("s.client.ResolvePathWithDivision in DocumentsAttachmentsEndpoint.Delete() returned error: %v, with url /api/v1/{division}/read/crm/DocumentsAttachments", e)
			}

			u, e2 := api.AddOdataKeyToURL(b, tt.args.id)
			if e2 != nil {
				t.Errorf("api.AddOdataKeyToURL in DocumentsAttachmentsEndpoint.Delete() returned error: %v", e2)
			}

			mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
				b, _ := json.Marshal(tt.want)
				fmt.Fprint(w, `{"d":`+string(b)+`}`)
			})

			got, err := s.DocumentsAttachments.Get(tt.args.ctx, tt.args.division, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocumentsAttachmentsEndpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentsAttachmentsEndpoint.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
