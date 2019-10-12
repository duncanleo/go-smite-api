# go-smite-api
[![GoDoc](https://godoc.org/github.com/duncanleo/go-smite-api/smite?status.svg)](https://godoc.org/github.com/duncanleo/go-smite-api/smite)
This is an API client library for the SMITE API (by Hi-Rez Studios).

### Usage
```go
import (
  "log"

  "github.com/duncanleo/go-smite-api/smite"
)

func main() {
  client := smite.Client{
    DevID: "1234",
    AuthKey: "abc123"
  }

  createSessionResponse, err := c.CreateSession()
	if err != nil {
		log.Panic(err)
  }
  log.Printf("Session created: %s", createSessionResponse.SessionID)
}
```
