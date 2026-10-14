package httpclient

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func sameQuery(t *testing.T, got, want string) bool {
	t.Helper()
	if got == "" || want == "" {
		return got == want
	}
	if !strings.HasPrefix(got, "?") || !strings.HasPrefix(want, "?") {
		return got == want
	}
	g, err := url.ParseQuery(strings.TrimPrefix(got, "?"))
	if err != nil {
		t.Fatalf("query inválida %q: %v", got, err)
	}
	w, err := url.ParseQuery(strings.TrimPrefix(want, "?"))
	if err != nil {
		t.Fatalf("query inválida %q: %v", want, err)
	}
	return reflect.DeepEqual(g, w)
}

func Test_castQueryParamsToString(t *testing.T) {
	type args struct {
		pathParams map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok with 1",
			args: args{
				pathParams: map[string]string{
					"path1": "um",
				},
			},
			want: "?path1=um",
		},
		{
			name: "ok with 2",
			args: args{
				pathParams: map[string]string{
					"path1": "um",
					"path2": "dois",
				},
			},
			want: "?path1=um&path2=dois",
		},
		{
			name: "no map",
			args: args{
				pathParams: nil,
			},
			want: "",
		},
		{
			name: "empty map",
			args: args{
				pathParams: map[string]string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// a ordem de iteração de map é aleatória: compara os parâmetros, não a string
			got := castQueryParamsToString(tt.args.pathParams)
			if !sameQuery(t, got, tt.want) {
				t.Errorf("castQueryParamsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
