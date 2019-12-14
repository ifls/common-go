package string

import (
	"log"
	"strings"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "the sky is blue"},
			want: "blue is sky the",
		},
		{
			args: args{s: "  hello world!  "},
			want: "world! hello",
		},
		{
			args: args{s: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	log.Println("[" + strings.TrimSpace("  ") + "]")
}
