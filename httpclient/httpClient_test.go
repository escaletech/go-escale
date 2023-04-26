package httpclient

import "testing"

func Test_castPathParamsToString(t *testing.T) {
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
			if got := castPathParamsToString(tt.args.pathParams); got != tt.want {
				t.Errorf("castPathParamsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
