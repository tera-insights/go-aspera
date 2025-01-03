package aspera

import (
	"testing"
)

func Test_endpoint_URL(t *testing.T) {
	type fields struct {
		Route  string
		Prefix string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "Endpoint with no prefix",
			fields: fields{
				Route:  "activity",
				Prefix: "",
			},
			want: "activity",
		},
		{
			name: "Endpoint with prefix",
			fields: fields{
				Route:  "activity",
				Prefix: "/connect/transfers/",
			},
			want: "/connect/transfers/activity",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &endpoint{
				Route:  tt.fields.Route,
				Prefix: tt.fields.Prefix,
			}
			if got := e.URL(); got != tt.want {
				t.Errorf("endpoint.URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_endpoint_URLWithParams(t *testing.T) {
	type fields struct {
		Route  string
		Prefix string
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Endpoint with no params",
			fields: fields{
				Route:  "activity",
				Prefix: "/connect/transfers/",
			},
			args: args{
				params: map[string]string{},
			},
			want: "/connect/transfers/activity",
		},
		{
			name: "Endpoint with params",
			fields: fields{
				Route:  "info/${id}",
				Prefix: "/connect/transfers/",
			},
			args: args{
				params: map[string]string{
					"id": "123",
				},
			},
			want: "/connect/transfers/info/123",
		},
		{
			name: "Endpoint with multiple params",
			fields: fields{
				Route:  "modify/${id}/${name}",
				Prefix: "/connect/transfers/",
			},
			args: args{
				params: map[string]string{
					"id":   "123",
					"name": "test",
				},
			},
			want: "/connect/transfers/modify/123/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &endpoint{
				Route:  tt.fields.Route,
				Prefix: tt.fields.Prefix,
			}
			if got := e.URLWithParams(tt.args.params); got != tt.want {
				t.Errorf("endpoint.URLWithParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
