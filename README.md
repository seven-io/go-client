![Sms77.io Logo](https://www.sms77.io/wp-content/uploads/2019/07/sms77-Logo-400x79.png "Sms77.io Logo")
# Official API Client for the Sms77.io SMS Gateway 

## Installation

Requires Go 1.11+ for modules support.

```go get github.com/sms77io/go-client```

### Usage

```go
package main

import (
	"fmt"
	"github.com/sms77io/go-client"
)

func main() {
	var client = Sms77API.New("MySuperSecretSms77ApiKey!")
	var balance, err = client.Balance.Get()
	if err == nil {
		fmt.Printf("Balance: %f", balance)
	}
}
```

#### Support

Need help? Feel free to [contact us](https://www.sms77.io/en/company/contact/).

[![MIT](https://img.shields.io/badge/License-MIT-teal.svg)](./LICENSE)
