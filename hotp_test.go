package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

// test cases are taken from https://www.ietf.org/rfc/rfc4226.txt. Additionally, the byte arrays have been
// generated using ./hextobytes.py.
var(
	b0 = []byte{204, 147, 207, 24, 80, 141, 148, 147, 76, 100, 182, 93, 139, 167, 102, 127, 183, 205, 228, 176}
	b1 = []byte{117, 164, 138, 25, 212, 203, 225, 0, 100, 78, 138, 193, 57, 126, 234, 116, 122, 45, 51, 171}
	b2 = []byte{11, 172, 183, 250, 8, 47, 239, 48, 120, 34, 17, 147, 139, 193, 197, 231, 4, 22, 255, 68}
	b3 = []byte{102, 194, 130, 39, 208, 58, 45, 85, 41, 38, 47, 240, 22, 161, 230, 239, 118, 85, 126, 206}
	b4 = []byte{169, 4, 201, 0, 166, 75, 53, 144, 152, 116, 179, 62, 97, 197, 147, 138, 142, 21, 237, 28}
	b5 = []byte{163, 126, 120, 61, 123, 114, 51, 192, 131, 212, 246, 41, 38, 199, 162, 95, 35, 141, 3, 22}
	b6 = []byte{188, 156, 210, 133, 97, 4, 44, 131, 242, 25, 50, 77, 60, 96, 114, 86, 192, 50, 114, 174}
	b7 = []byte{164, 251, 150, 12, 11, 192, 110, 30, 171, 184, 4, 229, 179, 151, 205, 196, 180, 85, 150, 250}
	b8 = []byte{27, 60, 137, 246, 94, 108, 158, 136, 48, 18, 5, 40, 35, 68, 63, 4, 139, 67, 50, 219}
	b9 = []byte{22, 55, 64, 152, 9, 166, 121, 220, 105, 130, 7, 49, 12, 140, 127, 192, 114, 144, 217, 229}

	d0 = 1284755224
	d1 = 1094287082
    d2 = 137359152
	d3 = 1726969429
    d4 = 1640338314
    d5 = 868254676
    d6 = 1918287922
    d7 = 82162583
    d8 = 673399871
    d9 = 645520489

	h0 = 755224
	h1 = 287082
	h2 = 359152
	h3 = 969429
	h4 = 338314
	h5 = 254676
	h6 = 287922
	h7 = 162583
	h8 = 399871
	h9 = 520489
)

func Test_decode_base32(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "Test 0", args: args{input: []byte("ORSXG5A=")}, want: []byte("test"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decode_base32(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("decode_base32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode_base32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hmac_sha_1(t *testing.T) {
	key := []byte("12345678901234567890")
	// Initialize the counter values as byte arrays
	inputs := make([][]byte, 10)
	for i := range inputs {
		inputs[i] = make([]byte, 8)
		binary.BigEndian.PutUint64(inputs[i], uint64(i))
	}
	// Initialize the expected values as byte arrays
	expected := make([][]byte, 20)
	expected[0] = b0
	expected[1] = b1
	expected[2] = b2
	expected[3] = b3
	expected[4] = b4
	expected[5] = b5
	expected[6] = b6
	expected[7] = b7
	expected[8] = b8
	expected[9] = b9
	type args struct {
		key   []byte
		value []byte
	}
	type test struct {
		name string
		args args
		want []byte
	}
	var tests [10]test
		for i := 0; i < 10; i++ {
		a := args{key: key, value: inputs[i]}
		tests[i] = test{name: fmt.Sprintf("Test %d", i), args: a, want: expected[i]}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hmac_sha_1(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hmac_sha_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamic_truncation(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Produces error (too short)", args: args{input: []byte("a")}, want: -1, wantErr: true},
		{name: "Produces error (too long)", args: args{input: []byte("abcdefghijklmnopqrstuvwxyz")}, want: -1, wantErr: true},
		{name: "Test 0", args: args{input: b0}, want: d0, wantErr: false},
		{name: "Test 1", args: args{input: b1}, want: d1, wantErr: false},
		{name: "Test 2", args: args{input: b2}, want: d2, wantErr: false},
		{name: "Test 3", args: args{input: b3}, want: d3, wantErr: false},
		{name: "Test 4", args: args{input: b4}, want: d4, wantErr: false},
		{name: "Test 5", args: args{input: b5}, want: d5, wantErr: false},
		{name: "Test 6", args: args{input: b6}, want: d6, wantErr: false},
		{name: "Test 7", args: args{input: b7}, want: d7, wantErr: false},
		{name: "Test 8", args: args{input: b8}, want: d8, wantErr: false},
		{name: "Test 9", args: args{input: b9}, want: d9, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dynamic_truncation(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("dynamic_truncation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("dynamic_truncation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_htop_value(t *testing.T) {
	digit := 6
	type args struct {
		value int
		digit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 0", args: args{value: d0, digit: digit},want: h0},	
		{name: "Test 1", args: args{value: d1, digit: digit},want: h1},
		{name: "Test 2", args: args{value: d2, digit: digit},want: h2},
		{name: "Test 3", args: args{value: d3, digit: digit},want: h3},
		{name: "Test 4", args: args{value: d4, digit: digit},want: h4},
		{name: "Test 5", args: args{value: d5, digit: digit},want: h5},
		{name: "Test 6", args: args{value: d6, digit: digit},want: h6},
		{name: "Test 7", args: args{value: d7, digit: digit},want: h7},
		{name: "Test 8", args: args{value: d8, digit: digit},want: h8},
		{name: "Test 9", args: args{value: d9, digit: digit},want: h9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hotp_value(tt.args.value, tt.args.digit); got != tt.want {
				t.Errorf("htop_value() = %v, want %v", got, tt.want)
			}
		})
	}
}
