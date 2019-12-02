package items

import (
	"api/app/mock"
	"api/app/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestGetItem(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	// Inject our mock into our handler.
	var is mock.ItemService
	Is = &is
	itemExpected := models.Item{ID: "100", Name: "Algo1", Description: "Algo2"}
	// Mock our User() call.
	is.ItemFn = func(id string) (*models.Item, error) {
		if id != "100" {
			t.Fatalf("unexpected id: %s", id)
		}
		return &itemExpected, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item/100", nil)
	router.ServeHTTP(w, r)
	var i models.Item
	err := json.Unmarshal(w.Body.Bytes(), &i)
	assert.Nil(t, err)
	assert.Equal(t, itemExpected, i)
}

func TestGetItemInvalidID(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	var is mock.ItemService
	Is = &is
	// Mock our User() call.
	is.ItemFn = func(id string) (*models.Item, error) {
		return nil, nil
	}
	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item/ ", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetItemFail(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	var is mock.ItemService
	Is = &is
	// Mock our User() call.
	is.ItemFn = func(id string) (*models.Item, error) {
		return nil, errors.New("Error")
	}
	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item/100", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestPostItem(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)
	var is mock.ItemService
	Is = &is
	mockItem := models.Item{Name: "Algo", Description: "algo2"}
	is.CreateItemFn = func(i *models.Item) error {
		return nil
	}
	// Inject our mock into our handler.
	itemBytes, _ := json.Marshal(mockItem)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/item", bytes.NewBuffer(itemBytes))
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostItemFail(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)
	var is mock.ItemService
	Is = &is
	is.CreateItemFn = func(i *models.Item) error {
		return nil
	}
	// Inject our mock into our handler.

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/item", bytes.NewBuffer([]byte(`{`)))
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestPostItemFailService(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)
	var is mock.ItemService
	Is = &is
	mockItem := models.Item{Name: "Algo", Description: "algo2"}
	is.CreateItemFn = func(i *models.Item) error {
		return errors.New("Error")
	}
	// Inject our mock into our handler.
	itemBytes, _ := json.Marshal(mockItem)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/item", bytes.NewBuffer(itemBytes))
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
func TestGetItemAll(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)
	var is mock.ItemService
	Is = &is
	itemsExpected := []*models.Item{&models.Item{ID: "100", Name: "Algo1", Description: "Algo2"}}
	is.ItemsFn = func() ([]*models.Item, error) {
		return itemsExpected, nil
	}
	// Inject our mock into our handler.

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item", nil) //buffer?
	router.ServeHTTP(w, r)
	var its []*models.Item
	err := json.Unmarshal(w.Body.Bytes(), &its)
	assert.Nil(t, err)
	assert.Equal(t, itemsExpected, its)
}
func TestGetItemAllFail(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)
	var is mock.ItemService
	Is = &is
	is.ItemsFn = func() ([]*models.Item, error) {
		return nil, errors.New("Error")
	}
	// Inject our mock into our handler.

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/item", nil) //buffer?
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteItem(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)
	var is mock.ItemService
	Is = &is

	is.DeleteItemFn = func(id string) error {
		return nil
	}
	// Inject our mock into our handler.

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/item/1", nil) //buffer?
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestDeleteItemInvalidID(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	var is mock.ItemService
	Is = &is
	// Mock our User() call.
	is.DeleteItemFn = func(id string) error {
		return nil
	}
	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/item/ ", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteItemFail(t *testing.T) {
	router := gin.Default()
	Configure(router, nil)

	var is mock.ItemService
	Is = &is
	// Mock our User() call.
	is.DeleteItemFn = func(id string) error {
		return errors.New("Error")
	}
	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/item/1", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
