[![Go Report Card](https://goreportcard.com/badge/github.com/antfie/veracode-go-api)](https://goreportcard.com/report/github.com/antfie/veracode-go-api) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/antfie/veracode-go-api/blob/master/LICENSE)

# Veracode Go API

A Go version of the Veracode API.

## Running

Create a Go file with the following example contents:

```go
package main

import (
	api "github.com/antfie/veracode-go-api"
)

func main() {
	var apiKeyID = "YOUR_VERACODE_API_KEY_ID"
	var apiKeySecret = "YOUR_VERACODE_API_KEY_SECRET"

	response, err := api.GetApplicationList(apiKeyID, apiKeySecret)

	if err != nil {
		panic(err)
	}

	print(response)
}
```

Configure your `apiKeyID` and `apiKeySecret` values and then run.
