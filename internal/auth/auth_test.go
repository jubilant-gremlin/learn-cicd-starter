package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name: "valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_token"},
			},
			wantToken: "valid_token",
			wantErr:   false,
		},
		{
			name:      "Missing auth header",
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "malformed auth header",
			headers: http.Header{
				"Authorization": []string{"InvalidApiKey token"},
			},
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("GetAPIKey() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}

}
