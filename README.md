# adyen
Encrypt secrets for the Adyen payment platform.

This library uses `crypto/rand` to generate cryptographically secure AES keys and nonces,
and re-uses the same key and nonce for each client. Other publicly available libraries
typically use `math/rand` which is **insecure** for generating secret keys and nonces.

## Example

```go
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/CrimsonAIO/adyen"
)

func main() {
	// create a public key from the public key bytes
	//
	// if you have a key that looks like "10001|...", then you need to
	// hex decode the part after "|".
	// an example of this is shown here, minus removing the front part.
	const plaintextKey = "..."
	b, err := hex.DecodeString(plaintextKey)
	if err != nil {
		panic(err)
	}

	// create new encrypter
	enc, err := adyen.NewEncrypter("0_1_18", adyen.PubKeyFromBytes(b))
	if err != nil {
		panic(err)
	}

	// encrypt card information
	//
	// the number and month are automatically formatted with FormatCardNumber and
	// FormatMonthYear, so formatting doesn't matter.
	payload, err := enc.Encrypt(
		"4871049999999910",
		"737",
		3,
		2030,
	)
	if err != nil {
		panic(err)
	}

	// print the payload to send to the server
	fmt.Println(payload)
}
```
Check it out on [The Go Playground](https://go.dev/play/p/9tL6ziE52aw).

## Contributing
Pull requests are welcome to add new version constants or other improvements.
Note that you don't need to use one of our version constants; you can use any string you like.

If you open a pull request, please use our MIT copyright header.
If you're using GoLand (or any JetBrains IDE) you can do this by going in Settings -> Editor -> Copyright
and selecting the copyright profile found in `.idea/copyright/MIT_Crimson_Technologies_LLC.xml`.
You are welcome to add your own name for your contributions.