<img src="https://www.seven.io/wp-content/uploads/Logo.svg" width="250" />


# Official API Client for [Go](https://golang.org/)

## Installation

Requires Go 1.13+.

```go get github.com/seven-io/go-client/seven```

### Usage

```go
package main

import (
	"fmt"
	"github.com/seven-io/go-client/seven"
)

func main() {
	var client = seven.New(seven.Options{
		ApiKey: "InsertSuperSecretSevenApiKey!",
	})
	var balance, err = client.Balance.Get()
	if err == nil {
		fmt.Println(fmt.Sprintf("%f", *balance))
	} else {
		fmt.Println(err.Error())
	}
}
```

#### Tests

Some basic tests are implemented. Set environment variable `SEVEN_API_KEY` for live API keys. Set environment
variable `SEVEN_API_KEY_SANDBOX` for sandbox API keys. The dummy key takes preference if both are set. Run all suites by
running `go test`.

##### Support

Need help? Feel free to [contact us](https://www.seven.io/en/company/contact/).

[![MIT](https://img.shields.io/badge/License-MIT-teal.svg)](LICENSE)
