package auth

import "testing"

func TestCreateRandomString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRandomString(); got != tt.want {
				t.Errorf("CreateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
