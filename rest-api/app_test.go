package restapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go get github.com/smartystreets/goconvey
// run goconvey

// 매번 조건문에서 검사하기 귀찮으니까
// go get github.com/stretchr/testify/assert
// asset.Equal 활용

func TestFooPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	// if res.Code != http.StatusOK {
	// 	t.Fatal("Failed", res.Code)
	// }

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello Foo", string(data))
}

func TestFooPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo?name=zini", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	// if res.Code != http.StatusOK {
	// 	t.Fatal("Failed", res.Code)
	// }

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello zini", string(data))
}

func TestBarPathHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/bar", strings.NewReader(`{
		"firstname": "heo",
		"lastname": "zini",
		"email": "zinirun@github.com"
	}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)

	assert.Nil(err)
	assert.Equal("zini", user.LastName)
	assert.Equal("heo", user.FirstName)
}

func TestBarPathHandler_WithBadBody(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/bar", strings.NewReader(`{
		"firstname": "heo",
		"lastname": "zini",
	}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}
