package gateway

import (
	"reflect"
	"testing"
)

func main() {
	// Write the test cases you want to imploy here...
	/*
		- test valid tokens are valid
		- test that the token is only valid for its user
		- test that outdated token is not accepted
		- test tokens in the future
	*/
}

func TestGenerateJWT(t *testing.T) {
	type args struct {
		serviceID string
		secret    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateJWT(tt.args.serviceID, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyJWT(t *testing.T) {
	key := []byte("abcdef012345678")
	j, _ := GenerateJWT("test", key)

	tests := []struct {
		name    string
		have    string
		want string
	}{
		{"test_successful_retrieval", "test", "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := VerifyJWT(j, []byte("abcdef012345678"))
			if !reflect.DeepEqual(got.Username, tt.want) {
				t.Errorf("VerifyJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}