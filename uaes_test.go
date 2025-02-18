package uaes

import (
	"fmt"
	"testing"
)

const (
	SECRET = "%4NmStvr4@NrVheI"
)

var uaes = NewAES(SECRET)

func TestAES_Decrypt(t *testing.T) {
	type args struct {
		enc string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			name:    "String",
			args:    args{enc: "U2FsdGVkX18ByFJzfAGimHYkG7OoT27I96sgOZMJom4="},
			wantRes: "Zokijda",
		},
		{
			name:    "Map",
			args:    args{enc: "U2FsdGVkX1+i6OiRLkMjvRm1+XNK+7x+g4YFVxC1BI6rs+vLimyIJilgqk3Mk4QJ"},
			wantRes: `{"City":"Rorujut"}`,
		},
		{
			name:    "Empty enc",
			args:    args{enc: ""},
			wantRes: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := uaes.Decrypt(tt.args.enc); gotRes != tt.wantRes {
				t.Errorf("Aes.Decrypt() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestAES_Encrypt(t *testing.T) {
	type args struct {
		target []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		{
			name:    "Normal",
			args:    args{target: []byte("Zokijda")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := uaes.Encrypt(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Aes.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("gotRes: %v\n", gotRes)
		})
	}
}

func TestAES_Encrypt_Any(t *testing.T) {
	type args struct {
		target any
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		{
			name:    "Normal",
			args:    args{target: map[string]any{"City": "Rorujut"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := uaes.EncryptAny(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Aes.Encrypt_Any() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("gotRes: %v\n", gotRes)
		})
	}
}
