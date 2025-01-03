package aspera

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		httpClient *http.Client
		baseUrl    *url.URL
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "Create new client",
			args: args{
				httpClient: &http.Client{},
				baseUrl:    &url.URL{},
			},
			want: &Client{
				Client:  &http.Client{},
				BaseURL: &url.URL{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.httpClient, tt.args.baseUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRequestWithParameters(t *testing.T) {
	type args struct {
		method string
		name   string
		params map[string]string
		body   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "Create new request with parameters",
			args: args{
				method: "GET",
				name:   "activity",
				params: map[string]string{
					"id": "123",
				},
				body: nil,
			},
			want: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "connect/transfers/activity"},
			},
			wantErr: false,
		},
		{
			name: "Create new request with parameters",
			args: args{
				method: "GET",
				name:   "getTransfer",
				params: map[string]string{
					"id": "123",
				},
				body: nil,
			},
			want: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "connect/transfers/info/123"},
			},
			wantErr: false,
		},
		{
			name: "Non existing endpoint",
			args: args{
				method: "GET",
				name:   "nonExisting",
				params: map[string]string{
					"id": "123",
				},
				body: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:  &http.Client{},
				BaseURL: &url.URL{},
			}
			got, err := c.NewRequestWithParameters(tt.args.method, tt.args.name, tt.args.params, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewRequestWithParameters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRequestWithParameters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	type args struct {
		method string
		name   string
		body   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "Create new request",
			args: args{
				method: "GET",
				name:   "activity",
				body:   nil,
			},
			want: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "connect/transfers/activity"},
			},
			wantErr: false,
		},
		{
			name: "Create new request",
			args: args{
				method: "GET",
				name:   "getChecksum",
				body:   nil,
			},
			want: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "connect/file/checksum/"},
			},
			wantErr: false,
		},
		{
			name: "Non existing endpoint",
			args: args{
				method: "GET",
				name:   "nonExisting",
				body:   nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:  &http.Client{},
				BaseURL: &url.URL{},
			}
			got, err := c.NewRequest(tt.args.method, tt.args.name, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sanitizeURL(t *testing.T) {
	type args struct {
		uri *url.URL
	}
	tests := []struct {
		name string
		args args
		want *url.URL
	}{
		{
			name: "Sanitize URL",
			args: args{
				uri: &url.URL{
					RawQuery: "client_secret=123",
				},
			},
			want: &url.URL{
				RawQuery: "client_secret=REDACTED",
			},
		},
		{
			name: "Sanitize URL",
			args: args{
				uri: &url.URL{
					RawQuery: "client_secret=123&client_id=123",
				},
			},
			want: &url.URL{
				RawQuery: "client_secret=REDACTED&client_id=123",
			},
		},
		{
			name: "Sanitize sanitized URL",
			args: args{
				uri: &url.URL{
					RawQuery: "client_secret=REDACT",
				},
			},
			want: &url.URL{
				RawQuery: "client_secret=REDACT",
			},
		},
		{
			name: "Sanitize regular URL",
			args: args{
				uri: &url.URL{
					RawQuery: "client_id=123",
				},
			},
			want: &url.URL{
				RawQuery: "client_id=123",
			},
		},
		{
			name: "Sanitize empty URL",
			args: args{
				uri: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeURL(tt.args.uri); got.Query().Get("client_secret") != "" && got.Query().Get("client_secret") != "REDACTED" {
				t.Errorf("sanitizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
