![](https://www.sms77.io/wp-content/uploads/2019/07/sms77-Logo-400x79.png "Sms77.io Logo")
# sms77io SMS Gateway API Client

## Installation

```go get github.com/sms77io/go-client```

### Usage
```go
package main

import (``
	"fmt"
	Client "github.com/sms77io/go-client"
)

func main() {
	var client = Client.New("MySuperSecretSms77ApiKey!")
	var balance, err = client.Balance()
	if err == nil {
		fmt.Sprintf("Balance: %f", balance)
	}
}
```