/*
 * MIT License
 *
 * Copyright (C) 2022 Crimson Technologies, LLC. All rights reserved.
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
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/CrimsonAIO/aesccm"
)

const (
	// GenerationTimeKey is the JSON key for the generated time.
	GenerationTimeKey = "generationtime"

	// GenerationTimeFormat is the time format.
	// This is identical to time.RFC3339Nano except there is only three trailing zeros.
	GenerationTimeFormat = "2006-01-02T15:04:05.000Z07:00"

	// KeyNumber is the card number field key.
	KeyNumber = "number"

	// KeyExpiryMonth is the expiry month field key.
	KeyExpiryMonth = "expiryMonth"

	// KeyExpiryYear is the expiry year field key.
	KeyExpiryYear = "expiryYear"

	// KeySecurityCode is the security code field key.
	KeySecurityCode = "cvc"
)

// Encrypt encrypts a card number, security code (CVV/CVC), expiry month and year
// into a map and correctly formats all values using FormatCardNumber and FormatMonthYear.
func (enc *Encrypter) Encrypt(number, securityCode string, month, year int) (string, error) {
	m, y := FormatMonthYear(month, year)
	return enc.EncryptFields(map[string]string{
		KeyNumber:       FormatCardNumber(number),
		KeyExpiryMonth:  m,
		KeyExpiryYear:   y,
		KeySecurityCode: securityCode,
	})
}

// EncryptField encrypts a single key and value.
func (enc *Encrypter) EncryptField(key, value string) (string, error) {
	return enc.EncryptFields(map[string]string{key: value})
}

// EncryptFields encrypts a map.
func (enc *Encrypter) EncryptFields(fields map[string]string) (string, error) {
	if _, ok := fields[GenerationTimeKey]; !ok {
		fields[GenerationTimeKey] = enc.GetGenerationTime().Format(GenerationTimeFormat)
	}

	encoded, err := json.Marshal(fields)
	if err != nil {
		return "", err
	}
	return enc.EncryptPlaintext(encoded)
}

// EncryptPlaintext seals the given plaintext and returns the sealed content in the Adyen format.
//
// Most callers should use Encrypt, EncryptField or EncryptFields instead.
func (enc *Encrypter) EncryptPlaintext(plaintext []byte) (string, error) {
	// generate random nonce
	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// create ccm cipher
	ccm, err := aesccm.NewCCM(enc.block, len(nonce), 8)
	if err != nil {
		return "", err
	}

	// create ciphertext
	ciphertext := ccm.Seal(nil, nonce, plaintext, nil)
	// nonce with ciphertext
	nonceWithCiphertext := append(nonce, ciphertext...)

	// encrypted key using public key
	sealedKey, err := rsa.EncryptPKCS1v15(rand.Reader, enc.pubKey, enc.key[:])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"adyenjs_%s$%s$%s",
		enc.Version,
		base64.StdEncoding.EncodeToString(sealedKey),
		base64.StdEncoding.EncodeToString(nonceWithCiphertext),
	), nil
}
