![Sms77.io Logo](https://www.sms77.io/wp-content/uploads/2019/07/sms77-Logo-400x79.png "Sms77.io Logo")
# Official API Client for [Go](https://golang.org/)

## Installation
Requires Go 1.11+ for modules support.

```go get github.com/sms77io/go-client/sms77api```

### Usage
```go
package main

import (
	"fmt"
	"github.com/sms77io/go-client/sms77api"
)

func main() {
	var client = sms77api.New(sms77api.Options{
		ApiKey: "InsertSuperSecretSms77ApiKey!",
	})
	var balance, err = client.Balance.Get()
	if err == nil {
		fmt.Println(fmt.Sprintf("%f", *balance))
	}
}
```

#### Tests
Some basic tests are implemented.
Set environment variable `SMS77_API_KEY` for live API keys.
Set environment variable `SMS77_DUMMY_API_KEY` for sandbox API keys.
The dummy key takes preference if both are set.
Run all suites by running `go test`.

##### Support
Need help? Feel free to [contact us](https://www.sms77.io/en/company/contact/).

[![MIT](https://img.shields.io/badge/License-MIT-teal.svg)](./LICENSE)
