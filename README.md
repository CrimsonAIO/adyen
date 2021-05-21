# adyen
Encrypt secrets compatible with the Adyen payment platform.

## Example

```go
package main

import (
	"fmt"
	"github.com/CrimsonAIO/adyen"
	"os"
)

func main() {
	// create new client with a specific site key.
	client, err := adyen.NewClient(os.Getenv("ADYEN_SITE_KEY"))
	if err != nil {
		panic(err)
    }
    
    // card number to encrypt
    cardNumber := "000"
    
    cardNumber, err := client.Encrypt(adyen.Version1_18, "number", cardNumber)
    if err != nil {
    	panic(err)
    }
    
    // print encrypted card number
    fmt.Println(cardNumber)
}
```