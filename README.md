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
	
	// encrypt card number
	encrypted, err := client.Encrypt(adyen.Version118, "number", "5555555555554444")
	if err != nil {
		panic(err)
	}
	
	// print encrypted card number 
	fmt.Println(encrypted)
}
```
Check it out on [The Go Playground](https://play.golang.org/p/6VqHCU6Fj50).