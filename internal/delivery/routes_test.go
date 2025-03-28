package delivery

import (
	"fiber-test/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	app := fiber.New()
	RegisterRoutes(app) 
	return app
}

func TestPublicHandler(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/public", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGenerateTokenHandlerSuccess(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/generate-token/123", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGenerateTokenHandlerInvalidUserID(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/generate-token/abc", nil) 
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestProtectedHandlerWithValidToken(t *testing.T) {
	app := setupTestApp()

	token, _ := middleware.GenerateToken(123)

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token) 

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestProtectedHandlerWithoutToken(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestProtectedHandlerWithInvalidToken(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer token_ngasal") 

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
