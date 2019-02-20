package dialogflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert := assert.New(t)
	c := NewClient("token", "project")

	assert.NotNil(c.GetAccessToken(), "Access token not set")
	assert.Equal(c.GetAPIVersion(), DefaultVersion, "client v2 version is not set")
	assert.Equal(c.GetProjectID(), "project", "client project is not set")
}
