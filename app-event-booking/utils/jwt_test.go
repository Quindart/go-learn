package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestVerifyTokenGeneratedToken(t *testing.T) {
	// Preserve int64 IDs that cannot be represented exactly by float64.
	const wantID int64 = 9007199254740993
	token, err := GenerateToken("user@example.com", wantID)
	if err != nil {
		t.Fatal(err)
	}
	gotID, err := VerifyToken(token)
	if err != nil || gotID != wantID {
		t.Fatalf("VerifyToken() = (%d, %v), want (%d, nil)", gotID, err, wantID)
	}
}

func TestVerifyTokenRejectsInvalidTokens(t *testing.T) {
	future := time.Now().Add(time.Hour).Unix()
	tests := []struct {
		name   string
		method jwt.SigningMethod
		key    string
		claims jwt.MapClaims
	}{
		{"wrong signature", jwt.SigningMethodHS256, "wrong-key", jwt.MapClaims{"userId": 1, "exp": future}},
		{"wrong algorithm", jwt.SigningMethodHS384, secretKey, jwt.MapClaims{"userId": 1, "exp": future}},
		{"expired", jwt.SigningMethodHS256, secretKey, jwt.MapClaims{"userId": 1, "exp": time.Now().Add(-time.Hour).Unix()}},
		{"missing expiration", jwt.SigningMethodHS256, secretKey, jwt.MapClaims{"userId": 1}},
		{"missing user ID", jwt.SigningMethodHS256, secretKey, jwt.MapClaims{"exp": future}},
		{"invalid user ID type", jwt.SigningMethodHS256, secretKey, jwt.MapClaims{"userId": "1", "exp": future}},
		{"fractional user ID", jwt.SigningMethodHS256, secretKey, jwt.MapClaims{"userId": 1.5, "exp": future}},
		{"zero user ID", jwt.SigningMethodHS256, secretKey, jwt.MapClaims{"userId": 0, "exp": future}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := jwt.NewWithClaims(tt.method, tt.claims).SignedString([]byte(tt.key))
			if err != nil {
				t.Fatal(err)
			}
			if id, err := VerifyToken(token); err == nil || id != 0 {
				t.Fatalf("VerifyToken() = (%d, %v), want (0, error)", id, err)
			}
		})
	}
	t.Run("malformed", func(t *testing.T) {
		if id, err := VerifyToken("not-a-jwt"); err == nil || id != 0 {
			t.Fatalf("VerifyToken() = (%d, %v), want (0, error)", id, err)
		}
	})
}
