package dialogflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	token     = "token"
	projectID = "project-id"
)

func TestEntityTypeBatchDelete(t *testing.T) {
	assert := assert.New(t)

	srv := testServer([]byte(`{"name":"projects/projectId/agent/agentID","done":true}`))
	defer func() { srv.Close() }()

	c := NewClient(token, projectID)
	c.apiURL = srv.URL
	op, err := c.EntityTypeBatchDelete()
	assert.NoError(err)
	assert.Equal(op.Done, true)
}

func TestEntityTypeBatchUpdate(t *testing.T) {
	assert := assert.New(t)

	srv := testServer([]byte(`{"name":"projects/projectId/operations/id","done":true}`))
	defer func() { srv.Close() }()

	c := NewClient(token, projectID)
	c.apiURL = srv.URL
	op, err := c.EntityTypeBatchUpdate()
	assert.NoError(err)
	assert.Equal(op.Done, true)
}

func TestEntityTypeGet(t *testing.T) {
	assert := assert.New(t)

	srv := testServer([]byte(`{"name":"test","displayName":"display name"}`))
	defer func() { srv.Close() }()

	c := NewClient(token, projectID)
	c.apiURL = srv.URL
	entityType, err := c.EntityTypeGet("test")
	assert.NoError(err)
	assert.Equal(entityType.String(), "display name")
}

func TestEntityTypeList(t *testing.T) {
	assert := assert.New(t)

	srv := testServer([]byte(`{"entityTypes": [{"name": "name", "displayName": "display name"}] }`))
	defer func() { srv.Close() }()

	c := NewClient(token, projectID)
	c.apiURL = srv.URL
	entityTypes, err := c.EntityTypeList()
	assert.NoError(err)
	assert.Equal(entityTypes[0].String(), "display name")
}

func TestEntityTypeCreate(t *testing.T) {
	assert := assert.New(t)

	srv := testServer([]byte(`{"name":"test","displayName":"test-1"}`))
	defer func() { srv.Close() }()

	c := NewClient(token, projectID)
	c.apiURL = srv.URL
	entity, err := c.EntityTypeCreate(&EntityType{
		DisplayName: "test-1",
		Kind:        "KIND_MAP",
	})
	assert.NoError(err)
	assert.Equal(entity.String(), "test-1")
	assert.Equal(entity.Name, "test")
}

func TestEntityTypeUpdate(t *testing.T) {
	assert := assert.New(t)

	srv := testServer([]byte(`{"name":"projects/project-id/agent/entityTypes/test","displayName":"test-1"}`))
	defer func() { srv.Close() }()

	c := NewClient(token, projectID)
	c.apiURL = srv.URL
	entity, err := c.EntityTypePatch(&EntityType{
		DisplayName: "test-1",
		Kind:        "KIND_MAP",
	})
	assert.NoError(err)
	assert.Equal(entity.String(), "test-1")
	assert.Equal(entity.Name, "projects/project-id/agent/entityTypes/test")
}

func TestEntityTypeDelete(t *testing.T) {
	assert := assert.New(t)

	srv := testServer([]byte(`""`))
	defer func() { srv.Close() }()

	c := NewClient(token, projectID)
	c.apiURL = srv.URL
	err := c.EntityTypeDelete("projects/project-id/agent/entityTypes/test")
	assert.NoError(err)
}
