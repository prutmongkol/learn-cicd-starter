package auth

import (
	"reflect"
	"net/http"
	"testing"
	"errors"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		inputHeaders http.Header
		want         string
		wantErr      error
	}{
		"valid api key": {
			inputHeaders: http.Header{
				"Authorization": []string{"ApiKey 1234"},
			},
			want: "1234",
		},
		"no auth header": {
			inputHeaders: http.Header{},
			wantErr:      ErrNoAuthHeaderIncluded,
		},
		"malformed auth header": {
			inputHeaders: http.Header{
				"Authorization": []string{"Bearer 1234"},
			},
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.inputHeaders)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("GetAPIKey() got = %v, want %v", got, tc.want)
			}
			if !reflect.DeepEqual(err, tc.wantErr) {
				t.Fatalf("GetAPIKey() err = %v, want %v", err, tc.wantErr)
			}
		})
	}
}
