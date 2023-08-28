package main

import "testing"

type MockCounterProvider struct {
	state int
}

func (m *MockCounterProvider) provide() uint64 {
	ret := m.state
	m.state = m.state + 1
	return uint64(ret)
}

func Test_get_totp_code(t *testing.T) {
	key := []byte("GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ")
	digits := 6
	counter_provider := &MockCounterProvider{state: 0}	
	type args struct {
		key              []byte
		digits           int
		counter_provider CounterProvider
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Test 0", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h0, wantErr: false},		
		{name: "Test 1", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h1, wantErr: false},
		{name: "Test 2", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h2, wantErr: false},		
		{name: "Test 3", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h3, wantErr: false},
		{name: "Test 4", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h4, wantErr: false},		
		{name: "Test 5", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h5, wantErr: false},
		{name: "Test 6", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h6, wantErr: false},		
		{name: "Test 7", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h7, wantErr: false},
		{name: "Test 8", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h8, wantErr: false},		
		{name: "Test 9", args: args{key: key, digits: digits, counter_provider: counter_provider}, want: h9, wantErr: false},
		{name: "Failing, illegal key", args: args{key: []byte("0123456789"), digits: digits, counter_provider: counter_provider}, want: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := get_totp_code(tt.args.key, tt.args.digits, tt.args.counter_provider)
			if (err != nil) != tt.wantErr {
				t.Errorf("get_totp_code() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("get_totp_code() = %v, want %v", got, tt.want)
			}
		})
	}
}
