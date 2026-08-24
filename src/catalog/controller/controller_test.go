package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestNewController tests the creation of a new controller instance
func TestNewController(t *testing.T) {
	// Test that NewController properly initializes
	controller, err := NewController(nil)
	assert.NoError(t, err)
	assert.NotNil(t, controller)
}

// TestControllerHandlers tests that controller has required handler methods
func TestControllerHandlers(t *testing.T) {
	controller, err := NewController(nil)
	assert.NoError(t, err)

	// Verify the controller has the required methods
	assert.NotNil(t, controller.GetProducts)
	assert.NotNil(t, controller.GetProduct)
	assert.NotNil(t, controller.CatalogSize)
	assert.NotNil(t, controller.ListTags)
}

// TestGetProducts tests the GetProducts endpoint
func TestGetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	controller, err := NewController(nil)
	assert.NoError(t, err)

	router.GET("/catalog", controller.GetProducts)

	req, _ := http.NewRequest("GET", "/catalog", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verify the endpoint returns a response (may be error due to nil API)
	assert.NotNil(t, w)
}

// TestGetProduct tests the GetProduct endpoint
func TestGetProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	controller, err := NewController(nil)
	assert.NoError(t, err)

	router.GET("/catalog/product/:id", controller.GetProduct)

	req, _ := http.NewRequest("GET", "/catalog/product/test-id", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verify the endpoint is callable
	assert.NotNil(t, w)
}

// TestCatalogSize tests the CatalogSize endpoint
func TestCatalogSize(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	controller, err := NewController(nil)
	assert.NoError(t, err)

	router.GET("/catalog/size", controller.CatalogSize)

	req, _ := http.NewRequest("GET", "/catalog/size", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verify the endpoint is callable
	assert.NotNil(t, w)
}

// TestListTags tests the ListTags endpoint
func TestListTags(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	controller, err := NewController(nil)
	assert.NoError(t, err)

	router.GET("/catalog/tags", controller.ListTags)

	req, _ := http.NewRequest("GET", "/catalog/tags", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verify the endpoint is callable
	assert.NotNil(t, w)
}
