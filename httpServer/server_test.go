package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	recorder := httptest.NewRecorder()

	// simulate request
	req, _ := http.NewRequest("GET", "", nil)
	req.RemoteAddr = "1.2.3.4"
	req.RequestURI = "/doit?arg1=1&arg2=two"
	req.Header.Set("Accept", "application/json")

	eh := echoHandler{true}
	eh.ServeHTTP(recorder, req)

	//  verify response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// decode json
	dec := json.NewDecoder(recorder.Body)
	var resp Response
	err := dec.Decode(&resp)
	assert.NoError(t, err)

	//  json body
	assert.Equal(t, "1.2.3.4", resp.Origin)
	assert.Equal(t, "/doit", resp.Url)
	assert.Equal(t, "1", resp.Args["arg1"][0])
	assert.Equal(t, "two", resp.Args["arg2"][0])
	assert.Equal(t, "application/json", resp.Headers["Accept"][0])

}
