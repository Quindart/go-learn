package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/app-event-booking/utils"
	"github.com/gin-gonic/gin"
)

func TestAuthenticate(t *testing.T) {
	token, err := utils.GenerateToken("user@example.com", 42)
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct {
		name, header string
		wantStatus   int
	}{
		{"missing", "", http.StatusUnauthorized},
		{"invalid", "Bearer invalid", http.StatusUnauthorized},
		{"wrong scheme", "Basic " + token, http.StatusUnauthorized},
		{"valid", "Bearer " + token, http.StatusNoContent},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			server := gin.New()
			called := false
			server.GET("/protected", Authenticate, func(c *gin.Context) {
				called = true
				if id := c.GetInt64("userId"); id != 42 {
					t.Errorf("userId = %d, want 42", id)
				}
				c.Status(http.StatusNoContent)
			})
			request := httptest.NewRequest(http.MethodGet, "/protected", nil)
			request.Header.Set("Authorization", tt.header)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			if response.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", response.Code, tt.wantStatus)
			}
			if called != (tt.wantStatus == http.StatusNoContent) {
				t.Errorf("handler called = %v", called)
			}
		})
	}
}
