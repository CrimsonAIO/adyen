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

## Explained Usage:
```go
// Helper Function
func AdyenEncrypt(data map[string]string, encryptionKey string) (map[string]string, error) {

	// encryption key should be scraped from a specific endpoint response,
	// e.g: 10001|A237060180D24CDEF3E4E27D828BDB6A13E12C6959820770D7F2C1671DD0AEF4729670C20C6C5967C664D18955058B69549FBE8BF3609EF64832D7C033008A818700A9B0458641C5824F5FCBB9FF83D5A83EBDF079E73B81ACA9CA52FDBCAD7CD9D6A337A4511759FA21E34CD166B9BABD512DB7B2293C0FE48B97CAB3DE8F6F1A8E49C08D23A98E986B8A995A8F382220F06338622631435736FA064AEAC5BD223BAF42AF2B66F1FEA34EF3C297F09C10B364B994EA287A5602ACF153D0B4B09A604B987397684D19DBC5E6FE7E4FFE72390D28D6E21CA3391FA3CAADAD80A729FEF4823F6BE9711D4D51BF4DFCB6A3607686B34ACCE18329D415350FD0654D
	var plaintextKey = encryptionKey
	key := strings.Split(plaintextKey, "|")[1]
	fmt.Println("key: ", key)
	b, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}

	// create new encrypter
	enc, err := adyen.NewEncrypter("0_1_25", adyen.PubKeyFromBytes(b))
	if err != nil {
		panic(err)
	}

	encryptedCardNumber, _ := enc.EncryptField("encryptedCardNumber", data["cardNumber"])
	encryptedExpiryMonth, _ := enc.EncryptField("encryptedExpiryMonth", data["cardMonth"])
	encryptedExpiryYear, _ := enc.EncryptField("encryptedExpiryYear", data["cardYear"])
	encryptedSecurityCode, _ := enc.EncryptField("encryptedSecurityCode", data["cardCVV"])

	encryptedData := map[string]string{
		"encryptedCardNumber":   encryptedCardNumber,
		"encryptedExpiryMonth":  encryptedExpiryMonth,
		"encryptedExpiryYear":   encryptedExpiryYear,
		"encryptedSecurityCode": encryptedSecurityCode,
	}

	return encryptedData, nil
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
