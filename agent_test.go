package dialogflow

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAgent(t *testing.T) {
	assert := assert.New(t)

	srv := testServer([]byte(`{"displayName":"agent name"}`))
	defer func() { srv.Close() }()

	c := NewClient("token", "projectID")
	agent, err := c.GetAgent()
	assert.Error(err)

	c.apiURL = srv.URL
	agent, err = c.GetAgent()
	assert.NoError(err)
	assert.Equal(agent.String(), "agent name")
}

func testServer(data []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write(data)
	}))
}
