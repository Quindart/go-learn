package router

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"example.com/app-event-booking/db"
	"example.com/app-event-booking/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const eventBody = `{"Name":"Updated","Description":"Description","Location":"Hanoi","Date":"2026-09-13T10:00:00Z","UserID":999}`

func setupEventRouter(t *testing.T) *gin.Engine {
	t.Helper()
	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	database.SetMaxOpenConns(1)
	previous := db.DB
	db.DB = database
	t.Cleanup(func() { db.DB = previous; database.Close() })
	_, err = database.Exec(`CREATE TABLE events (
		id INTEGER PRIMARY KEY, name TEXT, description TEXT, location TEXT,
		dateTime DATETIME, user_id INTEGER
	);
	INSERT INTO events VALUES (1, 'Original', 'Description', 'Hanoi', '2026-09-13 10:00:00', 42);`)
	if err != nil {
		t.Fatal(err)
	}
	gin.SetMode(gin.TestMode)
	server := gin.New()
	RegisterRouter(server)
	return server
}

func eventRequest(server *gin.Engine, method, path, authorization, body string) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorization)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)
	return response
}

func eventToken(t *testing.T, userID int64) string {
	t.Helper()
	token, err := utils.GenerateToken("user@example.com", userID)
	if err != nil {
		t.Fatal(err)
	}
	return "Bearer " + token
}

func assertOriginalEvent(t *testing.T) {
	t.Helper()
	var count, owner int
	var name string
	if err := db.DB.QueryRow("SELECT COUNT(*) FROM events").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("event count = %d, want 1", count)
	}
	if err := db.DB.QueryRow("SELECT name, user_id FROM events WHERE id = 1").Scan(&name, &owner); err != nil {
		t.Fatal(err)
	}
	if name != "Original" || owner != 42 {
		t.Fatalf("event changed: name=%q, owner=%d", name, owner)
	}
}

func TestEventMutationsRequireAuthorization(t *testing.T) {
	valid := eventToken(t, 42)
	// A well-formed token signed by another key must still be rejected.
	untrusted, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": 42, "exp": time.Now().Add(time.Hour).Unix(),
	}).SignedString([]byte("untrusted-key"))
	if err != nil {
		t.Fatal(err)
	}
	for _, method := range []string{http.MethodPost, http.MethodPut, http.MethodDelete} {
		for _, header := range []struct{ name, value string }{
			{"missing", ""}, {"invalid", "Bearer invalid"},
			{"raw token", strings.TrimPrefix(valid, "Bearer ")},
			{"wrong scheme", "Basic " + strings.TrimPrefix(valid, "Bearer ")},
			{"empty bearer", "Bearer"}, {"extra fields", valid + " extra"},
			{"untrusted", "Bearer " + untrusted},
		} {
			t.Run(method+"/"+header.name, func(t *testing.T) {
				server := setupEventRouter(t)
				path := "/events/1"
				if method == http.MethodPost {
					path = "/events"
				}
				response := eventRequest(server, method, path, header.value, eventBody)
				if response.Code != http.StatusUnauthorized {
					t.Errorf("status = %d, want 401; body=%s", response.Code, response.Body)
				}
				assertOriginalEvent(t)
			})
		}
	}
}

func TestEventAuthorizationPrecedesInputValidation(t *testing.T) {
	server := setupEventRouter(t)
	for _, method := range []string{http.MethodPost, http.MethodPut, http.MethodDelete} {
		path := "/events/invalid-id"
		if method == http.MethodPost {
			path = "/events"
		}
		response := eventRequest(server, method, path, "", "invalid-json")
		if response.Code != http.StatusUnauthorized {
			t.Errorf("%s status = %d, want 401", method, response.Code)
		}
	}
}

func TestCreateEventUsesTokenOwner(t *testing.T) {
	server := setupEventRouter(t)
	response := eventRequest(server, http.MethodPost, "/events", eventToken(t, 42), eventBody)
	if response.Code != http.StatusCreated {
		t.Fatalf("status = %d, want 201; body=%s", response.Code, response.Body)
	}
	var owner int
	if err := db.DB.QueryRow("SELECT user_id FROM events WHERE id = 2").Scan(&owner); err != nil {
		t.Fatal(err)
	}
	if owner != 42 {
		t.Fatalf("owner = %d, want 42", owner)
	}
}

func TestEventMutationsRequireOwnership(t *testing.T) {
	for _, method := range []string{http.MethodPut, http.MethodDelete} {
		t.Run(method, func(t *testing.T) {
			server := setupEventRouter(t)
			response := eventRequest(server, method, "/events/1", eventToken(t, 7), eventBody)
			if response.Code != http.StatusForbidden {
				t.Errorf("status = %d, want 403", response.Code)
			}
			assertOriginalEvent(t)
		})
	}
}

func TestEventOwnerCanUpdateAndDelete(t *testing.T) {
	server := setupEventRouter(t)
	response := eventRequest(server, http.MethodPut, "/events/1", eventToken(t, 42), eventBody)
	if response.Code != http.StatusOK {
		t.Fatalf("update status = %d, want 200", response.Code)
	}
	var name string
	var owner int
	if err := db.DB.QueryRow("SELECT name, user_id FROM events WHERE id = 1").Scan(&name, &owner); err != nil {
		t.Fatal(err)
	}
	if name != "Updated" || owner != 42 {
		t.Fatalf("updated event: name=%q, owner=%d", name, owner)
	}
	response = eventRequest(server, http.MethodDelete, "/events/1", eventToken(t, 42), "")
	if response.Code != http.StatusOK {
		t.Fatalf("delete status = %d, want 200", response.Code)
	}
	var count int
	if err := db.DB.QueryRow("SELECT COUNT(*) FROM events").Scan(&count); err != nil || count != 0 {
		t.Fatalf("count = %d, err = %v; want 0, nil", count, err)
	}
}
