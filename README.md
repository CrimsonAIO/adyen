# adyen
Encrypt secrets for the Adyen payment platform.

This library uses `crypto/rand` to generate cryptographically secure AES keys and nonces,
and re-uses the same key and nonce for each client. Other publicly available libraries
typically use `math/rand` which is **insecure** for generating secret keys and nonces.

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

## Contributing
Pull requests are welcome to add new version constants or other improvements.
Note that you don't need to use one of our version constants; you can use any string you like.

If you open a pull request, please use our MIT copyright header.
If you're using GoLand (or any JetBrains IDE) you can do this by going in Settings -> Editor -> Copyright
and selecting the copyright profile found in `.idea/copyright/MIT_Crimson_Technologies_LLC.xml`.
You are welcome to add your own name for your contributions.