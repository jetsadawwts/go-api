//go:build integration
package user

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"strconv"
)

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

func TestGetAllUser(t *testing.T) {
	seedUser(t)

	var us []User
	res := request(http.MethodGet, uri("users"), nil)
	err := res.Decode(&us)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(us), 0)
}


func TestCreateUser(t *testing.T) {
	body := bytes.NewBufferString(`{
		"name": "Jetsadawwts",
		"age": 19
	}`)
	var u User 
	res := request(http.MethodPost, uri("users"), body)
	err := res.Decode(&u)
	if err != nil {
		t.Fatal("can't create user", err)
	}
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, u.ID)
	assert.Equal(t, "Jetsadawwts", u.Name)
	assert.Equal(t, 19, u.Age)
}


func TestGetUserByID(t *testing.T) {
	c := seedUser(t)
	
	var latest User
	res := request(http.MethodGet, uri("users", strconv.Itoa(c.ID)), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, latest.ID)
	assert.NotEmpty(t, latest.Name)
	assert.NotEmpty(t, latest.Age)
}

func TestUpdateUserByID(t *testing.T) {
	t.Skip("TODO: implement me PATCH /users/:id")
}

func TestDeleteUserByID(t *testing.T) {
	t.Skip("TODO: implement me DELETE /users/:id")
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func seedUser(t *testing.T) User {
	var c User
	body := bytes.NewBufferString(`{
		"name": "Jetsadawwts",
		"age": 19
	}`)
	err := request(http.MethodPost, uri("users"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't create user", err)
	}

	return c

}


