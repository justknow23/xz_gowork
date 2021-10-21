package tools

import "testing"

func Test_getCallerName(t *testing.T) {
	type args struct {
		skips []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "getCallerName",
			args: args{skips: []int{1}},
			want: "func1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCallerName(tt.args.skips...); got != tt.want {
				t.Errorf("getCallerName() = %v, want %v", got, tt.want)
			}
		})
	}
}
