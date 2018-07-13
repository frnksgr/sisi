package server

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_byteSize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"byte", args{"42B"}, 42},
		{"kilo", args{"42K"}, 42 * 1024},
		{"mega", args{"42M"}, 42 * 1024 * 1024},
		{"giga", args{"42G"}, 42 * 1024 * 1024 * 1024},
		{"fail-1", args{"42F"}, 0},
		{"fail-2", args{"42"}, 0},
		{"fail-3", args{"42FG"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteSize(tt.args.s); got != tt.want {
				t.Errorf("byteSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helloHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helloHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_chunkedHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chunkedHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_dataHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestRunServer(t *testing.T) {
	type args struct {
		listenAddress string
	}
	tests := []struct {
		name string
		args args
		want *http.Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunServer(tt.args.listenAddress); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
