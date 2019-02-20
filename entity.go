package dialogflow

import (
	"fmt"
	"strings"
)

// EntityType represent dialogflow.com Entity type
type EntityType struct {
	Name              string   `json:"name"`
	DisplayName       string   `json:"displayName"`
	Kind              string   `json:"kind"`
	AutoExpansionMode string   `json:"autoExpansionMode,omitempty"`
	Entities          []Entity `json:"entities,omitempty"`
}

func (e *EntityType) String() string {
	return e.DisplayName
}

func (e *EntityType) getID() string {
	if e.Name == "" {
		return e.Name
	}
	ns := strings.Split(e.Name, "/")
	return ns[len(ns)-1]
}

// Entity represents dialogflow.com Entity type
type Entity struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}

// EntityTypeBatchUpdate Updates entity types in the specified agent
func (c *Client) EntityTypeBatchUpdate(types ...EntityType) (*Operation, error) {
	data := map[string]interface{}{
		"entityTypeBatchInline": map[string]interface{}{
			"entityTypes": types,
		},
	}

	var op Operation
	err := c.post("/entityTypes:batchUpdate", data, &op)
	return &op, err
}

// EntityTypeBatchDelete Deletes entity types in the specified agent
func (c *Client) EntityTypeBatchDelete(names ...string) (*Operation, error) {
	data := map[string][]string{
		"entityTypeNames": names,
	}
	var op Operation
	err := c.post("/entityTypes:batchDelete", data, &op)
	return &op, err
}

// EntityTypeGet Retrieves the specified entity type
func (c *Client) EntityTypeGet(name string) (*EntityType, error) {
	var entityType EntityType

	resource := fmt.Sprintf("/entityTypes/%s", name)
	err := c.get(resource, &entityType)

	return &entityType, err
}

// EntityTypeList Returns the list of all entity types in the specified agent
func (c *Client) EntityTypeList() ([]*EntityType, error) {
	resp := map[string][]*EntityType{
		"entityTypes": []*EntityType{},
	}
	err := c.get("/entityTypes", &resp)
	return resp["entityTypes"], err
}

// EntityTypeCreate Creates an entity type in the specified agent.
func (c *Client) EntityTypeCreate(entity *EntityType) (*EntityType, error) {
	err := c.post("/entityTypes", entity, entity)
	return entity, err
}

// EntityTypePatch Updates the specified entity type.
func (c *Client) EntityTypePatch(entity *EntityType) (*EntityType, error) {
	resource := fmt.Sprintf("/entityTypes/%s", entity.getID())

	err := c.put(resource, entity, entity)
	return entity, err
}

// EntityTypeDelete Deletes the specified entity type
func (c *Client) EntityTypeDelete(name string) error {
	return c.delete(fmt.Sprintf("/entityTypes/%s", name))
}
