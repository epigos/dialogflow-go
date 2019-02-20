# Go client for Dialogflow v2 API

[![Godoc](http://godoc.org/github.com/epigos/dialogflow-go?status.svg)](http://godoc.org/github.com/epigos/dialogflow-go)
[![Build Status](https://travis-ci.org/epigos/dialogflow-go.svg?branch=master)](https://travis-ci.org/epigos/dialogflow-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/epigos/dialogflow-go)](https://goreportcard.com/report/github.com/epigos/dialogflow-go)

This package allows integrating your Golang application with [Dialogflow v2](https://dialogflow.com/docs/)

## Installation

Standard go get:

    go get github.com/epigos/dialogflow-go

## Resource coverage

- Agents
- EntityTypes
- Entities
- Intents
- Sessions
- SessionContexts
- SessionEntityTypes

## Usage

Create a Client instance, providing your access token and the project ID you want to use:

```go
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
```

## Unit tests

    go test -race -v

## Documentation

- For full usage and examples see the [Godoc](http://godoc.org/github.com/epigos/dialogflow-go)
- [Dialogflow v2 API reference](https://cloud.google.com/dialogflow-enterprise/docs/reference/rest/v2-overview)

## Author

Philip Adzanoukpe [@epigos](https://twitter.com/@epigos)