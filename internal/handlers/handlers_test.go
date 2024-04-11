package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers(t *testing.T) {
	var resURL string
	type args struct {
		method string
		body   string
		target string
	}
	type want struct {
		statusCode int
		link       string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "add link",
			args: args{
				method: http.MethodPost,
				body:   "https://ya.ru",
				target: "localhost:8080",
			},
			want: want{
				statusCode: 201,
			},
		},
		{
			name: "get link",
			args: args{
				method: http.MethodGet,
				body:   "",
				target: resURL,
			},
			want: want{
				statusCode: 307,
				link:       "https://ya.ru",
			},
		},
		{
			name: "bad request POST",
			args: args{
				method: http.MethodPost,
				body:   "",
				target: "/",
			},
			want: want{
				statusCode: 400,
				link:       "",
			},
		},
		{
			name: "bad request GET",
			args: args{
				method: http.MethodGet,
				body:   "",
				target: "/",
			},
			want: want{
				statusCode: 400,
				link:       "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var target string
			if tt.args.target != "" {
				target = tt.args.target
			} else {
				target = resURL
			}

			r := httptest.NewRequest(tt.args.method, target, strings.NewReader(tt.args.body))
			w := httptest.NewRecorder()
			if tt.args.method == http.MethodGet {
				MainHandlerGetById(w, r)
			} else {
				MainHandlerSet(w, r)
			}
			response := w.Result()
			defer response.Body.Close()

			if tt.args.method == http.MethodPost {
				assert.Equal(t, tt.want.statusCode, response.StatusCode)
				resURL = w.Body.String()

			}
			if tt.args.method == http.MethodGet {
				assert.Equal(t, tt.want.statusCode, response.StatusCode)
				assert.Equal(t, tt.want.link, w.Header().Get("Location"))
			}
		})
	}
}
