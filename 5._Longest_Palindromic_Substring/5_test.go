package longestpalindromicsubstring

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "babad",
			},
			want: "aba",
		},
		{
			name: "",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "",
			args: args{
				s: "abb",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
