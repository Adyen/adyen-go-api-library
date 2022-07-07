/*
 * Adyen API
 *
 * Contact: support@adyen.com
 */

package common

import (
	"bytes"
	"testing"
)

func TestResultCode_String(t *testing.T) {
	tests := []struct {
		name string
		s    ResultCode
		want string
	}{
		{
			"test valid case",
			AuthenticationFinished,
			"AuthenticationFinished",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("ResultCode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultCode_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		s    ResultCode
		want string
	}{
		{
			"test valid case",
			Authorised,
			`"Authorised"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.s.MarshalJSON(); string(got) != tt.want {
				t.Errorf("ResultCode.MarshallJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestResultCode_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		s          ResultCode
		jsonString string
		want       error
	}{
		{
			"test valid case",
			Authorised,
			`"Authorised"`,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnmarshalJSON(bytes.NewBufferString(tt.jsonString).Bytes()); err != tt.want {
				t.Errorf("ResultCode.UnmarshallJSON(). Got error: %v", err)
			}
		})
	}
}
