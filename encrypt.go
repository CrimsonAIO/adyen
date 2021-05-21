/*
 * MIT License
 *
 * Copyright (C) 2021 Crimson Technologies LLC. All rights reserved.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package adyen

import (
	"crypto/aes"
	crand "crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/CrimsonAIO/aesccm"
	"time"
)

const (
	// Version118 is the text representation of Adyen v1.18.
	Version118 = "0_1_18"
)

// Encrypt encrypts the specified value and returns it in the correct format.
//
// The version should be a valid Adyen version, such as Version118.
//
// The name should be the JSON field name. For encrypting a card number,
// this should be "number". The plain text result, if the value was "0", would be:
// 		{
//			"number": "0",
//			"generatedtime": [generated time]
//		}
//
// The value is the actual content, for example, a card number.
// The value will be correctly encoded to JSON, regardless of its type.
// For example, if the type is a map, it would be formatted as a JSON map.
// The type is typically a string.
func (c *client) Encrypt(version, name string, value interface{}) (string, error) {
	// Create the JSON result from a custom field name and value.
	// encode the value to JSON
	encoded, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	// convert to adyen format JSON
	plainText := fmt.Sprintf("{\"%s\":%s,\"generationtime\":\"%s\"}", name, string(encoded), time.Now().Format("2006-01-02T15:04:05.000Z07:00"))

	// Encrypt the plain text content.

	// create new AES cipher block
	block, err := aes.NewCipher(c.AESKey)
	if err != nil {
		return "", err
	}

	// create new CCM mode cipher
	ccm, err := aesccm.NewCCM(block, len(c.AESNonce), 8)
	if err != nil {
		return "", err
	}

	// encrypt content
	cipherText := ccm.Seal(nil, c.AESNonce, []byte(plainText), nil)
	// nonce with cipher text
	nonceWithCipherText := append(c.AESNonce, cipherText...)

	// encrypt the AES key using the site's RSA public key
	encryptedKey, err := rsa.EncryptPKCS1v15(crand.Reader, c.SiteKey, c.AESKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("adyenjs_%s$%s$%s",
		version,
		base64.StdEncoding.EncodeToString(encryptedKey),
		base64.StdEncoding.EncodeToString(nonceWithCipherText),
	), nil
}
