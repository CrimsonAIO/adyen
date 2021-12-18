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
	"time"

	"github.com/CrimsonAIO/aesccm"
)

const (
	// Version118 is the text representation of Adyen v1.18.
	Version118 = "0_1_18"

	// Version121 is the text representation of Adyen v1.21.
	Version121 = "0_1_21"

	// Version121 is the text representation of Adyen v1.25.
	Version125 = "0_1_25"
)

// GenerationTimeFunc is a function responsible for returning the correct "generationtime" key
// given a version and value.
type GenerationTimeFunc func(version string, value map[string]interface{}) time.Time

// GenerationTimeNow returns time.Now as the generated time.
var GenerationTimeNow GenerationTimeFunc = func(version string, value map[string]interface{}) time.Time {
	return time.Now()
}

// Encrypt encrypts the specified value and returns it in the correct format.
//
// This function is similar to EncryptSingle, except the value type is a map.
// This allows encryption of multiple values.
// For example, you might have the following map:
//		{
//			"number": "0",
// 			"expiryMonth": 1
//		}
//
// This map will be encrypted as one string, with the "generationtime" key being inserted
// into the map before encryption.
//
// The getGenerationTime function is used to obtain the aforementioned "generationtime" key.
// In most cases, GenerationTimeNow will suffice.
// Some websites require different generation times.
func (c *client) Encrypt(version string, value map[string]interface{}, getGenerationTime GenerationTimeFunc) (string, error) {
	// insert generated time
	value["generationtime"] = getGenerationTime(version, value).Format("2006-01-02T15:04:05.000Z07:00")

	encoded, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

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
	cipherText := ccm.Seal(nil, c.AESNonce, encoded, nil)
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

// EncryptSingle is like Encrypt, but is easier to use when only one value is being encrypted.
func (c *client) EncryptSingle(version, name string, value interface{}, getGenerationTime GenerationTimeFunc) (string, error) {
	return c.Encrypt(version, map[string]interface{}{
		name: value,
	}, getGenerationTime)
}
