/*
Package dialogflow is a Go client for interacting with the Dialogflow.com version 2 API.

Example:

Create a Client instance, providing your access token and the project ID you want to use:
  package main

  import (
    "fmt"
    "os"
    dialogflow "github.com/epigos/dialogflow-go"
  )

  func main() {
    // dialogflow access token
	token := os.Getenv("DIALOGFLOW_ACCESS_TOKEN")
	projectID := os.Getenv("DIALOGFLOW_PROJECT_ID")

    // dialogflow client
    client := dialogflow.NewClient(token, projectID)

    // retrieve all entity types
    entityTypes, err := client.EntityTypeList()

    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(entityTypes)
  }
Learn more about Dialogflow API https://dialogflow.com/docs
*/
package dialogflow
