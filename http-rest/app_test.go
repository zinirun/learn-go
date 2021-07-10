package httprest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(RestHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id: 89")

	resp, err = http.Get(ts.URL + "/users/abc")
	assert.NoError(err)
	assert.Equal(http.StatusNotFound, resp.StatusCode)
}

func TestCreateUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(RestHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`
	{
		"firstname": "zini",
		"lastname": "heo",
		"email": "zini@willog.io"
	}`))

	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)

	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}
