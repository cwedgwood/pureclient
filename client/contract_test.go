package client_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	client "github.com/cwedgwood/pureclient/client"
)

//go:fix inline
func ptr[T any](value T) *T {
	return new(value)
}

func contractClient(t *testing.T, server *httptest.Server) *client.ClientWithResponses {
	t.Helper()

	c, err := client.NewClientWithResponses(server.URL, client.WithHTTPClient(server.Client()))
	if err != nil {
		t.Fatalf("NewClientWithResponses: %v", err)
	}
	return c
}

func TestContractLogin(t *testing.T) {
	const (
		apiToken     = "0f2e2884-9486-c6c2-438c-f50418f2aac3"
		requestID    = "contract-login-request"
		sessionToken = "session-token"
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want %q", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/api/2.26/login" {
			t.Errorf("path = %q, want %q", r.URL.Path, "/api/2.26/login")
		}
		if r.URL.RawQuery != "" {
			t.Errorf("query = %q, want empty", r.URL.RawQuery)
		}
		if got := r.Header.Get("api-token"); got != apiToken {
			t.Errorf("api-token = %q, want %q", got, apiToken)
		}
		if got := r.Header.Get("X-Request-ID"); got != requestID {
			t.Errorf("X-Request-ID = %q, want %q", got, requestID)
		}
		if got := r.Header.Get("Content-Type"); got != "" {
			t.Errorf("Content-Type = %q, want empty", got)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("x-auth-token", sessionToken)
		_, _ = io.WriteString(w, `{"items":[{"username":"pureuser"}]}`)
	}))
	defer server.Close()

	response, err := contractClient(t, server).PostApi226LoginWithResponse(
		context.Background(),
		&client.PostApi226LoginParams{
			ApiToken:   ptr(apiToken),
			XRequestID: ptr(requestID),
		},
	)
	if err != nil {
		t.Fatalf("PostApi226LoginWithResponse: %v", err)
	}
	if response.StatusCode() != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode(), http.StatusOK)
	}
	if got := response.HTTPResponse.Header.Get("x-auth-token"); got != sessionToken {
		t.Errorf("x-auth-token = %q, want %q", got, sessionToken)
	}
	if response.JSON200 == nil || response.JSON200.Items == nil || len(*response.JSON200.Items) != 1 {
		t.Fatalf("decoded response items = %#v, want one item", response.JSON200)
	}
	if got := (*response.JSON200.Items)[0].Username; got == nil || *got != "pureuser" {
		t.Errorf("decoded username = %v, want pureuser", got)
	}
}

func TestContractListVolumes(t *testing.T) {
	const filter = "name='db vol' and provisioned>1048576"
	names := []string{"vol one", "vol/two"}
	sortFields := []string{"provisioned-", "name"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want %q", r.Method, http.MethodGet)
		}
		if r.URL.Path != "/api/2.26/volumes" {
			t.Errorf("path = %q, want %q", r.URL.Path, "/api/2.26/volumes")
		}

		query := r.URL.Query()
		wantQuery := map[string]string{
			"destroyed":        "false",
			"filter":           filter,
			"limit":            "0",
			"names":            strings.Join(names, ","),
			"sort":             strings.Join(sortFields, ","),
			"total_item_count": "true",
			"total_only":       "false",
		}
		for key, want := range wantQuery {
			if got := query.Get(key); got != want {
				t.Errorf("query parameter %s = %q, want %q", key, got, want)
			}
		}
		if len(query) != len(wantQuery) {
			t.Errorf("query parameters = %#v, want exactly %#v", query, wantQuery)
		}
		for _, encoded := range []string{
			"filter=name%3D%27db+vol%27+and+provisioned%3E1048576",
			"names=vol+one,vol%2Ftwo",
			"sort=provisioned-,name",
		} {
			if !strings.Contains(r.URL.RawQuery, encoded) {
				t.Errorf("raw query %q does not contain %q", r.URL.RawQuery, encoded)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{
			"continuation_token":"next-page",
			"items":[{
				"id":"11111111-2222-3333-4444-555555555555",
				"name":"vol one",
				"destroyed":false,
				"provisioned":1073741824,
				"serial":"A_REALISTIC_SERIAL",
				"subtype":"regular"
			}],
			"more_items_remaining":false,
			"total_item_count":1
		}`)
	}))
	defer server.Close()

	response, err := contractClient(t, server).GetApi226VolumesWithResponse(
		context.Background(),
		&client.GetApi226VolumesParams{
			Destroyed:      new(false),
			Filter:         ptr(filter),
			Limit:          new(int32(0)),
			Names:          &names,
			Sort:           &sortFields,
			TotalItemCount: new(true),
			TotalOnly:      new(false),
		},
	)
	if err != nil {
		t.Fatalf("GetApi226VolumesWithResponse: %v", err)
	}
	if response.JSON200 == nil || response.JSON200.Items == nil || len(*response.JSON200.Items) != 1 {
		t.Fatalf("decoded response items = %#v, want one item", response.JSON200)
	}
	volume := (*response.JSON200.Items)[0]
	if volume.Name == nil || *volume.Name != "vol one" {
		t.Errorf("decoded name = %v, want vol one", volume.Name)
	}
	if volume.Id == nil || *volume.Id != "11111111-2222-3333-4444-555555555555" {
		t.Errorf("decoded id = %v, want fixture UUID", volume.Id)
	}
	if volume.Provisioned == nil || *volume.Provisioned != 1073741824 {
		t.Errorf("decoded provisioned = %v, want 1073741824", volume.Provisioned)
	}
	if volume.Destroyed == nil || *volume.Destroyed {
		t.Errorf("decoded destroyed = %v, want false", volume.Destroyed)
	}
	if response.JSON200.MoreItemsRemaining == nil || *response.JSON200.MoreItemsRemaining {
		t.Errorf("decoded more_items_remaining = %v, want false", response.JSON200.MoreItemsRemaining)
	}
	if response.JSON200.TotalItemCount == nil || *response.JSON200.TotalItemCount != 1 {
		t.Errorf("decoded total_item_count = %v, want 1", response.JSON200.TotalItemCount)
	}
	if response.JSON200.ContinuationToken == nil || *response.JSON200.ContinuationToken != "next-page" {
		t.Errorf("decoded continuation_token = %v, want next-page", response.JSON200.ContinuationToken)
	}
}

func TestContractVolumeWrites(t *testing.T) {
	tests := []struct {
		name      string
		method    string
		rawQuery  string
		call      func(*client.ClientWithResponses) (*client.VolumeResponse, error)
		wantBody  map[string]any
		replyName string
	}{
		{
			name:     "create",
			method:   http.MethodPost,
			rawQuery: "names=new+volume&overwrite=false&with_default_protection=false",
			call: func(c *client.ClientWithResponses) (*client.VolumeResponse, error) {
				response, err := c.PostApi226VolumesWithResponse(
					context.Background(),
					&client.PostApi226VolumesParams{
						Names:                 &[]string{"new volume"},
						Overwrite:             new(false),
						WithDefaultProtection: new(false),
					},
					client.VolumePost{
						Destroyed:   new(false),
						Provisioned: new(int64(0)),
					},
				)
				if err != nil {
					return nil, err
				}
				return response.JSON200, nil
			},
			wantBody: map[string]any{
				"destroyed":   false,
				"provisioned": float64(0),
			},
			replyName: "new volume",
		},
		{
			name:     "update",
			method:   http.MethodPatch,
			rawQuery: "names=new+volume&truncate=false",
			call: func(c *client.ClientWithResponses) (*client.VolumeResponse, error) {
				response, err := c.PatchApi226VolumesWithResponse(
					context.Background(),
					&client.PatchApi226VolumesParams{
						Names:    &[]string{"new volume"},
						Truncate: new(false),
					},
					client.VolumePatch{
						Destroyed:   new(false),
						Name:        new("renamed volume"),
						Provisioned: new(int64(0)),
					},
				)
				if err != nil {
					return nil, err
				}
				return response.JSON200, nil
			},
			wantBody: map[string]any{
				"destroyed":   false,
				"name":        "renamed volume",
				"provisioned": float64(0),
			},
			replyName: "renamed volume",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != test.method {
					t.Errorf("method = %q, want %q", r.Method, test.method)
				}
				if r.URL.Path != "/api/2.26/volumes" {
					t.Errorf("path = %q, want %q", r.URL.Path, "/api/2.26/volumes")
				}
				if r.URL.RawQuery != test.rawQuery {
					t.Errorf("raw query = %q, want %q", r.URL.RawQuery, test.rawQuery)
				}
				if got := r.Header.Get("Content-Type"); got != "application/json" {
					t.Errorf("Content-Type = %q, want application/json", got)
				}

				var gotBody map[string]any
				if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
					t.Errorf("decode request body: %v", err)
				}
				if !reflect.DeepEqual(gotBody, test.wantBody) {
					t.Errorf("request body = %#v, want exactly %#v", gotBody, test.wantBody)
				}

				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"items":[{"id":"11111111-2222-3333-4444-555555555555","name":`+
					jsonQuote(test.replyName)+`,"destroyed":false,"provisioned":0}]}`)
			}))
			defer server.Close()

			response, err := test.call(contractClient(t, server))
			if err != nil {
				t.Fatalf("%s volume: %v", test.name, err)
			}
			if response == nil || response.Items == nil || len(*response.Items) != 1 {
				t.Fatalf("decoded response = %#v, want one item", response)
			}
			if got := (*response.Items)[0].Name; got == nil || *got != test.replyName {
				t.Errorf("decoded name = %v, want %q", got, test.replyName)
			}
		})
	}
}

func TestContractStructuredFailureAndMalformedJSON(t *testing.T) {
	t.Run("oauth structured 401", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("method = %q, want %q", r.Method, http.MethodPost)
			}
			if r.URL.Path != "/oauth2/1.0/token" {
				t.Errorf("path = %q, want %q", r.URL.Path, "/oauth2/1.0/token")
			}
			if got := r.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
				t.Errorf("Content-Type = %q, want application/x-www-form-urlencoded", got)
			}
			if err := r.ParseForm(); err != nil {
				t.Errorf("ParseForm: %v", err)
			}
			wantForm := map[string]string{
				"grant_type":         "urn:ietf:params:oauth:grant-type:token-exchange",
				"subject_token":      "invalid.jwt.token",
				"subject_token_type": "urn:ietf:params:oauth:token-type:jwt",
			}
			for key, want := range wantForm {
				if got := r.PostForm.Get(key); got != want {
					t.Errorf("form field %s = %q, want %q", key, got, want)
				}
			}
			if len(r.PostForm) != len(wantForm) {
				t.Errorf("form fields = %#v, want exactly %#v", r.PostForm, wantForm)
			}

			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = io.WriteString(w, `{"error":"invalid_client","error_description":"Invalid identity: JWT validation failed."}`)
		}))
		defer server.Close()

		response, err := contractClient(t, server).PostOauth210TokenWithFormdataBodyWithResponse(
			context.Background(),
			nil,
			client.PostOauth210TokenFormdataRequestBody{
				GrantType:        "urn:ietf:params:oauth:grant-type:token-exchange",
				SubjectToken:     "invalid.jwt.token",
				SubjectTokenType: "urn:ietf:params:oauth:token-type:jwt",
			},
		)
		if err != nil {
			t.Fatalf("PostOauth210TokenWithFormdataBodyWithResponse: %v", err)
		}
		if response.StatusCode() != http.StatusUnauthorized {
			t.Fatalf("status = %d, want %d", response.StatusCode(), http.StatusUnauthorized)
		}
		if response.JSON401 == nil || response.JSON401.Error == nil || *response.JSON401.Error != "invalid_client" {
			t.Fatalf("decoded 401 error = %#v, want invalid_client", response.JSON401)
		}
		if response.JSON401.ErrorDescription == nil ||
			*response.JSON401.ErrorDescription != "Invalid identity: JWT validation failed." {
			t.Errorf("decoded 401 description = %v", response.JSON401.ErrorDescription)
		}
	})

	t.Run("malformed JSON success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/api/2.26/volumes" {
				t.Errorf("path = %q, want %q", r.URL.Path, "/api/2.26/volumes")
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"items":[`)
		}))
		defer server.Close()

		response, err := contractClient(t, server).GetApi226VolumesWithResponse(context.Background(), nil)
		if err == nil {
			t.Fatalf("malformed JSON returned response %#v and nil error", response)
		}
	})

	t.Run("non-JSON content type is not decoded", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, _ = io.WriteString(w, `{"items":[]}`)
		}))
		defer server.Close()

		response, err := contractClient(t, server).GetApi226VolumesWithResponse(context.Background(), nil)
		if err != nil {
			t.Fatalf("GetApi226VolumesWithResponse: %v", err)
		}
		if response.JSON200 != nil {
			t.Errorf("JSON200 = %#v for text/plain response, want nil", response.JSON200)
		}
		if string(response.Body) != `{"items":[]}` {
			t.Errorf("body = %q, want preserved raw body", response.Body)
		}
	})
}

func jsonQuote(value string) string {
	encoded, _ := json.Marshal(value)
	return string(encoded)
}
