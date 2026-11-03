package main

import (
	"reflect"
	"testing"
)

func TestEncodeRedisValues(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want interface{}
	}{
		{name: "nil", in: nil, want: nil},
		{name: "integer", in: int64(42), want: int64(42)},
		{name: "string", in: "hello", want: "aGVsbG8="},
		{name: "bytes", in: []byte{0, 1, 2}, want: "AAEC"},
		{name: "nested values", in: []interface{}{"ok", int64(3), []byte("x")}, want: []interface{}{"b2s=", int64(3), "eA=="}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encode(%#v) = %#v, want %#v", tt.in, got, tt.want)
			}
		})
	}
}
